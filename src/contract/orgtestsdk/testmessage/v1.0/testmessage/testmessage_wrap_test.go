package testmessage

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	"blockchain/smcsdk/utest"
	"fmt"
)

var (
	contractName       = "testmessage" //contract name
	contractMethods    = []string{"Transfer(types.Address,bn.Number)", "TestContract()", "TestMethodID()", "TestData()", "TestGasPrice()", "TestSender()", "TestOrigins()", "TestInputReceipts()", "TestGetTransferToMe()"}
	contractInterfaces = []string{"Transfer(types.Address,bn.Number)"}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *TestMessage
}

//FuncRecover recover panic by Assert
func FuncRecover(err *types.Error) {
	if rerr := recover(); rerr != nil {
		if _, ok := rerr.(types.Error); ok {
			err.ErrorCode = rerr.(types.Error).ErrorCode
			err.ErrorDesc = rerr.(types.Error).ErrorDesc
			fmt.Println(err)
		} else {
			panic(rerr)
		}
	}
}

//NewTestObject This is a function
func NewTestObject(sender sdk.IAccount) *TestObject {
	return &TestObject{&TestMessage{sdk: utest.UTP.ISmartContract}}
}

//transfer This is a method of TestObject
func (t *TestObject) transfer(balance bn.Number) *TestObject {
	contract := t.obj.sdk.Message().Contract()
	utest.Transfer(t.obj.sdk.Message().Sender(), t.obj.sdk.Helper().GenesisHelper().Token().Name(), contract.Account(), balance)
	t.obj.sdk = sdkhelper.OriginNewMessage(t.obj.sdk, contract, t.obj.sdk.Message().MethodID(), t.obj.sdk.Message().(*object.Message).OutputReceipts())
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

//TestContract This is a method of TestObject
func (t *TestObject) TestContract() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestContract()
	utest.Commit()
	return
}

//TestMethodID This is a method of TestObject
func (t *TestObject) TestMethodID() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestMethodID()
	utest.Commit()
	return
}

//TestData This is a method of TestObject
func (t *TestObject) TestItems() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestItems()
	utest.Commit()
	return
}

//TestGasPrice This is a method of TestObject
func (t *TestObject) TestGasPrice() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestGasPrice()
	utest.Commit()
	return
}

//TestSender This is a method of TestObject
func (t *TestObject) TestSender() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestSender()
	utest.Commit()
	return
}

//TestOrigins This is a method of TestObject
func (t *TestObject) TestOrigins() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestOrigins()
	utest.Commit()
	return
}

//TestInputReceipts This is a method of TestObject
func (t *TestObject) TestInputReceipts() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestInputReceipts()
	utest.Commit()
	return
}

//TestGetTransferToMe This is a method of TestObject
func (t *TestObject) TestGetTransferToMe() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestGetTransferToMe()
	utest.Commit()
	return
}
