package mybasictype

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of BasicType
func (bt *BasicType) SetSdk(sdk sdk.ISmartContract) {
	bt.sdk = sdk
}

//GetSdk This is a method of BasicType
func (bt *BasicType) GetSdk() sdk.ISmartContract {
	return bt.sdk
}
