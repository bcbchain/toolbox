package testmessage

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestMessage
func (tm *TestMessage) SetSdk(sdk sdk.ISmartContract) {
	tm.sdk = sdk
}

//GetSdk This is a method of TestMessage
func (tm *TestMessage) GetSdk() sdk.ISmartContract {
	return tm.sdk
}
