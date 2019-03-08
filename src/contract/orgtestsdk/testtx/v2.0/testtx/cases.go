package testtx

import (
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl"
	"blockchain/smcsdk/sdkimpl/object"
	"fmt"
)

func (t *TestTx) testNote() {
	fmt.Println("\nTest Case: testNote()")

	cases := []struct {
		note   string
		desc   string
		code   uint32
		errMsg string
	}{
		{"testNote", "正常用例", types.CodeOK, ""},
		{"", "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		tx := object.NewTx(t.sdk, c.note, 0, 0, t.sdk.Message().Sender().Address())
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetTx(tx)
		Assert(t.sdk.Tx().Note() == c.note)
		printPass()
	}
}

func (t *TestTx) testGasLimit() {
	fmt.Println("\nTest Case: testGasLimit()")

	cases := []struct {
		gasLimit int64
		desc     string
		code     uint32
		errMsg   string
	}{
		{0, "正常用例", types.CodeOK, ""},
		{1, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		tx := object.NewTx(t.sdk, t.sdk.Tx().Note(), c.gasLimit, 0, t.sdk.Message().Sender().Address())
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetTx(tx)
		Assert(t.sdk.Tx().GasLimit() == c.gasLimit)
		printPass()
	}
}

func (t *TestTx) testGasLeft() {
	fmt.Println("\nTest Case: testGasLeft()")

	cases := []struct {
		gasLeft int64
		desc    string
		code    uint32
		errMsg  string
	}{
		{0, "正常用例", types.CodeOK, ""},
		{1, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		tx := object.NewTx(t.sdk, t.sdk.Tx().Note(), t.sdk.Tx().GasLimit(), c.gasLeft, t.sdk.Message().Sender().Address())
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetTx(tx)
		Assert(t.sdk.Tx().GasLeft() == c.gasLeft)
		printPass()
	}
}

func (t *TestTx) testSigner() {
	fmt.Println("\nTest Case: testSigner()")

	cases := []struct {
		signer string
		desc   string
		code   uint32
		errMsg string
	}{
		{"", "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		tx := object.NewTx(t.sdk, t.sdk.Tx().Note(), t.sdk.Tx().GasLimit(), t.sdk.Tx().GasLeft(), c.signer)
		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetTx(tx)
		Assert(t.sdk.Tx().Signer().Address() == c.signer)
		printPass()
	}
}
