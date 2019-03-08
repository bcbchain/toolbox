package testblockchainhelper

import (
	"blockchain/smcsdk/sdk/types"
	"encoding/hex"
	"fmt"
)

func (t *TestblockChainHelper) testCalcAccountFromPubKey() {
	fmt.Println("\nTEST CASE: testCalcAccountFromPubKey")
	pk, _ := hex.DecodeString("F1EDF8F50848B8FA121A24E2A3A83CC5C8CBF85D6CE23A3A8413F46A717BEDA1")

	cases := []struct {
		pubKey types.PubKey
		desc   string
		code   uint32
		errMsg string
	}{
		{types.PubKey(pk), "正常用例", types.CodeOK, ""},
	}
	for i, c := range cases {
		t.runCalcAccountFromPubKey(i, c.pubKey, c.desc, c.code, c.errMsg)
	}
}

func (t *TestblockChainHelper) testCalcAccountFromName() {
	fmt.Println("\nTEST CASE: testCalcAccountFromName")

	cases := []struct {
		name   string
		orgID  string
		desc   string
		code   uint32
		errMsg string
	}{
		{"", "", "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		t.runCalcAccountFromName(i, c.name, c.orgID, c.desc, c.code, c.errMsg)
	}
}

func (t *TestblockChainHelper) testCalcContractAddressParamName() {
	fmt.Println("\nTEST CASE: testCalcContractAddressParamName")

	cases := []struct {
		name    string
		version string
		owner   types.Address
		desc    string
		code    uint32
		errMsg  string
	}{
		{},
	}

	for i, c := range cases {
		t.runCalcContractAddress(i, c.name, c.version, c.owner, c.desc, c.code, c.errMsg)
	}
}

func (t *TestblockChainHelper) testCalcContractAddressParamVersion() {
	fmt.Println("\nTEST CASE: testCalcContractAddressParamVersion")

	cases := []struct {
		name    string
		version string
		owner   types.Address
		desc    string
		code    uint32
		errMsg  string
	}{
		{},
	}

	for i, c := range cases {
		t.runCalcContractAddress(i, c.name, c.version, c.owner, c.desc, c.code, c.errMsg)
	}
}

func (t *TestblockChainHelper) testCalcContractAddressParamOwner() {
	fmt.Println("\nTEST CASE: testCalcContractAddressParamOwner")

	cases := []struct {
		name    string
		version string
		owner   types.Address
		desc    string
		code    uint32
		errMsg  string
	}{
		{},
	}

	for i, c := range cases {
		t.runCalcContractAddress(i, c.name, c.version, c.owner, c.desc, c.code, c.errMsg)
	}
}

func (t *TestblockChainHelper) testCalcOrgID() {
	fmt.Println("\nTEST CASE: testCalcOrgID")

	cases := []struct {
		name   string
		desc   string
		code   uint32
		errMsg string
	}{
		{},
	}

	for i, c := range cases {
		t.runCalcOrgID(i, c.name, c.desc, c.code, c.errMsg)
	}
}

func (t *TestblockChainHelper) testCheckAddress() {
	fmt.Println("\nTEST CASE: testCheckAddress")

	cases := []struct {
		addr   types.Address
		desc   string
		code   uint32
		errMsg string
	}{
		{"", "", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		err := t.sdk.Helper().BlockChainHelper().CheckAddress(c.addr)
		if c.code == types.CodeOK {
			Assert(err == nil)
		} else {
			AssertEquals(err.Error(), c.errMsg)
		}
	}
}

func (t *TestblockChainHelper) testGetBlock() {
	fmt.Println("\nTEST CASE: testGetBlock")

	cases := []struct {
		height int64
		desc   string
		code   uint32
		errMsg string
	}{
		{0, "", types.CodeOK, ""},
	}

	for i, c := range cases {
		fmt.Println(i, c)
		// todo 构造区块信息，检查获取的信息是否正确
	}
}
