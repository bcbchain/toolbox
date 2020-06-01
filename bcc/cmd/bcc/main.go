package main

import (
	"fmt"
	"github.com/bcbchain/bclib/types"
	"github.com/bcbchain/toolbox/bcc/cache"
	"github.com/bcbchain/toolbox/bcc/common"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	time2 "time"
)

func main() {
	err := common.LoadBCCConfig()
	if err != nil {
		panic(err)
	}

	cache.Init(".keystore")

	rand.Seed(time2.Now().Unix())

	err = Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "bcc",
	Short: "BlockChain client for access information and commit tx",
	Long: "BlockChain client that it can query data, call contract's method that you want, \n" +
		"and deploy contract and so on.",
}

var (
	keyStorePath string
	file         string

	// call flag
	name       string
	password   string
	orgName    string
	contract   string
	token      string
	method     string
	gasLimit   string
	splitBy    string
	params     string
	paramsFile string
	note       string
	pay        string

	// query flag
	height     string
	txHash     string
	accAddress types.Address
	tokenName  string
	tx         string
	chainID    string
	all        string

	// deploy contract flag
	contractName string
	version      string
	codeFile     string
	effectHeight string
	owner        string

	// register token
	tokenSymbol      string
	totalSupply      string
	addSupplyEnabled string
	burnEnabled      string
	gasPrice         string

	// transfer
	to    string
	value string

	pubKeys      string
	deployer     types.Address
	orgID        string
	contractAddr types.Address

	// Block params
	num  string
	time string

	// query params
	key string

	// BVM params
	tokenAddr  types.Address
	sourceFile string
	binFile    string
	abiFile    string
	paramsArr  []string
	library    string
)

func Execute() error {
	addFlags()
	addCommand()
	return RootCmd.Execute()
}

// call method
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Call contract method",
	Long:  "Call contract method with require params; require pay token if need transfer before run method.\nThen set pay option.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return call(name, password, orgName, contract, method, paramsFile, params, splitBy, pay, gasLimit, note, chainID, keyStorePath)
	},
}

// query block height
var blockHeightCmd = &cobra.Command{
	Use:   "blockHeight",
	Short: "Query BlockChain current block height",
	Long:  "Query BlockChain current block height",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return blockHeight(chainID)
	},
}

// query block information
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Query block information",
	Long:  "Query block information with height, must great than zero",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return block(height, time, num, chainID)
	},
}

// query transaction
var transactionCmd = &cobra.Command{
	Use:   "tx",
	Short: "Query transaction information",
	Long:  "Query transaction information with txHash and cannot be empty",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return transaction(chainID, txHash)
	},
}

// query balance with require params
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Query balance information",
	Long: `Query balance information with account address or wallet'a name and password.\n
if you want get all balance then set all flag.`,
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return balance(accAddress, name, password, tokenName, all, chainID, keyStorePath)
	},
}

// query nonce
var nonceCmd = &cobra.Command{
	Use:   "nonce",
	Short: "Query account nonce",
	Long:  "Query account nonce with it's address",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nonce(accAddress, name, password, chainID, keyStorePath)
	},
}

// commit transaction
var commitTxCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit transaction",
	Long:  "Commit transaction with tx's data",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return commitTx(tx, file, chainID)
	},
}

// query version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Query version information",
	Long:  "Query version information",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return versionF()
	},
}

// deploy contract
var deployContractCmd = &cobra.Command{
	Use:   "deployContract",
	Short: "Deploy smart contract",
	Long:  "Deploy smart contract with information",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return deployContract(name, password, contractName, version, orgName, codeFile, effectHeight, owner, keyStorePath, gasLimit, note, chainID)
	},
}

// register token
var registerTokenCmd = &cobra.Command{
	Use:   "registerToken",
	Short: "Register token",
	Long:  "Register token with information",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return registerToken(name, password, tokenName, tokenSymbol, totalSupply, gasPrice, gasLimit, note, keyStorePath, chainID, addSupplyEnabled, burnEnabled)
	},
}

// register organization
var registerOrgCmd = &cobra.Command{
	Use:   "registerOrg",
	Short: "Register organization",
	Long:  "Register organization by name",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return registerOrg(name, password, orgName, gasLimit, note, keyStorePath, chainID)
	},
}

// setOrgSigners set organization's signers
var setOrgSignersCmd = &cobra.Command{
	Use:   "setOrgSigners",
	Short: "Set organization's signers",
	Long:  "Set organization's signers",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return setOrgSigners(name, password, orgName, pubKeys, gasLimit, note, keyStorePath, chainID)
	},
}

// setOrgDeployer deploy smart contract account
var setOrgDeployerCmd = &cobra.Command{
	Use:   "setOrgDeployer",
	Short: "Authorize deploy smart contract account",
	Long:  "Authorize deploy smart contract account",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return setOrgDeployer(name, password, orgName, deployer, gasLimit, note, keyStorePath, chainID)
	},
}

// transfer
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer token to someone and unit is a token",
	Long:  "Transfer token to someone and unit is a token",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return transfer(name, password, tokenName, gasLimit, note, to, value, keyStorePath, chainID)
	},
}

// run as rpc service
var runAsRPCServiceCmd = &cobra.Command{
	Use:   "runAsRPCService",
	Short: "Run bcc to rpc service",
	Long:  "Run bcc to rpc service",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAsRPCService()
	},
}

// query contract info
var contractInfoCmd = &cobra.Command{
	Use:   "contractInfo",
	Short: "Query contract information",
	Long:  "Query contract information",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return ContractInfo(orgName, contractName, orgID, contractAddr)
	},
}

// query contract info
var tokenInfoCmd = &cobra.Command{
	Use:   "tokenInfo",
	Short: "Query token information",
	Long:  "Query token information",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return TokenInfo(tokenName, chainID)
	},
}

// query method
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the chain information through the RPC interface",
	Long:  "Query the chain information through the RPC interface",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return query(key, chainID)
	},
}

// Solidity contract deploy
var solDeployCmd = &cobra.Command{
	Use:   "solDeploy",
	Short: "Deploy solidity smart contract",
	Long:  "Deploy solidity smart contract with information",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return SolDeploy(name, password, tokenAddr, tokenName, sourceFile, library, contractName, binFile, abiFile, gasLimit, note, chainID, keyStorePath, paramsArr)
	},
}

// Solidity contract call
var solCallCmd = &cobra.Command{
	Use:   "solCall",
	Short: "Call solidity contract method",
	Long:  "Call solidity contract method with require params",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return SolCall(abiFile, name, password, contractAddr, value, gasLimit, note, chainID, keyStorePath, method, paramsArr)
	},
}

func callCmdFlags() {
	callCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	callCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	callCmd.PersistentFlags().StringVarP(&orgName, "orgName", "o", "", "name of organization")
	callCmd.PersistentFlags().StringVarP(&contract, "contract", "t", "", "name of contract that you want invoke")
	callCmd.PersistentFlags().StringVarP(&method, "method", "m", "", "method that it invoked with contract")
	callCmd.PersistentFlags().StringVarP(&params, "params", "a", "", "parameters with method")
	callCmd.PersistentFlags().StringVarP(&paramsFile, "paramsFile", "f", "", "parameters in file if it too long")
	callCmd.PersistentFlags().StringVarP(&splitBy, "splitBy", "s", "@", "char that split params")
	callCmd.PersistentFlags().StringVarP(&pay, "pay", "y", "", "pay token before invoke method ,unit is a token")
	callCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	callCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note for tx")
	callCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	callCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "the path with load account")
}
func blockHeightCmdFlags() {
	blockHeightCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
}
func blockCmdFlags() {
	blockCmd.PersistentFlags().StringVarP(&height, "height", "t", "", "height of blockchain")
	blockCmd.PersistentFlags().StringVarP(&num, "num", "n", "", "number of block")
	blockCmd.PersistentFlags().StringVarP(&time, "rt", "m", "", "time of blockTime, example : \"2006-01-02 15:04:05\"")
	blockCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
}
func transactionCmdFlags() {
	transactionCmd.PersistentFlags().StringVarP(&txHash, "txhash", "t", "", "hash of transaction")
	transactionCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
}
func balanceCmdFlags() {
	balanceCmd.PersistentFlags().StringVarP(&accAddress, "accAddress", "a", "", "address of account")
	balanceCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	balanceCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	balanceCmd.PersistentFlags().StringVarP(&tokenName, "tokenName", "t", "", "name of token")
	balanceCmd.PersistentFlags().StringVarP(&all, "all", "l", "false", "query all balance of token")
	balanceCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	balanceCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path with load account")
}
func nonceCmdFlags() {
	nonceCmd.PersistentFlags().StringVarP(&accAddress, "accAddress", "a", "", "address of account")
	nonceCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	nonceCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	nonceCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	nonceCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path with load account")
}
func commitTxCmdFlags() {
	commitTxCmd.PersistentFlags().StringVarP(&tx, "tx", "t", "", "data of transaction")
	commitTxCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	commitTxCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "transaction data file")
}
func versionCmdFlags() {
}
func deployContractFlags() {
	deployContractCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	deployContractCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	deployContractCmd.PersistentFlags().StringVarP(&contractName, "contractName", "t", "", "name of contract for deploy")
	deployContractCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "version for contract")
	deployContractCmd.PersistentFlags().StringVarP(&orgName, "orgName", "o", "", "tge name for organization")
	deployContractCmd.PersistentFlags().StringVarP(&codeFile, "codeFile", "f", "", "path of contract code file")
	deployContractCmd.PersistentFlags().StringVarP(&effectHeight, "effectHeight", "i", "", "this version contract effective after height")
	deployContractCmd.PersistentFlags().StringVarP(&owner, "owner", "r", "", "address of owner for contract")
	deployContractCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	deployContractCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note of transaction")
	deployContractCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	deployContractCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}
func registerTokenFlags() {
	registerTokenCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	registerTokenCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	registerTokenCmd.PersistentFlags().StringVarP(&tokenName, "tokenName", "t", "", "name of token")
	registerTokenCmd.PersistentFlags().StringVarP(&tokenSymbol, "tokenSymbol", "l", "", "symbol of token")
	registerTokenCmd.PersistentFlags().StringVarP(&totalSupply, "totalSupply", "y", "", "total supply of token")
	registerTokenCmd.PersistentFlags().StringVarP(&addSupplyEnabled, "addSupplyEnabled", "a", "", "add supply enabled of token")
	registerTokenCmd.PersistentFlags().StringVarP(&burnEnabled, "burnEnabled", "b", "", "burn enabled of token")
	registerTokenCmd.PersistentFlags().StringVarP(&gasPrice, "gasPrice", "i", "", "gas price of token")
	registerTokenCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	registerTokenCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note of transaction")
	registerTokenCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	registerTokenCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}
func registerOrgFlags() {
	registerOrgCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	registerOrgCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	registerOrgCmd.PersistentFlags().StringVarP(&orgName, "orgName", "o", "", "organization name")
	registerOrgCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	registerOrgCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note of transaction")
	registerOrgCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	registerOrgCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}
func setOrgSignersFlags() {
	setOrgSignersCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	setOrgSignersCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	setOrgSignersCmd.PersistentFlags().StringVarP(&orgName, "orgName", "o", "", "organization name")
	setOrgSignersCmd.PersistentFlags().StringVarP(&pubKeys, "pubKeys", "s", "", "signer's pubKey")
	setOrgSignersCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	setOrgSignersCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note of transaction")
	setOrgSignersCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	setOrgSignersCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}
func setOrgDeployerFlags() {
	setOrgDeployerCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	setOrgDeployerCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	setOrgDeployerCmd.PersistentFlags().StringVarP(&orgName, "orgName", "o", "", "organization name")
	setOrgDeployerCmd.PersistentFlags().StringVarP(&deployer, "deployer", "d", "", "deployer's address")
	setOrgDeployerCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	setOrgDeployerCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note of transaction")
	setOrgDeployerCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	setOrgDeployerCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}
func transferFlags() {
	transferCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	transferCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	transferCmd.PersistentFlags().StringVarP(&tokenName, "token", "t", "", "token name of transfer")
	transferCmd.PersistentFlags().StringVarP(&to, "to", "o", "", "address of transfer")
	transferCmd.PersistentFlags().StringVarP(&value, "value", "v", "", "value of transfer")
	transferCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	transferCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note of transaction")
	transferCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	transferCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}
func runAsRpcServiceFlags() {
}
func contractInfoFlags() {
	contractInfoCmd.PersistentFlags().StringVarP(&orgName, "orgName", "o", "", "organization name")
	contractInfoCmd.PersistentFlags().StringVarP(&contractName, "contractName", "t", "", "name of contract for deploy")
	contractInfoCmd.PersistentFlags().StringVarP(&orgID, "orgID", "i", "", "organization ID")
	contractInfoCmd.PersistentFlags().StringVarP(&contractAddr, "contractAddr", "a", "", "address of contract")
	contractInfoCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
}
func tokenInfoFlags() {
	tokenInfoCmd.PersistentFlags().StringVarP(&tokenName, "token", "t", "", "token name")
	tokenInfoCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
}

func queryFlags() {
	queryCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "rpc query param")
	queryCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
}
func solDeployFlags() {
	solDeployCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	solDeployCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	solDeployCmd.PersistentFlags().StringVarP(&tokenAddr, "tokenAddr", "a", "", "address of token")
	solDeployCmd.PersistentFlags().StringVarP(&tokenName, "tokenName", "t", "", "name of token")
	solDeployCmd.PersistentFlags().StringVarP(&sourceFile, "sourceFile", "f", "", "path of solidity contract source file")
	solDeployCmd.PersistentFlags().StringVarP(&library, "library", "l", "", "name and address of library, example : \"name:address\"")
	solDeployCmd.PersistentFlags().StringVarP(&contractName, "contractName", "o", "", "name of contract for sourceFile deploy")
	solDeployCmd.PersistentFlags().StringVarP(&binFile, "binFile", "b", "", "path of solidity contract binaryCode file")
	solDeployCmd.PersistentFlags().StringVarP(&abiFile, "abiFile", "i", "", "path of solidity contract abi file")
	solDeployCmd.PersistentFlags().StringArrayVarP(&paramsArr, "paramsArr", "r", []string{}, "parameters array with construction method")
	solDeployCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	solDeployCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note for tx")
	solDeployCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	solDeployCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}
func solCallFlags() {
	solCallCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of wallet")
	solCallCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of wallet")
	solCallCmd.PersistentFlags().StringVarP(&contractAddr, "contractAddr", "a", "", "address of contract")
	solCallCmd.PersistentFlags().StringVarP(&value, "value", "v", "", "value of transfer")
	solCallCmd.PersistentFlags().StringVarP(&method, "method", "m", "", "method that it invoked with contract")
	solCallCmd.PersistentFlags().StringVarP(&abiFile, "abiFile", "i", "", "path of solidity contract abi file")
	solCallCmd.PersistentFlags().StringArrayVarP(&paramsArr, "paramsArr", "r", []string{}, "parameters array with call method")
	solCallCmd.PersistentFlags().StringVarP(&gasLimit, "gasLimit", "g", "", "gas limit for now transaction")
	solCallCmd.PersistentFlags().StringVarP(&note, "note", "e", "", "note for tx")
	solCallCmd.PersistentFlags().StringVarP(&chainID, "chainid", "c", "", "chainid define blockchain for this invoke")
	solCallCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of keystore")
}

func addFlags() {
	callCmdFlags()
	blockHeightCmdFlags()
	blockCmdFlags()
	transactionCmdFlags()
	balanceCmdFlags()
	nonceCmdFlags()
	commitTxCmdFlags()
	versionCmdFlags()
	registerOrgFlags()
	setOrgSignersFlags()
	setOrgDeployerFlags()
	deployContractFlags()
	registerTokenFlags()
	transferFlags()
	runAsRpcServiceFlags()
	contractInfoFlags()
	tokenInfoFlags()
	queryFlags()
	solDeployFlags()
	solCallFlags()
}

func addCommand() {
	RootCmd.AddCommand(callCmd)
	RootCmd.AddCommand(blockHeightCmd)
	RootCmd.AddCommand(blockCmd)
	RootCmd.AddCommand(transactionCmd)
	RootCmd.AddCommand(balanceCmd)
	RootCmd.AddCommand(nonceCmd)
	RootCmd.AddCommand(commitTxCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(deployContractCmd)
	RootCmd.AddCommand(registerOrgCmd)
	RootCmd.AddCommand(setOrgSignersCmd)
	RootCmd.AddCommand(setOrgDeployerCmd)
	RootCmd.AddCommand(registerTokenCmd)
	RootCmd.AddCommand(transferCmd)
	RootCmd.AddCommand(runAsRPCServiceCmd)
	RootCmd.AddCommand(contractInfoCmd)
	RootCmd.AddCommand(tokenInfoCmd)
	RootCmd.AddCommand(queryCmd)
	RootCmd.AddCommand(solDeployCmd)
	RootCmd.AddCommand(solCallCmd)
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
