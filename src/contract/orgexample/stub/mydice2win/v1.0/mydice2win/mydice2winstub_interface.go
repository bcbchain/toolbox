package mydice2winstub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"
)

//InterfaceDice2WinStub interface stub
type InterfaceDice2WinStub struct {
	smc sdk.ISmartContract
}

var _ stubType.IContractIntfcStub = (*InterfaceDice2WinStub)(nil)

//NewInterStub new interface stub
func NewInterStub(smc sdk.ISmartContract) stubType.IContractIntfcStub {
	return &InterfaceDice2WinStub{smc: smc}
}

//GetSdk get sdk
func (inter *InterfaceDice2WinStub) GetSdk() sdk.ISmartContract {
	return inter.smc
}

//SetSdk set sdk
func (inter *InterfaceDice2WinStub) SetSdk(smc sdk.ISmartContract) {
	inter.smc = smc
}

// isCycle
func (inter *InterfaceDice2WinStub) isCycle(origins []types.Address) bool {
	m := make(map[string]struct{})
	for _, addr := range origins {
		if _, ok := m[addr]; ok {
			return true
		} else {
			m[addr] = struct{}{}
		}
	}

	return false
}

//Invoke invoke function
func (inter *InterfaceDice2WinStub) Invoke(methodID string, p interface{}) (response bcType.Response) {
	defer FuncRecover(&response)

	if inter.isCycle(inter.smc.Message().Origins()) {
		response.Code = types.ErrStubDefined
		response.Log = "invoke contract's interface cannot cycle"
		return
	}

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
	case "d517c92": // prototype: PlaceBet(bn.Number,int64,int64,[]byte,[]byte,types.Address)
		inter.placeBet(p)
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(inter.smc.Message(), nil, data, fee, gasUsed, inter.smc.Tx().GasLimit(), err)
	return
}

func (inter *InterfaceDice2WinStub) placeBet(p interface{}) {
	//contractObj := new(mydice2win.Dice2Win)
	//contractObj.SetSdk(inter.smc)
	//param := p.(mydice2win.PlaceBetParam)
	//contractObj.PlaceBet(param.BetMask, param.Modulo, param.CommitLastBlock, param.Commit, param.SignData, param.RefAddress)
}
