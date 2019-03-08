package mybasictype

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/utest"
	"gopkg.in/check.v1"
	"testing"
)

//Test: This is a function
func Test(t *testing.T) { check.TestingT(t) }

//MySuite: This is a struct
type MySuite struct{}

var _ = check.Suite(&MySuite{})

//TestBasicType_EchoAddress: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoAddress(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoAddress("")
}

//TestBasicType_EchoHash: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoHash(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoHash([]byte(""))
}

//TestBasicType_EchoHexBytes: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoHexBytes(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoHexBytes([]byte(""))
}

//TestBasicType_EchoPubKey: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoPubKey(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoPubKey([]byte(""))
}

//TestBasicType_EchoNumber: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoNumber(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoNumber(bn.N(0))
}

//TestBasicType_EchoInt: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoInt(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoInt(0)
}

//TestBasicType_EchoInt8: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoInt8(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoInt8(0)
}

//TestBasicType_EchoInt16: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoInt16(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoInt16(0)
}

//TestBasicType_EchoInt32: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoInt32(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoInt32(0)
}

//TestBasicType_EchoInt64: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoInt64(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoInt64(0)
}

//TestBasicType_EchoUint: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoUint(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoUint(0)
}

//TestBasicType_EchoUint8: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoUint8(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoUint8(0)
}

//TestBasicType_EchoUint16: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoUint16(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoUint16(0)
}

//TestBasicType_EchoUint32: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoUint32(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoUint32(0)
}

//TestBasicType_EchoUint64: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoUint64(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoUint64(0)
}

//TestBasicType_EchoBool: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoBool(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoBool(true)
}

//TestBasicType_EchoByte: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoByte(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoByte('a')
}

//TestBasicType_EchoBytes: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoBytes(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoBytes([]byte("abc"))
}

//TestBasicType_EchoMap: This is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	map_ := make(map[string]int)
	map_["abc"] = 1

	test.obj.EchoMap(map_)
}

//TestBasicType_EchoMap1 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap1(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}

//TestBasicType_EchoMap2 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap2(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}

//TestBasicType_EchoMap3 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap3(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}

//TestBasicType_EchoMap4 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap4(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}

//TestBasicType_EchoMap5 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap5(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}

//TestBasicType_EchoMap6 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap6(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}

//TestBasicType_EchoMap7 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap7(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}

//TestBasicType_EchoMap8 is a method of MySuite
func (mysuit *MySuite) TestBasicType_EchoMap8(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	test := NewTestObject(contractOwner)

	test.obj.EchoMap1(nil)
}
