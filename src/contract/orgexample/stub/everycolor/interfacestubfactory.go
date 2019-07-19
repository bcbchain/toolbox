package everycolor

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdkimpl/helper"
	v3_0 "contract/orgexample/stub/everycolor/v3.0/everycolor"
	"contract/stubcommon/types"
)

//NewInterfaceStub new interface stub
func NewInterfaceStub(smc sdk.ISmartContract, contractName string) types.IContractIntfcStub {
	//Get contract with ContractName
	ch := helper.ContractHelper{}
	ch.SetSMC(smc)
	contract := ch.ContractOfName(contractName)

	switch contract.Version() {
	case "3.0":
		return v3_0.NewInterStub(smc)
	}
	return nil
}
