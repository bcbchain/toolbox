package testcontracthelper

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//TestContractHelper This is struct of contract
//@:contract:testcontracthelper
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111
//@:author:8d9d958020a77c9d3e04467d1e4095e43916a948a36168078bd28d9eeae7cfb4
type TestContractHelper struct {
	sdk sdk.ISmartContract
}

//InitChain Constructor of this TestContractHelper
//@:constructor
func (t *TestContractHelper) InitChain() {

}

//TestContractOfAddress test ContractOfAddress() interface of sdk IContractHelper
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContractHelper) TestContractOfAddress() {
	//根据合约地址构造合约对象，检查合约信息是否正确
	t.runCasesContractOfAddress()
}

//TestContractOfName test ContractOfName() interface of sdk IContractHelper
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContractHelper) TestContractOfName() {
	//根据合约名称构造合约对象，检查合约信息是否正确
	t.runCasesContractOfName()
}

//TestContractOfToken test ContractOfToken() interface of sdk IContractHelper
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContractHelper) TestContractOfToken() {
	//根据合约代币地址构造合约对象，检查合约信息是否正确
	t.runCasesContractOfToken()
}

//Transfer declare transfer interface of token of the contract owned
//@:public:method:gas[500]
//@:public:interface:gas[500]
func (t *TestContractHelper) Transfer(to types.Address, value bn.Number) {
	t.sdk.Message().Sender().Transfer(to, value)
}

func (t *TestContractHelper) contractOfToken(token types.Address) (icontract sdk.IContract, err types.Error) {
	defer funcRecover(&err)
	icontract = t.sdk.Helper().ContractHelper().ContractOfToken(token)
	return
}

func (t *TestContractHelper) contractOfAddress(addr types.Address) (icontract sdk.IContract, err types.Error) {
	defer funcRecover(&err)
	icontract = t.sdk.Helper().ContractHelper().ContractOfAddress(addr)
	return
}

func (t *TestContractHelper) contractOfName(name string) (icontract sdk.IContract, err types.Error) {
	defer funcRecover(&err)
	icontract = t.sdk.Helper().ContractHelper().ContractOfName(name)
	return
}
