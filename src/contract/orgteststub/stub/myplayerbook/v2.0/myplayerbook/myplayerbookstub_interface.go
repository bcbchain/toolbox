package myplayerbookstub

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	types2 "blockchain/types"
	"contract/stubcommon/common"
	types1 "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"

	"contract/orgteststub/code/myplayerbook/v2.0/myplayerbook"
)

//IntfcMyPlayerBookStub interface stub
type IntfcMyPlayerBookStub struct {
	smc sdk.ISmartContract
}

var _ types1.IContractIntfcStub = (*IntfcMyPlayerBookStub)(nil)

//NewIntfcStub new interface stub
func NewIntfcStub(smc sdk.ISmartContract) types1.IContractIntfcStub {
	return &IntfcMyPlayerBookStub{smc: smc}
}

//GetSdk get sdk
func (intfc *IntfcMyPlayerBookStub) GetSdk() sdk.ISmartContract {
	return intfc.smc
}

//SetSdk set sdk
func (intfc *IntfcMyPlayerBookStub) SetSdk(smc sdk.ISmartContract) {
	intfc.smc = smc
}

//Invoke invoke function
//TODO: 跨合约调用是否可以不用返回response??
func (intfc *IntfcMyPlayerBookStub) Invoke(methodID string, p interface{}) (response types2.Response) {
	defer FuncRecover(&response)

	// 扣手续费并生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(intfc.smc, false)
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(intfc.smc.Message(), "", fee, gasUsed, intfc.smc.Tx().GasLimit())
		return
	}
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})

	// 根据MethodID调用func
	var data string
	switch methodID {
	case "f94d817e": // prototype: GetPlayer(types.Address)string
		data = intfc.getPlayer(p)

	case "e463fdb2": // prototype: RegisterName(string)(types.Error)
		intfc.registerName(p)

	case "cccccccc": // prototype: RegisterName(string)(types.Error)
		intfc.multiTypeParam(p)
	}
	response = common.CreateResponse(intfc.smc.Message(), data, fee, gasUsed, intfc.smc.Tx().GasLimit())
	return
}

func (intfc *IntfcMyPlayerBookStub) multiTypeParam(p interface{}) {
	plyrbk := new(myplayerbook.MyPlayerBook)
	plyrbk.SetSdk(intfc.smc)
	param := p.(myplayerbook.MultiTypesParam)
	plyrbk.MultiTypesParam(param.Index, param.Flt, param.Bl, param.Bt, param.Hash, param.Hb, param.Bi, param.Mp)
}

func (intfc *IntfcMyPlayerBookStub) getPlayer(p interface{}) (addr types.Address) {
	plyrbk := new(myplayerbook.MyPlayerBook)
	plyrbk.SetSdk(intfc.smc)
	param := p.(myplayerbook.GetPlayerParam)
	addr = plyrbk.GetPlayer(param.Addr)
	return
}

func (intfc *IntfcMyPlayerBookStub) registerName(p interface{}) {
	plyrbk := new(myplayerbook.MyPlayerBook)
	plyrbk.SetSdk(intfc.smc)
	param := p.(myplayerbook.RegisterNameParam)
	plyrbk.RegisterName(param.Index, param.Plyr)
}
