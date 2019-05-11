package main

import (
	"blockchain/smcsdk/sdk/std"
	"blockchain/types"
	"bytes"
	"cmd/bcc/core"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var Version = make(map[string][]string)

// command line method
func call(name, password, orgName, contractName, methodName, file, params, splitBy, pay, gasLimit, note, chainID, keyStorePath string) error {

	bccParams := core.CallParam{OrgName: orgName, Contract: contractName, Method: methodName, ParamsFile: file, Params: params,
		SplitBy: splitBy, Pay: pay, GasLimit: gasLimit, Note: note, ChainID: chainID, KeyStorePath: keyStorePath}

	result, err := core.Call(name, password, bccParams)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func blockHeight(chainID string) error {

	blkHeight, err := core.BlockHeight(chainID)
	if err != nil {
		Error(fmt.Sprintf("Query Block Height failed, %v", err.Error()))
		return err
	}
	if blkHeight.LastBlock == 0 {
		fmt.Println(" BlockHeight query failed. Please check the input parameters")
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&blkHeight, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func block(chainID, height string) error {

	// if height is empty, then set it current height
	if height == "" {
		blkResult, err := core.BlockHeight(chainID)
		if err != nil {
			return err
		}

		height = fmt.Sprintf("%d", blkResult.LastBlock)
	}
	iHeight, err := strconv.ParseInt(height, 10, 64)
	if err != nil {
		return err
	}

	blk, err := core.Block(iHeight, chainID)
	if err != nil {
		Error(fmt.Sprintf("Query Block \"%v\" information failed, %v", height, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&blk, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func transaction(chainID, txHash string) error {

	tx, err := core.Transaction(chainID, txHash, nil)
	if err != nil {
		Error(fmt.Sprintf("Query transaction \"%v\" information failed, %v", txHash, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&tx, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func balance(accAddress types.Address, name, password, tokenName, allStr string, chainID, keyStorePath string) error {

	all, err := strconv.ParseBool(allStr)
	if err != nil {
		return err
	}

	result, err := core.Balance(accAddress, name, password, tokenName, all, chainID, keyStorePath)
	if err != nil {
		Error(fmt.Sprintf("Query balance \"%v\" information failed, %v", accAddress, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func nonce(accAddress types.Address, name, password, chainID, keyStorePath string) error {

	if accAddress == "" && name == "" {
		fmt.Println("Need name or accAddress, cannot all be empty")
		return nil
	}
	result, err := core.Nonce(accAddress, name, password, chainID, keyStorePath)
	if err != nil {
		Error(fmt.Sprintf("Query nonce \"%v\" information failed, %v", accAddress, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func commitTx(chainID, tx string) error {

	result, err := core.CommitTx(chainID, tx)
	if err != nil {
		Error(fmt.Sprintf("Commit transaction \"%v\" information failed, %v", tx, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func versionF() error {
	result, err := core.Version()
	if err != nil {
		Error(fmt.Sprintf("Query Version information failed, %v", err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func deployContract(name, password, contractName, version, orgName, codeFile,
	effectHeight, owner, keyStorePath, gasLimit, note, chainID string) error {

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
	if err != nil {
		Error(err.Error())
	}

	// 记录合约迭代版本
	//myVersion := new(core.VersionOfContract)
	//myVersion.Version[name] = append(myVersion.Version[name], version)
	Version[name] = append(Version[name], version)

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func registerToken(name, password, tokenName, tokenSymbol, totalSupply, gasPrice, gasLimit, note, keyStorePath, chainID,
	addSupplyEnabled, burnEnabled string) error {

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
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func registerOrg(name, password, orgName, gasLimit, note, keyStorePath, chainID string) error {

	param := core.RegisterOrgParam{
		OrgName:      orgName,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
		GasLimit:     gasLimit,
		Note:         note,
	}

	result, err := core.RegisterOrg(name, password, param)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func setOrgSigners(name, password, orgName, pubKeys, gasLimit, note, keyStorePath, chainID string) error {

	param := core.SetOrgSignersParam{
		OrgName:      orgName,
		PubKeys:      pubKeys,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
		GasLimit:     gasLimit,
		Note:         note,
	}

	result, err := core.SetOrgSigners(name, password, param)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func setOrgDeployer(name, password, orgName, deployer, gasLimit, note, keyStorePath, chainID string) error {

	param := core.SetOrgDeployerParam{
		OrgName:      orgName,
		Deployer:     deployer,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
		GasLimit:     gasLimit,
		Note:         note,
	}

	result, err := core.SetOrgDeployer(name, password, param)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func transfer(name, password, token, gasLimit, note, to, value, keyStorePath, chainID string) error {

	param := core.TransferParam{
		Token:        token,
		GasLimit:     gasLimit,
		Note:         note,
		To:           to,
		Value:        value,
		ChainID:      chainID,
		KeyStorePath: keyStorePath,
	}

	result, err := core.Transfer(name, password, param)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
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

	if orgID != "" && contractName != "" && contractAddr == "" {
		contractList, err := core.ContractInfo(chainID, orgID, contractName)
		if err != nil {
			Error(err.Error())
		}

		for _, v := range contractList {

			// 校验其他输入参数
			if orgName != "" {
				OrgInfo, err := core.QueryOrgInfo(orgID, chainID)
				if err != nil {
					Error(err.Error())
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
			Error(err.Error())
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
			Error(err.Error())
		}

		// 校验其他输入参数
		if orgName != "" && orgID != "" {
			OrgInfo, err := core.QueryOrgInfo(orgID, chainID)
			if err != nil {
				Error(err.Error())
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
			Error(err.Error())
		}
		fmt.Println("OK")
		fmt.Println("Response:")
		for _, v := range ContractAddrList {
			jsIndent, _ := json.MarshalIndent(&v, "", "  ")
			fmt.Printf("  %s\n", string(jsIndent))
		}
	} else {
		fmt.Println("Insufficient input parameters")
		return err
	}

	return
}

func ParamsExample(contract *std.Contract) (err error) {

	address, _ := json.MarshalIndent(&contract.Address, "", "  ")
	account, _ := json.MarshalIndent(&contract.Account, "", "  ")
	orgid, _ := json.MarshalIndent(&contract.OrgID, "", "  ")
	name, _ := json.MarshalIndent(&contract.Name, "", "  ")
	owner, _ := json.MarshalIndent(&contract.Owner, "", "  ")
	codeHash, _ := json.MarshalIndent(&contract.CodeHash, "", "  ")
	version, _ := json.MarshalIndent(&contract.Version, "", "  ")
	EffectHeight, _ := json.MarshalIndent(&contract.EffectHeight, "", "  ")
	loseEffect, _ := json.MarshalIndent(&contract.LoseHeight, "", "  ")
	keyPrefix, _ := json.MarshalIndent(&contract.KeyPrefix, "", "  ")
	interfaces, _ := json.MarshalIndent(&contract.Interfaces, "", "  ")
	token, _ := json.MarshalIndent(&contract.Token, "", "  ")

	fmt.Println("OK")
	fmt.Printf("Response: \n")
	fmt.Printf("  Version: %s\n", string(version))
	fmt.Printf("  Name: %s\n", string(name))
	fmt.Printf("  OrgID: %s\n", string(orgid))
	fmt.Printf("  Address: %s\n", string(address))
	fmt.Printf("  Account: %s\n", string(account))
	fmt.Printf("  Owner: %s\n", string(owner))
	fmt.Printf("  CodeHash: %s\n", string(codeHash))
	fmt.Printf("  EffectHeight: %s\n", string(EffectHeight))
	fmt.Printf("  LoseHeight: %s\n", string(loseEffect))
	fmt.Printf("  KeyPrefix: %s\n", string(keyPrefix))
	fmt.Printf("  Token: %s\n", string(token))
	fmt.Printf("  Interfaces: %s\n", string(interfaces))
	fmt.Printf("  Method: \n")

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

		fmt.Printf("    %s\n    Params： %s\n\n", v.ProtoType, example2)
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
