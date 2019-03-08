package testaccount

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
	contractName    = "testaccount" //contract name
	contractMethods = []string{"TestTransfer()", "TestTransferByToken()", "TestTransferByName()", "TestTransferBySymbol()", "Transfer(types.Address,bn.Number)",
		"TestAccountPubKey()", "TestBalance()", "TestBalanceOfToken()", "TestBalanceOfName()", "TestBalanceOfSymbol()"}
	contractInterfaces = []string{"Transfer(types.Address,bn.Number)"}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzf"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *TestAccount
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
	return &TestObject{&TestAccount{sdk: utest.UTP.ISmartContract}}
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

//TestTransfer This is a method of TestObject
func (t *TestObject) TestTransfer() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTransfer()
	utest.Commit()
	return
}

//TestTransferByToken This is a method of TestObject
func (t *TestObject) TestTransferByToken() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTransferByToken()
	utest.Commit()
	return
}

//TestTransferByName This is a method of TestObject
func (t *TestObject) TestTransferByName() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTransferByName()
	utest.Commit()
	return
}

//TestTransferBySymbol This is a method of TestObject
func (t *TestObject) TestTransferBySymbol() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTransferBySymbol()
	utest.Commit()
	return
}

//Transfer This is a method of TestObject
func (t *TestObject) Transfer(to types.Address, value bn.Number) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.Transfer(to, value)
	utest.Commit()
	return
}

//TestAccountPubKey This is a method of TestObject
func (t *TestObject) TestAccountPubKey() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestAccountPubKey()
	utest.Commit()
	return
}

//TestBalance This is a method of TestObject
func (t *TestObject) TestBalance() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestBalance()
	utest.Commit()
	return
}

//TestBalanceOfName This is a method of TestObject
func (t *TestObject) TestBalanceOfName() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestBalanceOfName()
	utest.Commit()
	return
}

//TestBalanceOfSymbol This is a method of TestObject
func (t *TestObject) TestBalanceOfSymbol() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestBalanceOfSymbol()
	utest.Commit()
	return
}

//TestBalanceOfToken This is a method of TestObject
func (t *TestObject) TestBalanceOfToken() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestBalanceOfToken()
	utest.Commit()
	return
}
