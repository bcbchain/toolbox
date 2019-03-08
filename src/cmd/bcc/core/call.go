package core

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdkimpl/helper"
	"blockchain/tx2"
	"blockchain/types"
	"cmd/bcc/cache"
	common2 "cmd/bcc/common"
	"cmd/bcc/pvar"
	"common/wal"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tendermint/go-crypto"
	"github.com/tendermint/tmlibs/common"
)

// Call call contract's method with params, save params to paramsFile if it's length too long
func Call(name, password string, bccParams CallParam) (result *CommitTxResult, err error) {

	result, err = call(name, password, bccParams.OrgName, bccParams.Contract, bccParams.Method, bccParams.ParamsFile,
		bccParams.Params, bccParams.SplitBy, bccParams.Pay, bccParams.GasLimit, bccParams.Note, bccParams.KeyStorePath,
		bccParams.ChainID, false, false)
	if err != nil {
		return
	}

	nonceErrDesc := "Invalid nonce"
	smcErrDesc := "The contract has expired"
	if result.Code != types.CodeOK {
		if result.Log == nonceErrDesc {
			result, err = call(name, password, bccParams.OrgName, bccParams.Contract, bccParams.Method, bccParams.ParamsFile,
				bccParams.Params, bccParams.SplitBy, bccParams.Pay, bccParams.GasLimit, bccParams.Note, bccParams.KeyStorePath,
				bccParams.ChainID, true, false)
		} else if result.Log == smcErrDesc {
			result, err = call(name, password, bccParams.OrgName, bccParams.Contract, bccParams.Method, bccParams.ParamsFile,
				bccParams.Params, bccParams.SplitBy, bccParams.Pay, bccParams.GasLimit, bccParams.Note, bccParams.KeyStorePath,
				bccParams.ChainID, false, false)
		}
	}

	return
}

func call(name, password, orgName, contractName, methodName, file, params,
	splitBy, pay, gasLimit, note, keyStorePath, chainID string, bNonceErr, bSmcErr bool) (result *CommitTxResult, err error) {

	if splitBy == "" {
		splitBy = "@"
	}

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		chainID = common2.GetBCCConfig().DefaultChainID
	}
	tx2.Init(chainID)

	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return nil, errors.New("getNonce error: " + err.Error())
	}

	contract, err := getContract(orgName, contractName, chainID, bSmcErr, keyStorePath)
	if err != nil {
		return nil, errors.New("getContract error: " + err.Error())
	}

	var item std.Method
	for _, methodItem := range contract.Methods {
		if strings.HasPrefix(methodItem.ProtoType, methodName+"(") {
			item = methodItem
			break
		}
	}

	if len(item.MethodID) == 0 {
		return nil, errors.New("invalid method")
	}

	rlpBytes, err := encode(item, splitBy, file, params)
	if err != nil {
		return
	}

	uGasLimit, err := strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		return
	}

	methodID, err := strconv.ParseUint(item.MethodID, 16, 64)
	if err != nil {
		return
	}

	var msgList []types.Message

	if len(pay) != 0 {
		transferMsg, err := createTransferMsg(contract, pay, chainID)
		if err != nil {
			return nil, err
		}
		msgList = append(msgList, transferMsg)
	}

	msg := types.Message{Contract: contract.Address, MethodID: uint32(methodID), Items: rlpBytes}
	msgList = append(msgList, msg)

	pl := tx2.WrapPayload(nonce, int64(uGasLimit), note, msgList...)

	priKeyHex, err := getAccountPriKey(keyStorePath, name, password)
	if err != nil {
		return
	}
	txStr := tx2.WrapTx(pl, priKeyHex)

	result, err = CommitTx(chainID, txStr)
	if err != nil {
		return
	}

	return
}

func getNonce(keyStorePath, chainID, name, password string, bNonceErr bool) (nonce uint64, err error) {

	nonce, err = cache.Nonce(name, keyStorePath)
	if err != nil || bNonceErr {
		var nonceResult *NonceResult

		nonceResult, err = Nonce("", name, password, chainID, keyStorePath)
		if err != nil {
			return
		}
		nonce = nonceResult.Nonce

		err = cache.SetNonce(name, nonce, keyStorePath)
		if err != nil {
			return
		}
	}

	return
}

func getContract(orgName, contractName, chainID string, bSmcErr bool, keyStorePath string) (contract *std.Contract, err error) {
	bh := helper.BlockChainHelper{}
	orgID := bh.CalcOrgID(orgName)

	contract, err = cache.Contract(orgID, contractName, keyStorePath)
	if err != nil || bSmcErr {

		contract, err = contractOfName(chainID, orgID, contractName)
		if err != nil {
			return
		}

		err = cache.SetContract(contract, keyStorePath)
		if err != nil {
			return
		}
	}

	return
}

func encode(method std.Method, splitBy, file, params string) (rlpBytes []common.HexBytes, err error) {
	if len(file) != 0 {
		// 如果文件存在且能够正确读取信息，则优先使用文件中的内容
		var temp string
		temp, err = readParamFile(file)
		if err != nil {
			return
		}
		params = temp
	}

	varList, err := pvar.Create(method, params, splitBy)
	if err != nil {
		return
	}

	rlpBytes = tx2.WrapInvokeParams(varList...)

	return
}

func readParamFile(file string) (params string, err error) {

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	params = string(b)

	return
}

func createTransferMsg(contract *std.Contract, pay, chainID string) (transferMsg types.Message, err error) {

	leftBracketIndex := strings.Index(pay, "(")
	value := pay[:leftBracketIndex]
	valueLen := len(value)
	if valueLen != 0 {
		dotIndex := strings.Index(value, ".")
		value = strings.Replace(value, ".", "", -1)
		var zeroCount = 9
		if dotIndex > 0 {
			zeroCount = 9 - (valueLen - dotIndex - 1)
		}
		for zeroCount > 0 {
			value += "0"
			zeroCount--
		}
	}
	rlpBytes := tx2.WrapInvokeParams(contract.Account, bn.NewNumberStringBase(value, 10))

	tokenName := pay[leftBracketIndex+1 : len(pay)-1]
	tokenContract, err := contractOfTokenName(chainID, tokenName)
	if err != nil {
		return
	}

	transferMsg = types.Message{Contract: tokenContract.Address, MethodID: 0x44d8ca60, Items: rlpBytes}

	return
}

func getAccountPriKey(keyStorePath, name, password string) (priKeyHex string, err error) {

	acct, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		return
	}

	priKey := acct.PrivateKey.(crypto.PrivKeyEd25519)

	return "0x" + hex.EncodeToString(priKey[:]), nil
}
