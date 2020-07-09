package core

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bcbchain/bcbchain/hyperledger/burrow/execution/bvm/abi"
	"github.com/bcbchain/bclib/algorithm"
	"github.com/bcbchain/bclib/jsoniter"
	"github.com/bcbchain/bclib/tx/v2"
	"github.com/bcbchain/bclib/types"
	"github.com/bcbchain/sdk/sdk/bn"
	"github.com/bcbchain/sdk/sdk/crypto/sha3"
	"github.com/bcbchain/sdk/sdk/std"
	types2 "github.com/bcbchain/sdk/sdk/types"
	"github.com/bcbchain/sdk/sdkimpl/helper"
	"github.com/bcbchain/toolbox/bcc/common"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	genesisOrgName = "genesis"
)

func RegisterOrg(name, password string, bccParams RegisterOrgParam) (result *CommitTxResult, err error) {

	defer FuncRecover(&err)

	contractName := "organization"
	_, keyStorePath, chainID := prepare("", bccParams.KeyStorePath, bccParams.ChainID)

	values := make([]interface{}, 0)
	values = append(values, bccParams.OrgName)

	//查询方法ID
	methodID, err := QueryMethodID("genesis", contractName, "RegisterOrganization", chainID, keyStorePath, false)
	if err != nil {
		return
	}

	result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, false, false, methodID, values)
	if err != nil {
		return
	}

	var count = 0
	for result.Code != types.CodeOK && count < 2 {
		if result.Log == nonceErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, true, false, methodID, values)
		} else if result.Log == smcErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, false, true, methodID, values)
		}

		count++
	}

	if result.Code == types.CodeOK {
		orgIDs := new([]string)
		err = json.Unmarshal([]byte(result.Data), orgIDs)
		if err != nil || len(*orgIDs) != 1 {
			return
		}
		result.OrgID = (*orgIDs)[0]
		result.Data = ""
	}

	return
}

func SetOrgSigners(name, password string, bccParams SetOrgSignersParam) (result *CommitTxResult, err error) {

	defer FuncRecover(&err)

	contractName := "organization"
	_, keyStorePath, chainID := prepare("", bccParams.KeyStorePath, bccParams.ChainID)

	// require not empty
	requireNotEmpty("orgName", bccParams.OrgName)
	requireNotEmpty("pubKeys", bccParams.PubKeys)

	pubKeys := make([]types2.HexBytes, 0)
	var pubKeyStrs []string
	err = jsoniter.Unmarshal([]byte(bccParams.PubKeys), &pubKeyStrs)
	if err != nil {
		return
	}

	for _, item := range pubKeyStrs {
		temp, err := hex.DecodeString(item[2:])
		if err != nil {
			return nil, err
		}
		pubKeys = append(pubKeys, temp)
	}

	bh := helper.BlockChainHelper{}
	orgID := bh.CalcOrgID(bccParams.OrgName)

	values := make([]interface{}, 0)
	values = append(values, orgID)
	values = append(values, pubKeys)

	//查询方法ID
	methodID, err := QueryMethodID("genesis", contractName, "SetSigners", chainID, keyStorePath, false)
	if err != nil {
		return
	}

	result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, false, false, methodID, values)
	if err != nil {
		return
	}

	var count = 0
	for result.Code != types.CodeOK && count < 2 {
		if result.Log == nonceErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, true, false, methodID, values)
		} else if result.Log == smcErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, false, true, methodID, values)
		}

		count++
	}

	return
}

func SetOrgDeployer(name, password string, bccParams SetOrgDeployerParam) (result *CommitTxResult, err error) {

	defer FuncRecover(&err)

	contractName := "smartcontract"
	_, keyStorePath, chainID := prepare("", bccParams.KeyStorePath, bccParams.ChainID)

	// require not empty
	requireNotEmpty("orgName", bccParams.OrgName)

	err = algorithm.CheckAddress(chainID, bccParams.Deployer)
	if err != nil {
		return
	}

	bh := helper.BlockChainHelper{}
	orgID := bh.CalcOrgID(bccParams.OrgName)

	values := make([]interface{}, 0)
	values = append(values, bccParams.Deployer)
	values = append(values, orgID)

	//查询方法ID
	methodID, err := QueryMethodID("genesis", contractName, "Authorize", chainID, keyStorePath, false)
	if err != nil {
		return
	}

	result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, false, false, methodID, values)
	if err != nil {
		return
	}

	var count = 0
	for result.Code != types.CodeOK && count < 2 {
		if result.Log == nonceErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, false, false, methodID, values)
		} else if result.Log == smcErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID, false, false, methodID, values)
		}

		count++
	}

	return
}

func DeployContract(name, password string, bccParams DeployContractParam) (result *CommitTxResult, err error) {

	defer FuncRecover(&err)

	contractName := "smartcontract"
	_, keyStorePath, chainID := prepare("", bccParams.KeyStorePath, bccParams.ChainID)

	// require not empty
	requireNotEmpty("name", name)
	requireNotEmpty("password", password)
	requireNotEmpty("contractName", bccParams.ContractName)
	requireNotEmpty("orgName", bccParams.OrgName)

	// check arguments
	err = algorithm.CheckAddress(chainID, bccParams.Owner)
	if err != nil {
		return
	}

	err = checkVersion(bccParams.Version)
	if err != nil {
		return
	}

	effectHeight, orgID, codeHash, codeData, devSig, orgSig, err := getDeployContractData(bccParams, chainID)
	if err != nil {
		return
	}

	values := make([]interface{}, 0)
	values = append(values, bccParams.ContractName)
	values = append(values, bccParams.Version)
	values = append(values, orgID)
	values = append(values, codeHash)
	values = append(values, codeData)
	values = append(values, devSig)
	values = append(values, orgSig)
	values = append(values, effectHeight)
	values = append(values, bccParams.Owner)

	//查询方法ID
	methodID, err := QueryMethodID("genesis", contractName, "DeployContract", chainID, keyStorePath, false)
	if err != nil {
		return
	}

	result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID,
		false, false, methodID, values)
	if err != nil {
		return
	}

	var count = 0
	for result.Code != types.CodeOK && count < 2 {
		if result.Log == nonceErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID,
				true, false, methodID, values)
		} else if result.Log == smcErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID,
				false, true, methodID, values)
		}

		count++
	}

	if result.Code == types.CodeOK {
		addrList := new([]string)
		err = json.Unmarshal([]byte(result.Data), addrList)
		if err != nil || len(*addrList) != 1 {
			return
		}
		result.SmcAddress = (*addrList)[0]
		result.Data = ""
	}

	return
}

func RegisterToken(name, password string, bccParams RegisterTokenParam) (result *CommitTxResult, err error) {

	defer FuncRecover(&err)

	contractName := "token-issue"
	_, keyStorePath, chainID := prepare("", bccParams.KeyStorePath, bccParams.ChainID)

	// require not empty
	requireNotEmpty("tokenName", bccParams.TokenName)
	requireNotEmpty("tokenSymbol", bccParams.TokenSymbol)
	requireNotEmpty("totalSupply", bccParams.TotalSupply)

	totalSupply, addSupplyEnabled, burnEnabled, gasPrice, err := getRegisterTokenData(bccParams)
	if err != nil {
		return
	}

	values := make([]interface{}, 0)
	values = append(values, bccParams.TokenName)
	values = append(values, bccParams.TokenSymbol)
	values = append(values, totalSupply)
	values = append(values, addSupplyEnabled)
	values = append(values, burnEnabled)
	values = append(values, gasPrice)

	//查询方法ID
	methodID, err := QueryMethodID("genesis", contractName, "NewToken", chainID, keyStorePath, false)
	if err != nil {
		return
	}

	result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID,
		false, false, methodID, values)
	if err != nil {
		return
	}

	var count = 0
	for result.Code != types.CodeOK && count < 2 {
		if result.Log == nonceErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID,
				true, false, methodID, values)
		} else if result.Log == smcErrDesc {
			result, err = packAndCommitTx(name, password, contractName, bccParams.GasLimit, bccParams.Note, keyStorePath, chainID,
				false, true, methodID, values)
		}

		count++
	}

	if result.Code == types.CodeOK {
		addrList := new([]string)
		err = json.Unmarshal([]byte(result.Data), addrList)
		if err != nil || len(*addrList) != 1 {
			return
		}
		result.TokenAddress = (*addrList)[0]
		result.Data = ""
	}

	return
}

func Transfer(name, password string, bccParams TransferParam) (result *CommitTxResult, err error) {
	defer FuncRecover(&err)

	_, keyStorePath, chainID := prepare("", bccParams.KeyStorePath, bccParams.ChainID)

	// require not empty
	requireNotEmpty("token", bccParams.Token)
	requireNotEmpty("value", bccParams.Value)
	requireNotEmpty("password", password)
	requireNotEmpty("gasLimit", bccParams.GasLimit)

	//err = algorithm.CheckAddress(chainID, bccParams.To)
	//if err != nil {
	//	return
	//}

	value, err := checkTransfer(bccParams.Value)
	if err != nil {
		return
	}
	var method uint32 = 1155058272
	result, err = transfer(name, password, bccParams.Token, bccParams.GasLimit, bccParams.Note, bccParams.To, value, keyStorePath, chainID, false, method)
	if err != nil {
		return
	}

	if result.Code != types.CodeOK {
		if result.Log == nonceErrDesc {
			result, err = transfer(name, password, bccParams.Token, bccParams.GasLimit, bccParams.Note, bccParams.To, value, keyStorePath, chainID, true, method)
		}
	}

	return
}

func transfer(name, password, token, gasLimit, note, to, value, keyStorePath, chainID string, bNonceErr bool, method uint32) (result *CommitTxResult, err error) {
	defer FuncRecover(&err)
	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return
	}

	tokenAddr, err := tokenAddressFromName(chainID, token)
	if err != nil {
		return
	}

	uGasLimit, err := strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		return
	}

	privStr, err := getAccountPriKey(keyStorePath, name, password)
	if err != nil {
		return
	}
	v := bn.NewNumberStringBase(value, 10)
	bccParamss := makeParams(to, v)
	txStr := GenerateTx(tokenAddr, method, bccParamss, nonce, int64(uGasLimit), note, privStr)

	result, err = CommitTx(chainID, txStr)
	if err != nil {
		return
	}
	return
}

func packAndCommitTx(name, password, contractName, gasLimit, note, keyStorePath, chainID string,
	bNonceErr, bSmcErr bool, methodID uint32, values []interface{}) (result *CommitTxResult, err error) {
	defer FuncRecover(&err)

	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return
	}

	contract, err := getContract(genesisOrgName, contractName, chainID, bSmcErr, keyStorePath)
	if err != nil {
		return
	}

	uGasLimit, err := strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		return
	}

	privStr, err := getAccountPriKey(keyStorePath, name, password)
	if err != nil {
		return
	}
	txStr := GenerateTx(contract.Address, methodID, values, nonce, int64(uGasLimit), note, privStr)

	result, err = CommitTx(chainID, txStr)

	return
}

//GenerateTx generate tx with one contract method request
func GenerateTx(contract types.Address, method uint32, bccParamss []interface{}, nonce uint64, gaslimit int64, note string, privKey string) string {
	items := tx2.WrapInvokeParams(bccParamss...)
	message := types.Message{
		Contract: contract,
		MethodID: method,
		Items:    items,
	}
	payload := tx2.WrapPayload(nonce, gaslimit, note, message)
	return tx2.WrapTx(payload, privKey)
}

func getRegisterTokenData(bccParams RegisterTokenParam) (totalSupply bn.Number, addSupplyEnabled, burnEnabled bool, gasPrice int, err error) {

	addSupplyEnabled, err = strconv.ParseBool(bccParams.AddSupplyEnabled)
	if err != nil {
		return
	}
	burnEnabled, err = strconv.ParseBool(bccParams.BurnEnabled)
	if err != nil {
		return
	}

	totalSupply = bn.NewNumberStringBase(bccParams.TotalSupply, 10)
	if totalSupply.IsLEI(0) {
		err = errors.New("invalid totalSupply")
		return
	}

	gasPrice, err = strconv.Atoi(bccParams.GasPrice)
	if err != nil {
		return
	}

	return
}

func getDeployContractData(bccParams DeployContractParam, chainID string) (
	effectHeightInt int,
	orgID string,
	codeHash []byte,
	codeData []byte,
	devSigStr string,
	orgSigStr string,
	err error) {

	// setup data
	codeData, err = ioutil.ReadFile(bccParams.CodeFile)
	if err != nil {
		return
	}

	devSigStr, err = getSigStr(bccParams.CodeFile + ".sig")
	if err != nil {
		return
	}

	orgSigStr, err = getSigStr(bccParams.CodeFile + ".sig.sig")
	if err != nil {
		return
	}

	codeHash = sha3.Sum256(codeData)
	blh, err := BlockHeight(chainID)
	if err != nil {
		return
	}
	HeightInt, err := strconv.Atoi(bccParams.EffectHeight)
	if err != nil {
		return
	}
	effectHeightInt = int(blh.LastBlock) + HeightInt

	bh := helper.BlockChainHelper{}
	orgID = bh.CalcOrgID(bccParams.OrgName)

	return
}

func makeParams(values ...interface{}) []interface{} {
	bccParamss := make([]interface{}, 0)
	for _, v := range values {
		bccParamss = append(bccParamss, v)
	}

	return bccParamss
}

func getSigStr(path string) (s string, err error) {

	type signature struct {
		PubKey    string `json:"pubkey"`
		Signature string `json:"signature"`
	}

	sigData, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	devSig := new(signature)
	err = json.Unmarshal(sigData, devSig)
	if err != nil {
		return
	}

	signaBtyes, err := json.Marshal(devSig)
	if err != nil {
		return
	}

	return string(signaBtyes), nil
}

func SolDeploy(name, password string, bvmParam BVMDeployParam) (result *CommitTxResult, err error) {
	defer FuncRecover(&err)

	_, keyStorePath, chainID := prepare("", bvmParam.KeyStorePath, bvmParam.ChainID)

	// require not empty
	requireNotEmpty("name", name)
	requireNotEmpty("password", password)
	if len(bvmParam.SourceFile) != 0 {
		requireNotEmpty("When compile solidity contract, contract name", bvmParam.ContractName)
	}

	if bvmParam.TokenAddr == "" && bvmParam.TokenName == "" {
		err = errors.New("tokenAddr and tokenName can not all be empty")
		return
	}

	addrS := nodeAddrSlice(chainID)

	if bvmParam.TokenName != "" {
		var tokenAddr types.Address

		key := std.KeyOfTokenWithName(bvmParam.TokenName)

		err = common.DoHttpQueryAndParse(addrS, key, &tokenAddr)
		if err != nil {
			return nil, errors.New("tokenName is right? error: " + err.Error())
		}

		bvmParam.TokenAddr = tokenAddr
	} else {
		// 添加代币地址验证
		key := std.KeyOfToken(bvmParam.TokenAddr)
		result := new(std.Token)
		err = common.DoHttpQueryAndParse(addrS, key, &result)
		if err != nil {
			return nil, errors.New("tokenAddr is right? error: " + err.Error())
		}
	}

	var binStr, abiStr string

	// read the sol contract binFile
	if bvmParam.BinFile != "" {
		codeData, err := ioutil.ReadFile(bvmParam.BinFile)
		if err != nil {
			return nil, err
		}

		binStr = string(codeData)
	}

	// Compile the sol contract
	if bvmParam.SourceFile != "" && bvmParam.BinFile == "" {
		abiData, binData, err := CompileSol(bvmParam)
		if err != nil {
			return nil, errors.New("sol contract compile failed, error: " + err.Error())
		}

		binStr = string(binData)
		abiStr = string(abiData)
	}

	newCodeData, err := hex.DecodeString(binStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Array := make([]interface{}, 0)
	for i := 0; i < len(bvmParam.ParamsArray); i++ {
		Array = append(Array, bvmParam.ParamsArray[i])
	}

	abiStrs, err := GetAbiObject(bvmParam.AbiFile, "")
	if err != nil {
		return nil, err
	}

	if abiStr == "" {
		abiStr = abiStrs
	}

	newAbiData := []byte(abiStr)

	if len(bvmParam.ParamsArray) > 0 {
		param, _, err := abi.EncodeFunctionCall(abiStr, "", Array...)
		if err != nil {
			return nil, err
		}

		newCodeData = append(newCodeData, param...)
	}

	result, err = packAndCommitTxForBVM(name, password, "", bvmParam.TokenAddr, bvmParam.GasLimit, bvmParam.Note, keyStorePath, chainID,
		false, true, 0, newCodeData, nil, newAbiData)

	var count = 0
	for result.Code != types.CodeOK && count < 2 {
		if result.Log == nonceErrDesc {
			result, err = packAndCommitTxForBVM(name, password, "", bvmParam.TokenAddr, bvmParam.GasLimit, bvmParam.Note, keyStorePath, chainID,
				true, true, 0, newCodeData, nil, newAbiData)
		}

		count++
	}

	return
}

func SolCall(name, password string, bvmParam BVMCallParam) (result *BVMCallResult, err error) {
	defer FuncRecover(&err)

	_, keyStorePath, chainID := prepare("", bvmParam.KeyStorePath, bvmParam.ChainID)

	// require not empty
	requireNotEmpty("name", name)
	requireNotEmpty("password", password)
	requireNotEmpty("method", bvmParam.Method)
	requireNotEmpty("contractAddr", bvmParam.ContractAddr)

	bvmParams := make([]byte, 0)

	Array := make([]interface{}, 0)
	for i := 0; i < len(bvmParam.ParamsArray); i++ {
		Array = append(Array, bvmParam.ParamsArray[i])
	}

	abistr, err := GetAbiObject(bvmParam.AbiFile, bvmParam.ContractAddr)
	if err != nil {
		return nil, err
	}

	param, _, err := abi.EncodeFunctionCall(abistr, bvmParam.Method, Array...)
	if err != nil {
		return nil, err
	}

	bvmParams = append(bvmParams, param...)

	transParams := make([]byte, 0)
	transMethodID := uint32(0)

	value, _ := strconv.Atoi(bvmParam.Value)
	if value > 0 {
		transMethodID = 0x44d8ca60
		transParams = []byte(bvmParam.Value)
	}

	contract, err := GetBvmContract(bvmParam.ContractAddr, bvmParam.ChainID)
	if err != nil {
		return nil, err
	}

	res, err := packAndCommitTxForBVM(name, password, bvmParam.ContractAddr, contract.Token, bvmParam.GasLimit, bvmParam.Note, keyStorePath, chainID,
		false, false, transMethodID, bvmParams, transParams, nil)

	var count = 0
	for (res.Code != types.CodeOK || res.Code != types.CodeBVMQueryOK) && count < 2 {
		if res.Log == nonceErrDesc {
			res, err = packAndCommitTxForBVM(name, password, bvmParam.ContractAddr, contract.Token, bvmParam.GasLimit, bvmParam.Note, keyStorePath, chainID,
				true, false, transMethodID, bvmParams, transParams, nil)
		} else if res.Log == smcErrDesc {
			res, err = packAndCommitTxForBVM(name, password, bvmParam.ContractAddr, contract.Token, bvmParam.GasLimit, bvmParam.Note, keyStorePath, chainID,
				false, false, transMethodID, bvmParams, transParams, nil)
		}

		count++
	}

	if res == nil {
		return nil, err
	}

	result = new(BVMCallResult)

	result.Data = []byte(res.Data)
	if res.Height != 0 || res.Fee != 0 {
		result.Height = res.Height
		result.Fee = res.Fee
	}
	result.Log = res.Log
	result.Code = res.Code
	result.TxHash = res.TxHash

	return
}

func packAndCommitTxForBVM(name, password, contractAddr, tokenAddr, gasLimit, note, keyStorePath, chainID string,
	bNonceErr, IsCreateCall bool, TransMethodID uint32, BVMParams, transParams, BVMAbi []byte) (result *CommitTxResult, err error) {

	defer FuncRecover(&err)

	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return
	}

	tx2.Init(chainID)
	Messages := PrepareMessages(contractAddr, tokenAddr, TransMethodID, transParams, BVMParams, BVMAbi, IsCreateCall)

	uGasLimit, err := strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		return
	}

	privStr, err := getAccountPriKey(keyStorePath, name, password)
	if err != nil {
		return
	}

	payLoad := tx2.WrapPayload(nonce, int64(uGasLimit), note, Messages...)

	txStr := tx2.WrapTx(payLoad, privStr)

	result, err = CommitTx(chainID, txStr)

	return
}

func GetBvmContract(addr, chainID string) (contract *std.BvmContract, err error) {

	addrS := nodeAddrSlice(chainID)

	contract = new(std.BvmContract)
	if err := common.DoHttpQueryAndParse(addrS, "/bvm/contract/"+addr, &contract); err != nil {
		return nil, err
	}

	return
}

// CompileSol - compile sol contract source file
func CompileSol(bvmParam BVMDeployParam) (abiData, binData []byte, err error) {

	systemStr := runtime.GOOS
	if systemStr == "windows" {
		cmd := exec.Command("cmd", "/C", "--bin "+bvmParam.SourceFile, "solc.exe")
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println("CompileBin failed, please check ! Error: ", err)
			return nil, nil, err
		}
		str := out.String()
		strs := strings.Split(str, "Binary:")
		str2 := strings.Replace(strings.Replace(strs[1], " ", "", -1), "\n", "", -1)
		return nil, []byte(str2), err

	} else {
		cmd := exec.Cmd{}
		if bvmParam.Library == "" {
			cmd = *exec.Command("/bin/sh", "-c", "./solc -o outputDirectory --overwrite --abi --bin "+bvmParam.SourceFile)
		} else {
			cmd = *exec.Command("/bin/sh", "-c", "./solc -o outputDirectory --overwrite --abi --bin "+bvmParam.SourceFile+" --libraries "+bvmParam.Library)
		}
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println("CompileBin failed, please check ! Error: ", err)
			return nil, nil, err
		}

		abiData, err := ioutil.ReadFile("./outputDirectory/" + bvmParam.ContractName + ".abi")
		if err != nil {
			return nil, nil, err
		}

		binData, err := ioutil.ReadFile("./outputDirectory/" + bvmParam.ContractName + ".bin")
		if err != nil {
			return nil, nil, err
		}

		return abiData, binData, nil
	}
}
