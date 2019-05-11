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
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdk/rlp"
	"blockchain/smcsdk/sdkimpl"
	"blockchain/smcsdk/sdkimpl/helper"
	"blockchain/smcsdk/sdkimpl/llstate"
	"blockchain/smcsdk/sdkimpl/object"
	"contract/{{.OrgID}}/stub"

	"blockchain/algorithm"
	"golang.org/x/crypto/sha3"
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
	p               *socket.ConnectionPool
	header          *abci.Header
	
	bMtx			sync.Mutex
	balances        map[string]std.AccountInfo
)

func pool() *socket.ConnectionPool {
	if p == nil {
		var err error
		p, err = socket.NewConnectionPool(flagCallbackURL, 4, logger)
		if err != nil {
			panic(err)
		}
	}

	return p
}

//adapter回调函数
func set(transID, txID int64, value map[string][]byte) {

	// for Marshal result can UnMarshal, it necessary
	data := make(map[string]string)
	for k, v := range value {
		if v == nil {
			data[k] = string([]byte{})
		} else {
			data[k] = string(v)
		}
	}

	logger.Debugf("[transID=%d][txID=%d]set data=%v", transID, txID, data)
	cli, err := pool().GetClient()
	if err != nil {
		msg := fmt.Sprintf("[transID=%d][txID=%d]socket set error: %s", transID, txID, err.Error())
		logger.Errorf(msg)
		panic(err)
	}
	defer pool().ReleaseClient(cli)

	result, err := cli.Call("set", map[string]interface{}{"transID": transID, "txID": txID, "data": data}, 10)
	if err != nil {
		msg := fmt.Sprintf("[transID=%d][txID=%d]socket set error: %s", transID, txID, err.Error())
		logger.Errorf(msg)
		panic(err)
	}
	logger.Debugf("[transID=%d][txID=%d]set return is %t", transID, txID, result.(bool))

	if result.(bool) == false {
		msg := fmt.Sprintf("[transID=%d][txID=%d]socket set error: return false", transID, txID)
		logger.Errorf(msg)
		panic(msg)
	}
}

func getBalance(key string) []byte {
	bMtx.Lock()
	defer bMtx.Unlock()
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

	return nil
}

func get(transID, txID int64, key string) []byte {
	b := getBalance(key)
	if len(b) > 0 {
		return b
	}

	logger.Debugf("[transID=%d][txID=%d]get key=%s", transID, txID, key)
	cli, err := pool().GetClient()
	if err != nil {
		msg := fmt.Sprintf("[transID=%d][txID=%d]socket get error: %s", transID, txID, err.Error())
		logger.Errorf(msg)
		panic(err)
	}
	defer pool().ReleaseClient(cli)

	result, err := cli.Call("get", map[string]interface{}{"transID": transID, "txID": txID, "key": key}, 10)
	if err != nil {
		msg := fmt.Sprintf("[transID=%d][txID=%d]socket get error: %s", transID, txID, err.Error())
		logger.Errorf(msg)
		panic(err)
	}
	logger.Debugf("[transID=%d][txID=%d]get key=%s, result=%v", transID, txID, key, result)

	return []byte(result.(string))
}

func build(transID int64, txID int64, contractMeta std.ContractMeta) std.BuildResult {

	resBytes, _ := jsoniter.Marshal(contractMeta)
	logger.Debugf("[transID=%d][txID=%d]build orgID=%s contract=%s version=%s", transID, txID, contractMeta.OrgID, contractMeta.Name, contractMeta.Version)
	cli, err := pool().GetClient()
	if err != nil {
		msg := fmt.Sprintf("[transID=%d][txID=%d]socket build error: %s", transID, txID, err.Error())
		logger.Errorf(msg)
		panic(err)
	}	
	defer pool().ReleaseClient(cli)

	result, err := cli.Call("build", map[string]interface{}{"transID": transID, "txID": txID, "contractMeta": string(resBytes)}, 180)
	if err != nil {
		msg := fmt.Sprintf("[transID=%d][txID=%d]socket build error: %s", transID, txID, err.Error())
		logger.Errorf(msg)
		panic(err)
	}
	logger.Debugf("[transID=%d][txID=%d]build result=%v", transID, txID, result)

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
		block.BlockHash = blockHash(block)

		return block
	}

	logger.Debugf("get block height=%d", height)
	cli, err := pool().GetClient()
	if err != nil {
		msg := fmt.Sprintf("socket getBlock error: %s", err.Error())
		logger.Errorf(msg)
		panic(err)
	}
	defer pool().ReleaseClient(cli)

	result, err := cli.Call("block", map[string]interface{}{"height": height}, 10)
	if err != nil {
		msg := fmt.Sprintf("socket getBlock error: %s", err.Error())
		logger.Errorf(msg)
		panic(err)
	}
	resBytes, _ := jsoniter.Marshal(result.(map[string]interface{}))
	logger.Debugf("get block height=%d, result=%s", height, string(resBytes))

	var blockResult std.Block
	err = jsoniter.Unmarshal(resBytes, &blockResult)

	return blockResult
}

func blockHash(block std.Block) sdkType.HexBytes {
	sha256 := sha3.New256()
	sha256.Write([]byte(block.ChainID))
	sha256.Write(algorithm.IntToBytes(int(block.Height)))
	sha256.Write(algorithm.IntToBytes(int(block.Time)))
	sha256.Write(algorithm.IntToBytes(int(block.NumTxs)))
	sha256.Write(block.DataHash)
	sha256.Write([]byte(block.ProposerAddress))
	sha256.Write([]byte(block.RewardAddress))
	sha256.Write(block.RandomNumber)
	sha256.Write([]byte(block.Version))
	sha256.Write(block.LastBlockHash)
	sha256.Write(block.LastCommitHash)
	sha256.Write(block.LastAppHash)
	sha256.Write(algorithm.IntToBytes(int(block.LastFee)))

	return sha256.Sum(nil)
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
	logger.AllowLevel("info")
	logger.SetOutputAsync(true)
	logger.SetOutputToFile(true)
	logger.SetOutputToScreen(false)
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
	logger.Tracef("Invoke starting")

	transID := int64(req["transID"].(float64))
	txID := int64(req["txID"].(float64))

	// setup call parameter
	mCallParam := req["callParam"].(map[string]interface{})
	var callParam types.RPCInvokeCallParam
	jsonStr, _ := jsoniter.Marshal(mCallParam)
	logger.Debugf("[transID=%d][txID=%d]callParam=%s", transID, txID, string(jsonStr))
	err := jsoniter.Unmarshal(jsonStr, &callParam)
	if err != nil {
		logger.Errorf("[transID=%d][txID=%d]callParam Unmarshal error", transID, txID, err.Error())
		panic(err)
	}

	// setup block header
	mBlockHeader := req["blockHeader"].(map[string]interface{})
	var blockHeader abci.Header
	jsonStr, _ = jsoniter.Marshal(mBlockHeader)
	logger.Debugf("[transID=%d][txID=%d]blockHeader=%s", transID, txID, string(jsonStr))
	err = jsoniter.Unmarshal(jsonStr, &blockHeader)
	if err != nil {
		logger.Errorf("[transID=%d][txID=%d]invoke error=%s", transID, txID, err.Error())
		panic(err)
	}
	header = &blockHeader

	bals := make([]std.AccountInfo, 0)
	err = jsoniter.Unmarshal(callParam.Balances, &bals)
	if err != nil {
		logger.Errorf("[transID=%d][txID=%d]balances Unmarshal error=%s", transID, txID, err.Error())
		panic(err)
	}
	for _, v := range bals {
		logger.Debugf("[transID=%d][txID=%d]Invoke sender balance, token=%s, balance=%s", transID, txID, v.Address, v.Balance.String())
		key := std.KeyOfAccountToken(callParam.Sender, v.Address)
		bMtx.Lock()
		balances[key] = v
		bMtx.Unlock()
		defer deleteMapKey(key)
	}

	if callParam.To != "" {
		var bal std.AccountInfo
		err = jsoniter.Unmarshal(callParam.ToBalance, &bal)
		if err != nil {
			logger.Errorf("[transID=%d][txID=%d]toBalance Unmarshal error=%s", transID, txID, err.Error())
			panic(err)
		}
		key := std.KeyOfAccountToken(callParam.To, bal.Address)
		bMtx.Lock()
		balances[key] = bal
		bMtx.Unlock()
		defer deleteMapKey(key)
	}

	logger.Infof("[transID=%d][txID=%d]invoke", transID, txID)

	sdkReceipts := make([]sdkType.KVPair, 0)
	for _, v := range callParam.Receipts {
		sdkReceipts = append(sdkReceipts, sdkType.KVPair{Key: v.Key, Value: v.Value})
	}

	items := make([]sdkType.HexBytes, 0)
	for _, item := range callParam.Message.Items {
		items = append(items, []byte(item))
	}

	logger.Debugf("[transID=%d][txID=%d]invoke sdkhelper New", transID, txID)
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
	var response types.Response
	if contractStub == nil {
		response.Code = sdkType.ErrInvalidAddress
		response.Log = fmt.Sprintf("[transID=%d][txID=%d]Call contract=%s,version=%s is not exist or lost",
			transID, txID, smc.Message().Contract().Address(), smc.Message().Contract().Version())
	} else {
		logger.Debugf("[transID=%d][txID=%d]contractStub Invoke", transID, txID)
		response = contractStub.Invoke(smc)

		logger.Debugf("[transID=%d][txID=%d]contractStub Commit", transID, txID)
		smc.(*sdkimpl.SmartContract).Commit()
	}

	resBytes, _ := jsoniter.Marshal(response)

	return string(resBytes), nil
}

func deleteMapKey(key string) {
	bMtx.Lock()
	defer bMtx.Unlock()
	delete(balances, key)
}

//McCommitTrans commit transaction data of memory cache
func McCommitTrans(req map[string]interface{}) (interface{}, error) {

	transID := int64(req["transID"].(float64))
	logger.Infof("[transID=%d]McCommitTrans", transID)

	sdkhelper.McCommit(transID)
	return true, nil
}

//McDirtyTrans dirty transaction data of memory cache
func McDirtyTrans(req map[string]interface{}) (interface{}, error) {

	transID := int64(req["transID"].(float64))
	logger.Infof("[transID=%d]McDirtyTrans", transID)

	sdkhelper.McDirtyTrans(transID)
	return true, nil
}

//McDirtyTransTx dirty tx data of transaction of memory cache
func McDirtyTransTx(req map[string]interface{}) (interface{}, error) {

	transID := int64(req["transID"].(float64))
	txID := int64(req["txID"].(float64))
	logger.Infof("[transID=%d][txID=%d]McDirtyTransTx", transID)

	sdkhelper.McDirtyTransTx(transID, txID)
	return true, nil
}

//McDirtyToken dirty token data of memory cache
func McDirtyToken(req map[string]interface{}) (interface{}, error) {

	tokenAddr := req["tokenAddr"].(string)
	logger.Infof("McDirtyToken tokenAddr=%s", tokenAddr)

	sdkhelper.McDirtyToken(tokenAddr)
	return true, nil
}

//McDirtyContract dirty contract data of memory cache
func McDirtyContract(req map[string]interface{}) (interface{}, error) {

	contractAddr := req["contractAddr"].(string)
	logger.Infof("McDirtyToken contractAddr=%s", contractAddr)

	sdkhelper.McDirtyContract(contractAddr)
	return true, nil
}

//SetLogLevel sets log level
func SetLogLevel(req map[string]interface{}) (interface{}, error) {

	level := req["level"].(string)
	logger.Infof("SetLogLevel level=%s", level)

	logger.AllowLevel(level)
	return true, nil
}

// Health return health message
func Health(req map[string]interface{}) (interface{}, error) {
	return "health", nil
}

// InitChain initial smart contract
func InitChain(req map[string]interface{}) (interface{}, error) {
	logger.Info("contract InitChain")

	smc := newSMC(req)
	contractStub := stub.NewStub(smc, logger)

	logger.Debugf("Invoke contractStub InitChain")
	response := contractStub.InitChain(smc)

	logger.Debugf("Invoke contractStub Commit")
	smc.(*sdkimpl.SmartContract).Commit()

	resBytes, _ := jsoniter.Marshal(response)
	return string(resBytes), nil
}

// UpdateChain initial smart contract
func UpdateChain(req map[string]interface{}) (interface{}, error) {
	logger.Info("contract UpdateChain")

	smc := newSMC(req)
	contractStub := stub.NewStub(smc, logger)

	logger.Debugf("Invoke contractStub UpdateChain")
	response := contractStub.UpdateChain(smc)

	logger.Debugf("Invoke contractStub Commit")
	smc.(*sdkimpl.SmartContract).Commit()

	resBytes, _ := jsoniter.Marshal(response)
	return string(resBytes), nil
}

func newSMC(req map[string]interface{}) sdk.ISmartContract {

	transID := int64(req["transID"].(float64))
	txID := int64(req["txID"].(float64))

	mCallParam := req["callParam"].(map[string]interface{})
	var callParam types.RPCInvokeCallParam
	jsonStr, _ := jsoniter.Marshal(mCallParam)
	logger.Debugf("[transID=%d][txID=%d]callParam=%s", transID, txID, string(jsonStr))
	err := jsoniter.Unmarshal(jsonStr, &callParam)
	if err != nil {
		logger.Errorf("[transID=%d][txID=%d]callParam Unmarshal error", transID, txID, err.Error())
		panic(err)
	}

	mBlockHeader := req["blockHeader"].(map[string]interface{})
	var blockHeader abci.Header
	jsonStr, _ = jsoniter.Marshal(mBlockHeader)
	logger.Debugf("[transID=%d][txID=%d]callParam=%s", transID, txID, string(jsonStr))
	err = jsoniter.Unmarshal(jsonStr, &blockHeader)
	if err != nil {
		logger.Errorf("[transID=%d][txID=%d]invoke error=%s", transID, txID, err.Error())
		panic(err)
	}

	logger.Debugf("[transID=%d][txID=%d]create smart contract object", transID, txID)
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

	block := object.NewBlock(&smc, blockHeader.ChainID, blockHeader.Version, sdkType.Hash{}, blockHeader.DataHash,
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

	mID := "44d8ca60"
	newSdk := sdkhelper.OriginNewMessage(sdk, contract, mID, nil)

	items := wrapInvokeParams(to, value)

	newmsg := object.NewMessage(newSdk, newSdk.Message().Contract(), mID, items, newSdk.Message().Sender().Address(), newSdk.Message().Origins(), nil)
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

// wrapInvokeParams - wrap contract parameters
func wrapInvokeParams(params ...interface{}) []sdkType.HexBytes {
	paramsRlp := make([]sdkType.HexBytes, 0)
	for _, param := range params {
		var paramRlp []byte
		var err error

		paramRlp, err = rlp.EncodeToBytes(param)
		if err != nil {
			panic(err)
		}
		paramsRlp = append(paramsRlp, paramRlp)
	}
	return paramsRlp
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
	go func() {
		if e := http.ListenAndServe(":2019", nil); e != nil {
			fmt.Println("pprof cannot start!!!")
		}
	}()

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
