package mymixtype

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of Mymixtype
func (m *Mymixtype) SetSdk(sdk sdk.ISmartContract) {
	m.sdk = sdk
}

//GetSdk This is a method of Mymixtype
func (m *Mymixtype) GetSdk() sdk.ISmartContract {
	return m.sdk
}
