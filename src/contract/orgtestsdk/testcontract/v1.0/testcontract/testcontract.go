package testcontract

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//TestContract This is struct of contract
//@:contract:testcontract
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxew
//@:author:8d9d958020a77c9d3e04467d1e4095e43916a948a36168078bd28d9eeae7cfb4
type TestContract struct {
	sdk sdk.ISmartContract
}

//InitChain Constructor of this TestContract
//@:constructor
func (t *TestContract) InitChain() {

}

//TestAddress test Address() interface of sdk IContract
//@:public:interface:gas[400]
//@:public:method:gas[400]
func (t *TestContract) TestAddress() {
	t.testAddress()
}

//TestAccount test Account() interface of sdk IContract
//@:public:method:gas[400]
func (t *TestContract) TestAccount() {
	t.testAccount()
}

//TestOwner test Owner() interface of sdk IContract
//@:public:method:gas[400]
func (t *TestContract) TestOwner() {
	t.testOwner()
}

//TestName test Name() interface of sdk IContract
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContract) TestName() {
	t.testName()
}

//TestVersion test Version() interface of sdk IContract
//@:public:method:gas[400]
func (t *TestContract) TestVersion() {
	t.testVersion()
}

//TestCodeHash test CodeHash() interface of sdk IContract
//@:public:method:gas[400]
func (t *TestContract) TestCodeHash() {
	t.testCodeHash()
}

//TestEffectHeight test EffectHeight() interface of sdk IContract
//@:public:method:gas[400]
func (t *TestContract) TestEffectHeight() {
	t.testEffectHeight()
}

//TestLoseHeight test LoseHeight() interface of sdk IContract
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContract) TestLoseHeight() {
	t.testLoseHeight()
}

//TestKeyPrefix test LoseHeight() interface of sdk IContract
//@:public:method:gas[400]
func (t *TestContract) TestKeyPrefix() {
	t.testKeyPrefix()
}

//TestMethods test Methods() interface of sdk IContract
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContract) TestMethods() {
	t.testMethods()
}

//TestInterfaces test Interfaces() interface of sdk IContract
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContract) TestInterfaces() {
	t.testInterfaces()
}

//TestToken test Token() interface of sdk IContract
//@:public:method:gas[400]
func (t *TestContract) TestToken() {
	t.testToken()
}

//TestOrgID test OrgID() interface of sdk IContract
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContract) TestOrgID() {
	t.testOrgID()
}

//TestSetOwner test SetOwner() interface of sdk IContract
//@:public:method:gas[400]
//@:public:interface:gas[400]
func (t *TestContract) TestSetOwner() {
	t.testSetOwner()
}

func (t *TestContract) setOwner(newowner types.Address) (err types.Error) {
	defer funcRecover(&err)
	t.sdk.Message().Contract().SetOwner(newowner)
	return
}

//Transfer declare transfer interface of token of the contract owned
//@:public:method:gas[500]
//@:public:interface:gas[500]
func (t *TestContract) Transfer(to types.Address, value bn.Number) {
	t.sdk.Message().Sender().Transfer(to, value)
}
