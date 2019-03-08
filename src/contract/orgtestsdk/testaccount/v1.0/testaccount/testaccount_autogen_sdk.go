package testaccount

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestAccount
func (ta *TestAccount) SetSdk(sdk sdk.ISmartContract) {
	ta.sdk = sdk
}

//GetSdk This is a method of TestAccount
func (ta *TestAccount) GetSdk() sdk.ISmartContract {
	return ta.sdk
}
