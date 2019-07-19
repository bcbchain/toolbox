package mystoragestub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"

	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/rlp"
	"contract/orgexample/code/mystorage/v1.0/mystorage"
	"github.com/tendermint/tmlibs/log"
)

//MyStorageStub an object
type MyStorageStub struct {
	logger log.Logger
}

var _ stubType.IContractStub = (*MyStorageStub)(nil)

//New generate a stub
func New(logger log.Logger) stubType.IContractStub {
	return &MyStorageStub{logger: logger}
}

//FuncRecover recover panic by Assert
func FuncRecover(response *bcType.Response) {
	if err := recover(); err != nil {
		if _, ok := err.(types.Error); ok {
			error := err.(types.Error)
			response.Code = error.ErrorCode
			response.Log = error.Error()
		} else {
			panic(err)
		}
	}
}

// InitChain initial smart contract
func (pbs *MyStorageStub) InitChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	contractObj := new(mystorage.MyStorage)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (pbs *MyStorageStub) UpdateChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	response.Code = types.CodeOK
	return response
}

// Mine update smart contract
func (pbs *MyStorageStub) Mine(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	response.Code = types.CodeOK
	return response
}

//Invoke invoke function
func (pbs *MyStorageStub) Invoke(smc sdk.ISmartContract) (response bcType.Response) {
	return pbs.InvokeInternal(smc, true)
}

//InvokeInterface invoke function
func (pbs *MyStorageStub) InvokeInternal(smc sdk.ISmartContract, feeFlag bool) (response bcType.Response) {
	defer FuncRecover(&response)

	// 生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(smc, feeFlag)
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(smc.Message(), response.Tags, "", fee, gasUsed, smc.Tx().GasLimit(), err)
		return
	}

	var data string
	err = types.Error{ErrorCode: types.CodeOK}
	switch smc.Message().MethodID() {
	case "a21de13a": // prototype: Set(uint64)
		set(smc)
	case "b8b06eb1": // prototype: Get()
		data = get(smc)
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(smc.Message(), response.Tags, data, fee, gasUsed, smc.Tx().GasLimit(), err)
	return
}

func set(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 uint64
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mystorage.MyStorage)
	contractObj.SetSdk(smc)
	contractObj.Set(v0)
}

func get(smc sdk.ISmartContract) string {
	items := smc.Message().Items()
	sdk.Require(len(items) == 0, types.ErrStubDefined, "Invalid message data")

	contractObj := new(mystorage.MyStorage)
	contractObj.SetSdk(smc)
	rst0 := contractObj.Get()
	resultList := make([]interface{}, 0)
	resultList = append(resultList, rst0)

	resBytes, _ := jsoniter.Marshal(resultList)
	return string(resBytes)
}
