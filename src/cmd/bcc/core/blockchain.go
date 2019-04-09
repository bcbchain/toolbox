package core

import (
	"blockchain/algorithm"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/crypto/sha3"
	types2 "blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/helper"
	"blockchain/tx2"
	"blockchain/types"
	"common/jsoniter"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
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

	var methodID uint32 = 0x9e922f48
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

	var methodID uint32 = 0x62191292
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

	var methodID uint32 = 0xd7596e75
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

	effectHeight, orgID, codeHash, codeData, devSig, orgSig, err := getDeployContractData(bccParams)
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

	var methodID uint32 = 0xe0da7827
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

	var methodID uint32 = 0xed1d1d9a
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

	err = algorithm.CheckAddress(chainID, bccParams.To)
	if err != nil {
		return
	}

	var method uint32 = 1155058272
	result, err = transfer(name, password, bccParams.Token, bccParams.GasLimit, bccParams.Note, bccParams.To, bccParams.Value, keyStorePath, chainID, false, method)
	if err != nil {
		return
	}

	if result.Code != types.CodeOK {
		if result.Log == nonceErrDesc {
			result, err = transfer(name, password, bccParams.Token, bccParams.GasLimit, bccParams.Note, bccParams.To, bccParams.Value, keyStorePath, chainID, true, method)
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

	contract, err := contractOfTokenName(chainID, token)
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
	txStr := GenerateTx(contract.Address, method, bccParamss, nonce, int64(uGasLimit), note, privStr)

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

func getDeployContractData(bccParams DeployContractParam) (
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
	effectHeightInt, err = strconv.Atoi(bccParams.EffectHeight)
	if err != nil {
		return
	}

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
