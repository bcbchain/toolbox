package mymixtype

import (
	"blockchain/smcsdk/utest"
	"gopkg.in/check.v1"
	"testing"
)

//Test is a function
func Test(t *testing.T) { check.TestingT(t) }

//MySuite is a struct
type MySuite struct{}

var _ = check.Suite(&MySuite{})

var (
	orgId = ""
)

//TestMymixtype_Basic is a method of MySuite
func (mysuit *MySuite) TestMymixtype_Basic(c *check.C) {
	utest.Init(orgId)
	contractOwner := utest.DeployContract(c, contractName, orgId, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.Basic(nil)
}

//TestMymixtype_Slice is a method of MySuite
func (mysuit *MySuite) TestMymixtype_Slice(c *check.C) {
	utest.Init(orgId)
	contractOwner := utest.DeployContract(c, contractName, orgId, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.Slice(nil)
}

//TestMymixtype_MapString is a method of MySuite
func (mysuit *MySuite) TestMymixtype_MapString(c *check.C) {
	utest.Init(orgId)
	contractOwner := utest.DeployContract(c, contractName, orgId, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.MapString(nil)
}

//TestMymixtype_MapOther is a method of MySuite
func (mysuit *MySuite) TestMymixtype_MapOther(c *check.C) {
	utest.Init(orgId)
	contractOwner := utest.DeployContract(c, contractName, orgId, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.MapOther(nil)
}

//TestMymixtype_Complex is a method of MySuite
func (mysuit *MySuite) TestMymixtype_Complex(c *check.C) {
	utest.Init(orgId)
	contractOwner := utest.DeployContract(c, contractName, orgId, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.Complex(nil)
}

//TestMymixtype_Long is a method of MySuite
func (mysuit *MySuite) TestMymixtype_Long(c *check.C) {
	utest.Init(orgId)
	contractOwner := utest.DeployContract(c, contractName, orgId, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.Long(nil)
}

//TestMymixtype_MapSlice is a method of MySuite
func (mysuit *MySuite) TestMymixtype_MapSlice(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	//TODO
	test.obj.MapSlice(nil)
}
