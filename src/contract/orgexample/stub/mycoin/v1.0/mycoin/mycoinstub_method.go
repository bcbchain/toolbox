package mycoinstub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"

	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	"contract/orgexample/code/mycoin/v1.0/mycoin"
	"github.com/tendermint/tmlibs/log"
)

//MycoinStub an object
type MycoinStub struct {
	logger log.Logger
}

var _ stubType.IContractStub = (*MycoinStub)(nil)

//New generate a stub
func New(logger log.Logger) stubType.IContractStub {
	return &MycoinStub{logger: logger}
}

//FuncRecover recover panic by Assert
func FuncRecover(response *bcType.Response) {
	if err := recover(); err != nil {
		if _, ok := err.(types.Error); ok {
			error := err.(types.Error)
			response.Code = error.ErrorCode
			response.Log = error.Error()
		} else {
			panic(err)
		}
	}
}

// InitChain initial smart contract
func (pbs *MycoinStub) InitChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	contractObj := new(mycoin.Mycoin)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (pbs *MycoinStub) UpdateChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	response.Code = types.CodeOK
	return response
}

//Invoke invoke function
func (pbs *MycoinStub) Invoke(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	// 生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(smc, true)
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(smc.Message(), response.Tags, "", fee, gasUsed, smc.Tx().GasLimit(), err)
		return
	}

	var data string
	err = types.Error{ErrorCode: types.CodeOK}
	switch smc.Message().MethodID() {
	case "44d8ca60": // prototype: Transfer(types.Address,bn.Number)
		transfer(smc)
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(smc.Message(), response.Tags, data, fee, gasUsed, smc.Tx().GasLimit(), err)
	return
}

func transfer(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 2, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 types.Address
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v1 bn.Number
	err = rlp.DecodeBytes(items[1], &v1)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mycoin.Mycoin)
	contractObj.SetSdk(smc)
	contractObj.Transfer(v0, v1)
}
