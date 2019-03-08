package mycoin

import (
	"reflect"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	types2 "blockchain/types"
	stubTypes "contract/stubcommon/types"

	mycoin_v1_0 "contract/orgexample/code/mycoin/v1.0/mycoin"
)

type IntfcCoinStub struct {
	smcApi sdk.ISmartContract
}

var _ stubTypes.IContractIntfcStub = (*IntfcCoinStub)(nil)

func NewIntfcStub(smcApi sdk.ISmartContract) stubTypes.IContractIntfcStub {
	return &IntfcCoinStub{smcApi: smcApi}
}

func (inter *IntfcCoinStub) GetSdk() sdk.ISmartContract {
	return inter.smcApi
}

func (inter *IntfcCoinStub) SetSdk(smc sdk.ISmartContract) {
	inter.smcApi = smc
}

func (inter *IntfcCoinStub) Invoke(methodID string, p interface{}) types2.Response {
	// TODO 扣手续费
	//
	switch methodID {
	case "23445656":
		return inter.core_transfer(p)
	}
	return types2.Response{}
}

func (inter *IntfcCoinStub) core_transfer(p interface{}) types2.Response {
	response := types2.Response{}

	param := reflect.ValueOf(p).Elem()
	if param.NumField() != 2 {
		smcError := types.Error{types.ErrInvalidParameter, ""}
		response.Code = smcError.ErrorCode
		response.Log = smcError.Error()
		return response
	}

	to := param.Field(0).Interface().(types.Address)
	value := param.Field(1).Interface().(bn.Number)
	sdk.RequireAddress(inter.GetSdk(), to)

	coin := new(mycoin_v1_0.Mycoin)
	coin.SetSdk(inter.smcApi)

	coin.Transfer(to, value)

	return response
}
