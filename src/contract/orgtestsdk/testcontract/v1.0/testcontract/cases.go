package testcontract

import (
	"blockchain/algorithm"
	"blockchain/smcsdk/sdk/types"
	"encoding/hex"
	"fmt"
)

var (
	contractMethods    = []string{"Transfer(types.Address,bn.Number)", "TestAddress()", "TestAccount()", "TestOwner()", "TestName()", "TestVersion()", "TestCodeHash()", "TestEffectHeight()", "TestLoseHeight()", "TestKeyPrefix()", "TestMethods()", "TestInterfaces()", "TestToken()", "TestOrgID()", "TestSetOwner()"}
	contractInterfaces = []string{"Transfer(types.Address,bn.Number)", "TestAddress()", "TestName()", "TestLoseHeight()", "TestMethods()", "TestInterfaces()", "TestOrgID()", "TestSetOwner()"}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxew"
)

func (t *TestContract) testAddress() {
	fmt.Println("\nTest Case: TestAddress()")
	printTestCase(0, "计算合约地址，并对比")
	addr := algorithm.CalcContractAddress(
		t.sdk.Helper().GenesisHelper().ChainID(),
		t.sdk.Message().Contract().Owner().Address(),
		t.sdk.Message().Contract().Name(),
		t.sdk.Message().Contract().Version())
	AssertEquals(addr, t.sdk.Message().Contract().Address())
	printPass()
}

func (t *TestContract) testAccount() {
	fmt.Println("\nTest Case: TestAccount()")

	printTestCase(0, "计算合约账户地址，并对比")
	addr := algorithm.CalcContractAddress(
		t.sdk.Helper().GenesisHelper().ChainID(),
		"",
		t.sdk.Message().Contract().Name(),
		"")
	AssertEquals(addr, t.sdk.Message().Contract().Account().Address())
	printPass()
}

func (t *TestContract) testOwner() {
	fmt.Println("\nTest Case: TestOwner()")
	printTestCase(0, "Message Sender是合约Owner")
	AssertEquals(t.sdk.Message().Sender().Address(), t.sdk.Message().Contract().Owner().Address())
	printPass()
}

func (t *TestContract) testName() {
	fmt.Println("\nTest Case: TestName()")
	printTestCase(0, "合约名称为testcontract")
	AssertEquals("testcontract", t.sdk.Message().Contract().Name())
	printPass()
}

func (t *TestContract) testVersion() {
	fmt.Println("\nTest Case: TestVersion()")
	printTestCase(0, "合约版本号为1.0")
	AssertEquals("1.0", t.sdk.Message().Contract().Version())
	printPass()
}

func (t *TestContract) testCodeHash() {
	fmt.Println("\nTest Case: TestCodeHash()")
	printTestCase(0, "")
	//AssertEquals(int64(1), t.sdk.Message().Contract().EffectHeight())
	printPass()
}

func (t *TestContract) testEffectHeight() {
	fmt.Println("\nTest Case: TestEffectHeight()")

	printTestCase(0, "初始版本合约生效高度为1")
	AssertEquals(int64(1), t.sdk.Message().Contract().EffectHeight())
	printPass()
}

func (t *TestContract) testLoseHeight() {
	fmt.Println("\nTest Case: TestLoseHeight()")

	printTestCase(0, "初始版本合约失效高度为0")
	AssertEquals(int64(0), t.sdk.Message().Contract().LoseHeight())
	printPass()
}

func (t *TestContract) testKeyPrefix() {
	fmt.Println("\nTest Case: TestKeyPrefix()")

	printTestCase(0, "KeyPrefix为合约名称")
	AssertEquals("/testcontract", t.sdk.Message().Contract().KeyPrefix())
	printPass()
}

func (t *TestContract) testMethods() {
	fmt.Println("\nTest Case: TestMethods()")

	printTestCase(0, "对比合约Method数量")
	ms := t.sdk.Message().Contract().Methods()
	AssertEquals(len(contractMethods), len(ms))
	printPass()
	printTestCase(1, "检查各个Method原型")
	for i, m := range ms {
		AssertEquals(contractMethods[i], m.ProtoType)
	}
	printPass()

	printTestCase(2, "检查各个Method 的MethodID")
	for _, m := range ms {
		id := algorithm.ConvertMethodID(algorithm.CalcMethodId(m.ProtoType))
		AssertEquals(id, m.MethodID)
	}
	printPass()

	printTestCase(3, "检查各个Method 的 gas")
	for _, m := range ms {
		AssertEquals(int64(200), m.Gas)
	}
	printPass()
}

func (t *TestContract) testInterfaces() {
	fmt.Println("\nTest Case: TestInterfaces()")

	printTestCase(0, "对比合约Interface数量")
	ms := t.sdk.Message().Contract().Interfaces()
	AssertEquals(len(contractInterfaces), len(ms))
	printPass()
	printTestCase(1, "检查各个Interface原型")
	for i, m := range ms {
		AssertEquals(contractInterfaces[i], m.ProtoType)
	}
	printPass()

	printTestCase(2, "检查各个Interface 的MethodID")
	for _, m := range ms {
		id := algorithm.ConvertMethodID(algorithm.CalcMethodId(m.ProtoType))
		AssertEquals(id, m.MethodID)
	}
	printPass()
	printTestCase(3, "检查各个Interface的 gas")
	for _, m := range ms {
		AssertEquals(int64(100), m.Gas)
	}
	printPass()
}

func (t *TestContract) testToken() {
	fmt.Println("\nTest Case: TestToken()")
	printTestCase(0, "未注册代币的合约，代币为空")
	AssertEquals("", t.sdk.Message().Contract().Token())
	printPass()
}

func (t *TestContract) testOrgID() {
	fmt.Println("\nTest Case: TestOrgID()")

	printTestCase(0, "对比组织ID")
	AssertEquals(orgID, t.sdk.Message().Contract().OrgID())
	printPass()
}

//nolint unhandled
func (t *TestContract) testSetOwner() {
	fmt.Println("\nTest Case: TestSetOwner()")
	printTestCase(0, "合约原始Owner为Message Sender")
	AssertEquals(t.sdk.Message().Sender().Address(), t.sdk.Message().Contract().Owner().Address())
	printPass()

	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey).Address()
	chainID := t.sdk.Helper().GenesisHelper().ChainID()

	var cases = []struct {
		newowner types.Address
		desc     string
		errcode  uint32
		errmsg   string
	}{
		{desc: "新Owner地址为空", errcode: types.ErrInvalidAddress, errmsg: "Address chainID is error! "},
		{"h", "新Owner地址为 h", types.ErrInvalidAddress, "Address chainID is error! "},
		{chainID + "h", "新Owner地址为chainID+h", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{"hjsd43878@#*9", "新Owner地址为 hjsd43878@#*9", types.ErrInvalidAddress, "Address chainID is error! "},
		{chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4", "新Owner地址被截短", types.ErrInvalidAddress, "Address checksum is error! "},
		{chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b8", "新Owner地址尾数被修改", types.ErrInvalidAddress, "Address checksum is error! "},
		{chainID + "KG7Y7hWLhjxiBNZ8UBgja40cLcSHKW4b4", "新Owner地址中间o转0", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{recv + "a", "新Owner地址增加一个字符a", types.ErrInvalidAddress, "Address checksum is error! "},
		{newowner: recv + "ab4", desc: "新Owner地址增加三个字符ab4", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{newowner: recv, desc: "新Owner地址为" + recv, errcode: types.CodeOK},
	}
	for i, c := range cases {
		printTestCase(i+1, c.desc)
		err := t.setOwner(c.newowner)
		AssertError(err, c.errcode, c.errmsg)
		if c.errcode == types.CodeOK {
			AssertEquals(c.newowner, t.sdk.Message().Contract().Owner().Address())
			if t.sdk.Message().Contract().Token() != "" {
				token := t.sdk.Helper().TokenHelper().TokenOfAddress(t.sdk.Message().Contract().Token())
				AssertEquals(c.newowner, token.Owner())
			}
		}
		printPass()
	}
}
