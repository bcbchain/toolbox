package gencmd

import (
	"blockchain/smccheck/parsecode"
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

var templateText = `package main

import (
	"common/socket"
	"fmt"
	"os"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdkimpl"
	"blockchain/smcsdk/sdkimpl/helper"
	"blockchain/smcsdk/sdkimpl/llstate"
	"blockchain/smcsdk/sdkimpl/object"
	"contract/{{.OrgID}}/stub"

	"blockchain/smcsdk/sdk/jsoniter"
	sdkType "blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	"blockchain/types"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/abci/types"
	tmcommon "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"
)

var (
	logger          log.Loggerf
	flagRPCPort     int
	flagCallbackURL string
	//rpc             *rpcclient.JSONRPCClient
	cli             *socket.Client
	header          *abci.Header
	balances        map[string]std.AccountInfo
)

//adapter回调函数
func set(transID, txID int64, value map[string][]byte) {
	var err error
	if cli == nil {
		logger.Infof(flagCallbackURL)
		cli, err = socket.NewClient(flagCallbackURL, 10, false, logger)
		if err != nil {
			panic(err)
		}
	}

	data := make(map[string]string)
	for k, v := range value {
		data[k] = string(v)
	}

	result, err := cli.Call("set", map[string]interface{}{"transID": transID, "txID": txID, "data": data})
	if err != nil {
		logger.Errorf("socket set error: " + err.Error())
		logger.Flush()
		panic("socket set error: " + err.Error())
	}

	logger.Debugf("set return is %t", result.(bool))

	if result.(bool) == false {
		msg := "socket set error: return false"
		logger.Errorf(msg)
		panic(msg)
	}
}

func get(transID, txID int64, key string) []byte {
	if v, ok := balances[key]; ok {
		b, _ := jsoniter.Marshal(v)

		res := std.GetResult{
			Code: types.CodeOK,
			Msg:  "",
			Data: b,
		}
		resByte, _ := jsoniter.Marshal(res)
		return resByte
	}
	var err error
	if cli == nil {
		logger.Infof(flagCallbackURL)
		cli, err = socket.NewClient(flagCallbackURL, 0, false, logger)
		if err != nil {
			panic(err)
		}
	}

	result, err := cli.Call("get", map[string]interface{}{"transID": transID, "txID": txID, "key": key})
	if err != nil {
		logger.Errorf("socket get error: " + err.Error())
		logger.Flush()
		panic("socket get error: " + err.Error())
	}

	return []byte(result.(string))
}

func build(transID int64, txID int64, contractMeta std.ContractMeta) std.BuildResult {

	var err error
	if cli == nil {
		logger.Infof(flagCallbackURL)
		cli, err = socket.NewClient(flagCallbackURL, 0, false, logger)
		if err != nil {
			panic(err)
		}
	}

	resBytes, _ := jsoniter.Marshal(contractMeta)
	result, err := cli.Call("build", map[string]interface{}{"transID": transID, "txID": txID, "contractMeta": string(resBytes)})
	if err != nil {
		logger.Errorf("socket build error: " + err.Error())
		logger.Flush()
		panic("socket get error: " + err.Error())
	}

	var buildResult std.BuildResult
	err = jsoniter.Unmarshal([]byte(result.(string)), &buildResult)
	if err != nil {
		panic(err)
	}

	return buildResult
}

func getBlock(height int64) std.Block {

	if height == 0 {
		block := std.Block{
			ChainID:         header.ChainID,
			BlockHash:       header.LastBlockID.Hash, //todo
			Height:          header.Height,
			Time:            header.Time,
			NumTxs:          header.NumTxs,
			DataHash:        header.DataHash,
			ProposerAddress: header.ProposerAddress,
			RewardAddress:   header.RewardAddress,
			RandomNumber:    header.RandomeOfBlock,
			Version:         header.Version,
			LastBlockHash:   header.LastBlockID.Hash,
			LastCommitHash:  header.LastCommitHash,
			LastAppHash:     header.LastAppHash,
			LastFee:         int64(header.LastFee),
		}

		return block
	}
	var err error
	if cli == nil {
		logger.Infof(flagCallbackURL)
		cli, err = socket.NewClient(flagCallbackURL, 0, false, logger)
		if err != nil {
			panic(err)
		}
	}

	result, err := cli.Call("block", map[string]interface{}{"height": height})
	if err != nil {
		logger.Errorf("socket getBlock error: " + err.Error())
		logger.Flush()
		panic("rpc get error: " + err.Error())
	}

	resBytes, _ := jsoniter.Marshal(result.(map[string]interface{}))
	var blockResult std.Block
	err = jsoniter.Unmarshal(resBytes, &blockResult)

	return blockResult
}

//Routes routes map
//NOTE: Amino is registered in rpc/core/types/wire.go.
var Routes = map[string]socket.CallBackFunc{
	"Invoke":          Invoke,
	"McCommitTrans":   McCommitTrans,
	"McDirtyTrans":    McDirtyTrans,
	"McDirtyTransTx":  McDirtyTransTx,
	"McDirtyToken":    McDirtyToken,
	"McDirtyContract": McDirtyContract,
	"SetLogLevel":     SetLogLevel,
	"Health":          Health,
	"InitChain":       InitChain,
	"UpdateChain":     UpdateChain,
}

//RunRPC starts RPC service
func RunRPC(port int) error {
	logger = log.NewTMLogger(".", "smcsvc")
	logger.AllowLevel("debug")
	logger.SetOutputAsync(true)
	logger.SetOutputToFile(true)
	logger.SetOutputToScreen(true)
	logger.SetOutputFileSize(20000000)

	sdkhelper.Init(transfer, build, set, get, getBlock, &logger)
	balances = make(map[string]std.AccountInfo)

	// start server and wait forever
	svr, err := socket.NewServer("tcp://0.0.0.0:"+fmt.Sprintf("%d", port), Routes, 0, logger)
	if err != nil {
		tmcommon.Exit(err.Error())
	}

	// start server and wait forever
	err = svr.Start()
	if err != nil {
		tmcommon.Exit(err.Error())
	}


	return nil
}

//Invoke invoke function
func Invoke(req map[string]interface{}) (interface{}, error) {

	logger.Debugf("Invoke starting")
	transID := int64(req["transID"].(float64))
	txID := int64(req["txID"].(float64))
	mCallParam := req["callParam"].(map[string]interface{})
	var callParam types.RPCInvokeCallParam
	jsonStr, _ := jsoniter.Marshal(mCallParam)
	err := jsoniter.Unmarshal(jsonStr, &callParam)
	if err != nil {
		logger.Errorf("callParam Unmarshal error: " + err.Error())
		panic(err)
	}
	mBlockHeader := req["blockHeader"].(map[string]interface{})
	var blockHeader abci.Header
	jsonStr, _ = jsoniter.Marshal(mBlockHeader)
	logger.Debugf(string(jsonStr))
	err = jsoniter.Unmarshal(jsonStr, &blockHeader)
	if err != nil {
		logger.Errorf("Invoke error: " + err.Error())
		panic(err)
	}

	logger.Debug("Invoke", "transID", transID, "txID", txID)
	logger.Trace("smcRunSvc Invoke", "callParam", callParam)
	header = &blockHeader

	bals := make([]std.AccountInfo, 0)
	jsoniter.Unmarshal(callParam.Balances, &bals)
	for _, v := range bals {
		logger.Debug("Invoke sender balance", "token", v.Address, "balance", v.Balance)
		key := std.KeyOfAccountToken(callParam.Sender, v.Address)
		balances[key] = v
		defer delete(balances, key)
	}

	if callParam.To != ""{
        var bal std.AccountInfo
        jsoniter.Unmarshal(callParam.ToBalance, &bal)
        key := std.KeyOfAccountToken(callParam.To, bal.Address)
		 balances[key] = bal
		 defer delete(balances, key)
    }

	sdkReceipts := make([]sdkType.KVPair, 0)
	for _, v := range callParam.Receipts {
		sdkReceipts = append(sdkReceipts, sdkType.KVPair{Key: v.Key, Value: v.Value})
	}

	items := make([]sdkType.HexBytes, 0)
	for _, item := range callParam.Message.Items {
		items = append(items, []byte(item))
	}

	logger.Debugf("Invoke sdkhelper New")
	smc := sdkhelper.New(
		transID,
		txID,
		callParam.Sender,
		callParam.Tx.GasLimit,
		callParam.GasLeft,
		callParam.Tx.Note,
		callParam.Message.Contract,
		fmt.Sprintf("%x", callParam.Message.MethodID),
		items,
		sdkReceipts,
	)

	contractStub := stub.NewStub(smc, logger)
	logger.Debugf("Invoke contractStub Invoke")
	response := contractStub.Invoke(smc)
	smc.(*sdkimpl.SmartContract).Commit()
	logger.Debugf("Invoke contractStub Commit")

	resBytes, _ := jsoniter.Marshal(response)

	return string(resBytes), nil
}

//McCommitTrans commit transaction data of memory cache
func McCommitTrans(req map[string]interface{}) (interface{}, error) {
	transID := int64(req["transID"].(float64))
	sdkhelper.McCommit(transID)
	return true, nil
}

//McDirtyTrans dirty transaction data of memory cache
func McDirtyTrans(req map[string]interface{}) (interface{}, error) {
	transID := int64(req["transID"].(float64))
	sdkhelper.McDirtyTrans(transID)
	return true, nil
}

//McDirtyTransTx dirty tx data of transaction of memory cache
func McDirtyTransTx(req map[string]interface{}) (interface{}, error) {
	transID := int64(req["transID"].(float64))
	txID := int64(req["txID"].(float64))
	sdkhelper.McDirtyTransTx(transID, txID)
	return true, nil
}

//McDirtyToken dirty token data of memory cache
func McDirtyToken(req map[string]interface{}) (interface{}, error) {
	tokenAddr := req["tokenAddr"].(string)
	sdkhelper.McDirtyToken(tokenAddr)
	return true, nil
}

//McDirtyContract dirty contract data of memory cache
func McDirtyContract(req map[string]interface{}) (interface{}, error) {
	contractAddr := req["contractAddr"].(string)
	sdkhelper.McDirtyContract(contractAddr)
	return true, nil
}

//SetLogLevel sets log level
func SetLogLevel(req map[string]interface{}) (interface{}, error) {
	level := req["level"].(string)
	logger.AllowLevel(level)
	return true, nil
}

// Health return health message
func Health(req map[string]interface{}) (interface{}, error) {
	return "health", nil
}

// InitChain initial smart contract
func InitChain(req map[string]interface{}) (interface{}, error) {
	logger.Info("genesis contract InitChain")

	smc := newSMC(req)

	contractStub := stub.NewStub(smc, logger)
	logger.Debugf("Invoke contractStub InitChain")
	response := contractStub.InitChain(smc)
	smc.(*sdkimpl.SmartContract).Commit()
	logger.Debugf("Invoke contractStub Commit")

	resBytes, _ := jsoniter.Marshal(response)

	return string(resBytes), nil
}

// UpdateChain initial smart contract
func UpdateChain(req map[string]interface{}) (interface{}, error) {
	logger.Info("genesis contract UpdateChain")

	smc := newSMC(req)

	contractStub := stub.NewStub(smc, logger)
	logger.Debugf("Invoke contractStub UpdateChain")
	response := contractStub.UpdateChain(smc)
	smc.(*sdkimpl.SmartContract).Commit()
	logger.Debugf("Invoke contractStub Commit")

	resBytes, _ := jsoniter.Marshal(response)

	return string(resBytes), nil
}

func newSMC(req map[string]interface{}) sdk.ISmartContract {
	transID := int64(req["transID"].(float64))
	txID := int64(req["txID"].(float64))
	mCallParam := req["callParam"].(map[string]interface{})
	var callParam types.RPCInvokeCallParam
	jsonStr, _ := jsoniter.Marshal(mCallParam)
	err := jsoniter.Unmarshal(jsonStr, &callParam)
	if err != nil {
		logger.Errorf(err.Error())
		panic(err)
	}
	mBlockHeader := req["blockHeader"].(map[string]interface{})
	var blockHeader abci.Header
	jsonStr, _ = jsoniter.Marshal(mBlockHeader)
	logger.Debugf(string(jsonStr))
	err = jsoniter.Unmarshal(jsonStr, &blockHeader)
	if err != nil {
		logger.Errorf("Invoke error: " + err.Error())
		panic(err)
	}

	logger.Debug("InitChain/UpdateChain", "transID", transID, "txID", txID)
	logger.Trace("InitChain/UpdateChain", "callParam", callParam)

	sdkReceipts := make([]sdkType.KVPair, 0)
	for _, v := range callParam.Receipts {
		sdkReceipts = append(sdkReceipts, sdkType.KVPair{Key: v.Key, Value: v.Value})
	}

	items := make([]sdkType.HexBytes, 0)
	for _, item := range callParam.Message.Items {
		items = append(items, []byte(item))
	}

	smc := sdkimpl.SmartContract{}
	llState := llstate.NewLowLevelSDB(&smc, transID, txID)
	smc.SetLlState(llState)

	block := object.NewBlock(&smc, blockHeader.ChainID,  blockHeader.Version, sdkType.Hash{}, blockHeader.DataHash,
		blockHeader.Height, blockHeader.Time, blockHeader.NumTxs, blockHeader.ProposerAddress, blockHeader.RewardAddress,
		blockHeader.RandomeOfBlock, blockHeader.LastBlockID.Hash, blockHeader.LastCommitHash, blockHeader.LastAppHash,
		int64(blockHeader.LastFee))
	smc.SetBlock(block)

	helperObj := helper.NewHelper(&smc)
	smc.SetHelper(helperObj)

	contract := object.NewContractFromAddress(&smc, callParam.Message.Contract)
	msg := object.NewMessage(&smc, contract, "", items, callParam.Sender,
		nil, nil)
	smc.SetMessage(msg)

	return &smc
}

//TransferFunc is used to transfer token for crossing contract invoking.
// nolint unhandled
func transfer(sdk sdk.ISmartContract, tokenAddr, to types.Address, value bn.Number) ([]sdkType.KVPair, sdkType.Error) {
	logger.Debug("TransferFunc", "tokenAddress", tokenAddr, "to", to, "value", value)
	contract := sdk.Helper().ContractHelper().ContractOfToken(tokenAddr)
	logger.Info("Contract", "address", contract.Address(), "name", contract.Name(), "version", contract.Version())
	originMessage := sdk.Message()

	//todo 改成计算的方式
	mID := "af0228bc"
	newSdk := sdkhelper.OriginNewMessage(sdk, contract, mID, nil)

	//todo: 打包参数到message data
	// or 寻找一种方法生成InterfaceStub, 现在的设计会导致循环调用，不可使用。
	tobyte, _ := rlp.EncodeToBytes(to)
	valuebyte, _ := rlp.EncodeToBytes(value.Bytes())

	itemsbyte := make([]sdkType.HexBytes, 0)
	itemsbyte = append(itemsbyte, tobyte)
	itemsbyte = append(itemsbyte, valuebyte)

	newmsg := object.NewMessage(newSdk, newSdk.Message().Contract(), mID, itemsbyte, newSdk.Message().Sender().Address(), newSdk.Message().Origins(), nil)
	newSdk.(*sdkimpl.SmartContract).SetMessage(newmsg)
	contractStub := stub.NewStub(newSdk, logger)
	response := contractStub.Invoke(newSdk)
	logger.Debug("Invoke response", "code", response.Code, "tags", response.Tags)
	if response.Code != sdkType.CodeOK {
		return nil, sdkType.Error{ErrorCode: response.Code, ErrorDesc: response.Log}
	}

	// read receipts from response and append to original sdk message
	recKV := make([]sdkType.KVPair, 0)
	for _, v := range response.Tags {
		recKV = append(recKV, sdkType.KVPair{Key: v.Key, Value: v.Value})
	}
	newSdk.(*sdkimpl.SmartContract).SetMessage(originMessage)
	return recKV, sdkType.Error{ErrorCode: sdkType.CodeOK}
}

//RootCmd cmd
var RootCmd = &cobra.Command{
	Use:   "smcrunsvc",
	Short: "grpc",
	Long:  "smcsvc rpc console",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunRPC(flagRPCPort)
	},
}

func main() {
	err := excute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func excute() error {
	addFlags()
	addCommand()
	return RootCmd.Execute()
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the smc_service",
	Long:  "start the smc_service",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunRPC(flagRPCPort)
	},
}

func addStartFlags() {
	startCmd.PersistentFlags().IntVarP(&flagRPCPort, "port", "p", 8080, "The port of the smc rpc service")
	startCmd.PersistentFlags().StringVarP(&flagCallbackURL, "callbackUrl", "c", "http://localhost:8081", "The url of the adapter callback")
}

func addFlags() {
	addStartFlags()
}

func addCommand() {
	RootCmd.AddCommand(startCmd)
}

`

type Cmd struct {
	OrgID string
}

// GenStubCommon - generate the stub common go source
func GenCmd(rootDir, orgID string) error {

	newPath := filepath.Join(rootDir, "cmd/smcrunsvc")
	if err := os.MkdirAll(newPath, os.FileMode(0750)); err != nil {
		return err
	}
	filename := filepath.Join(newPath, "smcrunsvc.go")

	tmpl, err := template.New("smcrunsvc").Parse(templateText)
	if err != nil {
		return err
	}

	orgCmd := Cmd{OrgID: orgID}
	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, orgCmd); err != nil {
		return err
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}

	return nil
}
