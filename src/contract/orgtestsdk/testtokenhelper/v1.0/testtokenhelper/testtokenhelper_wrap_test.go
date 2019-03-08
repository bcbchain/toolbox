package testtokenhelper

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
	contractName       = "testtokenhelper" //contract name
	contractMethods    = []string{"TestRegisterToken()", "TestToken()", "TestTokenOfAddress()", "TestTokenOfName()", "TestTokenOfSymbol()", "TestTokenOfContract()", "TestBaseGasPrice()", "Transfer(types.Address,bn.Number)"}
	contractInterfaces = []string{"Transfer(types.Address,bn.Number)"}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzf"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *TestTokenHelper
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
	return &TestObject{&TestTokenHelper{sdk: utest.UTP.ISmartContract}}
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

//TestRegisterTokenName This is a method of TestObject
func (t *TestObject) TestRegisterTokenName() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRegisterTokenName()
	utest.Commit()
	return
}

//TestRegisterTokenSymbol This is a method of TestObject
func (t *TestObject) TestRegisterTokenSymbol() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRegisterTokenSymbol()
	utest.Commit()
	return
}

//TestRegisterTokenTotalSupply This is a method of TestObject
func (t *TestObject) TestRegisterTokenTotalSupply() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRegisterTokenTotalSupply()
	utest.Commit()
	return
}

//TestRegisterTokenAddSupplyEnabled This is a method of TestObject
func (t *TestObject) TestRegisterTokenAddSupplyEnabled() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRegisterTokenAddSupplyEnabled()
	utest.Commit()
	return
}

//TestRegisterTokenBurnEnabled This is a method of TestObject
func (t *TestObject) TestRegisterTokenBurnEnabled() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRegisterTokenBurnEnabled()
	utest.Commit()
	return
}

//TestRegisterTokenDuplicate This is a method of TestObject
func (t *TestObject) TestRegisterTokenDuplicate() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRegisterTokenDuplicate()
	utest.Commit()
	return
}

//TestToken This is a method of TestObject
func (t *TestObject) TestToken() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestToken()
	utest.Commit()
	return
}

//TestTokenOfAddress This is a method of TestObject
func (t *TestObject) TestTokenOfAddress() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTokenOfAddress()
	utest.Commit()
	return
}

//TestTokenOfName This is a method of TestObject
func (t *TestObject) TestTokenOfName() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTokenOfName()
	utest.Commit()
	return
}

//TestTokenOfSymbol This is a method of TestObject
func (t *TestObject) TestTokenOfSymbol() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTokenOfSymbol()
	utest.Commit()
	return
}

//TestTokenOfContract This is a method of TestObject
func (t *TestObject) TestTokenOfContract() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTokenOfContract()
	utest.Commit()
	return
}

//TestBaseGasPrice This is a method of TestObject
func (t *TestObject) TestBaseGasPrice() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestBaseGasPrice()
	utest.Commit()
	return
}
