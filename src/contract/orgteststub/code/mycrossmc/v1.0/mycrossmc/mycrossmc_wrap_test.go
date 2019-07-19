package mycrossmc

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/utest"
)

var (
	contractName    = "mycrossmc" //contract name
	contractMethods = []string{"Register(uint64)", "Set(uint64)", "Get()TYPE"}
	orgID           = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJH"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *MyCrossmc
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
	return &TestObject{&MyCrossmc{sdk: utest.UTP.ISmartContract}}
}

//transfer This is a method of TestObject
func (t *TestObject) transfer(balance bn.Number) *TestObject {
	t.obj.sdk.Message().Sender().TransferByName("LOC", t.obj.sdk.Message().Contract().Account().Address(), balance)
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

//Register This is a method of TestObject
func (t *TestObject) Register(data uint64) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.Register(data)
	utest.Commit()
	return
}

//Set This is a method of TestObject
func (t *TestObject) Set(data uint64) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.Set(data)
	utest.Commit()
	return
}

//Get This is a method of TestObject
func (t *TestObject) Get() (result0 uint64, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.Get()
	utest.Commit()
	return
}

//Get This is a method of TestObject
func (t *TestObject) MultiParam() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.MultParam()
	utest.Commit()
	return
}
