package common

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	rpcclient "github.com/bcbchain/bclib/rpc/lib/client"
	"github.com/bcbchain/sdk/sdk/types"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
	"strings"
	"time"
)

//网络请求和结果解析
func DoHttpRequestAndParse(nodeAddrSlice []string, methodName string, params map[string]interface{}, result interface{}) (err error) {

	for i, nodeAddr := range nodeAddrSlice {
		rpc := rpcclient.NewJSONRPCClientEx(nodeAddr, "", true)
		_, err = rpc.Call(methodName, params, result)
		if err == nil {
			break
		} else {
			if i == len(nodeAddrSlice)-1 {
				splitErr := strings.Split(err.Error(), ":")
				return errors.New(strings.Trim(splitErr[len(splitErr)-1], " "))
			}
		}
	}

	return
}

// 网络请求和结果解析
func DoHttpCommitTxAndParseAsync(nodeAddrSlice []string, txStr string) (result *core_types.ResultTx, err error) {

	bct := new(core_types.ResultBroadcastTx)

	for i, nodeAddr := range nodeAddrSlice {
		rpc := rpcclient.NewJSONRPCClientEx(nodeAddr, "", true)
		_, err = rpc.Call("broadcast_tx_async", map[string]interface{}{"tx": []byte(txStr)}, bct)
		if err == nil {
			result = new(core_types.ResultTx)
			for {
				err = DoHttpRequestAndParse(nodeAddrSlice, "tx", map[string]interface{}{"hash": strings.ToUpper(hex.EncodeToString(bct.Hash))}, result)
				if err != nil {
					return
				}

				if result.CheckResult.Code != 0 && result.CheckResult.Code != types.CodeOK {
					return
				}

				if result.DeliverResult.Code != 0 {
					return
				}

				time.Sleep(1 * time.Second)
			}
		} else {
			if i == len(nodeAddrSlice)-1 {
				splitErr := strings.Split(err.Error(), ":")
				return nil, errors.New(strings.Trim(splitErr[len(splitErr)-1], " "))
			}
		}
	}

	return result, nil
}

func DoHttpQueryAndParse(nodeAddrSlice []string, key string, data interface{}) (err error) {
	// 如果网络问题，则重复尝试访问
	var value []byte
	for {
		value, err = DoHttpQuery(nodeAddrSlice, key)
		if err == nil {
			break
		}
		fmt.Println("http query by urls fail...")
	}

	if len(value) == 0 {
		return errors.New("return value is empty, please check key=" + key)
	}

	err = json.Unmarshal(value, data)

	return
}

func DoHttpQuery(nodeAddrSlice []string, key string) (value []byte, err error) {

	result := new(core_types.ResultABCIQuery)
	for i, nodeAddr := range nodeAddrSlice {
		rpc := rpcclient.NewJSONRPCClientEx(nodeAddr, "", true)
		_, err = rpc.Call("abci_query", map[string]interface{}{"path": key}, result)
		if err == nil {
			break
		} else {
			if i == len(nodeAddrSlice)-1 {
				splitErr := strings.Split(err.Error(), ":")
				return nil, errors.New(strings.Trim(splitErr[len(splitErr)-1], " "))
			}
		}
	}
	value = result.Response.Value

	return
}
