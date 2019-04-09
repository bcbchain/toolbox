package common

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/tx2"
	"cmd/bcc/common"
	"common/wal"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"blockchain/smcsdk/sdk/std"
	"blockchain/types"
	"bytes"
	"cmd/bcc/core"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/tendermint/go-crypto"
)

const (
	nonceErrDesc = "Invalid nonce"
	smcErrDesc   = "The contract has expired"
)

var keyStorePath = ".keystore"

func nodeAddrSlice(chainID string) []string {
	if len(chainID) == 0 {
		chainID = common.GetBCCConfig().DefaultChainID
	}

	switch chainID {
	case "bcb":
		return common.GetBCCConfig().Bcb
	case "bcbtest":
		return common.GetBCCConfig().Bcbtest
	case "devtest":
		return common.GetBCCConfig().Devtest
	case "local":
		return common.GetBCCConfig().Local
	default:
		return []string{}
	}
}

func getAccountPriKey(keyStorePath, name, password string) (priKeyHex string, err error) {

	acct, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		return
	}

	priKey := acct.PrivateKey.(crypto.PrivKeyEd25519)

	return "0x" + hex.EncodeToString(priKey[:]), nil
}

func FuncRecover(errPtr *error) {
	if err := recover(); err != nil {
		msg := ""
		if errInfo, ok := err.(error); ok {
			msg = errInfo.Error()
		}

		if errInfo, ok := err.(string); ok {
			msg = errInfo
		}

		*errPtr = errors.New(msg)
	}
}

func prepare(splitBy, keyStorePath, chainID string) (string, string, string) {
	if splitBy == "" {
		splitBy = "@"
	}

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
	}
	crypto.SetChainId(chainID)
	tx2.Init(chainID)

	return splitBy, keyStorePath, chainID
}

func RequireNotEmpty(key, data string) {

	if len(data) == 0 {
		panic(errors.New(fmt.Sprintf("%s cannot be emtpy", key)))
	}
}

func requireUint64(key, valueStr string, base int) (uint64, error) {
	value, err := strconv.ParseUint(valueStr, base, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("%s error=%s", key, err.Error()))
	}

	return value, nil
}

func checkPay(pay string) (value bn.Number, token string, err error) {

	token = ""
	value = bn.N(0)
	if len(pay) > 0 {

		// step 1. check format
		firstIndex := strings.Index(pay, "(")
		lastIndex := strings.Index(pay, ")")
		if firstIndex == -1 || lastIndex < firstIndex {
			err = errors.New("pay option's format error, right format example: 1.02(bcb)")
			return
		}

		// step 2. check token
		token = pay[firstIndex+1 : lastIndex]
		if len(token) <= 0 {
			err = errors.New("pay option's format error, token cannot be empty")
			return
		}

		// step 3. check value
		valueStr := pay[:firstIndex]
		potIndex := strings.Index(valueStr, ".")
		if potIndex != -1 && len(strings.TrimRight(valueStr[potIndex+1:], "0")) > 9 {
			err = errors.New("pay option's format error, value's decimals cannot great than 9 chars")
			return
		}

		valueStr = strings.Replace(valueStr, ".", "", -1)
		zeroCount := 9 - (len(valueStr) - potIndex)
		for zeroCount > 0 {
			valueStr += "0"
			zeroCount--
		}
		value = bn.NewNumberStringBase(valueStr, 10)
		if value.IsLEI(0) {
			err = errors.New("pay option's format error, value must be number and greater than zero")
		}
	}

	return
}

func checkVersion(version string) (err error) {
	if len(version) < 3 {
		return errors.New("invalid version")
	}

	if len(strings.Trim(version, ".")) != len(version) {
		return errors.New("invalid version")
	}

	verStr := strings.Replace(version, ".", "", -1)
	verN := bn.NewNumberStringBase(verStr, 10)
	if verN.IsLessThanI(0) {
		return errors.New("invalid version")
	}

	return
}

// CheckUTF8 check format
func CheckUTF8(buf []byte) bool {
	nBytes := 0
	for i := 0; i < len(buf); i++ {
		if nBytes == 0 {
			if (buf[i] & 0x80) != 0 {
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1
					nBytes++
				}

				if nBytes < 2 || nBytes > 6 {
					return false
				}

				nBytes--
			}
		} else {
			if buf[i]&0xc0 != 0x80 {
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}

// --------------------------------------------------------------------------------

var Version = make(map[string][]string)

// command line method
func Call(name, password, orgName, contractName, methodName, file, params, splitBy, pay, gasLimit, note, chainID, keyStorePath string) (result *core.CommitTxResult, err error) {

	bccParams := core.CallParam{OrgName: orgName, Contract: contractName, Method: methodName, ParamsFile: file, Params: params,
		SplitBy: splitBy, Pay: pay, GasLimit: gasLimit, Note: note, ChainID: chainID, KeyStorePath: keyStorePath}

	result, err = core.Call(name, password, bccParams)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return
}

func BlockHeight(chainID string) (blkHeight *core.BlockHeightResult, err error) {

	blkHeight, err = core.BlockHeight(chainID)
	if err != nil {
		Error(fmt.Sprintf("Query Block Height failed, %v", err.Error()))
		return nil, err
	}

	return blkHeight, nil
}

func Block(chainID, height string) (blk string, err error) {
	defer FuncRecover(&err)
	if height == "" {
		blkResult, err := core.BlockHeight(chainID)
		if err != nil {
			return "", err
		}

		height = fmt.Sprintf("%d", blkResult.LastBlock)
	}
	iHeight, err := strconv.ParseInt(height, 10, 64)
	if err != nil {
		return
	}

	block, err := core.Block(iHeight, chainID)
	if err != nil {
		fmt.Println("Query Block \"%v\" information failed, %v", height, err.Error())
		return "", err
	}

	jsIndent, _ := json.MarshalIndent(&block, "", "\t")

	return string(jsIndent), err
}

func Transaction(chainID, txHash string) (txx string, err error) {

	tx, err := core.Transaction(chainID, txHash, nil)

	jsIndent, _ := json.MarshalIndent(&tx, "", "\t")

	return string(jsIndent), nil
}

func Balance(accAddress types.Address, name, password, tokenName, allStr string, chainID, keyStorePath string) (resultStr string, err error) {

	all, err := strconv.ParseBool(allStr)
	if err != nil {
		return
	}
	result, err := core.Balance(accAddress, name, password, tokenName, all, chainID, keyStorePath)

	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	return string(jsIndent), nil
}

func Nonce(accAddress types.Address, name, password, chainID, keyStorePath string) (nonceStr string, err error) {

	if accAddress == "" && name == "" {
		fmt.Println("Need name or accAddress, cannot all be empty")
		return
	}
	result, err := core.Nonce(accAddress, name, password, chainID, keyStorePath)
	if err != nil {
		Error(fmt.Sprintf("Query nonce \"%v\" information failed, %v", accAddress, err.Error()))
		return
	}

	jsIndent, _ := json.MarshalIndent(&result.Nonce, "", "\t")
	return string(jsIndent), nil
}

func commitTx(chainID, tx string) error {

	result, err := core.CommitTx(chainID, tx)
	if err != nil {
		Error(fmt.Sprintf("Commit transaction \"%v\" information failed, %v", tx, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func VersionF() (versionStr string, err error) {
	result, err := core.Version()
	if err != nil {
		Error(fmt.Sprintf("Query Version information failed, %v", err.Error()))
		return
	}

	jsIndent, _ := json.MarshalIndent(&result.Version, "", "\t")

	return string(jsIndent), nil
}

func DeployContract(name, password, contractName, version, orgName, codeFile,
	effectHeight, owner, keyStorePath, gasLimit, note, chainID string) (res string /*ult *core.CommitTxResult*/, err error) {
	defer FuncRecover(&err)

	param := core.DeployContractParam{
		ContractName: contractName,
		Version:      version,
		OrgName:      orgName,
		CodeFile:     codeFile,
		EffectHeight: effectHeight,
		Owner:        owner,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
		GasLimit:     gasLimit,
		Note:         note,
	}

	result, err := core.DeployContract(name, password, param)

	Version[name] = append(Version[name], version)

	jsIndent, _ := json.MarshalIndent(&result.Code, "", "\t")
	return string(jsIndent), nil
}

func RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, gasPrice, gasLimit, note, keyStorePath, chainID,
	addSupplyEnabled, burnEnabled string) (res string /*ult *core.CommitTxResult*/, err error) {
	defer FuncRecover(&err)
	param := core.RegisterTokenParam{
		TokenName:        tokenName,
		TokenSymbol:      tokenSymbol,
		TotalSupply:      totalSupply,
		AddSupplyEnabled: addSupplyEnabled,
		BurnEnabled:      burnEnabled,
		GasPrice:         gasPrice,
		ChainID:          chainID,
		KeyStorePath:     keyStorePath,
		GasLimit:         gasLimit,
		Note:             note,
	}

	result, err := core.RegisterToken(name, password, param)

	jsIndent, _ := json.MarshalIndent(&result.TokenAddress, "", "\t")
	return string(jsIndent), nil
}

func RegisterOrg(name, password, orgName, gasLimit, note, keyStorePath, chainID string) (resu string, err error) {
	defer FuncRecover(&err)
	param := core.RegisterOrgParam{
		OrgName:      orgName,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
		GasLimit:     gasLimit,
		Note:         note,
	}

	result, err := core.RegisterOrg(name, password, param)

	jsIndent, _ := json.MarshalIndent(&result.TxHash, "", "\t")
	return string(jsIndent), nil
}

func SetOrgSigners(name, password, orgName, pubKeys, gasLimit, note, keyStorePath, chainID string) (re string /*sult *core.CommitTxResult*/, err error) {
	defer FuncRecover(&err)
	param := core.SetOrgSignersParam{
		OrgName:      orgName,
		PubKeys:      pubKeys,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
		GasLimit:     gasLimit,
		Note:         note,
	}

	result, err := core.SetOrgSigners(name, password, param)

	jsIndent, _ := json.MarshalIndent(&result.Code, "", "\t")
	return string(jsIndent), err
}

func SetOrgDeployer(name, password, orgName, deployer, gasLimit, note, keyStorePath, chainID string) (res string /*ult *core.CommitTxResult*/, err error) {
	defer FuncRecover(&err)

	param := core.SetOrgDeployerParam{
		OrgName:      orgName,
		Deployer:     deployer,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
		GasLimit:     gasLimit,
		Note:         note,
	}

	result, err := core.SetOrgDeployer(name, password, param)
	jsIndent, _ := json.MarshalIndent(&result.Code, "", "\t")
	return string(jsIndent), nil
}

func Transfer(name, password, token, gasLimit, note, to, value, keyStorePath, chainID string) (result *core.CommitTxResult, err error) {

	param := core.TransferParam{
		Token:        token,
		GasLimit:     gasLimit,
		Note:         note,
		To:           to,
		Value:        value,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
	}

	result, err = core.Transfer(name, password, param)

	return result, err
}

func runAsRPCService() (err error) {
	cmd := exec.Command("/bin/bash", "-c", "./bccrpcservice")

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Start()
	if err != nil {
		return
	}

	fmt.Println("Execute finished")

	return nil
}

// Query the contract information based on the parameters
func ContractInfo(orgName, contractName, orgID, contractAddr string) (err error) {
	defer FuncRecover(&err)
	if orgID != "" && contractName != "" && contractAddr == "" {
		contractList, err := core.ContractInfo(chainID, orgID, contractName)
		if err != nil {
			fmt.Println("Query ContractInfo faild")
			return err
		}

		for _, v := range contractList {

			// 校验其他输入参数
			if orgName != "" {
				OrgInfo, err := core.QueryOrgInfo(orgID, chainID)
				if err != nil {
					fmt.Println("Query ContractInfo faild")
					return err
				}

				if orgName != OrgInfo.Name {
					fmt.Println("Error: Input orgName is wrong.")
					return err
				}
			}

			err = ParamsExample(&v)
			if err != nil {
				Error(err.Error())
			}
		}

	} else if orgName != "" && contractName != "" && contractAddr == "" {
		contract, err := core.QueryContractInfo(orgName, contractName, chainID, keyStorePath)
		if err != nil {
			fmt.Println("Query ContractInfo faild")
			return err
		}

		// 校验其他输入参数
		if orgID != "" && orgID != contract.OrgID {
			fmt.Println("Error: Input orgID is wrong.")
			return err
		}

		err = ParamsExample(contract)
		if err != nil {
			Error(err.Error())
		}

	} else if contractAddr != "" {
		contract, err := core.ContractInfoWithAddr(chainID, contractAddr)
		if err != nil {
			fmt.Println("Query ContractInfo faild")
			return err
		}

		// 校验其他输入参数
		if orgName != "" && orgID != "" {
			OrgInfo, err := core.QueryOrgInfo(orgID, chainID)
			if err != nil {
				fmt.Println("Query ContractInfo faild")
				return err
			}

			if orgName != OrgInfo.Name {
				fmt.Println("Error: Input orgName is wrong.")
				return err
			}
		}
		if orgID != "" && orgID != contract.OrgID {
			fmt.Println("Error: orgID orgName is wrong.")
			return err
		}

		err = ParamsExample(contract)
		if err != nil {
			Error(err.Error())
		}

	} else if orgName == "" && contractName == "" && orgID == "" && contractAddr == "" {
		ContractAddrList, err := core.AllContractInfo(chainID)
		if err != nil {
			fmt.Println("Query ContractInfo faild")
			return err
		}
		for _, v := range ContractAddrList {
			fmt.Println("OK")
			jsIndent, _ := json.MarshalIndent(&v, "", "\t")
			fmt.Printf("Response: %s\n", string(jsIndent))
		}

	} else {
		fmt.Println("Insufficient input parameters")
		return err
	}

	return
}

func ParamsExample(contract *std.Contract) (err error) {

	address, _ := json.MarshalIndent(&contract.Address, "", "\t")
	account, _ := json.MarshalIndent(&contract.Account, "", "\t")
	orgid, _ := json.MarshalIndent(&contract.OrgID, "", "\t")
	name, _ := json.MarshalIndent(&contract.Name, "", "\t")
	owner, _ := json.MarshalIndent(&contract.Owner, "", "\t")
	codeHash, _ := json.MarshalIndent(&contract.CodeHash, "", "\t")
	version, _ := json.MarshalIndent(&contract.Version, "", "\t")
	//version = strings.Join(versionList, ",")
	EffectHeight, _ := json.MarshalIndent(&contract.EffectHeight, "", "\t")
	loseEffect, _ := json.MarshalIndent(&contract.LoseHeight, "", "\t")
	keyPrefix, _ := json.MarshalIndent(&contract.KeyPrefix, "", "\t")
	interfaces, _ := json.MarshalIndent(&contract.Interfaces, "", "\t")
	token, _ := json.MarshalIndent(&contract.Token, "", "\t")

	fmt.Println("OK")
	fmt.Printf("Response: \n")
	fmt.Printf("    Version: %s\n", string(version))
	fmt.Printf("    Name: %s\n", string(name))
	fmt.Printf("    OrgID: %s\n", string(orgid))
	fmt.Printf("    Address: %s\n", string(address))
	fmt.Printf("    Account: %s\n", string(account))
	fmt.Printf("    Owner: %s\n", string(owner))
	fmt.Printf("    CodeHash: %s\n", string(codeHash))
	fmt.Printf("    EffectHeight: %s\n", string(EffectHeight))
	fmt.Printf("    LoseEffect: %s\n", string(loseEffect))
	fmt.Printf("    KeyPrefix: %s\n", string(keyPrefix))
	fmt.Printf("    Token: %s\n", string(token))
	fmt.Printf("    Interfaces: %s\n", string(interfaces))
	fmt.Printf("    Method: \n")

	var example2 = ""
	for _, v := range contract.Methods {

		leftBracketIndex := strings.Index(v.ProtoType, "(")
		rightBracketIndex := strings.Index(v.ProtoType, ")")
		splitTypes := strings.Split(v.ProtoType[leftBracketIndex+1:rightBracketIndex], ",")

		example := make([]string, 0)
		for _, v := range splitTypes {
			v = checkType(v)
			example = append(example, v)
			example2 = strings.Join(example, "@")
		}

		fmt.Printf("          %s\n          Params： %s\n\n", v.ProtoType, example2)
	}

	fmt.Println("PS: If the string is just a string, Example: \"example\"\n " +
		"If the string is a special string, Example: \"recvFeeRatio\":[500,500], \"recvFeeAddr\":[\"localKrHJUVGAt4R9gcfsBthu3dWJR7bAYq1c8\",\"localNwdwjpDotDDLGiB9pARk1CcSM71bdgTef\"]")

	return
}

func checkType(Type interface{}) string {

	switch Type {

	case "int", "int8", "int16", "int32", "int64":
		return "200000"

	case "uint", "uint8", "uint16", "uint32", "uint64":
		return "200000"

	case "float32", "float64":
		return "20.11"

	case "types.Address":
		return "localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j"

	case "bn.Number":
		return "1000000000000"

	case "bool":
		return "true"

	case "byte":
		return "0x01bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e"

	case "types.Hash", "types.HexBytes", "types.PubKey", "[]byte":
		return "0x01bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e"

	case "string":
		return "example"

	default:
		return ""
	}
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}

// testnonce
func Noncce(accAddress types.Address, name, password, chainID, keyStorePath string) (result *core.NonceResult, err error) {

	if accAddress == "" && name == "" {
		fmt.Println("Need name or accAddress, cannot all be empty")
		return
	}
	result, err = core.Nonce(accAddress, name, password, chainID, keyStorePath)
	if err != nil {
		fmt.Println("Query nonce information failed")
		return
	}

	return result, err
}
