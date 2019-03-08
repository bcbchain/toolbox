package testblockchainhelper

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
	contractName       = "testblockchainhelper" //contract name
	contractMethods    = []string{"TestCalcAccountFromPubKey()", "TestCalcAccountFromName()", "TestCalcContractAddress()", "TestCalcOrgID()", "TestCheckAddress()", "TestGetBlock()"}
	contractInterfaces = []string{}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *TestblockChainHelper
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
	return &TestObject{&TestblockChainHelper{sdk: utest.UTP.ISmartContract}}
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

//TestCalcAccountFromPubKey This is a method of TestObject
func (t *TestObject) TestCalcAccountFromPubKey() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestCalcAccountFromPubKey()
	utest.Commit()
	return
}

//TestCalcAccountFromName This is a method of TestObject
func (t *TestObject) TestCalcAccountFromName() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestCalcAccountFromName()
	utest.Commit()
	return
}

//TestCalcContractAddress This is a method of TestObject
func (t *TestObject) TestCalcContractAddress() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestCalcContractAddress()
	utest.Commit()
	return
}

//TestCalcOrgID This is a method of TestObject
func (t *TestObject) TestCalcOrgID() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestCalcOrgID()
	utest.Commit()
	return
}

//TestCheckAddress This is a method of TestObject
func (t *TestObject) TestCheckAddress() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestCheckAddress()
	utest.Commit()
	return
}

//TestGetBlock This is a method of TestObject
func (t *TestObject) TestGetBlock() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestGetBlock()
	utest.Commit()
	return
}
