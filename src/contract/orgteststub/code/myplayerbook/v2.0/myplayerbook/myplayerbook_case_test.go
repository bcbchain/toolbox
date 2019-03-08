package myplayerbook

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

//TestMyPlayerBook_GetPlayer This is a method of MySuite
func (mysuit *MySuite) TestMyPlayerBook_GetPlayer(c *check.C) {
	//utest.Init(orgID)
	//contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	//test := NewTestObject(contractOwner)

	//TODO
}

//TestMyPlayerBook_RegisterName This is a method of MySuite
func (mysuit *MySuite) TestMyPlayerBook_RegisterName(c *check.C) {
	//utest.Init(orgID)
	//contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods)
	//test := NewTestObject(contractOwner)

	//TODO
}

//TestMyPlayerBook_MultiTypesParam is a method of MySuite
func (mysuit *MySuite) TestMyPlayerBook_MultiTypesParam(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, nil) // TODO 添了個nil消除語法錯誤，請改正
	test := NewTestObject(contractOwner)
	_ = test
	//TODO
}
