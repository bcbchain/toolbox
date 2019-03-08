package myplayerbook

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
	contractName    = "myplayerbook" //contract name
	contractMethods = []string{"GetPlayer(types.Address)TYPE", "RegisterName(int64,Player)", "MultiTypesParam(uint64,float64,bool,byte,types.Hash,types.HexBytes,bn.Number,map[int]string)"}
	orgID           = "orgGyRrMVF7ukfHNwaZhgWMTbQAYz7d7RcBh"
)

//TestObject This is a struct for test
type TestObject struct {
	obj *MyPlayerBook
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
	return &TestObject{&MyPlayerBook{sdk: utest.UTP.ISmartContract}}
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

//GetPlayer This is a method of TestObject
func (t *TestObject) GetPlayer(addr types.Address) (result0 string, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.GetPlayer(addr)
	utest.Commit()
	return
}

//RegisterName This is a method of TestObject
func (t *TestObject) RegisterName(index int64, plyr Player) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.RegisterName(index, plyr)
	utest.Commit()
	return
}

//MultiTypesParam This is a method of TestObject
func (t *TestObject) MultiTypesParam(index uint64, flt float64, bl bool, bt byte, hash types.Hash, hb types.HexBytes, bi bn.Number, mp map[int]string) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.MultiTypesParam(index, flt, bl, bt, hash, hb, bi, mp)
	utest.Commit()
	return
}
