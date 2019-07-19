package mymixtype

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/rlp"
	sdktypes "blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/types"
	"contract/orgexample/code/mydice2win/v1.0/mydice2win"
	stubTypes "contract/stubcommon/types"

	"contract/orgteststub/code/mymixtype/v1.0/mymixtype"
	common2 "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"
)

const msg = "Message can not be nil."

// MyBasicTypeStub stub
type MyMixTypeStub struct {
	logger log.Logger
}

var _ stubTypes.IContractStub = (*MyMixTypeStub)(nil)

// New new stub
func New(logger log.Logger) stubTypes.IContractStub {

	var stub MyMixTypeStub
	stub.logger = logger

	return &stub
}

// InitChain initial smart contract
func (mc *MyMixTypeStub) InitChain(smc sdk.ISmartContract) (response types.Response) {

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (mc *MyMixTypeStub) UpdateChain(smc sdk.ISmartContract) (response types.Response) {
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.UpdateChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (mc *MyMixTypeStub) Mine(smc sdk.ISmartContract) (response types.Response) {
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.UpdateChain()

	response.Code = types.CodeOK
	return response
}

// Invoke invoke method
func (mc *MyMixTypeStub) Invoke(smc sdk.ISmartContract) types.Response {
	return mc.InvokeInternal(smc, true)
}

// InvokeInternal invoke method
func (mc *MyMixTypeStub) InvokeInternal(smc sdk.ISmartContract, feeFlag bool) types.Response {

	// TODO 手续费
	//emitFee(smc)

	switch smc.Message().MethodID() {
	case "1":
		return basic(smc)
	case "2":
		return slice(smc)
	case "3":
		return mapString(smc)
	case "4":
		return mapOther(smc)
	case "5":
		return mapSlice(smc)
	case "6":
		return complexDefine(smc)
	case "7":
		return longType(smc)
	}
	return types.Response{}
}

func basic(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()
	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v mymixtype.BasicTypes
	if err := rlp.DecodeBytes(itemsBytes[0], &v); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	myMixType := new(mymixtype.Mymixtype)
	myMixType.SetSdk(smc)

	data := myMixType.Basic(&v)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Data = string(data)
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}

func slice(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()
	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v mymixtype.SliceTypes
	if err := rlp.DecodeBytes(itemsBytes[0], &v); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	myMixType := new(mymixtype.Mymixtype)
	myMixType.SetSdk(smc)

	data := myMixType.Slice(&v)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Data = string(data)
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}

func mapString(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()
	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v mymixtype.MapStringTypes
	if err := rlp.DecodeBytes(itemsBytes[0], &v); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	myMixType := new(mymixtype.Mymixtype)
	myMixType.SetSdk(smc)

	data := myMixType.MapString(&v)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Data = string(data)
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}

func mapOther(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()
	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v mymixtype.MapOtherTypes
	if err := rlp.DecodeBytes(itemsBytes[0], &v); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	myMixType := new(mymixtype.Mymixtype)
	myMixType.SetSdk(smc)

	data := myMixType.MapOther(&v)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Data = string(data)
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}

func mapSlice(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()
	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v mymixtype.MapSliceTypes
	if err := rlp.DecodeBytes(itemsBytes[0], &v); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	myMixType := new(mymixtype.Mymixtype)
	myMixType.SetSdk(smc)

	data := myMixType.MapSlice(&v)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Data = string(data)
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}

func complexDefine(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()
	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v mymixtype.ComplexDefine
	if err := rlp.DecodeBytes(itemsBytes[0], &v); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	myMixType := new(mymixtype.Mymixtype)
	myMixType.SetSdk(smc)

	//v1_ := new(mymixtype.ComplexDefine)
	//jsoniter.Unmarshal(v, v1_)

	data := myMixType.Complex(&v)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Data = string(data)
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}

func longType(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()
	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v mymixtype.LongType
	if err := rlp.DecodeBytes(itemsBytes[0], &v); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	myMixType := new(mymixtype.Mymixtype)
	myMixType.SetSdk(smc)

	data := myMixType.Long(&v)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Data = string(data)
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}
