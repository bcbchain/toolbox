package types

import (
	"blockchain/smcsdk/sdk"
	"blockchain/types"
)

type IContractStub interface {
	Invoke(smcapi sdk.ISmartContract) types.Response
}

type IContractIntfcStub interface {
	Invoke(methodid string, p interface{}) types.Response
	GetSdk() sdk.ISmartContract
	SetSdk(smc sdk.ISmartContract)
}
