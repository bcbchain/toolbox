package testtokenhelper

import (
	"blockchain/smcsdk/utest"
	"testing"

	"gopkg.in/check.v1"
)

//Test This is a function
func Test(t *testing.T) { check.TestingT(t) }

//MySuite This is a struct
type MySuite struct{}

var _ = check.Suite(&MySuite{})

//TestTestTokenHelper_TestRegisterToken This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestRegisterTokenName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestRegisterTokenName()
}

//TestTestTokenHelper_TestRegisterTokenSymbol This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestRegisterTokenSymbol(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestRegisterTokenSymbol()
}

//TestTestTokenHelper_TestRegisterTokenTotalSupply This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestRegisterTokenTotalSupply(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestRegisterTokenTotalSupply()
}

//TestTestTokenHelper_TestRegisterTokenAddSupplyEnabled This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestRegisterTokenAddSupplyEnabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestRegisterTokenAddSupplyEnabled()
}

//TestTestTokenHelper_TestRegisterTokenBurnEnabled This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestRegisterTokenBurnEnabled(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestRegisterTokenBurnEnabled()
}

//TestTestTokenHelper_TestRegisterTokenDuplicate This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestRegisterTokenDuplicate(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestRegisterTokenDuplicate()
}

//TestTestTokenHelper_TestToken This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestToken(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestToken()
}

//TestTestTokenHelper_TestTokenOfAddress This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestTokenOfAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestTokenOfAddress()
}

//TestTestTokenHelper_TestTokenOfName This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestTokenOfName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestTokenOfName()
}

//TestTestTokenHelper_TestTokenOfSymbol This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestTokenOfSymbol(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestTokenOfSymbol()
}

//TestTestTokenHelper_TestTokenOfContract This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestTokenOfContract(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestTokenOfContract()
}

//TestTestTokenHelper_TestBaseGasPrice This is a method of MySuite
func (mysuit *MySuite) TestTestTokenHelper_TestBaseGasPrice(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.run().TestBaseGasPrice()
}
