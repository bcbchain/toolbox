package everycolorstub

import (
	bcType "blockchain/types"
	"strings"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"

	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	"contract/orgexample/code/everycolor/v3.0/everycolor"
	"github.com/tendermint/tmlibs/log"
)

//EverycolorStub an object
type EverycolorStub struct {
	logger log.Logger
}

var _ stubType.IContractStub = (*EverycolorStub)(nil)

//New generate a stub
func New(logger log.Logger) stubType.IContractStub {
	return &EverycolorStub{logger: logger}
}

//FuncRecover recover panic by Assert
func FuncRecover(response *bcType.Response) {
	if err := recover(); err != nil {
		if e, ok := err.(types.Error); ok {
			response.Code = e.ErrorCode
			response.Log = e.Error()
		} else if e, ok := err.(error); ok {
			if strings.HasPrefix(e.Error(), "runtime error") {
				response.Code = types.ErrStubDefined
				response.Log = e.Error()
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

// InitChain initial smart contract
func (pbs *EverycolorStub) InitChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (pbs *EverycolorStub) UpdateChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (pbs *EverycolorStub) Mine(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	response.Code = types.CodeOK
	return response
}

//Invoke invoke function
func (pbs *EverycolorStub) Invoke(smc sdk.ISmartContract) (response bcType.Response) {
	return pbs.InvokeInternal(smc, true)
}

//InvokeInternal invoke function
func (pbs *EverycolorStub) InvokeInternal(smc sdk.ISmartContract, feeFlag bool) (response bcType.Response) {
	defer FuncRecover(&response)

	// 生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(smc, feeFlag)
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(smc.Message(), response.Tags, "", fee, gasUsed, smc.Tx().GasLimit(), err)
		return
	}

	var data string
	err = types.Error{ErrorCode: types.CodeOK}
	switch smc.Message().MethodID() {
	case "d373a935": // prototype: SetSecretSigner(types.PubKey)
		setSecretSigner(smc)
	case "44f1b25c": // prototype: SetSettings(string)
		setSettings(smc)
	case "e4214bb2": // prototype: SetRecvFeeInfo(string)
		setRecvFeeInfo(smc)
	case "948c4d24": // prototype: WithdrawFunds(string,types.Address,bn.Number)
		withdrawFunds(smc)
	case "86307b80": // prototype: PlaceBet(string,bn.Number,string,int64,[]byte,[]byte,types.Address)
		placeBet(smc)
	case "1e89776f": // prototype: SettleBets([]byte,int64)
		settleBets(smc)
	case "d32a5fe6": // prototype: WithdrawWin([]byte)
		withdrawWin(smc)
	case "ac8151c9": // prototype: RefundBets([]byte,int64)
		refundBets(smc)
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(smc.Message(), response.Tags, data, fee, gasUsed, smc.Tx().GasLimit(), err)
	return
}

func setSecretSigner(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 types.PubKey
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.SetSecretSigner(v0)
}

func setSettings(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.SetSettings(v0)
}

func setRecvFeeInfo(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.SetRecvFeeInfo(v0)
}

func withdrawFunds(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 3, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v1 types.Address
	err = rlp.DecodeBytes(items[1], &v1)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v2 bn.Number
	err = rlp.DecodeBytes(items[2], &v2)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.WithdrawFunds(v0, v1, v2)
}

func placeBet(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 7, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v1 bn.Number
	err = rlp.DecodeBytes(items[1], &v1)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v2 string
	err = rlp.DecodeBytes(items[2], &v2)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v3 int64
	err = rlp.DecodeBytes(items[3], &v3)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v4 []byte
	err = rlp.DecodeBytes(items[4], &v4)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v5 []byte
	err = rlp.DecodeBytes(items[5], &v5)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v6 types.Address
	err = rlp.DecodeBytes(items[6], &v6)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.PlaceBet(v0, v1, v2, v3, v4, v5, v6)
}

func settleBets(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 2, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 []byte
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v1 int64
	err = rlp.DecodeBytes(items[1], &v1)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.SettleBets(v0, v1)
}

func withdrawWin(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 []byte
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.WithdrawWin(v0)
}

func refundBets(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 2, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 []byte
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v1 int64
	err = rlp.DecodeBytes(items[1], &v1)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(everycolor.Everycolor)
	contractObj.SetSdk(smc)
	contractObj.RefundBets(v0, v1)
}
