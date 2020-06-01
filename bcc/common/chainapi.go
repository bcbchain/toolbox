package common

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bcbchain/bclib/rpc/lib/client"
	rpctypes "github.com/bcbchain/bclib/rpc/lib/types"
	"github.com/bcbchain/bclib/tendermint/go-amino"
	"github.com/bcbchain/bclib/types"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"time"
)

//网络请求和结果解析-故障队列版
func DoHttpRequestAndParseExBlock(nodeAddrSlice []string, methodName string, params map[string]interface{}, result interface{}) (err error) {

	for {

		RWLock.Lock()
		if len(nodeAddrSlice) == 0 {
			RWLock.Unlock()
			return errors.New("no available nodes can to connect")
		}

		length := len(nodeAddrSlice)

		var rnd int
		if length > 1 {
			rnd = rand.Intn(length - 1)
		} else {
			rnd = 0
		}

		url := nodeAddrSlice[rnd]
		RWLock.Unlock()

		err = CallChainApi(url, methodName, params, result)
		if err == nil {
			break
		} else {
			RWLock.Lock()
			if _, ok := FaultCounterMap[url]; !ok {
				FaultCounterMap[url] = 0
			}
			FaultCounterMap[url] += 1

			if FaultCounterMap[url] > 10 {
				if rnd == length-1 {
					nodeAddrSlice = append(nodeAddrSlice[:rnd])
				} else {
					nodeAddrSlice = append(nodeAddrSlice[:rnd], nodeAddrSlice[rnd+1:]...)
				}
				length -= 1
			}
			RWLock.Unlock()

			if length <= len(nodeAddrSlice)/3 {
				go DealFaultUrls(nodeAddrSlice, methodName, params, result)
			}

			if length == 0 {
				splitErr := strings.Split(err.Error(), ":")
				return errors.New(strings.Trim(splitErr[len(splitErr)-1], " "))
			}
		}
	}

	return
}

func CallChainApi(url string, methodName string, params map[string]interface{}, result interface{}) (err error) {

	rpc := NewJSONRPCClientEx(url, "", true)
	_, err = rpc.Call(methodName, params, result)
	return
}

func DealFaultUrls(nodeAddrSlice []string, methodName string, params map[string]interface{}, result interface{}) {
	RWLock.Lock()
	FaultUrls2 := FaultCounterMap
	RWLock.Unlock()

	for k, _ := range FaultUrls2 {
		err := CallChainApi(k, methodName, params, result)
		if err == nil {
			RWLock.Lock()
			nodeAddrSlice = append(nodeAddrSlice, k)
			FaultCounterMap[k] = 0
			RWLock.Unlock()
		}
	}
}

type JSONRPCClient struct {
	address string
	client  *http.Client
	cdc     *amino.Codec
}

func NewJSONRPCClientEx(remote, certFile string, disableKeepAlive bool) *JSONRPCClient {
	var pool *x509.CertPool
	if certFile != "" {
		pool = x509.NewCertPool()
		caCert, err := ioutil.ReadFile(certFile)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		pool.AppendCertsFromPEM(caCert)
	}

	address, client := makeHTTPSClient(remote, pool, disableKeepAlive)

	return &JSONRPCClient{
		address: address,
		client:  client,
		cdc:     rpcclient.CDC,
	}
}

func makeHTTPSClient(remoteAddr string, pool *x509.CertPool, disableKeepAlive bool) (string, *http.Client) {
	//_, dialer := makeHTTPDialer(remoteAddr)

	tr := new(http.Transport)
	tr.DisableKeepAlives = disableKeepAlive
	tr.IdleConnTimeout = time.Second * 120
	if pool != nil {
		tr.TLSClientConfig = &tls.Config{RootCAs: pool}
	} else {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return remoteAddr, &http.Client{Transport: tr, Timeout: time.Duration(time.Second * 3)}
}

func (c *JSONRPCClient) Call(method string, params map[string]interface{}, result interface{}) (interface{}, error) {
	//request, err := types.MapToRequest("jsonrpc-client", method, params)
	request, err := rpctypes.MapToRequest("jsonrpc-client", method, params)
	if err != nil {
		return nil, err
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println("lib client http_client error to json.Marshal(request)")
		return nil, err
	}

	// log.Info(string(requestBytes))
	requestBuf := bytes.NewBuffer(requestBytes)
	// log.Info(Fmt("RPC request to %v (%v): %v", c.remote, method, string(requestBytes)))
	httpResponse, err := c.client.Post(c.address, "text/json", requestBuf)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close() // nolint: errcheck

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	// 	log.Info(Fmt("RPC response: %v", string(responseBytes)))
	return unmarshalResponseBytes(c.cdc, responseBytes, result)
}

func unmarshalResponseBytes(cdc *amino.Codec, responseBytes []byte, result interface{}) (interface{}, error) {
	// Read response.  If rpc/core/types is imported, the result will unmarshal
	// into the correct type.
	// log.Notice("response", "response", string(responseBytes))
	var err error
	//response := &types.RPCResponse{}
	response := &rpctypes.RPCResponse{}
	err = json.Unmarshal(responseBytes, response)
	if err != nil {
		return nil, errors.New("Error unmarshalling rpc response: " + err.Error())
	}
	if response.Error != nil {
		return nil, errors.New("Response error: " + response.Error.Error())
	}
	// Unmarshal the RawMessage into the result.
	err = cdc.UnmarshalJSON(response.Result, result)
	if err != nil {
		return nil, errors.New("Error unmarshalling rpc response result: " + err.Error())
	}
	return result, nil
}

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
func DoHttpCommitTxAndParse(nodeAddrSlice []string, txStr string) (result *core_types.ResultBroadcastTxCommit, err error) {

	result = new(core_types.ResultBroadcastTxCommit)

	for i, nodeAddr := range nodeAddrSlice {
		rpc := rpcclient.NewJSONRPCClientEx(nodeAddr, "", true)
		_, err = rpc.Call("broadcast_tx_commit", map[string]interface{}{"tx": []byte(txStr)}, result)
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

func DoBulkHttpQueryAndParse(nodeAddrSlice []string, key string, data interface{}) (result []interface{}, err error) {
	values, err := DoBulkHttpQuery(nodeAddrSlice, key)
	if err != nil {
		return
	}

	if len(values) == 0 {
		return nil, errors.New("return value is empty, please check key=" + key)
	}

	for _, value := range values {
		err = json.Unmarshal(value, data)
		temp := reflect.ValueOf(data).Elem().Interface()
		result = append(result, temp)
	}
	return
}
func DoBulkHttpQuery(nodeAddrSlice []string, key string) (value [][]byte, err error) {

	result := new(core_types.ResultABCIQueryEx)
	for i, nodeAddr := range nodeAddrSlice {
		rpc := rpcclient.NewJSONRPCClientEx(nodeAddr, "", true)
		_, err = rpc.Call("abci_query_ex", map[string]interface{}{"path": key}, result)
		if err == nil {
			break
		} else {
			if i == len(nodeAddrSlice)-1 {
				splitErr := strings.Split(err.Error(), ":")
				return nil, errors.New(strings.Trim(splitErr[len(splitErr)-1], " "))
			}
		}
	}

	for _, KeyValues := range result.Response.KeyValues {
		value = append(value, KeyValues.Value)
	}
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

func DoHttpQueryForRpc(nodeAddrSlice []string, key string) (result *core_types.ResultABCIQuery, err error) {

	result = new(core_types.ResultABCIQuery)
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

	return
}
