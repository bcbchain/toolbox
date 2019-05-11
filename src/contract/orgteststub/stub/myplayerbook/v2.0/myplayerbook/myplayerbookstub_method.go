package myplayerbookstub

import (
	"blockchain/algorithm"
	"blockchain/smcsdk/sdk/rlp"
	types2 "blockchain/types"
	"contract/orgexample/code/mydice2win/v1.0/mydice2win"
	tmcommon "github.com/tendermint/tmlibs/common"
	"strconv"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	types1 "contract/stubcommon/types"

	"contract/orgteststub/code/myplayerbook/v2.0/myplayerbook"
	"github.com/tendermint/tmlibs/log"
)

//MyPlayerBookStub an object
type MyPlayerBookStub struct {
	logger log.Logger
}

var _ types1.IContractStub = (*MyPlayerBookStub)(nil)

//New generate a stub
func New(logger log.Logger) types1.IContractStub {
	return &MyPlayerBookStub{logger: logger}
}

//FuncRecover recover panic by Assert
func FuncRecover(response *types2.Response) {
	if rerr := recover(); rerr != nil {
		if _, ok := rerr.(types.Error); ok {
			response.Code = rerr.(types.Error).ErrorCode
			response.Log = rerr.(types.Error).ErrorDesc
		} else {
			panic(rerr)
		}
	}
}

// InitChain initial smart contract
func (mc *MyPlayerBookStub) InitChain(smc sdk.ISmartContract) (response types2.Response) {

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types2.CodeOK
	return response
}

// UpdateChain update smart contract
func (mc *MyPlayerBookStub) UpdateChain(smc sdk.ISmartContract) (response types2.Response) {
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.UpdateChain()

	response.Code = types2.CodeOK
	return response
}

//Invoke invoke function
func (pbs *MyPlayerBookStub) Invoke(smc sdk.ISmartContract) (response types2.Response) {
	defer FuncRecover(&response)

	// 扣手续费并生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(smc, true)
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(smc.Message(), nil, "", fee, gasUsed, smc.Tx().GasLimit(), types.Error{})
		return
	}
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})

	var data string
	switch smc.Message().MethodID() {
	case "23445656": // prototype: GetPlayer(types.Address)*Player
		data = getPlayer(smc)

	case "e463fdb2": // prototype: RegisterName(string)(types.Error)
		registerName(smc)
	}
	response = common.CreateResponse(smc.Message(), nil, data, fee, gasUsed, smc.Tx().GasLimit(), types.Error{})
	return
}

//nolint unhandled
func getPlayer(smc sdk.ISmartContract) (addr string) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")

	var param types2.Address
	rlp.DecodeBytes(items[0], param)
	sdk.Require(len(param) > 0, types.ErrStubDefined, "Invalid address")

	plyrbk := new(myplayerbook.MyPlayerBook)
	plyrbk.SetSdk(smc)
	addr = plyrbk.GetPlayer(param)
	return
}

// nolint unhandled
func registerName(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 2, types.ErrStubDefined, "Invalid message data")

	var byteInt []byte
	rlp.DecodeBytes(items[0], &byteInt)
	sdk.Require(len(byteInt) > 0, types.ErrStubDefined, "Invalid index")
	index := algorithm.BytesToInt64(byteInt)

	var plyrbyte []byte
	rlp.DecodeBytes(items[1], &plyrbyte)
	sdk.Require(len(plyrbyte) > 0, types.ErrStubDefined, "Invalid parameter")
	var player myplayerbook.Player
	err := jsoniter.Unmarshal(plyrbyte, &player)
	sdk.RequireNotError(err, types.ErrStubDefined)

	plyrbk := new(myplayerbook.MyPlayerBook)
	plyrbk.SetSdk(smc)
	plyrbk.RegisterName(index, player)
}

//todo 讨论 数据类型的RLP编码
// int, float, byte, bool， map 类型不能直接进行rlp编码，需进行转换
//建议将 int, float, byte, bool 类型统一转换为string, map 类型进行marshal转换为[]byte
//nolint unhandled
func multiTypeParam(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 8, types.ErrStubDefined, "Invalid message data")

	//todo : 此处的int64先按照转换为[]byte进行解码
	//the first parameter - index
	var bytes []byte
	rlp.DecodeBytes(items[0], &bytes)
	sdk.Require(len(bytes) > 0, types.ErrStubDefined, "Invalid index")
	index := algorithm.BytesToUint64(bytes)

	//todo float按照转换为string进行解码
	//the second parameter - flt
	var temp string
	rlp.DecodeBytes(items[1], &temp)
	sdk.Require(len(temp) > 0, types.ErrStubDefined, "Invalid parameter")
	flt, err := strconv.ParseFloat(temp, 64)
	sdk.Require(err == nil, types.ErrStubDefined, "Invalid parameter")
	//the third parameter - bl
	rlp.DecodeBytes(items[2], &temp)
	sdk.Require(len(temp) > 0, types.ErrStubDefined, "Invalid parameter")
	bl, err := strconv.ParseBool(temp)
	sdk.Require(err == nil, types.ErrStubDefined, "Invalid parameter")
	//the fourth parameter - bt (byte)
	rlp.DecodeBytes(items[3], &temp)
	sdk.Require(len(temp) == 1, types.ErrStubDefined, "Invalid parameter")
	bt := byte(temp[0])
	//the fifth parameter - hash
	rlp.DecodeBytes(items[4], &bytes)
	sdk.Require(len(bytes) > 0, types.ErrStubDefined, "Invalid parameter")
	hash := bytes
	//the sixth parameter - hb
	rlp.DecodeBytes(items[5], &bytes)
	sdk.Require(len(bytes) > 0, types.ErrStubDefined, "Invalid parameter")
	hexhash := bytes
	//the seventh parameter - hash
	rlp.DecodeBytes(items[6], &bytes)
	sdk.Require(len(bytes) > 0, types.ErrStubDefined, "Invalid parameter")
	bi := bn.NBytes(bytes)
	//the eighth parameter - hash
	rlp.DecodeBytes(items[7], &bytes)
	sdk.Require(len(bytes) > 0, types.ErrStubDefined, "Invalid parameter")
	mp := make(map[int]string)
	err = jsoniter.Unmarshal(bytes, &mp)
	sdk.Require(err == nil, types.ErrStubDefined, "Invalid parameter")

	plyrbk := new(myplayerbook.MyPlayerBook)
	plyrbk.SetSdk(smc)
	plyrbk.MultiTypesParam(index, flt, bl, bt, hash, hexhash, bi, mp)
}
