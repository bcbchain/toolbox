package testblock

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

//TestTestBlock_TestChainID This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestChainID(c *check.C) {
	utest.Init(orgID)
	SetChecker(c)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestChainID()
}

//TestTestBlock_TestBlockHash This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestBlockHash(c *check.C) {
	utest.Init(orgID)
	SetChecker(c)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestBlockHash()
}

//TestTestBlock_TestHeight This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestHeight(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestHeight()
	//TODO
}

//TestTestBlock_TestTime This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestTime(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestTime()
	//TODO
}

//TestTestBlock_TestNow This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestNow(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestNow()
}

//TestTestBlock_TestNumTxs This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestNumTxs(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestNumTxs()
}

//TestTestBlock_TestDataHash This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestDataHash(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestDataHash()
}

//TestTestBlock_TestProposerAddress This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestProposerAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestProposerAddress()
}

//TestTestBlock_TestRewardAddress This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestRewardAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestRewardAddress()
}

//TestTestBlock_TestRandomNumber This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestRandomNumber(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestRandomNumber()
}

//TestTestBlock_TestVersion This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestVersion(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestVersion()
}

//TestTestBlock_TestLastBlockHash This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestLastBlockHash(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestLastBlockHash()
}

//TestTestBlock_TestCommitHash This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestCommitHash(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestLastCommitHash()
}

//TestTestBlock_TestAppHash This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestAppHash(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestAppHash()
}

//TestTestBlock_TestFee This is a method of MySuite
func (mysuit *MySuite) TestTestBlock_TestFee(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.run().TestLastFee()
}
