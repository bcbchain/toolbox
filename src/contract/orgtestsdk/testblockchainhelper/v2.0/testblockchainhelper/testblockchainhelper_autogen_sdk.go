package testblockchainhelper

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestblockChainHelper
func (t *TestblockChainHelper) SetSdk(sdk sdk.ISmartContract) {
	t.sdk = sdk
}

//GetSdk This is a method of TestblockChainHelper
func (t *TestblockChainHelper) GetSdk() sdk.ISmartContract {
	return t.sdk
}
