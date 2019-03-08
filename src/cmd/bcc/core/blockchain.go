package core

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/crypto/sha3"
	"blockchain/smcsdk/sdkimpl/helper"
	"blockchain/tx2"
	"blockchain/types"
	"cmd/bcc/common"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

const (
	genesisOrgName  = "genesis"
	smcContractName = "smartcontract"
	tokenIssue      = "token-issue"
	orgContractName = "organization"
)

func DeployContract(name, password string, param DeployContractParam) (result *CommitTxResult, err error) {

	var methodID uint32 = 3772413991
	result, err = deployContract(name, password, param.ContractName, param.Version, param.OrgName, param.CodeFile,
		param.EffectHeight, param.Owner, param.KeyStorePath, param.GasLimit, param.Note, param.ChainID, false, false, methodID)
	if err != nil {
		return
	}

	nonceErrDesc := "Invalid nonce"
	smcErrDesc := "The contract has expired"
	if result.Code != types.CodeOK {
		if result.Log == nonceErrDesc {
			result, err = deployContract(name, password, param.ContractName, param.Version, param.OrgName, param.CodeFile,
				param.EffectHeight, param.Owner, param.KeyStorePath, param.GasLimit, param.Note, param.ChainID, true, false, methodID)
		} else if result.Log == smcErrDesc {
			result, err = deployContract(name, password, param.ContractName, param.Version, param.OrgName, param.CodeFile,
				param.EffectHeight, param.Owner, param.KeyStorePath, param.GasLimit, param.Note, param.ChainID, false, true, methodID)
		}
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

func RegisterToken(name, password string, param RegisterTokenParam) (result *CommitTxResult, err error) {

	var methodID uint32 = 3978108314
	result, err = registerToken(name, password, param.TokenName, param.TokenSymbol, param.TotalSupply, param.GasPrice,
		param.GasLimit, param.Note, param.KeyStorePath, param.ChainID,
		param.AddSupplyEnabled, param.BurnEnabled, false, false, methodID)
	if err != nil {
		return
	}

	nonceErrDesc := "Invalid nonce"
	smcErrDesc := "The contract has expired"
	if result.Code != types.CodeOK {
		if result.Log == nonceErrDesc {
			result, err = registerToken(name, password, param.TokenName, param.TokenSymbol, param.TotalSupply, param.GasPrice,
				param.GasLimit, param.Note, param.KeyStorePath, param.ChainID,
				param.AddSupplyEnabled, param.BurnEnabled, true, false, methodID)
		} else if result.Log == smcErrDesc {
			result, err = registerToken(name, password, param.TokenName, param.TokenSymbol, param.TotalSupply, param.GasPrice,
				param.GasLimit, param.Note, param.KeyStorePath, param.ChainID,
				param.AddSupplyEnabled, param.BurnEnabled, false, true, methodID)
		}
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

func RegisterOrg(name, password string, param RegisterOrgParam) (result *CommitTxResult, err error) {

	var methodID uint32 = 2660380488
	result, err = registerOrg(name, password, param.OrgName, param.GasLimit, param.Note, param.KeyStorePath, param.ChainID, false, false, methodID)
	if err != nil {
		return
	}

	nonceErrDesc := "Invalid nonce"
	smcErrDesc := "The contract has expired"
	if result.Code != types.CodeOK {
		if result.Log == nonceErrDesc {
			result, err = registerOrg(name, password, param.OrgName, param.GasLimit, param.Note, param.KeyStorePath, param.ChainID, true, false, methodID)
		} else if result.Log == smcErrDesc {
			result, err = registerOrg(name, password, param.OrgName, param.GasLimit, param.Note, param.KeyStorePath, param.ChainID, false, true, methodID)
		}
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

func Transfer(name, password string, param TransferParam) (result *CommitTxResult, err error) {

	var method uint32 = 1155058272
	result, err = transfer(name, password, param.Token, param.GasLimit, param.Note, param.To, param.Value, param.KeyStorePath, param.ChainID, false, method)
	if err != nil {
		return
	}

	nonceErrDesc := "Invalid nonce"
	if result.Code != types.CodeOK {
		if result.Log == nonceErrDesc {
			result, err = transfer(name, password, param.Token, param.GasLimit, param.Note, param.To, param.Value, param.KeyStorePath, param.ChainID, true, method)
		}
	}

	return
}

func transfer(name, password, token, gasLimit, note, to, value, keyStorePath, chainID string, bNonceErr bool, method uint32) (result *CommitTxResult, err error) {

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
	}
	tx2.Init(chainID)

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
	params := makeParams(to, v)
	txStr := GenerateTx(contract.Address, method, params, nonce, int64(uGasLimit), note, privStr)

	result, err = CommitTx(chainID, txStr)
	if err != nil {
		return
	}
	return
}

func registerOrg(name, password, orgName, gasLimit, note, keyStorePath, chainID string,
	bNonceErr, bSmcErr bool, methodID uint32) (result *CommitTxResult, err error) {

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
	}
	tx2.Init(chainID)

	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return
	}

	contract, err := getContract(genesisOrgName, orgContractName, chainID, bSmcErr, keyStorePath)
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
	params := makeParams(orgName)
	txStr := GenerateTx(contract.Address, methodID, params, nonce, int64(uGasLimit), note, privStr)

	result, err = CommitTx(chainID, txStr)
	if err != nil {
		return
	}
	return
}

func registerToken(name, password, tokenName, tokenSymbol, totalSupply, gasPrice, gasLimit, note, keyStorePath, chainID,
	addSupplyEnabled, burnEnabled string, bNonceErr, bSmcErr bool, methodID uint32) (result *CommitTxResult, err error) {

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
	}
	tx2.Init(chainID)

	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return
	}

	contract, err := getContract(genesisOrgName, tokenIssue, chainID, bSmcErr, keyStorePath)
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

	a, err := strconv.ParseBool(addSupplyEnabled)
	if err != nil {
		return
	}
	b, err := strconv.ParseBool(burnEnabled)
	if err != nil {
		return
	}

	t := bn.NewNumberStringBase(totalSupply, 10)

	gp, err := strconv.Atoi(gasPrice)
	if err != nil {
		return
	}

	params := makeParams(tokenName, tokenSymbol, t, a, b, int64(gp))
	txStr := GenerateTx(contract.Address, methodID, params, nonce, int64(uGasLimit), note, privStr)

	result, err = CommitTx(chainID, txStr)
	if err != nil {
		return
	}
	return
}

//GenerateTx generate tx with one contract method request
func GenerateTx(contract types.Address, method uint32, params []interface{}, nonce uint64, gaslimit int64, note string, privKey string) string {
	items := tx2.WrapInvokeParams(params...)
	message := types.Message{
		Contract: contract,
		MethodID: method,
		Items:    items,
	}
	payload := tx2.WrapPayload(nonce, gaslimit, note, message)
	return tx2.WrapTx(payload, privKey)
}

func deployContract(name, password, contractName, version, orgName, codeFile,
	effectHeight, owner, keyStorePath, gasLimit, note, chainID string, bNonceErr, bSmcErr bool, methodID uint32) (result *CommitTxResult, err error) {

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
	}
	tx2.Init(chainID)

	nonce, err := getNonce(keyStorePath, chainID, name, password, bNonceErr)
	if err != nil {
		return
	}

	contract, err := getContract(genesisOrgName, smcContractName, chainID, bSmcErr, keyStorePath)
	if err != nil {
		return
	}

	uGasLimit, err := strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		return
	}

	codeData, err := ioutil.ReadFile(codeFile)
	if err != nil {
		return
	}

	devSigStr, err := getSigStr(codeFile + ".sig")
	if err != nil {
		return
	}

	orgSigStr, err := getSigStr(codeFile + ".sig.sig")
	if err != nil {
		return
	}

	codeHash := sha3.Sum256(codeData)
	effectHeightInt, err := strconv.Atoi(effectHeight)
	if err != nil {
		return
	}

	privStr, err := getAccountPriKey(keyStorePath, name, password)
	if err != nil {
		return
	}
	bh := helper.BlockChainHelper{}
	orgID := bh.CalcOrgID(orgName)
	params := makeParams(contractName, version, orgID, codeHash, codeData, devSigStr, orgSigStr, int64(effectHeightInt), owner)
	txStr := GenerateTx(contract.Address, methodID, params, nonce, int64(uGasLimit), note, privStr)

	result, err = CommitTx(chainID, txStr)
	if err != nil {
		return
	}
	return
}

func makeParams(values ...interface{}) []interface{} {
	params := make([]interface{}, 0)
	for _, v := range values {
		params = append(params, v)
	}

	return params
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
