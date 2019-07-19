package types

import (
	"blockchain/smcsdk/sdk"
	"blockchain/types"
)

type IContractStub interface {
	InitChain(smcapi sdk.ISmartContract) types.Response
	UpdateChain(smcapi sdk.ISmartContract) types.Response
	Mine(smcapi sdk.ISmartContract) types.Response
	Invoke(smcapi sdk.ISmartContract) types.Response
	InvokeInternal(smcapi sdk.ISmartContract, feeFlag bool) types.Response
}

type IContractIntfcStub interface {
	Invoke(methodid string, p interface{}) types.Response
	GetSdk() sdk.ISmartContract
	SetSdk(smc sdk.ISmartContract)
}
