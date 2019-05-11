package myplayerbook

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/rlp"
	types2 "blockchain/types"
	"contract/orgexample/code/mydice2win/v1.0/mydice2win"
	stubTypes "contract/stubcommon/types"
	"github.com/tendermint/tmlibs/log"
)

//PlayerBookStub player book stub
type PlayerBookStub struct {
	logger log.Logger
}

var _ stubTypes.IContractStub = (*PlayerBookStub)(nil)

//New new stub
func New(logger log.Logger) stubTypes.IContractStub {
	return &PlayerBookStub{logger: logger}
}

// InitChain initial smart contract
func (mc *PlayerBookStub) InitChain(smc sdk.ISmartContract) (response types2.Response) {

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types2.CodeOK
	return response
}

// UpdateChain update smart contract
func (mc *PlayerBookStub) UpdateChain(smc sdk.ISmartContract) (response types2.Response) {
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.UpdateChain()

	response.Code = types2.CodeOK
	return response
}

//Invoke invoke function
func (pbs *PlayerBookStub) Invoke(smcAPI sdk.ISmartContract) types2.Response {
	// TODO 扣手续费
	//
	switch smcAPI.Message().MethodID() {
	case "23445656": // prototype: GetPlayer(types.Address)*Player
		return getPlayer(smcAPI)

	case "f94d817e": // prototype: RegisterName(string,string,string)(types.Error)
		return registerName(smcAPI)
	}

	return types2.Response{}
}

func getPlayer(smcAPI sdk.ISmartContract) (response types2.Response) {

	itemsBytes := smcAPI.Message().Items()

	if len(itemsBytes) != 1 {
		response.Code = 4001
		response.Log = "Message can not be nil."

		return response
	}

	var param int64
	if err := rlp.DecodeBytes(itemsBytes[0], param); err != nil {
		response.Code = 4001
		response.Log = err.Error()

		return response
	}

	inter := IntfcPlayerBookStub{smcAPI}

	return inter.coreGetPlayer(param)
}

func registerName(smcAPI sdk.ISmartContract) (response types2.Response) {

	itemsBytes := smcAPI.Message().Items()

	if len(itemsBytes) != 2 {
		response.Code = 4001
		response.Log = "Message can not be nil."

		return response
	}
	//
	//param := RegisterName{}
	//if err := rlp.DecodeBytes(itemsBytes[0], param.Name); err != nil {
	//	response.Code = 4001
	//	response.Log = err.Error()
	//
	//	return response
	//}
	//
	//if err := rlp.DecodeBytes(itemsBytes[1], param.Sex); err != nil {
	//	response.Code = 4004
	//	response.Log = err.Error()
	//
	//	return response
	//}
	//if err := rlp.DecodeBytes(itemsBytes[2], param.BirthDate); err != nil {
	//	response.Code = 4004
	//	response.Log = err.Error()
	//
	//	return response
	//}
	//
	//inter := Intfc_PlayerBookStub{smcAPI}
	//
	//return inter.core_registerName(&param)
	return response
}
