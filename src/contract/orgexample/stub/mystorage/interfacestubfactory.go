package mystorage

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdkimpl/helper"
	v1_0 "contract/orgexample/stub/mystorage/v1.0/mystorage"
	"contract/stubcommon/types"
)

//NewInterfaceStub new interface stub
func NewInterfaceStub(smc sdk.ISmartContract, contractName string) types.IContractIntfcStub {
	//Get contract with ContractName
	ch := helper.ContractHelper{}
	ch.SetSMC(smc)
	contract := ch.ContractOfName(contractName)

	switch contract.Version() {
	case "1.0":
		return v1_0.NewInterStub(smc)
	}
	return nil
}
