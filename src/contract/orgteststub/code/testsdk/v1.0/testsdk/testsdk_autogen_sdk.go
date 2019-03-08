package testsdk

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk set sdk
func (ts *TestSdk) SetSdk(sdk sdk.ISmartContract) {
	ts.sdk = sdk
}

//GetSdk get sdk
func (ts *TestSdk) GetSdk() sdk.ISmartContract {
	return ts.sdk
}
