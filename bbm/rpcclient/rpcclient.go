//rpcclient.go 创建RPC client, 并发送RPC request到tmcore
package rpcclient

import (
	rpcclient "github.com/bcbchain/bclib/rpc/lib/client"
	"github.com/bcbchain/bclib/tendermint/tmlibs/log"
	"github.com/bcbchain/bclib/types"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
	"strings"
	"time"
)

//rpcclient client
type RPCClient struct {
	remoteUrl string
	logger    log.Logger
}

const errMaxRetry = 40

//InitRPCClient create and init RPC client
func InitClient(remoteUrl string, logger log.Logger) RPCClient {

	return RPCClient{remoteUrl: remoteUrl, logger: logger}
}

func (client *RPCClient) Query(key string) []byte {
	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	result := new(core_types.ResultABCIQuery)
	var err error
	for i := 0; i < errMaxRetry; i++ {
		_, err = rpc.Call("abci_query", map[string]interface{}{"path": key}, result)
		if err == nil {
			break
		}
		// retry
		client.logger.Error("abci_query error", "key", key, "error", err)
		time.Sleep(time.Second * 2)
		continue
	}

	//something error
	if err != nil {
		client.logger.Error("abci_query error, application panic", "error", err)
		panic(err)
	}

	return result.Response.Value
}

func (client *RPCClient) SendTx(tx []byte) types.Response {

	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	result := new(core_types.ResultBroadcastTxCommit)

	var err error
	for i := 0; i < errMaxRetry; i++ {
		_, err = rpc.Call("broadcast_tx_commit", map[string]interface{}{"tx": tx}, result)
		if err == nil {
			break
		}

		// tx already send
		if strings.Contains(err.Error(), "Tx already exists in cache") {
			err = nil
			client.logger.Error("broadcast_tx_commit recovery", "time", time.Now())
			break
		}

		// retry
		client.logger.Error("broadcast_tx_commit failed", "time", time.Now(), "error", err.Error())
		time.Sleep(time.Second * 2)
		continue
	}
	//something error
	if err != nil {
		client.logger.Error("broadcast_tx_commit failed, application panic", "time", time.Now(), "error", err.Error())
		panic(err)
	}

	var resp types.Response
	if result.CheckTx.Code != types.CodeOK {
		resp.Code = result.CheckTx.Code
		resp.Log = result.CheckTx.Log
	} else {
		resp.Code = result.DeliverTx.Code
		resp.Log = result.DeliverTx.Log
	}
	resp.Fee = int64(result.DeliverTx.Fee)
	resp.GasLimit = int64(result.DeliverTx.GasLimit)
	resp.TxHash = result.Hash
	resp.Height = result.Height
	resp.GasUsed = int64(result.DeliverTx.GasUsed)
	resp.Tags = result.DeliverTx.Tags
	resp.Info = result.DeliverTx.Info
	resp.Data = result.DeliverTx.Data

	return resp
}

func (client *RPCClient) InitChain() {

}

func (client *RPCClient) EndAndCommit() {

}

func (client *RPCClient) Status() core_types.ResultStatus {
	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	par := make(map[string]interface{})
	result := new(core_types.ResultStatus)
	_, err := rpc.Call("status", par, result)
	if err != nil {
		client.logger.Error("status error", "error", err)
		return core_types.ResultStatus{}
	}

	return *result
}

func (client *RPCClient) Health() core_types.ResultHealth {
	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	par := make(map[string]interface{})
	result := new(core_types.ResultHealth)

	var err error
	for i := 0; i < errMaxRetry; i++ {
		_, err = rpc.Call("health", par, result)
		if err == nil {
			break
		}

		// retry
		client.logger.Error("status error ", "error", err)
		time.Sleep(time.Second * 2)
		continue
	}

	if err != nil {
		client.logger.Error("status error, application panic", "error", err)
		panic(err)
	}

	return *result
}

func (client *RPCClient) NumUnconfirmedTxs() (core_types.ResultUnconfirmedTxs, error) {
	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	par := make(map[string]interface{})
	result := new(core_types.ResultUnconfirmedTxs)
	var err error
	for i := 0; i < errMaxRetry; i++ {
		_, err = rpc.Call("num_unconfirmed_txs", par, result)
		if err == nil {
			break
		}

		// retry
		client.logger.Error("num_unconfirmed_txs error ", "error", err)
		time.Sleep(time.Second * 2)
		continue
	}

	if err != nil {
		client.logger.Error("status error, application panic", "error", err)
		panic(err)
	}
	return *result, nil
}

func (client *RPCClient) SendTxAsync(tx []byte) core_types.ResultBroadcastTxCommit {
	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	result := new(core_types.ResultBroadcastTxCommit)
	var err error
	for i := 0; i < errMaxRetry; i++ {
		_, err = rpc.Call("broadcast_tx_async", map[string]interface{}{"tx": tx}, result)
		if err == nil {
			break
		}

		// tx already send
		if strings.Contains(err.Error(), "Tx already exists in cache") {
			client.logger.Error("broadcast_tx_async recovery", "time", time.Now())
			break
		}

		// retry
		client.logger.Error("broadcast_tx_async error ", "error", err)
		time.Sleep(time.Second * 2)
		continue
	}

	if err != nil {
		client.logger.Error("broadcast_tx_async error, application panic ", "error", err)
		panic(err)
	}
	var resp types.Response
	resp.TxHash = result.Hash

	//fmt.Println("当前交易hash:", result.Hash)
	return *result
}

func (client *RPCClient) BlockWithHeight(height int64) (*core_types.ResultBlock, error) {
	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	result := new(core_types.ResultBlock)
	var err error
	for i := 0; i < errMaxRetry; i++ {
		_, err = rpc.Call("block", map[string]interface{}{"height": height}, result)
		if err == nil {
			break
		}

		// retry
		client.logger.Error("block error ", "error", err)
		time.Sleep(time.Second * 2)
		continue
	}

	if err != nil {
		client.logger.Error("block error, application panic ", "error", err)
		panic(err)
	}
	return result, nil
}

//validators?height=_
func (client *RPCClient) Validators(height int64) *core_types.ResultValidators {
	rpc := rpcclient.NewJSONRPCClientEx(client.remoteUrl, "", true)

	result := new(core_types.ResultValidators)
	_, err := rpc.Call("validators", map[string]interface{}{"height": height}, result)
	if err != nil {
		panic(err)
	}
	return result
}
