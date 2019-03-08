package testtx

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

//TestTestTx_Transfer This is a method of MySuite
func (mysuit *MySuite) TestTestTx_Transfer(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTestTx_TestNote This is a method of MySuite
func (mysuit *MySuite) TestTestTx_TestNote(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestNote()
}

//TestTestTx_TestGasLimit This is a method of MySuite
func (mysuit *MySuite) TestTestTx_TestGasLimit(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestGasLimit()
}

//TestTestTx_TestGasLeft This is a method of MySuite
func (mysuit *MySuite) TestTestTx_TestGasLeft(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestGasLeft()
}

//TestTestTx_TestSigner This is a method of MySuite
func (mysuit *MySuite) TestTestTx_TestSigner(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestSigner()
}
