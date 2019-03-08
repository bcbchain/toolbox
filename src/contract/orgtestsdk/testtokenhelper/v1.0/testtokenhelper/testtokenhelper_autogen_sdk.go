package testtokenhelper

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestTokenHelper
func (tth *TestTokenHelper) SetSdk(sdk sdk.ISmartContract) {
	tth.sdk = sdk
}

//GetSdk This is a method of TestTokenHelper
func (tth *TestTokenHelper) GetSdk() sdk.ISmartContract {
	return tth.sdk
}
