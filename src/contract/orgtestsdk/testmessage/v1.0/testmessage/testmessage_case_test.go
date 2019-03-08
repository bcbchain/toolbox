package testmessage

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

//TestTestMessage_TestContract This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestContract(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestContract()
}

//TestTestMessage_TestMethodID This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestMethodID(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.run().TestMethodID()
}

//TestTestMessage_TestData This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestData(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTestMessage_TestGasPrice This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestGasPrice(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTestMessage_TestSender This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestSender(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTestMessage_TestOrigins This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestOrigins(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestOrigins()
}

//TestTestMessage_TestInputReceipts This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestInputReceipts(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestInputReceipts()
}

//TestTestMessage_TestGetTransferToMe This is a method of MySuite
func (mysuit *MySuite) TestTestMessage_TestGetTransferToMe(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	test.run().TestGetTransferToMe()
}
