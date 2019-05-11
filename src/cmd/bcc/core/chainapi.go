package core

import (
	"common/rpc/lib/client"
	"encoding/json"
	"errors"
	"github.com/tendermint/tendermint/rpc/core/types"
	"strings"
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

//网络请求和结果解析
func DoHttpCommitTxAndParse(nodeAddrSlice []string, txStr string) (result *core_types.ResultBroadcastTxCommit, err error) {

	result = new(core_types.ResultBroadcastTxCommit)

	for i, nodeAddr := range nodeAddrSlice {
		rpc := rpcclient.NewJSONRPCClientEx(nodeAddr, "", true)
		_, err := rpc.Call("broadcast_tx_commit", map[string]interface{}{"tx": []byte(txStr)}, result)
		if err == nil {
			break
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

	value, err := DoHttpQuery(nodeAddrSlice, key)
	if err != nil {
		return
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
