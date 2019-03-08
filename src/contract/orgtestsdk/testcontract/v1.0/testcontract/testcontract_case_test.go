package testcontract

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

//TestTestContract_TestAddress This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestAddress()
}

//TestTestContract_TestAccount This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestAccount(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestAccount()
}

//TestTestContract_TestOwner This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestOwner(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestOwner()
}

//TestTestContract_TestName This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestName()
}

//TestTestContract_TestVersion This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestVersion(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestVersion()
}

//TestTestContract_TestCodeHash This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestCodeHash(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestCodeHash()
}

//TestTestContract_TestEffectHeight This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestEffectHeight(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestEffectHeight()
}

//TestTestContract_TestLoseHeight This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestLoseHeight(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestLoseHeight()
}

//TestTestContract_TestKeyPrefix This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestKeyPrefix(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestKeyPrefix()
}

//TestTestContract_TestMethods This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestMethods(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestMethods()
}

//TestTestContract_TestInterfaces This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestInterfaces(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestInterfaces()
}

//TestTestContract_TestToken This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestToken(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestToken()
}

//TestTestContract_TestOrgID This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestOrgID(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().setSender(contractOwner).TestOrgID()
}

//TestTestContract_TestSetOwner This is a method of MySuite
func (mysuit *MySuite) TestTestContract_TestSetOwner(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//注册代币后测试SetOwner
	test.obj.sdk.Helper().TokenHelper().RegisterToken("mycoin", "mycn", bn.N(9E18), true, true)
	test.run().setSender(contractOwner).TestSetOwner()
}
