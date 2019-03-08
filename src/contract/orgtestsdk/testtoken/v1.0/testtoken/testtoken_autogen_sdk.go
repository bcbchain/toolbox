package testtoken

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestToken
func (tt *TestToken) SetSdk(sdk sdk.ISmartContract) {
	tt.sdk = sdk
}

//GetSdk This is a method of TestToken
func (tt *TestToken) GetSdk() sdk.ISmartContract {
	return tt.sdk
}
