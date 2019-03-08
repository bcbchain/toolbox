package mycrossmc

import (
	"math"
	"testing"

	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/utest"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

const registrationFee int64 = 100000000

func (mysuit *MySuite) TestMyCrossmc_Register(c *C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.run().transfer(bn.N(registrationFee)).Register(200)
	test.run().MultiParam()
	//	utest.AssertSDB("/mycrossmc/storedData", 200)
	//v, err := test.run().Get()
	//utest.AssertOK(err)
	//utest.Assert(v == 200)
}

func (mysuit *MySuite) TestMyCrossmc_Set(c *C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.run().transfer(bn.N(registrationFee))
	utest.Commit()
	utest.AssertBalance(utest.UTP.Helper().AccountHelper().AccountOf(utest.UTP.Message().Contract().Account()), utest.UTP.Helper().TokenHelper().TokenOfName("LOC").Name(), bn.N(registrationFee))

	test.run().transfer(bn.N(registrationFee)).Set(0)
	v := uint64(0)
	utest.AssertSDB("/mycrossmc/storedData", v)
	//utest.AssertBalance(utest.UTP.Helper().AccountHelper().AccountOf(utest.UTP.Contract().Account()), utest.UTP.Helper().TokenHelper().Token().Name(),bn.N(registrationFee))
	v, err := test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == 0)

	test.run().transfer(bn.N(registrationFee)).Set(2000348989)
	utest.AssertSDB("/mycrossmc/storedData", 2000348989)
	v, err = test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == 2000348989)

	test.run().transfer(bn.N(registrationFee)).Set(math.MaxInt64)
	utest.AssertSDB("/mycrossmc/storedData", math.MaxInt64)
	v, err = test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == math.MaxInt64)
}

func (mysuit *MySuite) TestMyCrossmc_Get(c *C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	test.run().transfer(bn.N(registrationFee)).Set(0)
	v, err := test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == 0)

	test.run().transfer(bn.N(registrationFee)).Set(1)
	v, err = test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == 1)

	test.run().transfer(bn.N(registrationFee)).Set(2)
	v, err = test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == 2)

	test.run().transfer(bn.N(registrationFee)).Set(3)
	test.run().transfer(bn.N(registrationFee)).Set(4)
	test.run().transfer(bn.N(registrationFee)).Set(6)
	test.run().transfer(bn.N(registrationFee)).Set(9)
	test.run().transfer(bn.N(registrationFee)).Set(10)
	test.run().transfer(bn.N(registrationFee)).Set(11)
	v, err = test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == 11)

	test.run().transfer(bn.N(registrationFee)).Set(math.MaxInt64)
	v, err = test.run().Get()
	utest.AssertOK(err)
	utest.Assert(v == math.MaxInt64)
}

func TestMyCrossmc_Calc(t *testing.T) {
	//fmt.Println(hex.EncodeToString(algorithm.CalcMethodId("RegisterName(string)(types.Error)")))
	//fmt.Println(algorithm.CalcContractAddress("local",
	//	"localETK7Zh9hNSPrEKdmCgnHDtFPatcs9WwVL","myplayerbook", "2.0"))
	//fmt.Println(algorithm.CalcContractAddress("local",
	//	"","myplayerbook", ""))
}
