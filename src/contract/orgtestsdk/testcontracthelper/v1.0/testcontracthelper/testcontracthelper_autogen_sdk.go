package testcontracthelper

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestContractHelper
func (tch *TestContractHelper) SetSdk(sdk sdk.ISmartContract) {
	tch.sdk = sdk
}

//GetSdk This is a method of TestContractHelper
func (tch *TestContractHelper) GetSdk() sdk.ISmartContract {
	return tch.sdk
}
