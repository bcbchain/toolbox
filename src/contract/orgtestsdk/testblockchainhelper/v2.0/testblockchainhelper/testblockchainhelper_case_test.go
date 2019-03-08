package testblockchainhelper

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

//TestTestblockChainHelper_TestCalcAccountFromPubKey This is a method of MySuite
func (mysuit *MySuite) TestTestblockChainHelper_TestCalcAccountFromPubKey(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestCalcAccountFromPubKey()

	//TODO
}

//TestTestblockChainHelper_TestCalcAccountFromName This is a method of MySuite
func (mysuit *MySuite) TestTestblockChainHelper_TestCalcAccountFromName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestCalcAccountFromName()

	//TODO
}

//TestTestblockChainHelper_TestCalcContractAddress This is a method of MySuite
func (mysuit *MySuite) TestTestblockChainHelper_TestCalcContractAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestCalcContractAddress()

	//TODO
}

//TestTestblockChainHelper_TestCalcOrgID This is a method of MySuite
func (mysuit *MySuite) TestTestblockChainHelper_TestCalcOrgID(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestCalcOrgID()

	//TODO
}

//TestTestblockChainHelper_TestCheckAddress This is a method of MySuite
func (mysuit *MySuite) TestTestblockChainHelper_TestCheckAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestCheckAddress()
	//TODO
}

//TestTestblockChainHelper_TestGetBlock This is a method of MySuite
func (mysuit *MySuite) TestTestblockChainHelper_TestGetBlock(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestGetBlock()

	//TODO
}
