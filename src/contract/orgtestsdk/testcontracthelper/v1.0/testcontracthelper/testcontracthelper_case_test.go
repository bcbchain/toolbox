package testcontracthelper

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

//TestTestContractHelper_TestContractOfAddress This is a method of MySuite
func (mysuit *MySuite) TestTestContractHelper_TestContractOfAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.run().setSender(contractOwner).TestContractOfAddress()
}

//TestTestContractHelper_TestContractOfName This is a method of MySuite
func (mysuit *MySuite) TestTestContractHelper_TestContractOfName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.run().setSender(contractOwner).TestContractOfName()
}

//TestTestContractHelper_TestContractOfToken This is a method of MySuite
func (mysuit *MySuite) TestTestContractHelper_TestContractOfToken(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.run().setSender(contractOwner).TestContractOfToken()
	//注册代币
	test.obj.sdk.Helper().TokenHelper().RegisterToken("mycoin", "mycn", bn.N(9E18), true, true)
	test.run().setSender(contractOwner).TestContractOfToken()
}

//TestTestContractHelper_Transfer This is a method of MySuite
func (mysuit *MySuite) TestTestContractHelper_Transfer(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}
