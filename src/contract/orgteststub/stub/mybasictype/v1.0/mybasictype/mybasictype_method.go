package mybasictype

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	sdkTypes "blockchain/smcsdk/sdk/types"
	"blockchain/types"
	"contract/orgexample/code/mydice2win/v1.0/mydice2win"
	"contract/stubcommon/common"
	stubTypes "contract/stubcommon/types"
	"fmt"

	"contract/orgteststub/code/mybasictype/v1.0/mybasictype"
	"github.com/tendermint/tmlibs/log"
)

const msg = "Message can not be nil."

// MyBasicTypeStub stub
type MyBasicTypeStub struct {
	logger log.Logger
}

var _ stubTypes.IContractStub = (*MyBasicTypeStub)(nil)

// New new stub
func New(logger log.Logger) stubTypes.IContractStub {

	var stub MyBasicTypeStub
	stub.logger = logger

	return &stub
}

//FuncRecover recover panic by Assert
func FuncRecover(response *types.Response) {
	if rerr := recover(); rerr != nil {
		if _, ok := rerr.(sdkTypes.Error); ok {
			response.Code = rerr.(sdkTypes.Error).ErrorCode
			response.Log = rerr.(sdkTypes.Error).ErrorDesc
		} else {
			panic(rerr)
		}
	}
}

// InitChain initial smart contract
func (mc *MyBasicTypeStub) InitChain(smc sdk.ISmartContract) (response types.Response) {

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (mc *MyBasicTypeStub) UpdateChain(smc sdk.ISmartContract) (response types.Response) {
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.UpdateChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (mc *MyBasicTypeStub) Mine(smc sdk.ISmartContract) (response types.Response) {
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.UpdateChain()

	response.Code = types.CodeOK
	return response
}

// Invoke invoke method
func (mc *MyBasicTypeStub) Invoke(smc sdk.ISmartContract) (response types.Response) {
	return mc.InvokeInternal(smc, true)
}

// Invoke invoke method
func (mc *MyBasicTypeStub) InvokeInternal(smc sdk.ISmartContract, feeFlag bool) (response types.Response) {
	defer FuncRecover(&response)

	// 生成手续费收据
	fee, gasUsed, _, err := common.FeeAndReceipt(smc, feeFlag)
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(smc.Message(), nil, "", fee, gasUsed, smc.Tx().GasLimit(), sdkTypes.Error{})
		return
	}

	var data string
	switch smc.Message().MethodID() {
	case "1":
		data = echoAddress(smc)
	case "2":
		data = echoHash(smc)
	case "3":
		data = echoHexBytes(smc)
	case "4":
		data = echoPubKey(smc)
	case "5":
		data = echoNumber(smc)
	case "6":
		data = echoInt(smc)
	case "7":
		data = echoInt8(smc)
	case "8":
		data = echoInt16(smc)
	case "9":
		data = echoInt32(smc)
	case "10":
		data = echoInt64(smc)
	case "11":
		data = echoUint(smc)
	case "12":
		data = echoUint8(smc)
	case "13":
		data = echoUint16(smc)
	case "14":
		data = echoUint32(smc)
	case "15":
		data = echoUint64(smc)
	case "18":
		data = echoBool(smc)
	case "19":
		data = echoBool(smc)
	case "20":
		data = echoByte(smc)
	case "21":
		data = echoSlice(smc)
	case "22":
		data = echoMap(smc)
	case "23":
		data = echoMap1(smc)
	case "24":
		data = echoMap2(smc)
	case "25":
		data = echoMap3(smc)
	case "26":
		data = echoMap4(smc)
	case "27":
		data = echoMap5(smc)
	case "28":
		data = echoMap6(smc)
	case "29":
		data = echoMap7(smc)
	case "30":
		data = echoMap8(smc)
	}
	response = common.CreateResponse(smc.Message(), nil, data, fee, gasUsed, smc.Tx().GasLimit(), sdkTypes.Error{})
	return
}

func echoAddress(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v types.Address
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoAddress(v)

	return data
}

func echoHash(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v sdkTypes.Hash
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoHash(v)

	return string(data)
}

func echoHexBytes(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v sdkTypes.HexBytes
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoHexBytes(v)

	return string(data)
}

func echoPubKey(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v sdkTypes.PubKey
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoPubKey(v)

	return string(data)
}

func echoNumber(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v []byte
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoNumber(new(bn.Number).SetBytes(v))

	return data.String()
}

func echoInt(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v int
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoInt(v)

	return fmt.Sprintf("%d", data)
}

func echoInt8(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v int8
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoInt8(v)

	return fmt.Sprintf("%d", data)
}

func echoInt16(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v int16
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoInt16(v)

	return fmt.Sprintf("%d", data)
}

func echoInt32(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v int32
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoInt32(v)

	return fmt.Sprintf("%d", data)
}

func echoInt64(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v int64
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoInt64(v)

	return fmt.Sprintf("%d", data)
}

func echoUint(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v uint
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoUint(v)

	return fmt.Sprintf("%d", data)
}

func echoUint8(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v uint8
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoUint8(v)

	return fmt.Sprintf("%d", data)
}

func echoUint16(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v uint16
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoUint16(v)

	return fmt.Sprintf("%d", data)
}

func echoUint32(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v uint32
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoUint32(v)

	return fmt.Sprintf("%d", data)
}

func echoUint64(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v uint64
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoUint64(v)

	return fmt.Sprintf("%d", data)
}

func echoBool(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v bool
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoBool(v)

	return fmt.Sprintf("%v", data)
}

func echoByte(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v byte
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoByte(v)

	return string(data)
}

func echoSlice(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v []byte
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoBytes(v)

	return string(data)
}

func echoMap(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	//v map[string]int)
	var v map[string]int
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap(v)

	return string(data)
}

func echoMap1(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[uint]map[string]sdkTypes.Address
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap1(v)

	return string(data)
}

func echoMap2(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[int]map[int8]sdkTypes.Hash
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap2(v)

	return string(data)
}

func echoMap3(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[int]map[uint64]sdkTypes.HexBytes
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap3(v)

	return string(data)
}

func echoMap4(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[bool]map[byte]sdkTypes.PubKey
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap4(v)

	return string(data)
}

func echoMap5(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[bool]map[bn.Number]bool
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap5(v)

	return string(data)
}

func echoMap6(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[byte]map[string]bn.Number
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap6(v)

	return string(data)
}

func echoMap7(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[string]map[types.Address]byte
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap7(v)

	return string(data)
}

func echoMap8(smc sdk.ISmartContract) string {
	itemsBytes := smc.Message().Items()
	sdk.Require(len(itemsBytes) == 1, sdkTypes.ErrStubDefined, msg)

	var v map[types.Address]map[bn.Number]string
	errDecode := rlp.DecodeBytes(itemsBytes[0], &v)
	sdk.RequireNotError(errDecode, sdkTypes.ErrStubDefined)

	myBasicType := new(mybasictype.BasicType)
	myBasicType.SetSdk(smc)

	data := myBasicType.EchoMap8(v)

	return string(data)
}
