package main

import (
	"blockchain/types"
	"bytes"
	"cmd/bcc/core"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

// command line method
func call(name, password, orgName, contractName, methodName, file, params, splitBy, pay, gasLimit, note, chainID, keyStorePath string) error {

	bccParams := core.CallParam{OrgName: orgName, Contract: contractName, Method: methodName, ParamsFile: file, Params: params,
		SplitBy: splitBy, Pay: pay, GasLimit: gasLimit, Note: note, ChainID: chainID, KeyStorePath: keyStorePath}

	result, err := core.Call(name, password, bccParams)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func blockHeight(chainID string) error {

	blkHeight, err := core.BlockHeight(chainID)
	if err != nil {
		Error(fmt.Sprintf("Query Block Height failed, %v", err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&blkHeight, "", "\t")
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
	jsIndent, _ := json.MarshalIndent(&blk, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func transaction(chainID, txHash string) error {
	if err := core.RequireNoEmpty("txHash", txHash); err != nil {
		return err
	}

	tx, err := core.Transaction(chainID, txHash, nil)
	if err != nil {
		Error(fmt.Sprintf("Query transaction \"%v\" information failed, %v", txHash, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&tx, "", "\t")
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
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func nonce(accAddress types.Address, name, password, chainID, keyStorePath string) error {

	result, err := core.Nonce(accAddress, name, password, chainID, keyStorePath)
	if err != nil {
		Error(fmt.Sprintf("Query nonce \"%v\" information failed, %v", accAddress, err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func commitTx(chainID, tx string) error {
	if err := core.RequireNoEmpty("transaction information", tx); err != nil {
		return err
	}

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

func versionF() error {
	result, err := core.Version()
	if err != nil {
		Error(fmt.Sprintf("Query Version information failed, %v", err.Error()))
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return nil
}

func deployContract(name, password, contractName, version, orgName, codeFile,
	effectHeight, owner, keyStorePath, gasLimit, note, chainID string) error {

	if contractName == "" || version == "" || orgName == "" || codeFile == "" ||
		effectHeight == "" || owner == "" || gasLimit == "" {
		Error("Invalid value.")
	}

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

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func registerToken(name, password, tokenName, tokenSymbol, totalSupply, gasPrice, gasLimit, note, keyStorePath, chainID,
	addSupplyEnabled, burnEnabled string) error {

	if tokenName == "" || tokenSymbol == "" || totalSupply == "" || gasPrice == "" || addSupplyEnabled == "" || burnEnabled == "" {
		Error("Invalid value.")
	}

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
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func registerOrg(name, password, orgName, gasLimit, note, keyStorePath, chainID string) error {

	if orgName == "" || gasLimit == "" {
		Error("Invalid value.")
	}

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
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return err
}

func transfer(name, password, token, gasLimit, note, to, value, keyStorePath, chainID string) error {

	if token == "" || gasLimit == "" {
		Error("Invalid value.")
	}

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
	jsIndent, _ := json.MarshalIndent(&result, "", "\t")
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
