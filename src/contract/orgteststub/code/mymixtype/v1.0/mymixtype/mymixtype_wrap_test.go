package mymixtype

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	"blockchain/smcsdk/utest"
)

var (
	contractName    = "mymixtype" //contract name
	contractMethods = []string{"Basic(*BasicTypes)ARRAY_OR_SLICE_TYPE", "Slice(*SliceTypes)ARRAY_OR_SLICE_TYPE", "MapString(*MapStringTypes)ARRAY_OR_SLICE_TYPE", "MapOther(*MapOtherTypes)ARRAY_OR_SLICE_TYPE", "MapSlice(*MapSliceTypes)ARRAY_OR_SLICE_TYPE", "Complex(*ComplexDefine)ARRAY_OR_SLICE_TYPE", "Long(*LongType)ARRAY_OR_SLICE_TYPE"}
	orgID           = "org5xSEfgNDAaPU6f5fxai5BaaUW3FL1gsTV"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *Mymixtype
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
	return &TestObject{&Mymixtype{sdk: utest.UTP.ISmartContract}}
}

//transfer This is a method of TestObject
func (t *TestObject) transfer(balance bn.Number) *TestObject {
	t.obj.sdk.Message().Sender().TransferByName(t.obj.sdk.Helper().GenesisHelper().Token().Name(), t.obj.sdk.Message().Contract().Account(), balance)
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

//Basic This is a method of TestObject
func (t *TestObject) Basic(basic_ *BasicTypes) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.Basic(basic_)
	utest.Commit()
	return
}

//Slice This is a method of TestObject
func (t *TestObject) Slice(slice_ *SliceTypes) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.Slice(slice_)
	utest.Commit()
	return
}

//MapString This is a method of TestObject
func (t *TestObject) MapString(map_ *MapStringTypes) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.MapString(map_)
	utest.Commit()
	return
}

//MapOther This is a method of TestObject
func (t *TestObject) MapOther(map_ *MapOtherTypes) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.MapOther(map_)
	utest.Commit()
	return
}

//MapSlice This is a method of TestObject
func (t *TestObject) MapSlice(map_ *MapSliceTypes) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.MapSlice(map_)
	utest.Commit()
	return
}

//Complex This is a method of TestObject
func (t *TestObject) Complex(complex_ *ComplexDefine) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.Complex(complex_)
	utest.Commit()
	return
}

//Long This is a method of TestObject
func (t *TestObject) Long(long_ *LongType) (result0 []byte, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.Long(long_)
	utest.Commit()
	return
}
