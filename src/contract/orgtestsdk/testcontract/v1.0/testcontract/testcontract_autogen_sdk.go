package testcontract

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestContract
func (tc *TestContract) SetSdk(sdk sdk.ISmartContract) {
	tc.sdk = sdk
}

//GetSdk This is a method of TestContract
func (tc *TestContract) GetSdk() sdk.ISmartContract {
	return tc.sdk
}
