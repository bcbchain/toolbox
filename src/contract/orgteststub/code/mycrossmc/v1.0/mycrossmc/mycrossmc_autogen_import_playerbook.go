package mycrossmc

import (
	"blockchain/smcsdk/common/gls"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	types3 "blockchain/types"
	"common/jsoniter"
	myplayerbookv1 "contract/orgteststub/code/myplayerbook/v1.0/myplayerbook"
	myplayerbookv2 "contract/orgteststub/code/myplayerbook/v2.0/myplayerbook"
	myplayerbook3 "contract/orgteststub/stub/myplayerbook"
	types2 "contract/stubcommon/types"
)

//InterfaceStub interface stub of playerbook
type InterfaceStub struct {
	stub     types2.IContractIntfcStub
	receipts []types.KVPair
}

const importContractName = "myplayerbook"

func (s *MyCrossmc) myplayerbook() *InterfaceStub {
	return &InterfaceStub{
		stub:     myplayerbook3.NewInterfaceStub(s.sdk, importContractName),
		receipts: make([]types.KVPair, 0),
	}
}

func (intfc *InterfaceStub) run(f func()) *InterfaceStub {
	// step 1. save old all receipts
	oldReceipts := intfc.stub.GetSdk().Message().(*object.Message).OutputReceipts()

	// step 2. run function
	f()

	// step 3. save new all receipts
	newReceipts := intfc.stub.GetSdk().Message().(*object.Message).OutputReceipts()

	// step 4. sub new all receipts off old all receipts
	if len(newReceipts) > len(oldReceipts) {
		intfc.receipts = newReceipts[len(oldReceipts):]
	}
	return intfc
}

func (intfc *InterfaceStub) contract() sdk.IContract {
	return intfc.stub.GetSdk().Helper().ContractHelper().ContractOfName(importContractName)
}

//RegisterName register name by calling playerbook contract
func (intfc *InterfaceStub) RegisterName(index int64, plyr Player) {

	methodID := "e463fdb2" // prototype: RegisterName(string)(types.Error)
	oldSmc := intfc.stub.GetSdk()
	defer intfc.stub.SetSdk(oldSmc)
	//合约调用时的输入收据，同时可作为跨合约调用的输入收据
	contract := oldSmc.Helper().ContractHelper().ContractOfName(importContractName)
	sdk.Require(contract != nil, types.ErrExpireContract, "")

	newSmc := sdkhelper.OriginNewMessage(oldSmc, contract, methodID, intfc.receipts)
	intfc.stub.SetSdk(newSmc)

	//TODO 编译时builder从数据库已获取合约版本和失效高度，直接使用
	height := newSmc.Block().Height()
	var rn interface{}

	if height < 1000 {
		panic("Invalid parameter") // if parameters are not matched to specified version, panic
	} else {
		var ply myplayerbookv2.Player
		resBytes, err := jsoniter.Marshal(plyr)
		if err != nil {
			panic(err)
		}
		err = jsoniter.Unmarshal(resBytes, &ply)
		if err != nil {
			panic(err)
		}

		rn = myplayerbookv2.RegisterNameParam{
			Index: index,
			Plyr:  ply,
		}
	}

	var response types3.Response
	gls.Mgr.SetValues(gls.Values{gls.SDKKey: newSmc}, func() {
		response = intfc.stub.Invoke(methodID, rn)
	})
	if response.Code != types.CodeOK {
		panic(response.Log)
	}
	oldmsg := oldSmc.Message()
	oldmsg.(*object.Message).AppendOutput(intfc.stub.GetSdk().Message().(*object.Message).OutputReceipts())
	intfc.receipts = nil
}

//GetPlayer get player information by calling playerbook contract
func (intfc *InterfaceStub) GetPlayer(addr types.Address) string {
	methodID := "f94d817e" // prototype: GetPlayer(types.Address)*Player

	oldSmc := intfc.stub.GetSdk()
	defer intfc.stub.SetSdk(oldSmc)

	//合约调用时的输入收据，同时可作为跨合约调用的输入收据
	contract := oldSmc.Helper().ContractHelper().ContractOfName(importContractName)
	newSmc := sdkhelper.OriginNewMessage(oldSmc, contract, methodID, oldSmc.Message().(*object.Message).OutputReceipts())
	intfc.stub.SetSdk(newSmc)
	height := newSmc.Block().Height()
	var rn interface{}
	if height <= 1000 {
		rn = myplayerbookv1.GetPlayerParam{
			Addr: addr,
		}
	} else {
		rn = myplayerbookv2.GetPlayerParam{
			Addr: addr,
		}
	}

	response := intfc.stub.Invoke(methodID, rn)
	if response.Code != types.CodeOK {
		return ""
	}
	oldmsg := oldSmc.Message()
	oldmsg.(*object.Message).AppendOutput(intfc.stub.GetSdk().Message().(*object.Message).OutputReceipts())
	return response.Data
}

//MultiTypeParam test func
func (intfc *InterfaceStub) MultiTypeParam(index uint64, flt float64, bl bool, bt byte, hash types.Hash, hb types.HexBytes, bi bn.Number, mp map[int]string) {

	methodID := "cccccccc" // prototype: RegisterName(string)(types.Error)
	oldSmc := intfc.stub.GetSdk()
	defer intfc.stub.SetSdk(oldSmc)

	//合约调用时的输入收据，同时可作为跨合约调用的输入收据
	contract := oldSmc.Helper().ContractHelper().ContractOfName(importContractName)
	newSmc := sdkhelper.OriginNewMessage(oldSmc, contract, methodID, oldSmc.Message().(*object.Message).OutputReceipts())
	intfc.stub.SetSdk(newSmc)

	height := newSmc.Block().Height()
	var rn interface{}
	if height <= 1000 {
		panic("Invalid parameter") // if parameters are not matched to specified version, panic
	} else {
		rn = myplayerbookv2.MultiTypesParam{
			Index: index,
			Flt:   flt,
			Bl:    bl,
			Bt:    bt,
			Hash:  hash,
			Hb:    hb,
			Bi:    bi,
			Mp:    mp,
		}
	}
	response := intfc.stub.Invoke(methodID, rn)
	if response.Code != types.CodeOK {
		panic(response.Log)
	}
	oldmsg := oldSmc.Message()
	oldmsg.(*object.Message).AppendOutput(intfc.stub.GetSdk().Message().(*object.Message).OutputReceipts())
}
