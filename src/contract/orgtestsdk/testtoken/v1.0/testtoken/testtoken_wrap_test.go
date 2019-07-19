package testtoken

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
	contractName       = "testtoken" //contract name
	contractMethods    = []string{"Transfer(types.Address,bn.Number)"}
	contractInterfaces = []string{"Transfer(types.Address,bn.Number)"}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzf"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *TestToken
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
	return &TestObject{&TestToken{sdk: utest.UTP.ISmartContract}}
}

//transfer This is a method of TestObject
func (t *TestObject) transfer(balance bn.Number) *TestObject {
	contract := t.obj.sdk.Message().Contract()
	utest.Transfer(t.obj.sdk.Message().Sender(), t.obj.sdk.Helper().GenesisHelper().Token().Name(), contract.Account().Address(), balance)
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

//TestSetOwnerForCurrentToken This is a method of TestObject
func (t *TestObject) TestSetTotalSupply(s int) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestSetTotalSupply(s)
	utest.Commit()
	return
}

//TestSetOwnerForCurrentToken This is a method of TestObject
func (t *TestObject) TestSetGasPrice() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestSetGasPrice()
	utest.Commit()
	return
}

//TestAddress This is a method of TestObject
func (t *TestObject) TestAddress(address types.Address) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestAddress(address)
	utest.Commit()
	return
}

//TestAddress This is a method of TestObject
func (t *TestObject) TestOwner(owner types.Address) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestOwner(owner)
	utest.Commit()
	return
}

//TestName This is a method of TestObject
func (t *TestObject) TestName(name string) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestName(name)
	utest.Commit()
	return
}

//TestName This is a method of TestObject
func (t *TestObject) TestSymbol(symbol string) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestSymbol(symbol)
	utest.Commit()
	return
}

//TestTotalSupply This is a method of TestObject
func (t *TestObject) TestTotalSupply(total bn.Number) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	utest.Commit()
	t.obj.TestTotalSupply(total)
	return
}

//TestAddSupplyEnabled This is a method of TestObject
func (t *TestObject) TestAddSupplyEnabled(enabled bool) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	utest.Commit()
	t.obj.TestAddSupplyEnabled(enabled)
	return
}

//TestAddSupplyEnabled This is a method of TestObject
func (t *TestObject) TestBurnEnabled(enabled bool) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	utest.Commit()
	t.obj.TestBurnEnabled(enabled)
	return
}

//TestGasPrice This is a method of TestObject
func (t *TestObject) TestGasPrice(gasPrice int64) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	utest.Commit()
	t.obj.TestGasPrice(gasPrice)
	return
}
