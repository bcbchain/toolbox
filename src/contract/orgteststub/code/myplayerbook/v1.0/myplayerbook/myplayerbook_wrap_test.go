package myplayerbook

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	ut "blockchain/smcsdk/utest"
)

type TestObject struct {
	obj *MyPlayerBook
}

var (
	contractName    = "MyPlayerBook" //contract name
	contractMethods = []string{"GetPlayerID(types.Address)int64",
		"GetPlayerName(int64)string",
		"GetPlayerLAff(int64)int64",
		"GetPlayerAddr(int64)types.Address",
		"RegisterNameXid(string,nt64)types.Error",
	}
	orgID = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJH"
)

func NewTestObject(sender sdk.IAccount) *TestObject {
	return &TestObject{&MyPlayerBook{sdk: ut.UTP.ISmartContract}}
}

func (t *TestObject) run() *TestObject {
	t.obj.sdk = ut.ResetMsg()
	return t
}

func (t *TestObject) Init() {
	ut.NextBlock(1)

	t.obj.InitChain()
}

func (t *TestObject) registerNameXid(_name string, _affId int64) types.Error {
	ut.NextBlock(1)

	return t.obj.RegisterNameXid(_name, _affId)
}

func (t *TestObject) getPlayerAddr(_pID int64) types.Address {
	ut.NextBlock(1)

	return t.obj.GetPlayerAddr(_pID)
}

func (t *TestObject) getPlayerLAff(_pID int64) int64 {
	ut.NextBlock(1)

	return t.obj.GetPlayerLAff(_pID)
}

func (t *TestObject) getPlayerName(_pID int64) string {
	ut.NextBlock(1)

	return t.obj.GetPlayerName(_pID)
}

func (t *TestObject) getPlayerID(addr types.Address) int64 {
	ut.NextBlock(1)

	return t.obj.GetPlayerID(addr)
}
