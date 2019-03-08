package mybasictype

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdkimpl/helper"
	"contract/stubcommon/common"
	"contract/stubcommon/types"
)

func NewIntfcfaceStub(smc sdk.ISmartContract, contractName string) types.IContractIntfcStub {
	//Get contract with ContractName
	ch := helper.ContractHelper{}
	ch.SetSMC(smc)
	//contract := ch.ContractOfName(contractName)

	switch common.CalcKey("mybasictype", "1.0") {
	case "mybasictype_1_0":
		//return mybasictype.NewIntfcStub(smc.(*sdkimpl.SmartContract))
	}
	return nil
}
