package testsdk

import (
	"blockchain/smcsdk/sdk"
	. "blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	ut "blockchain/smcsdk/utest"
)

var (
	contractName    = "TestSdk" //contract name
	contractMethods = []string{"SetDatai16(int16)", "GetDatai16()int16", "SetOwner(types.Address)types.Error", "SetMydata(Mystruct)types.Error"}
)

type TestObject struct {
	obj *TestSdk
}

func NewTestObject(sender sdk.IAccount) *TestObject {
	return &TestObject{&TestSdk{sdk: ut.UTP.ISmartContract}}
}
func (t *TestObject) transfer(balance Number) *TestObject {
	t.obj.sdk.Message().Sender().Transfer(t.obj.sdk.Message().Contract().Account().Address(), balance)
	return t
}
func (t *TestObject) setSender(sender sdk.IAccount) *TestObject {
	t.obj.sdk = ut.SetSender(sender.Address())
	return t
}
func (t *TestObject) run() *TestObject {
	t.obj.sdk = ut.ResetMsg()
	return t
}
func (t *TestObject) SetDatai16(d int16) {
	ut.NextBlock(1)
	t.obj.SetDatai16(d)
	ut.Commit()
}

func (t *TestObject) GetDatai16() int16 {
	ut.NextBlock(1)
	return t.obj.GetDatai16()
}

func (t *TestObject) SetOwner(owner types.Address) types.Error {
	ut.NextBlock(1)
	return t.obj.SetOwner(owner)
}

func (t *TestObject) SetMydata(v Mystruct) types.Error {
	ut.NextBlock(1)
	t.obj.SetMydata(v)
	ut.Commit()
	return types.Error{ErrorCode: types.CodeOK}
}

func (t *TestObject) GetMydata() (Mystruct, types.Error) {
	ut.NextBlock(1)
	v, err := t.obj.GetMydata()

	ut.Commit()
	return v, err
}
