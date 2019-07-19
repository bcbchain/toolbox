package testblock

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
	contractName       = "testblock" //contract name
	contractMethods    = []string{"TestChainID()", "TestBlockHash()", "TestHeight()", "TestTime()", "TestNow()", "TestNumTxs()", "TestDataHash()", "TestProposerAddress()", "TestRewardAddress()", "TestRandomNumber()", "TestVersion()", "TestLastBlockHash()", "TestCommitHash()", "TestAppHash()", "TestFee()"}
	contractInterfaces = []string{}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *TestBlock
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
	return &TestObject{&TestBlock{sdk: utest.UTP.ISmartContract}}
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

//TestChainID This is a method of TestObject
func (t *TestObject) TestChainID() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestChainID()
	utest.Commit()
	return
}

//TestBlockHash This is a method of TestObject
func (t *TestObject) TestBlockHash() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestBlockHash()
	utest.Commit()
	return
}

//TestHeight This is a method of TestObject
func (t *TestObject) TestHeight() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestHeight()
	utest.Commit()
	return
}

//TestTime This is a method of TestObject
func (t *TestObject) TestTime() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestTime()
	utest.Commit()
	return
}

//TestNow This is a method of TestObject
func (t *TestObject) TestNow() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestNow()
	utest.Commit()
	return
}

//TestNumTxs This is a method of TestObject
func (t *TestObject) TestNumTxs() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestNumTxs()
	utest.Commit()
	return
}

//TestDataHash This is a method of TestObject
func (t *TestObject) TestDataHash() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestDataHash()
	utest.Commit()
	return
}

//TestProposerAddress This is a method of TestObject
func (t *TestObject) TestProposerAddress() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestProposerAddress()
	utest.Commit()
	return
}

//TestRewardAddress This is a method of TestObject
func (t *TestObject) TestRewardAddress() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRewardAddress()
	utest.Commit()
	return
}

//TestRandomNumber This is a method of TestObject
func (t *TestObject) TestRandomNumber() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestRandomNumber()
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

//TestLastBlockHash This is a method of TestObject
func (t *TestObject) TestLastBlockHash() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestLastBlockHash()
	utest.Commit()
	return
}

//TestCommitHash This is a method of TestObject
func (t *TestObject) TestLastCommitHash() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestLastCommitHash()
	utest.Commit()
	return
}

//TestAppHash This is a method of TestObject
func (t *TestObject) TestAppHash() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestAppHash()
	utest.Commit()
	return
}

//TestFee This is a method of TestObject
func (t *TestObject) TestLastFee() (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.TestLastFee()
	utest.Commit()
	return
}
