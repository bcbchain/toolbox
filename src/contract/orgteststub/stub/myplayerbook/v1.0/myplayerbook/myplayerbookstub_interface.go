package myplayerbook

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
	types2 "blockchain/types"
	stubTypes "contract/stubcommon/types"

	playerbook2 "contract/orgteststub/code/myplayerbook/v1.0/myplayerbook"
)

//IntfcPlayerBookStub struct
type IntfcPlayerBookStub struct {
	smcAPI sdk.ISmartContract
}

var _ stubTypes.IContractIntfcStub = (*IntfcPlayerBookStub)(nil)

//NewIntfcStub new stub
func NewIntfcStub(smcAPI sdk.ISmartContract) stubTypes.IContractIntfcStub {
	return &IntfcPlayerBookStub{smcAPI: smcAPI}
}

//GetSdk get sdk
func (inter *IntfcPlayerBookStub) GetSdk() sdk.ISmartContract {
	return inter.smcAPI
}

//SetSdk set sdk
func (inter *IntfcPlayerBookStub) SetSdk(smc sdk.ISmartContract) {
	inter.smcAPI = smc
}

//Invoke invoke function
func (inter *IntfcPlayerBookStub) Invoke(methodID string, p interface{}) types2.Response {
	// stub里生成燃料收据（不扣费），但是要检查账户余额是否足够支付燃料消耗，交易完成后GIChain根据燃料收据合并为一条手续费收据，并生成手续费转账交易，执行手续费扣费。
	// 跨合约调用的燃料费由Sender支付
	switch methodID {
	case "23445656": // prototype: GetPlayer(types.Address)*Player
		return inter.coreGetPlayer(p.(int64))

		//case 0xf94d817e:   // prototype: RegisterName(string,string,string)(types.Error)
		//	return inter.core_registerName(p.(*myplayerbook.RegisterName))
	}

	return types2.Response{}
}

func (inter *IntfcPlayerBookStub) coreGetPlayer(p int64) types2.Response {
	response := types2.Response{}

	plyrbk := new(playerbook2.MyPlayerBook)
	plyrbk.SetSdk(inter.smcAPI)

	player := plyrbk.GetPlayer(p)

	response.Code = types.CodeOK
	response.Log = ""

	b, err := jsoniter.Marshal(*player)
	if err != nil {
		panic(err)
	}
	response.Data = string(b)

	return response
}

//func (inter *Intfc_PlayerBookStub)core_registerName(p *myplayerbook.RegisterName) types2.Response {
//	response := types2.Response{}
//
//	plyrbk := new(playerbook2.MyPlayerBook)
//	plyrbk.SetSdk( inter.smcAPI )
//
//	smcError := plyrbk.RegisterName(p.Name, p.Sex, p.BirthDate)
//
//	response.Code = smcError.ErrorCode
//	response.Log = smcError.Error()
//
//	return response
//}
