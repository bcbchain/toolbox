package mycrossmc

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of MyCrossmc
func (mc *MyCrossmc) SetSdk(sdk sdk.ISmartContract) {
	mc.sdk = sdk
}

//GetSdk This is a method of MyCrossmc
func (mc *MyCrossmc) GetSdk() sdk.ISmartContract {
	return mc.sdk
}
