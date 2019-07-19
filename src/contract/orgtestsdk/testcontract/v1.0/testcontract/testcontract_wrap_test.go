package testcontract

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
	contractName = "testcontract" //contract name
	//contractMethods    = []string{"Transfer(types.Address,bn.Number)", "TestAddress()", "TestAccount()", "TestOwner()", "TestName()", "TestVersion()", "TestCodeHash()", "TestEffectHeight()", "TestLoseHeight()", "TestKeyPrefix()", "TestMethods()", "TestInterfaces()", "TestToken()", "TestOrgID()", "TestSetOwner()"}
	//contractInterfaces = []string{"Transfer(types.Address,bn.Number)", "TestAddress()", "TestName()", "TestLoseHeight()", "TestMethods()", "TestInterfaces()", "TestOrgID()", "TestSetOwner()"}
	//orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxew"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *TestContract
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
	return &TestObject{&TestContract{sdk: utest.UTP.ISmartContract}}
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

//TestAddress This is a method of TestObject
func (t *TestObject) TestAddress() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestAddress()
	utest.Commit()
	return
}

//TestAccount This is a method of TestObject
func (t *TestObject) TestAccount() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestAccount()
	utest.Commit()
	return
}

//TestOwner This is a method of TestObject
func (t *TestObject) TestOwner() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestOwner()
	utest.Commit()
	return
}

//TestName This is a method of TestObject
func (t *TestObject) TestName() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestName()
	utest.Commit()
	return
}

//TestVersion This is a method of TestObject
func (t *TestObject) TestVersion() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestVersion()
	utest.Commit()
	return
}

//TestCodeHash This is a method of TestObject
func (t *TestObject) TestCodeHash() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestCodeHash()
	utest.Commit()
	return
}

//TestEffectHeight This is a method of TestObject
func (t *TestObject) TestEffectHeight() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestEffectHeight()
	utest.Commit()
	return
}

//TestLoseHeight This is a method of TestObject
func (t *TestObject) TestLoseHeight() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestLoseHeight()
	utest.Commit()
	return
}

//TestKeyPrefix This is a method of TestObject
func (t *TestObject) TestKeyPrefix() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestKeyPrefix()
	utest.Commit()
	return
}

//TestMethods This is a method of TestObject
func (t *TestObject) TestMethods() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestMethods()
	utest.Commit()
	return
}

//TestInterfaces This is a method of TestObject
func (t *TestObject) TestInterfaces() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestInterfaces()
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

//TestOrgID This is a method of TestObject
func (t *TestObject) TestOrgID() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestOrgID()
	utest.Commit()
	return
}

//TestSetOwner This is a method of TestObject
func (t *TestObject) TestSetOwner() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestSetOwner()
	utest.Commit()
	return
}
