package mycoinstub

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdkimpl"
	"blockchain/smcsdk/sdkimpl/helper"
	"contract/stubcommon/common"
	"contract/stubcommon/types"

	"contract/orgexample/stub/mycoin/v1.0/mycoin"
)

func NewIntfcfaceStub(smc sdk.ISmartContract, contractName string) types.IContractIntfcStub {
	//Get contract with ContractName
	ch := helper.ContractHelper{}
	ch.SetSMC(smc)
	//contract := ch.ContractOfName(contractName)

	switch common.CalcKey("mycoin", "1.0") {
	case "mycoin_1_0":
		return mycoin.NewIntfcStub(smc.(*sdkimpl.SmartContract))
		//case "mycoin_1_1":
		//	return mycoin_1_1_stub.NewIntfcStub(smc.(*sdk.SmartContract))

	}
	return nil
}
