package everycolorstub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"
)

//InterfaceEverycolorStub interface stub
type InterfaceEverycolorStub struct {
	smc sdk.ISmartContract
}

var _ stubType.IContractIntfcStub = (*InterfaceEverycolorStub)(nil)

//NewInterStub new interface stub
func NewInterStub(smc sdk.ISmartContract) stubType.IContractIntfcStub {
	return &InterfaceEverycolorStub{smc: smc}
}

//GetSdk get sdk
func (inter *InterfaceEverycolorStub) GetSdk() sdk.ISmartContract {
	return inter.smc
}

//SetSdk set sdk
func (inter *InterfaceEverycolorStub) SetSdk(smc sdk.ISmartContract) {
	inter.smc = smc
}

//Invoke invoke function
func (inter *InterfaceEverycolorStub) Invoke(methodID string, p interface{}) (response bcType.Response) {
	defer FuncRecover(&response)

	if len(inter.smc.Message().Origins()) > 8 {
		response.Code = types.ErrStubDefined
		response.Log = "invoke contract's interface steps beyond 8 step"
		return
	}

	// 生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(inter.smc, false)
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(inter.smc.Message(), response.Tags, "", fee, gasUsed, inter.smc.Tx().GasLimit(), err)
		return
	}

	var data string
	err = types.Error{ErrorCode: types.CodeOK}
	switch methodID {
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(inter.smc.Message(), nil, data, fee, gasUsed, inter.smc.Tx().GasLimit(), err)
	return
}