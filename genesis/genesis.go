package main

import (
	"github.com/bcbchain/bclib/algorithm"
	"github.com/bcbchain/bclib/fs"
	"github.com/bcbchain/bclib/wal"

	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/bcbchain/bclib/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}

func FileNameOfNodes(chainID string) string {
	return chainID + string("-nodes") + ".json"
}
func FileNameOfAccountOwner() string {
	return "wal/owner/owner.json"
}
func FileNameOfAccountReward(chainID, name string) string {
	return chainID + string("-account-") + name + ".json"
}
func FileNameOfNodeAccount(chainID string, index int) string {
	return fmt.Sprintf("%v-node%v.json", chainID, index)
}
func FileNameOfCharter(chainID string) string {
	return chainID + string("-charter") + ".json"
}
func FileNameOfGenesis(chainID string) string {
	return chainID + string("-genesis") + ".json"
}
func FileNameOfValidators(chainID string) string {
	return chainID + string("-validators") + ".json"
}

func checkFileExist(fn string) (bool, error) {
	exist, err := fs.PathExists(fn)
	if err != nil {
		return false, err
	} else if !exist {
		return false, errors.New(fmt.Sprintf("File \"%v\" is not exist", fn))
	}
	return true, nil
}

func CheckFileExist(fn string) {
	_, err := checkFileExist(fn)
	if err != nil {
		Error(err.Error())
	}
}

func CheckInputFileExist(chainID, path string) {
	CheckFileExist(filepath.Join(path, FileNameOfCharter(chainID)))
	CheckFileExist(filepath.Join(path, FileNameOfNodes(chainID)))
	CheckFileExist(filepath.Join(path, FileNameOfAccountOwner()))
}

func LoadCharterDefFile(chainID, path string) *GenesisCharter {
	jsonBytes, err := ioutil.ReadFile(filepath.Join(path, FileNameOfCharter(chainID)))
	if err != nil {
		Error(fmt.Sprintf("Read file \"%v\" failed, %v", FileNameOfCharter(chainID), err.Error()))
	}

	genesisCharter := new(GenesisCharter)
	err = cdc.UnmarshalJSON(jsonBytes, genesisCharter)
	if err != nil {
		Error(fmt.Sprintf("UnmarshalJSON from file \"%v\" failed, %v", FileNameOfCharter(chainID), err.Error()))
	}
	return genesisCharter
}

func LoadOwnerAccountFile(chainID, path string) *Account {
	accounts := findWalAccount(filepath.Join(path, "wal/owner"))
	if len(accounts) != 1 {
		Error(fmt.Sprintf("Owner account number must be one"))
	}
	return &accounts[0]
}

func LoadRewardAccountFiles(path string) []Account {
	return findWalAccount(filepath.Join(path, "wal/rewards"))
}

func LoadNodeAccountFiles(path string) []Account {
	return findWalAccount(filepath.Join(path, "wal/validators"))
}

func LoadNodesDefFile(chainID, path string) []NodeDef {
	jsonBytes, err := ioutil.ReadFile(filepath.Join(path, FileNameOfNodes(chainID)))
	if err != nil {
		Error(fmt.Sprintf("Read file \"%v\" failed, %v", FileNameOfNodes(chainID), err.Error()))
	}

	nodes := make([]NodeDef, 0)
	err = cdc.UnmarshalJSON(jsonBytes, &nodes)
	if err != nil {
		Error(fmt.Sprintf("UnmarshalJSON from file \"%v\" failed, %v", FileNameOfNodes(chainID), err.Error()))
	}
	return nodes
}

func SaveNodesDefFile(chainId, pathOfOutput string, nodes []NodeDef) {
	var jsonBytes []byte
	var err error

	if jsonBytes, err = cdc.MarshalJSONIndent(nodes, "", "  "); err != nil {
		Error(fmt.Sprintf("MarshalJSON to file \"%v\" failed, %v", FileNameOfNodes(chainId), err.Error()))
	}

	fn := filepath.Join(pathOfOutput, FileNameOfNodes(chainId))
	if err := ioutil.WriteFile(fn, jsonBytes, 0775); err != nil {
		Error(fmt.Sprintf("Write file \"%v\" failed, %v", fn, err.Error()))
	}
}

func SaveContractFiles(pathOfOutput string, contractInfos []contractInfo, releaseVersion string) {
	for _, contract := range contractInfos {
		fn := filepath.Join(pathOfOutput, contract.Code)
		if err := ioutil.WriteFile(fn, contract.code, 0775); err != nil {
			Error(fmt.Sprintf("Write file \"%v\" failed, %v", fn, err.Error()))
		}
	}

	fn := filepath.Join(pathOfOutput, "genesis-smart-contract-release-version.txt")
	if err := ioutil.WriteFile(fn, []byte(releaseVersion), 0775); err != nil {
		Error(fmt.Sprintf("Write file \"%v\" failed, %v", fn, err.Error()))
	}
}

func loadContractCodeFile(fileName string) string {

	datas, err := ioutil.ReadFile(fileName)
	if err != nil {
		Error(fmt.Sprintf("Read file \"%v\" failed, %v", fileName, err.Error()))
		return ""
	}
	return hex.EncodeToString(algorithm.SHA3256(datas))
}

func loadContractCodeSignFile(fileName string) codeSign {
	datas, err := ioutil.ReadFile(fileName)
	sign := codeSign{}
	if err != nil {
		Error(fmt.Sprintf("Read file \"%v\" failed, %v", fileName, err.Error()))
		return sign
	}

	err = json.Unmarshal(datas, &sign)
	if err != nil {
		Error(fmt.Sprintf("UnmarshalJSON from file \"%v\" failed, %v", fileName, err.Error()))
	}

	return sign
}

//创世文件自定义内容
func genAppStateOptions(genesisCharter *GenesisCharter, ownerAddr crypto.Address, genesisContracts []contractInfo, org string) GenesisAppState {
	genesisToken := genesisCharter.Token
	genesisRewards := genesisCharter.Rewards

	//基础代币
	tokenBasic := tokenInfo{
		Address:          "",
		Owner:            ownerAddr,
		Name:             genesisToken.TokenName,
		Symbol:           genesisToken.TokenSymbol,
		TotalSupply:      genesisToken.TotalSupply,
		AddSupplyEnabled: genesisToken.AddSupplyEnabled,
		BurnEnabled:      genesisToken.BurnEnabled,
		GasPrice:         genesisToken.GasPrice,
	}

	return GenesisAppState{Organization: org, Token: tokenBasic, RewardStrategy: genesisRewards, Contracts: genesisContracts}
}

func genValidators(nodes []NodeDef) []GenesisValidator {
	GenesisValidators := make([]GenesisValidator, 0)
	for _, node := range nodes {
		validator := GenesisValidator{node.Name, node.RewardAddr, int64(node.Power)}
		GenesisValidators = append(GenesisValidators, validator)
	}

	return GenesisValidators
}

func genesis(chainID, password, pathOfCharter, pathOfContracts, pathOfOutput, org string) error {

	if chainID == "" {
		return errors.New("Need chain id")
	}
	if pathOfCharter == "" {
		return errors.New("Need root path of charter files ")
	}
	if pathOfContracts == "" {
		return errors.New("Need path of contract files ")
	}
	if pathOfOutput == "" {
		return errors.New("Need path of output genesis files ")
	}
	if _, err := fs.MakeDir(pathOfOutput); err != nil {
		Error(err.Error())
	}

	CheckInputFileExist(chainID, pathOfCharter)
	ownerAcct := LoadOwnerAccountFile(chainID, pathOfCharter)
	genesisCharter := LoadCharterDefFile(chainID, pathOfCharter)
	nodes := LoadNodesDefFile(chainID, pathOfCharter)

	nodeAccounts := LoadNodeAccountFiles(pathOfCharter)
	matchNodeCounts := 0
	for i, node := range nodes {
		for _, acct := range nodeAccounts {
			if node.Name == acct.name {
				nodes[i].RewardAddr = acct.Addr
				matchNodeCounts++
			}
		}
	}
	if matchNodeCounts != len(nodes) || matchNodeCounts != len(nodeAccounts) {
		Error("Node definition file is not match with validators account")
	}

	rewardAccounts := LoadRewardAccountFiles(pathOfCharter)
	matchRewardCounts := 0
	for i, reward := range genesisCharter.Rewards {
		if reward.Name != "validators" {
			for _, acct := range rewardAccounts {
				if reward.Name == acct.name {
					genesisCharter.Rewards[i].Address = acct.Addr
					matchRewardCounts++
				}
			}
		}
	}
	if matchRewardCounts != len(genesisCharter.Rewards)-1 || matchRewardCounts != len(rewardAccounts) {
		Error("Charter definition file is not match with rewards account")
	}

	genesisContracts, _ := findContract(pathOfContracts)
	appState := genAppStateOptions(genesisCharter, ownerAcct.Addr, genesisContracts, org)
	validators := genValidators(nodes)

	//构造创世对象
	doc := GenesisDoc{
		ChainID:      chainID,
		ChainVersion: "2",
		GenesisTime:  time.Now(),
		AppHash:      "",
		AppState:     appState,
		Validators:   validators,
	}

	//输出创世文件
	out, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		Error(fmt.Sprintf("Marshal genesis failed, %v", err.Error()))
	}
	if err = ioutil.WriteFile(filepath.Join(pathOfOutput, FileNameOfGenesis(chainID)), out, 0600); err != nil {
		Error(fmt.Sprintf("Write genesis file failed, %v", err.Error()))
	}

	//输出验证者信息文件
	validatorsBytes, err := json.MarshalIndent(validators, "", "  ")
	if err != nil {
		Error(fmt.Sprintf("Marshal validators failed, %v", err.Error()))
	}
	if err = ioutil.WriteFile(filepath.Join(pathOfOutput, FileNameOfValidators(chainID)), validatorsBytes, 0600); err != nil {
		Error(fmt.Sprintf("Write validators file failed, %v", err.Error()))
	}

	//输出节点信息文件
	SaveNodesDefFile(chainID, pathOfOutput, nodes)

	//输出合约打包文件
	//SaveContractFiles(pathOfOutput, genesisContracts, releaseVersion)

	//如果存在owner钱包并指定了密码，直接对创世文件进行签名（采用文本文件签名方式）
	if password != "" {
		if ok, _ := checkFileExist(filepath.Join(pathOfCharter, string("wal/owner/")+ownerAcct.name+".wal")); ok {
			ownerWal, err := wal.LoadAccount(filepath.Join(pathOfCharter, string("wal/owner")), ownerAcct.name, password)
			if err != nil {
				Error(fmt.Sprintf("Load owner's wallet failed, %v", err.Error()))
			}
			ownerWal.SignTextFile(
				filepath.Join(pathOfOutput, FileNameOfGenesis(chainID)),
				filepath.Join(pathOfOutput, FileNameOfGenesis(chainID))+".sig",
			)
		}
	}

	fmt.Println("OK")
	return nil
}
