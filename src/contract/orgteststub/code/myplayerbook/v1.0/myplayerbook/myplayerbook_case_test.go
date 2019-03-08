package myplayerbook

import (
	ut "blockchain/smcsdk/utest"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (mysuit *MySuite) TestPlayerBook_GetPlayerAddr(c *C) {

}

func (mysuit *MySuite) TestPlayerBook_RegisterNameXid(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.run().Init()

	player := "bob"
	test.run().registerNameXid(player, 0)

}
