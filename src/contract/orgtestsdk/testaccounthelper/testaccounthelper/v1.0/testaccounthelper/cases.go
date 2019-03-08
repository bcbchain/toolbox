package testaccounthelper

import (
	"blockchain/smcsdk/sdk/types"
)

func (tah *TestAccountHelper) testAccountOf() {

	chainID := tah.sdk.Helper().GenesisHelper().ChainID()
	cases := []struct {
		addr   types.Address
		desc   string
		code   uint32
		errmsg string
	}{
		{chainID + "", "异常用例-地址只包括chainID", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{chainID + "a", "异常用例-地址太短", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{chainID + "JgaGConUyK81zibntUB", "异常用例-地址过短", types.ErrInvalidAddress, "Address checksum is error! "},
		{chainID + "JgaGConUyK81zibntUBJgaGConUyK81zibntUBjQ33PKctpk1K1G", "异常用例-地址过长", types.ErrInvalidAddress, "Address checksum is error! "},
		{chainID + "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "异常用例-无效的地址", types.ErrInvalidAddress, "Address checksum is error! "},
		{chainID + "你好，测试地址", "异常用例-无效的地址包括汉字", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{"JgaGConUyK81zibntUBjQ33PKctpk1K1G", "异常用例-chainID不对", types.ErrInvalidAddress, "Address chainID is error! "},
		{chainID + "JgaGConUyK81zibntUBjQ33PKctpk1K1G", "正常用例", types.CodeOK, ""},
		{chainID + "7pf84qStBtG7xjDtPzus15yQtMeWAP27M", "正常用例", types.CodeOK, ""},
	}
	for i, c := range cases {
		tah.runTestAccountOf(i, c.addr, c.desc, c.code, c.errmsg)
	}
}

func (tah *TestAccountHelper) testAccountOfPubKey() {

	cases := []struct {
		pk     types.PubKey
		desc   string
		code   uint32
		errmsg string
	}{
		{[]byte{}, "异常用例-pubkey为空", types.ErrInvalidParameter, "Invalid PubKey"},
		{[]byte{'0'}, "异常用例-pubkey长度为1", types.ErrInvalidParameter, "Invalid PubKey"},
		{[]byte("ababababababababababababababababa"), "异常用例-pubkey长度为33", types.ErrInvalidParameter, "Invalid PubKey"},
		{[]byte("abababababababababababababababab"), "正常用例", types.CodeOK, ""},
	}
	for i, c := range cases {
		tah.runTestAccountOfPubKey(i, c.pk, c.desc, c.code, c.errmsg)
	}
}
