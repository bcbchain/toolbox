package myplayerbook

import "blockchain/smcsdk/sdk"

//SetSdk set sdk
func (pb *MyPlayerBook) SetSdk(sdk sdk.ISmartContract) {
	pb.sdk = sdk
}

//GetSdk get sdk
func (pb *MyPlayerBook) GetSdk() sdk.ISmartContract {
	return pb.sdk
}
