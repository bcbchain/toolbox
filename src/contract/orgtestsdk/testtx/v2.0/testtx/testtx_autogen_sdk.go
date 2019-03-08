package testtx

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of TestTx
func (tt *TestTx) SetSdk(sdk sdk.ISmartContract) {
	tt.sdk = sdk
}

//GetSdk This is a method of TestTx
func (tt *TestTx) GetSdk() sdk.ISmartContract {
	return tt.sdk
}
