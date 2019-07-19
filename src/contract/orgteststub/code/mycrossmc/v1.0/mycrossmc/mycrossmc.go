package mycrossmc

import (
	"encoding/hex"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//MyCrossmc a demo for crossing contracts calling
//@:contract:mycrossmc
//@:version:1.0
//@:organization:orgGyRrMVF7ukfHNwaZhgWMTbQAYz7d7RcBh
//@:author:b37e7627431feb18123b81bcf1f41ffd37efdb90513d48ff2c7f8a0c27a9d06c
type MyCrossmc struct {
	sdk sdk.ISmartContract

	//@:public:store
	storedData uint64
}

type Player struct {
	Address types.Address
	Name    string
}

//nolint unused
//@:import:myplayerbook
type myplayerbook interface {
	RegisterName(index int64, plyr Player)
	GetPlayer(addr types.Address) string
	MultiTypesParam(index uint64, flt float64, bl bool, bt byte, hash types.Hash, hb types.HexBytes, bi bn.Number, mp map[int]string)
}

//InitChain init once only when deploy contract
//@:constructor
func (ms *MyCrossmc) InitChain() {
}

//Register register a value
//@:public:method:gas[500]
func (ms *MyCrossmc) Register(data uint64) {
	_to := ms.sdk.Helper().ContractHelper().ContractOfName(importContractName).Account()
	_from := ms.sdk.Message().Contract().Account()
	_from.TransferByName("LOC", _to.Address(), bn.N(1000000))
	plyr := Player{
		Address: "locNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJH",
		Name:    "bob",
	}

	ms.myplayerbook().run(func() {
		_from.TransferByName("LOC", _to.Address(), bn.N(1000000))
	}).RegisterName(55, plyr)

	ms.myplayerbook().RegisterName(55, plyr)
	ms._setStoredData(data)
}

//MultParam multi param
//nolint unhandled
func (ms *MyCrossmc) MultParam() {
	b, _ := hex.DecodeString("aabbccdd")
	mv := make(map[int]string)
	mv[999] = "testmap"

	ms.myplayerbook().MultiTypeParam(100, 20.22, true, 50, []byte("testhash"), b, bn.N(-888888), mv)
}

//Set sets data
//@:public:method:gas[100]
func (ms *MyCrossmc) Set(data uint64) {
	ms._setStoredData(data)
}

//Get gets data
//@:public:method:gas[100]
func (ms *MyCrossmc) Get() uint64 {
	return ms._storedData()
}
