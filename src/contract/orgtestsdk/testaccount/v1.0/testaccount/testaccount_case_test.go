package testaccount

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

//TestTestAccount_TestTransfer This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestTransfer(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	test.setSender(contractOwner).InitChain()
	SetChecker(c)
	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestTransfer()
	utest.AssertOK(err)
}

//TestTestAccount_TestTransferByToken This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestTransferByToken(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	//为合约账户转入本币
	owner := test.obj.sdk.Helper().GenesisHelper().Token().Owner()
	ownacc := test.obj.sdk.Helper().AccountHelper().AccountOf(owner)
	test.run().setSender(ownacc).transfer(bn.N(1E12))

	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestTransferByToken()
	utest.AssertOK(err)

}

//TestTestAccount_TestTransferByName This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestTransferByName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	//为合约账户转入本币
	owner := test.obj.sdk.Helper().GenesisHelper().Token().Owner()
	ownacc := test.obj.sdk.Helper().AccountHelper().AccountOf(owner)
	test.run().setSender(ownacc).transfer(bn.N(1E12))

	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestTransferByName()
	utest.AssertOK(err)
}

//TestTestAccount_TestTransferBySymbol This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestTransferBySymbol(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	//为合约账户转入本币
	owner := test.obj.sdk.Helper().GenesisHelper().Token().Owner()
	ownacc := test.obj.sdk.Helper().AccountHelper().AccountOf(owner)
	test.run().setSender(ownacc).transfer(bn.N(1E12))

	test.setSender(contractOwner).InitChain()
	SetChecker(c)

	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestTransferBySymbol()
	utest.AssertOK(err)
}

//TestTestAccount_Transfer This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_Transfer(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO

}

//TestTestAccount_TestAccoutPubKey This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestAccoutPubKey(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
	err := test.run().setSender(contractOwner).TestAccountPubKey()
	utest.AssertOK(err)
}

//TestTestAccount_TestBalance This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestBalance(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestBalance()
	utest.AssertOK(err)
}

//TestTestAccount_TestBalanceOfToken This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestBalanceOfToken(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestBalanceOfToken()
	utest.AssertOK(err)
}

//TestTestAccount_TestBalanceOfName This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestBalanceOfName(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestBalanceOfName()
	utest.AssertOK(err)
}

//TestTestAccount_TestBalanceOfSymbol This is a method of MySuite
func (mysuit *MySuite) TestTestAccount_TestBalanceOfSymbol(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()

	test.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, bn.N(9E18), true, true)
	err := test.run().setSender(contractOwner).TestBalanceOfSymbol()
	utest.AssertOK(err)
}

//func (mysuit *MySuite) TestTestAccount_TestDecode(c *check.C) {
//	fmt.Println(hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B108"))
//	fmt.Println(hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083G"))
//	fmt.Println(hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083"))
//	fmt.Println(hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B108345"))
//}
