package mycoinstub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"

	"contract/orgexample/code/mycoin/v1.0/mycoin"
)

//InterfaceMycoinStub interface stub
type InterfaceMycoinStub struct {
	smc sdk.ISmartContract
}

var _ stubType.IContractIntfcStub = (*InterfaceMycoinStub)(nil)

//NewInterStub new interface stub
func NewInterStub(smc sdk.ISmartContract) stubType.IContractIntfcStub {
	return &InterfaceMycoinStub{smc: smc}
}

//GetSdk get sdk
func (inter *InterfaceMycoinStub) GetSdk() sdk.ISmartContract {
	return inter.smc
}

//SetSdk set sdk
func (inter *InterfaceMycoinStub) SetSdk(smc sdk.ISmartContract) {
	inter.smc = smc
}

//Invoke invoke function
func (inter *InterfaceMycoinStub) Invoke(methodID string, p interface{}) (response bcType.Response) {
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
	case "44d8ca60": // prototype: Transfer(types.Address,bn.Number)
		inter.transfer(p)
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(inter.smc.Message(), nil, data, fee, gasUsed, inter.smc.Tx().GasLimit(), err)
	return
}

func (inter *InterfaceMycoinStub) transfer(p interface{}) {
	contractObj := new(mycoin.Mycoin)
	contractObj.SetSdk(inter.smc)
	param := p.(mycoin.TransferParam)
	contractObj.Transfer(param.To, param.Value)
}
