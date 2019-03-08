package testtokenhelper

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//TestTokenHelper This is struct of contract
//@:contract:testtokenhelper
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzf
//@:author:011b8728b2e5fd42c769bffe46cc92348f3f14f24a42a860d3d7760784f97869
type TestTokenHelper struct {
	sdk sdk.ISmartContract
}

//InitChain Constructor of this TestTokenHelper
//@:constructor
func (t *TestTokenHelper) InitChain() {

}

//TestRegisterTokenName This is a sample method
//@:public:method:gas[500]
func (t *TestTokenHelper) TestRegisterTokenName() {
	t.testRegisterTokenName()
}

//TestRegisterTokenSymbol This is a sample method
//@:public:method:gas[500]
func (t *TestTokenHelper) TestRegisterTokenSymbol() {
	t.testRegisterTokenSymbol()
}

//TestRegisterTokenTotalSupply This is a sample method
//@:public:method:gas[500]
func (t *TestTokenHelper) TestRegisterTokenTotalSupply() {
	t.testRegisterTokenTotalSupply()
}

//TestRegisterTokenAddSupplyEnabled This is a sample method
//@:public:method:gas[500]
func (t *TestTokenHelper) TestRegisterTokenAddSupplyEnabled() {
	t.testRegisterTokenAddSupplyEnabled()
}

//TestRegisterTokenBurnEnabled This is a sample method
//@:public:method:gas[500]
func (t *TestTokenHelper) TestRegisterTokenBurnEnabled() {
	t.testRegisterTokenBurnEnabled()
}

//TestRegisterTokenDuplicate This is a sample method
//@:public:method:gas[500]
func (t *TestTokenHelper) TestRegisterTokenDuplicate() {
	t.testRegisterTokenDuplicate()
}

// TestToken test sdk tokenHelper token
//@:public:method:gas[500]
func (t *TestTokenHelper) TestToken() {
	t.testToken()
}

// TestTokenOfAddress test
//@:public:method:gas[500]
func (t *TestTokenHelper) TestTokenOfAddress() {
	t.testTokenOfAddress()
}

//TestTokenOfName test
//@:public:method:gas[500]
func (t *TestTokenHelper) TestTokenOfName() {
	t.testTokenOfName()
}

//TestTokenOfSymbol test
//@:public:method:gas[500]
func (t *TestTokenHelper) TestTokenOfSymbol() {
	t.testTokenOfSymbol()
}

//TestTokenOfContract test
//@:public:method:gas[500]
func (t *TestTokenHelper) TestTokenOfContract() {
	t.testTokenOfContract()
}

//TestBaseGasPrice test
//@:public:method:gas[500]
func (t *TestTokenHelper) TestBaseGasPrice() {
	t.testBaseGasPrice()
}

func (t *TestTokenHelper) runRegisterToken(
	index int,
	desc string,
	code uint32,
	errMsg string, name,
	symbol string,
	totalSupply bn.Number,
	addSupplyEnabled,
	burnEnabled bool) {

	printTestCase(index, desc)

	err := t.sdkRegisterToken(name, symbol, totalSupply, addSupplyEnabled, burnEnabled)
	AssertError(err, code, errMsg)
	if code == types.CodeOK {
		Assert(t.sdk.Helper().TokenHelper().Token().Name() == name)
		Assert(t.sdk.Helper().TokenHelper().Token().Symbol() == symbol)
		Assert(t.sdk.Helper().TokenHelper().Token().TotalSupply() == totalSupply)
		Assert(t.sdk.Helper().TokenHelper().Token().AddSupplyEnabled() == addSupplyEnabled)
		Assert(t.sdk.Helper().TokenHelper().Token().BurnEnabled() == burnEnabled)
	}

	printPass()
}

func (t *TestTokenHelper) sdkRegisterToken(name, symbol string, totalSupply bn.Number, addSupplyEnabled, burnEnabled bool) (err types.Error) {
	defer funcRecover(&err)
	t.sdk.Helper().TokenHelper().RegisterToken(name, symbol, totalSupply, addSupplyEnabled, burnEnabled)
	return
}
