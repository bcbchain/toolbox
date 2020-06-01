package query

import (
	rpcclient "github.com/bcbchain/bclib/rpc/lib/client"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
	"github.com/bcbchain/tendermint/types"
)

func GetHeader(client *rpcclient.JSONRPCClient, h int64) (*types.Header, error) {
	resultBlock := new(core_types.ResultBlock)
	_, err := client.Call("block", map[string]interface{}{"height": h}, resultBlock)
	if err != nil {
		return nil, err
	}
	return resultBlock.Block.Header, nil
}

func GetTx(client *rpcclient.JSONRPCClient, h int64) ([]string, error) {
	resultBlock := new(core_types.ResultBlock)
	_, err := client.Call("block", map[string]interface{}{"height": h}, resultBlock)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(resultBlock.Block.Txs))
	for i, v := range resultBlock.Block.Txs {
		result[i] = string([]byte(v))
	}
	return result, nil
}

func GetCurrentHeight(url string) (int64, error) {
	result, err := abciInfoQuery(url)
	if err != nil {
		return 0, err
	}
	return result.Response.LastBlockHeight, nil
}

func abciInfoQuery(url string) (resultABCIInfo *core_types.ResultABCIInfo, err error) {
	resultABCIInfo = new(core_types.ResultABCIInfo)
	client := GetClient(url, false)
	_, err = client.Call("abci_info", nil, resultABCIInfo)
	return
}

func GetClient(url string, keepAlive bool) *rpcclient.JSONRPCClient {

	return rpcclient.NewJSONRPCClientExWithTimeout(url, "", !keepAlive, 5)
}
