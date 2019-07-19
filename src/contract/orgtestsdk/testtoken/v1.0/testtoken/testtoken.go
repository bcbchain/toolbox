package testtoken

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"fmt"
)

//TestToken This is struct of contract
//@:contract:testtoken
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzf
//@:author:8d9d958020a77c9d3e04467d1e4095e43916a948a36168078bd28d9eeae7cfb4
type TestToken struct {
	sdk sdk.ISmartContract

	//This is a sample field which is to store in db
	//@:public:store
	sampleStore string
}

const (
	// BurnAndAddEnabled burn enabled and add supply enabled too
	BurnAndAddEnabled = iota
	// BurnEnabledAndAddDisabled burn enabled but add supply disabled
	BurnEnabledAndAddDisabled
	// BurnDisableAndAddEnabled burn disabled but add supply enabled
	BurnDisableAndAddEnabled
	// BurnAndAddDisabled burn and add supply both disabled
	BurnAndAddDisabled
)

//InitChain Constructor of this TestToken
//@:constructor
func (t *TestToken) InitChain() {

}

func printPass() {
	fmt.Printf("---- PASS\n")
}
func printTestCase(index int, desc string) {
	fmt.Printf("Case:%3d (%s) \t", index, desc)
}

// TestSetTotalSupply total
//@:public:method:gas[500]
func (t *TestToken) TestSetTotalSupply(status int) {
	switch status {
	case BurnAndAddEnabled:
		t.setTotalSupplyBurnAndAddEnabled()
	case BurnDisableAndAddEnabled:
		t.setTotalSupplyBurnDisableAndAddEnabled()
	case BurnEnabledAndAddDisabled:
		t.setTotalSupplyBurnEnabledAndAddDisabled()
	case BurnAndAddDisabled:
		t.setTotalSupplyBurnAndAddDisabled()
	}
}

// TestSetGasPrice test token set gas price
//@:public:method:gas[500]
func (t *TestToken) TestSetGasPrice() {
	t.setGasPriceNormal()
}

// TestAddress test token address interface
//@:public:method:gas[500]
func (t *TestToken) TestAddress(address types.Address) {
	t.testAddress(address)
}

// TestOwner test token owner
//@:public:method:gas[500]
func (t *TestToken) TestOwner(owner types.Address) {
	t.testOwner(owner)
}

// TestName test token name
//@:public:method:gas[500]
func (t *TestToken) TestName(name string) {
	t.testName(name)
}

// TestSymbol test token symbol
//@:public:method:gas[500]
func (t *TestToken) TestSymbol(symbol string) {
	t.testSymbol(symbol)
}

// TestTotalSupply test token total supply
//@:public:method:gas[500]
func (t *TestToken) TestTotalSupply(total bn.Number) {
	t.testTotalSupply(total)
}

// TestAddSupplyEnabled test token get add supply enabled
//@:public:method:gas[500]
func (t *TestToken) TestAddSupplyEnabled(enabled bool) {
	t.testAddSupplyEnabled(enabled)
}

// TestBurnEnabled test token get add supply enabled
//@:public:method:gas[500]
func (t *TestToken) TestBurnEnabled(enabled bool) {
	t.testBurnEnabled(enabled)
}

// TestGasPrice test token get add supply enabled
//@:public:method:gas[500]
func (t *TestToken) TestGasPrice(gasPrice int64) {
	t.testGasPrice(gasPrice)
}

func (t *TestToken) runSetTotalSupply(index int, token sdk.IToken, totalSupply bn.Number, code uint32, desc, errMsg string) {
	printTestCase(index, desc)
	acc := token.Owner()
	oldBalance := acc.Balance()
	oldTotalSupply := token.TotalSupply()
	err := t.setTotalSupply(token, totalSupply)
	AssertError(err, code, errMsg)
	if code == types.CodeOK {
		acc := token.Owner()
		Assert(acc.Balance().Sub(oldBalance).IsEqual(totalSupply.Sub(oldTotalSupply)))
		Assert(token.TotalSupply().IsEqual(totalSupply))
	}

	printPass()
}

func (t *TestToken) setTotalSupply(token sdk.IToken, totalSupply bn.Number) (err types.Error) {
	defer funcRecover(&err)
	token.SetTotalSupply(totalSupply)
	return
}

func (t *TestToken) runSetGasPrice(index int, token sdk.IToken, gasPrice int64, code uint32, desc, errMsg string) {
	printTestCase(index, desc)
	err := t.setGasPrice(token, gasPrice)
	AssertError(err, code, errMsg)
	if code == types.CodeOK {
		Assert(token.GasPrice() == gasPrice)
	}
	printPass()
}

func (t *TestToken) setGasPrice(token sdk.IToken, gasPrice int64) (err types.Error) {
	defer funcRecover(&err)
	token.SetGasPrice(gasPrice)
	return
}
