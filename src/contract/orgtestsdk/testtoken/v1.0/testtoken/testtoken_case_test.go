package testtoken

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/utest"
	"testing"

	"gopkg.in/check.v1"
)

//Test This is a function
func Test(t *testing.T) { check.TestingT(t) }

//MySuite This is a struct
type MySuite struct{}

var _ = check.Suite(&MySuite{})

// 都可以
func (mysuit *MySuite) TestTestToken_TestSetTotalSupplyBurnAndAddEnabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestSetTotalSupply(BurnAndAddEnabled)
}

// 可以增发不可以燃烧
func (mysuit *MySuite) TestTestToken_TestSetTotalSupplyBurnDisableAndAddEnabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, false)
	test.run().setSender(contractOwner).TestSetTotalSupply(BurnDisableAndAddEnabled)
}

// 可以燃烧不可以增发
func (mysuit *MySuite) TestTestToken_TestSetTotalSupplyBurnEnabledAndAddDisabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), false, true)
	test.run().setSender(contractOwner).TestSetTotalSupply(BurnEnabledAndAddDisabled)
}

// 都不可以
func (mysuit *MySuite) TestTestToken_TestSetTotalSupplyBurnAndAddDisabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), false, false)
	test.run().setSender(contractOwner).TestSetTotalSupply(BurnAndAddDisabled)
}

// set gas price
func (mysuit *MySuite) TestTestToken_TestSetGasPrice(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestSetGasPrice()
}

// address
func (mysuit *MySuite) TestTestToken_TestAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestAddress(test.obj.sdk.Message().Contract().Address())
}

// owner
func (mysuit *MySuite) TestTestToken_TestOwner(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestOwner(test.obj.sdk.Message().Contract().Owner())
}

// name
func (mysuit *MySuite) TestTestToken_TestName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestName("token-test")
}

// symbol
func (mysuit *MySuite) TestTestToken_TestSymbol(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestSymbol("test")
}

// total supply
func (mysuit *MySuite) TestTestToken_TestTotalSupply(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestTotalSupply(bn.N(1E15))
}

// addSupplyEnabled
func (mysuit *MySuite) TestTestToken_TestAddSupplyEnabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, true)
	test.run().setSender(contractOwner).TestAddSupplyEnabled(true)
}

// burnEnabled
func (mysuit *MySuite) TestTestToken_TestBurnEnabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, false)
	test.run().setSender(contractOwner).TestBurnEnabled(false)
}

// gasPrice
func (mysuit *MySuite) TestTestToken_TestGasPrice(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken("token-test",
		"test", bn.N(1E15), true, false)
	test.obj.sdk.Helper().TokenHelper().Token().SetGasPrice(5000)
	test.run().setSender(contractOwner).TestGasPrice(5000)
}
