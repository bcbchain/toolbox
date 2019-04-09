package core

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdkimpl/helper"
	"blockchain/tx2"
	"blockchain/types"
	"cmd/bcc/cache"
	"cmd/bcc/pvar"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/tendermint/tmlibs/common"
)

// Call call contract's method with params, save params to paramsFile if it's length too long
func Call(name, password string, bccParams CallParam) (result *CommitTxResult, err error) {

	defer FuncRecover(&err)

	// reset value or not
	splitBy, keyStorePath, chainID := prepare(bccParams.SplitBy, bccParams.KeyStorePath, bccParams.ChainID)

	// require not empty
	requireNotEmpty("orgName", bccParams.OrgName)
	requireNotEmpty("contractName", bccParams.Contract)
	requireNotEmpty("name", name)
	requireNotEmpty("password", password)
	requireNotEmpty("methodName", bccParams.Method)

	// check pay
	value, token, err := checkPay(bccParams.Pay)
	if err != nil {
		return
	}

	result, err = call(name, password, bccParams.OrgName, bccParams.Contract, bccParams.Method, bccParams.ParamsFile,
		bccParams.Params, splitBy, token, value, bccParams.GasLimit, bccParams.Note, keyStorePath,
		chainID, false, false)
	if err != nil {
		return
	}

	var count = 0
	for result.Code != types.CodeOK && count < 2 {
		if result.Log == nonceErrDesc {
			result, err = call(name, password, bccParams.OrgName, bccParams.Contract, bccParams.Method, bccParams.ParamsFile,
				bccParams.Params, splitBy, token, value, bccParams.GasLimit, bccParams.Note, keyStorePath,
				chainID, true, false)
		} else if result.Log == smcErrDesc {
			result, err = call(name, password, bccParams.OrgName, bccParams.Contract, bccParams.Method, bccParams.ParamsFile,
				bccParams.Params, splitBy, token, value, bccParams.GasLimit, bccParams.Note, keyStorePath,
				chainID, false, false)
		}

		count++
	}

	return
}

func call(name, password, orgName, contractName, methodName, file, params, splitBy, token string, value bn.Number,
	gasLimit, note, keyStorePath, chainID string, bNonceErr, bSmcErr bool) (result *CommitTxResult, err error) {

	// get account transaction nonce
	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return nil, errors.New("getNonce error: " + err.Error())
	}

	// get contract information with orgName and contractName
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

	// encode method parameters
	rlpBytes, err := encode(item, splitBy, file, params)
	if err != nil {
		return
	}

	uGasLimit, err := requireUint64("gasLimit", gasLimit, 10)
	if err != nil {
		return
	}
	methodID, _ := requireUint64("methodID", item.MethodID, 16)

	var msgList []types.Message

	// pack tx
	// if pay option not empty, then create transfer message
	if value.IsGreaterThanI(0) {
		transferMsg, err := createTransferMsg(contract, value, token, chainID)
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

	// commit transaction
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

// 生成转账message
func createTransferMsg(contract *std.Contract, value bn.Number, token, chainID string) (transferMsg types.Message, err error) {

	rlpBytes := tx2.WrapInvokeParams(contract.Account, value)

	tokenContract, err := contractOfTokenName(chainID, token)
	if err != nil {
		return
	}

	transferMsg = types.Message{Contract: tokenContract.Address, MethodID: 0x44d8ca60, Items: rlpBytes}

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

// 查询合约信息
func QueryContractInfo(OrgName, ContractName, chainID, keyStorePath string) (contract *std.Contract, err error) {

	contract, err = getContract(OrgName, ContractName, chainID, false, keyStorePath)
	if err != nil {
		return
	}

	if contract.Methods == nil {
		contract, err = getContract(OrgName, ContractName, chainID, true, keyStorePath)
		if err != nil {
			return
		}
	}

	return
}
