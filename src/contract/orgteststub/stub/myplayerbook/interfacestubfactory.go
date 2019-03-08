package myplayerbook

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdkimpl/helper"
	"contract/stubcommon/types"

	v1_0 "contract/orgteststub/stub/myplayerbook/v1.0/myplayerbook"
	v2_0 "contract/orgteststub/stub/myplayerbook/v2.0/myplayerbook"
)

//NewInterfaceStub new interface stub
func NewInterfaceStub(smc sdk.ISmartContract, contractName string) types.IContractIntfcStub {
	//Get contract with ContractName
	ch := helper.ContractHelper{}
	ch.SetSMC(smc)
	contract := ch.ContractOfName(contractName)

	switch contract.Version() {
	case "1.0":
		return v1_0.NewIntfcStub(smc)
	case "2.0":
		return v2_0.NewIntfcStub(smc)
	}
	return nil
}
