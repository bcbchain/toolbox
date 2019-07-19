package mybasictype

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	"blockchain/smcsdk/utest"
)

var (
	contractName    = "mybasictype" //contract name
	contractMethods = []string{"EchoAddress(types.Address)TYPE", "EchoHash(types.Hash)TYPE", "EchoHexBytes(types.HexBytes)TYPE", "EchoPubKey(types.PubKey)TYPE", "EchoNumber(bn.Number)TYPE", "EchoInt(int)TYPE", "EchoInt8(int8)TYPE", "EchoInt16(int16)TYPE", "EchoInt32(int32)TYPE", "EchoInt64(int64)TYPE", "EchoUint(uint)TYPE", "EchoUint8(uint8)TYPE", "EchoUint16(uint16)TYPE", "EchoUint32(uint32)TYPE", "EchoUint64(uint64)TYPE", "EchoBool(bool)TYPE", "EchoByte(byte)TYPE", "EchoBytes([]byte)ARRAY_OR_SLICE_TYPE", "EchoMap(map[string]int)ARRAY_OR_SLICE_TYPE", "EchoMap1(map[uint]map[string]types.Address)ARRAY_OR_SLICE_TYPE", "EchoMap2(map[int]map[int8]types.Hash)ARRAY_OR_SLICE_TYPE", "EchoMap3(map[int]map[int8]types.Hash)ARRAY_OR_SLICE_TYPE", "EchoMap4(map[int]map[int8]types.Hash)ARRAY_OR_SLICE_TYPE", "EchoMap5(map[int]map[int8]types.Hash)ARRAY_OR_SLICE_TYPE", "EchoMap6(map[int]map[int8]types.Hash)ARRAY_OR_SLICE_TYPE", "EchoMap7(map[int]map[int8]types.Hash)ARRAY_OR_SLICE_TYPE", "EchoMap8(map[int]map[int8]types.Hash)ARRAY_OR_SLICE_TYPE"}
	orgID           = "org5xSEfgNDAaPU6f5fxai5BaaUW3FL1gsTV"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *BasicType
}

//FuncRecover recover panic by Assert
func FuncRecover(err *types.Error) {
	if rerr := recover(); rerr != nil {
		if _, ok := rerr.(types.Error); ok {
			err.ErrorCode = rerr.(types.Error).ErrorCode
			err.ErrorDesc = rerr.(types.Error).ErrorDesc
		} else {
			panic(rerr)
		}
	}
}

//NewTestObject This is a function
func NewTestObject(sender sdk.IAccount) *TestObject {
	return &TestObject{&BasicType{sdk: utest.UTP.ISmartContract}}
}

//transfer This is a method of TestObject
func (t *TestObject) transfer(balance bn.Number) *TestObject {
	t.obj.sdk.Message().Sender().TransferByName(t.obj.sdk.Helper().GenesisHelper().Token().Name(), t.obj.sdk.Message().Contract().Account().Address(), balance)
	t.obj.sdk = sdkhelper.OriginNewMessage(t.obj.sdk, t.obj.sdk.Message().Contract(), t.obj.sdk.Message().MethodID(), t.obj.sdk.Message().(*object.Message).OutputReceipts())
	return t
}

//setSender This is a method of TestObject
func (t *TestObject) setSender(sender sdk.IAccount) *TestObject {
	t.obj.sdk = utest.SetSender(sender.Address())
	return t
}

//run This is a method of TestObject
func (t *TestObject) run() *TestObject {
	t.obj.sdk = utest.ResetMsg()
	return t
}

//InitChain This is a method of TestObject
func (t *TestObject) InitChain() {
	utest.NextBlock(1)
	t.obj.InitChain()
	utest.Commit()
	return
}

//EchoAddress This is a method of TestObject
func (t *TestObject) EchoAddress(v types.Address) (result0 types.Address, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoAddress(v)
	utest.Commit()
	return
}

//EchoHash This is a method of TestObject
func (t *TestObject) EchoHash(v types.Hash) (result0 types.Hash, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoHash(v)
	utest.Commit()
	return
}

//EchoHexBytes This is a method of TestObject
func (t *TestObject) EchoHexBytes(v types.HexBytes) (result0 types.HexBytes, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoHexBytes(v)
	utest.Commit()
	return
}

//EchoPubKey This is a method of TestObject
func (t *TestObject) EchoPubKey(v types.PubKey) (result0 types.PubKey, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoPubKey(v)
	utest.Commit()
	return
}

//EchoNumber This is a method of TestObject
func (t *TestObject) EchoNumber(v bn.Number) (result0 bn.Number, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoNumber(v)
	utest.Commit()
	return
}

//EchoInt This is a method of TestObject
func (t *TestObject) EchoInt(v int) (result0 int, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoInt(v)
	utest.Commit()
	return
}

//EchoInt8 This is a method of TestObject
func (t *TestObject) EchoInt8(v int8) (result0 int8, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoInt8(v)
	utest.Commit()
	return
}

//EchoInt16 This is a method of TestObject
func (t *TestObject) EchoInt16(v int16) (result0 int16, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoInt16(v)
	utest.Commit()
	return
}

//EchoInt32 This is a method of TestObject
func (t *TestObject) EchoInt32(v int32) (result0 int32, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoInt32(v)
	utest.Commit()
	return
}

//EchoInt64 This is a method of TestObject
func (t *TestObject) EchoInt64(v int64) (result0 int64, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoInt64(v)
	utest.Commit()
	return
}

//EchoUint This is a method of TestObject
func (t *TestObject) EchoUint(v uint) (result0 uint, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoUint(v)
	utest.Commit()
	return
}

//EchoUint8 This is a method of TestObject
func (t *TestObject) EchoUint8(v uint8) (result0 uint8, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoUint8(v)
	utest.Commit()
	return
}

//EchoUint16 This is a method of TestObject
func (t *TestObject) EchoUint16(v uint16) (result0 uint16, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoUint16(v)
	utest.Commit()
	return
}

//EchoUint32 This is a method of TestObject
func (t *TestObject) EchoUint32(v uint32) (result0 uint32, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoUint32(v)
	utest.Commit()
	return
}

//EchoUint64 This is a method of TestObject
func (t *TestObject) EchoUint64(v uint64) (result0 uint64, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoUint64(v)
	utest.Commit()
	return
}

//EchoBool This is a method of TestObject
func (t *TestObject) EchoBool(v bool) (result0 bool, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoBool(v)
	utest.Commit()
	return
}

//EchoByte This is a method of TestObject
func (t *TestObject) EchoByte(v byte) (result0 byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoByte(v)
	utest.Commit()
	return
}

//EchoBytes This is a method of TestObject
func (t *TestObject) EchoBytes(v []byte) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoBytes(v)
	utest.Commit()
	return
}

//EchoMap This is a method of TestObject
func (t *TestObject) EchoMap(v map[string]int) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap(v)
	utest.Commit()
	return
}

//EchoMap1 This is a method of TestObject
func (t *TestObject) EchoMap1(v map[uint]map[string]types.Address) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap1(v)
	utest.Commit()
	return
}

//EchoMap2 This is a method of TestObject
func (t *TestObject) EchoMap2(v map[int]map[int8]types.Hash) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap2(v)
	utest.Commit()
	return
}

//EchoMap3 This is a method of TestObject
func (t *TestObject) EchoMap3(v map[int]map[uint64]types.HexBytes) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap3(v)
	utest.Commit()
	return
}

//EchoMap4 This is a method of TestObject
func (t *TestObject) EchoMap4(v map[bool]map[byte]types.PubKey) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap4(v)
	utest.Commit()
	return
}

//EchoMap5 This is a method of TestObject
func (t *TestObject) EchoMap5(v map[bool]map[bn.Number]bool) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap5(v)
	utest.Commit()
	return
}

//EchoMap6 This is a method of TestObject
func (t *TestObject) EchoMap6(v map[byte]map[string]bn.Number) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap6(v)
	utest.Commit()
	return
}

//EchoMap7 This is a method of TestObject
func (t *TestObject) EchoMap7(v map[string]map[types.Address]byte) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap7(v)
	utest.Commit()
	return
}

//EchoMap8 This is a method of TestObject
func (t *TestObject) EchoMap8(v map[types.Address]map[bn.Number]string) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.EchoMap8(v)
	utest.Commit()
	return
}
