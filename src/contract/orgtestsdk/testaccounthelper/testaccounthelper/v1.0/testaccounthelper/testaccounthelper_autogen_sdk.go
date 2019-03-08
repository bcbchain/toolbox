package testaccounthelper

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestAccountHelper
func (tah *TestAccountHelper) SetSdk(sdk sdk.ISmartContract) {
	tah.sdk = sdk
}

//GetSdk This is a method of TestAccountHelper
func (tah *TestAccountHelper) GetSdk() sdk.ISmartContract {
	return tah.sdk
}
