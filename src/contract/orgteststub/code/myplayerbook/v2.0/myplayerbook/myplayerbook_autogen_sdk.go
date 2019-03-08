package myplayerbook

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of MyPlayerBook
func (mpb *MyPlayerBook) SetSdk(sdk sdk.ISmartContract) {
	mpb.sdk = sdk
}

//GetSdk This is a method of MyPlayerBook
func (mpb *MyPlayerBook) GetSdk() sdk.ISmartContract {
	return mpb.sdk
}
