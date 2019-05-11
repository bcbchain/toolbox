package testmessage

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl"
	"blockchain/smcsdk/sdkimpl/object"
	"encoding/hex"
	"fmt"
)

//nolint unhandled
func (t *TestMessage) testContract() {
	fmt.Println("\nTest Case: testContract()")
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	tempAddr := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)

	conMethods := []std.Method{
		{"0x1234", 5000, "Transfer(types.Address,bn.Number)"},
	}

	token := t.sdk.Helper().TokenHelper().RegisterToken("test", "test", bn.N(1E15), true, true)

	cases := []struct {
		ContractOwnerAddr    types.Address
		ContractName         string
		ContractVersion      string
		ContractPrefix       string
		ContractCodeHash     types.Hash
		ContractEffectHeight int64
		ContractLoseHeight   int64
		ContractMethods      []std.Method
		ContractInterface    []std.Method
		ContractToken        types.Address
		desc                 string
		code                 uint32
		errMsg               string
	}{
		{tempAddr.Address(), "contract", "1.0", "", []byte("hello"), 100, 200, conMethods, conMethods, token.Address(), "正常用例", types.CodeOK, ""},
		{"", "", "", "", []byte{}, 0, 0, nil, nil, "", "正常用例-contract所有值都为空", types.CodeOK, ""},
		{tempAddr.Address(), "", "", "", []byte{}, 0, 0, nil, nil, "", "正常用例-合约地址不为空", types.CodeOK, ""},
		{"", "test", "", "", []byte{}, 0, 0, nil, nil, "", "正常用例-合约名字不为空", types.CodeOK, ""},
		{"", "", "2.0", "", []byte{}, 0, 0, nil, nil, "", "正常用例-合约版本号不为空", types.CodeOK, ""},
		{"", "", "", "a", []byte{}, 0, 0, nil, nil, "", "正常用例-合约前缀不为空", types.CodeOK, ""},
		{"", "", "", "a", []byte("hello"), 0, 0, nil, nil, "", "正常用例-合约codeHash不为空", types.CodeOK, ""},
		{"", "", "", "a", []byte{}, 10, 0, nil, nil, "", "正常用例-合约生效高度不为空", types.CodeOK, ""},
		{"", "", "", "a", []byte{}, 0, 10, nil, nil, "", "正常用例-合约失效高度不为空", types.CodeOK, ""},
		{"", "", "", "a", []byte{}, 0, 0, conMethods, nil, "", "正常用例-合约Methos不为空", types.CodeOK, ""},
		{"", "", "", "a", []byte{}, 0, 0, nil, conMethods, "", "正常用例-合约Interface不为空", types.CodeOK, ""},
		{"", "", "", "a", []byte{}, 0, 0, nil, conMethods, token.Address(), "正常用例-合约token地址不为空", types.CodeOK, ""},
	}
	for i, c := range cases {
		contract := object.NewContract(t.sdk,
			t.sdk.Helper().GenesisHelper().OrgID(),
			c.ContractOwnerAddr,
			c.ContractName,
			c.ContractVersion,
			c.ContractPrefix,
			c.ContractCodeHash,
			c.ContractEffectHeight,
			c.ContractLoseHeight,
			c.ContractMethods,
			c.ContractInterface,
			c.ContractToken)

		m := object.NewMessage(t.sdk, contract, "123123", nil, t.sdk.Message().Sender().Address(), nil, nil)
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetMessage(m)

		printTestCase(i, c.desc)
		Assert(t.sdk.Message().Contract().Owner() == c.ContractOwnerAddr)
		Assert(t.sdk.Message().Contract().Name() == c.ContractName)
		Assert(t.sdk.Message().Contract().Version() == c.ContractVersion)
		Assert(t.sdk.Message().Contract().KeyPrefix() == c.ContractPrefix)
		Assert(t.sdk.Message().Contract().EffectHeight() == c.ContractEffectHeight)
		Assert(t.sdk.Message().Contract().LoseHeight() == c.ContractLoseHeight)
		Assert(t.sdk.Message().Contract().Token() == c.ContractToken)

		for index, v := range c.ContractCodeHash {
			Assert(v == t.sdk.Message().Contract().CodeHash()[index])
		}

		for index, v := range c.ContractMethods {
			Assert(v == t.sdk.Message().Contract().Methods()[index])
		}

		for index, v := range c.ContractInterface {
			Assert(v == t.sdk.Message().Contract().Interfaces()[index])
		}
		printPass()
	}

}

func (t *TestMessage) testMethodID() {
	fmt.Println("\nTest Case: testMethodID()")
	cases := []struct {
		MethodID string
		desc     string
		code     uint32
		errMsg   string
	}{
		{},
	}
	for i, c := range cases {
		printTestCase(i, c.desc)

		m := object.NewMessage(t.sdk, t.sdk.Message().Contract(), c.MethodID, t.sdk.Message().Items(), t.sdk.Message().Sender().Address(), t.sdk.Message().Origins(), t.sdk.Message().InputReceipts())
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetMessage(m)
		Assert(t.sdk.Message().MethodID() == c.MethodID)

		printPass()
	}
}

func (t *TestMessage) testItems() {
	fmt.Println("\nTest Case: testMethodID()")
	cases := []struct {
		Items  []types.HexBytes
		desc   string
		code   uint32
		errMsg string
	}{
		{},
	}
	for i, c := range cases {
		printTestCase(i, c.desc)

		m := object.NewMessage(t.sdk, t.sdk.Message().Contract(), t.sdk.Message().MethodID(), c.Items, t.sdk.Message().Sender().Address(), t.sdk.Message().Origins(), t.sdk.Message().InputReceipts())
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetMessage(m)
		//Assert(t.sdk.Message().Items == c.Items)

		printPass()
	}
}

func (t *TestMessage) testGasPrice() {
	//fmt.Println("\nTest Case: testGasPrice()")
	//cases := []struct {
	//	GasPrice int64
	//	desc     string
	//	code     uint32
	//	errMsg   string
	//}{
	//	{},
	//}
	//for i, c := range cases {
	//	printTestCase(i, c.desc)
	//
	//	m := object.NewMessage(t.sdk, t.sdk.Message().Contract(), t.sdk.Message().MethodID(), c.Items, t.sdk.Message().Sender().Address(), t.sdk.Message().Origins(), t.sdk.Message().InputReceipts())
	//	smc := t.sdk.(*sdkimpl.SmartContract)
	//	smc.SetMessage(m)
	//	//Assert(t.sdk.Message().Items == c.Items)
	//
	//	printPass()
	//}
}

func (t *TestMessage) testSender() {
	fmt.Println("\nTest Case: testSender()")
	cases := []struct {
		Sender types.Address
		desc   string
		code   uint32
		errMsg string
	}{
		{},
	}
	for i, c := range cases {
		printTestCase(i, c.desc)

		m := object.NewMessage(t.sdk, t.sdk.Message().Contract(), t.sdk.Message().MethodID(), t.sdk.Message().Items(), c.Sender, t.sdk.Message().Origins(), t.sdk.Message().InputReceipts())
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetMessage(m)
		Assert(t.sdk.Message().Sender().Address() == c.Sender)

		printPass()
	}
}

func (t *TestMessage) testOrigins() {
	fmt.Println(t.sdk.Message().Origins())
	fmt.Println(t.sdk.Helper().GenesisHelper().Contracts())
	fmt.Println(t.sdk.Helper().GenesisHelper().Token().Owner())
	//for _, v := range t.sdk.Helper().GenesisHelper().Contracts() {
	//	fmt.Println(v.Address())
	//}
}

func (t *TestMessage) testInputReceipts() {
	fmt.Println(t.sdk.Message().InputReceipts())
}

func (t *TestMessage) testGetTransferToMe() {
	fmt.Println(t.sdk.Message().GetTransferToMe())
}
