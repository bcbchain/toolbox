package genstub

import (
	"blockchain/smccheck/parsecode"
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

var templateText1 = `package types

import (
	"blockchain/smcsdk/sdk"
	"blockchain/types"
)

type IContractStub interface {
	InitChain(smcapi sdk.ISmartContract) types.Response
	UpdateChain(smcapi sdk.ISmartContract) types.Response
	Invoke(smcapi sdk.ISmartContract) types.Response
}

type IContractIntfcStub interface {
	Invoke(methodid string, p interface{}) types.Response
	GetSdk() sdk.ISmartContract
	SetSdk(smc sdk.ISmartContract)
}
`

var templateText2 = `package common

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/crypto/sha3"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl"
	"blockchain/smcsdk/sdkimpl/object"
	types2 "blockchain/types"
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/tendermint/tmlibs/common"
)

//CreateResponse create response data
func CreateResponse(message sdk.IMessage, oriTags []common.KVPair, data string, fee, gasUsed, gasLimit int64, err types.Error) (response types2.Response) {
	response.Code = err.ErrorCode
	response.Data = data
	response.Fee = fee
	response.Log = err.Error()
	response.GasLimit = gasLimit
	response.GasUsed = gasUsed
	if oriTags != nil{
		response.Tags = oriTags
	}
	for _, v := range message.(*object.Message).OutputReceipts() {
		tag := common.KVPair{}
		tag.Value = v.Value
		tag.Key = v.Key
		response.Tags = append(response.Tags, tag)
	}
	return
}

//FeeAndReceipt pay fee for the calling and emit fee receipt
func FeeAndReceipt(smc sdk.ISmartContract, bMethod bool) (fee, gasUsed int64, receipt types.KVPair, err types.Error) {

	err.ErrorCode = types.CodeOK
	//Get gas price
	var gasprice int64
	if smc.Message().Contract().Token() == "" {
		gasprice = smc.Helper().TokenHelper().BaseGasPrice()
	} else {
		gasprice = smc.Helper().TokenHelper().Token().GasPrice()
	}
	//calculate fee
	var methods []std.Method
	if bMethod {
		methods = smc.Message().Contract().Methods()
	} else {
		methods = smc.Message().Contract().Interfaces()
	}
	var gas int64
	for _, m := range methods {
		if m.MethodID == smc.Message().MethodID() {
			gas = m.Gas
			break
		}
	}
	gasAbs := int64(math.Abs(float64(gas))) //abs number

	gasLeft := smc.Tx().GasLeft()
	if gasLeft < gasAbs {
		gasUsed = gasLeft
		err.ErrorCode = types.ErrGasNotEnough
	} else {
		gasUsed = gasAbs
	}
	fee = gasprice * gasUsed

	//negative gas means contract account is the payer
	//positive gas means tx signer is the payer
	//check and set payer's balance
	payer := smc.Tx().Signer()
	if gas < 0 {
		payer = smc.Helper().AccountHelper().AccountOf(smc.Message().Contract().Account())
	}
	token := smc.Helper().GenesisHelper().Token().Address()
	balance := payer.BalanceOfToken(token)
	if balance.IsLessThanI(fee) {
		fee = balance.V.Int64()
		balance = bn.N(0)
		gasUsed = fee/gasprice
		err.ErrorCode = types.ErrInsufficientBalance
	} else {
		balance = balance.SubI(fee)
	}
	payer.(*object.Account).SetBalanceOfToken(token, balance)

	//Set gasLeft to tx
	gasLeft = gasLeft - gasUsed
	smc.Tx().(*object.Tx).SetGasLeft(gasLeft)
	//emit receipt
	feeReceipt := std.Fee{
		Token: smc.Helper().GenesisHelper().Token().Address(),
		From:  payer.Address(),
		Value: fee,
	}
	receipt = emitFeeReceipt(smc, feeReceipt)

	return
}

func CalcKey(name, version string) string {
	if strings.HasPrefix(name, "token-template-") {
		name = "token-issue"
	}
	return strings.Replace(name, "-", "", -1) + strings.Replace(version, ".", "", -1)
}

func emitFeeReceipt(smc sdk.ISmartContract,receipt std.Fee) types.KVPair {
	bz, err := jsoniter.Marshal(receipt)
	if err != nil {
		sdkimpl.Logger.Fatalf("[sdk]Cannot marshal receipt data=%v", receipt)
		sdkimpl.Logger.Flush()
		panic(err)
	}

	rcpt := std.Receipt{
		Name:         receiptName(receipt),
		ContractAddr: smc.Message().Contract().Address(),
		Bytes:        bz,
		Hash:         nil,
	}
	rcpt.Hash = sha3.Sum256([]byte(rcpt.Name), []byte(rcpt.ContractAddr), bz)
	resBytes, _ := jsoniter.Marshal(rcpt) // nolint unhandled

	result := types.KVPair{
		Key:   []byte(fmt.Sprintf("/%d/%s", len(smc.Message().(*object.Message).OutputReceipts()), rcpt.Name)),
		Value: resBytes,
	}

	return result
}

func receiptName(receipt interface{}) string {
	typeOfInterface := reflect.TypeOf(receipt).String()

	if strings.HasPrefix(typeOfInterface, "std.") {
		prefixLen := len("std.")
		return "std::" + strings.ToLower(typeOfInterface[prefixLen:prefixLen+1]) + typeOfInterface[prefixLen+1:]
	}

	return typeOfInterface
}
`

// GenStubCommon - generate the stub common go source
func GenStubCommon(rootDir string) error {

	err := genTypes(rootDir)
	if err != nil {
		return err
	}

	err = genCommon(rootDir)

	return err
}

func genTypes(rootDir string) error {
	newPath := filepath.Join(rootDir, "types")
	if err := os.MkdirAll(newPath, os.FileMode(0750)); err != nil {
		return err
	}
	filename := filepath.Join(newPath, "types.go")

	tmpl, err := template.New("types").Parse(templateText1)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, nil); err != nil {
		return err
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}

	return nil
}

func genCommon(rootDir string) error {
	newPath := filepath.Join(rootDir, "common")
	if err := os.MkdirAll(newPath, os.FileMode(0750)); err != nil {
		return err
	}
	filename := filepath.Join(newPath, "common.go")

	tmpl, err := template.New("common").Parse(templateText2)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, nil); err != nil {
		return err
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}

	return nil
}
