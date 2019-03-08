package testblock

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestBlock
func (tb *TestBlock) SetSdk(sdk sdk.ISmartContract) {
	tb.sdk = sdk
}

//GetSdk This is a method of TestBlock
func (tb *TestBlock) GetSdk() sdk.ISmartContract {
	return tb.sdk
}
