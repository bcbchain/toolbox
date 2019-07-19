package testcontracthelper

import (
	"blockchain/algorithm"
	"blockchain/smcsdk/sdk/types"
	"fmt"
	"strings"
)

func (t *TestContractHelper) runCasesContractOfName() {
	fmt.Println("\nTest Case: 根据名字构造合约对象，runCasesContractOfName()")

	chainID := t.sdk.Helper().GenesisHelper().ChainID()

	var cases = []struct {
		name    string
		isvalid bool //是否返回有效对象
		desc    string
		errcode uint32
		errmsg  string
	}{
		{name: "", desc: "名字为空", errcode: types.CodeOK},
		{name: "h", desc: "名字为 h", errcode: types.CodeOK},
		{name: chainID, desc: "名字为 " + chainID, errcode: types.CodeOK},
		{name: chainID + "h", desc: "名字为chainID+h", errcode: types.CodeOK},
		{name: "hjsd43878@#*9", desc: "名字为 hjsd43878@#*9", errcode: types.CodeOK},
		{name: chainID + "KG7Y7hWLhjxiBNZ8UBgja4ocLcSHKW4b4", desc: "名字一个地址", errcode: types.CodeOK},
		{name: "tokenbasica", desc: "名字为字符tokenbasica", errcode: types.CodeOK},
		{name: "tokenbasicab4", desc: "名字为tokenbasicab4", errcode: types.CodeOK},
		{name: t.sdk.Helper().GenesisHelper().Token().Name(), desc: "本币名字", errcode: types.CodeOK},
		{name: t.sdk.Helper().GenesisHelper().Contracts()[0].Name(), isvalid: true, desc: "系统合约名字", errcode: types.CodeOK},
		{name: strings.ToUpper(t.sdk.Helper().GenesisHelper().Contracts()[0].Name()), desc: "系统合约名字全大写", errcode: types.CodeOK},
		{name: t.sdk.Message().Contract().Name(), isvalid: true, desc: "本合约名字", errcode: types.CodeOK},
		{name: strings.ToUpper(t.sdk.Message().Contract().Name()), desc: "本合约名字全大写", errcode: types.CodeOK},
	}
	for i, c := range cases {
		printTestCase(i+1, c.desc)
		icontract, err := t.contractOfName(c.name)
		AssertError(err, c.errcode, c.errmsg)
		if c.isvalid == false {
			Assert(icontract == nil)
		} else if c.errcode == types.CodeOK {
			AssertEquals(strings.ToLower(icontract.Name()), strings.ToLower(c.name))
		}
		printPass()
	}
}

func (t *TestContractHelper) runCasesContractOfAddress() {
	fmt.Println("\nTest Case: 根据地址构造合约对象，runCasesContractOfAddress()")

	chainID := t.sdk.Helper().GenesisHelper().ChainID()
	tempAddr := algorithm.CalcContractAddress(chainID, t.sdk.Helper().GenesisHelper().Token().Owner().Address(), "test", "1.0")

	var cases = []struct {
		addr    types.Address
		isvalid bool //是否返回有效对象
		desc    string
		errcode uint32
		errmsg  string
	}{
		{addr: "", desc: "地址为空", errcode: types.ErrInvalidAddress, errmsg: "Address chainID is error! "},
		{addr: "h", desc: "地址为 h", errcode: types.ErrInvalidAddress, errmsg: "Address chainID is error! "},
		{addr: chainID + "h", desc: "地址为chainID+h", errcode: types.ErrInvalidAddress, errmsg: "Base58Addr parse error! "},
		{addr: "hjsd43878@#*9", desc: "地址为 hjsd43878@#*9", errcode: types.ErrInvalidAddress, errmsg: "Address chainID is error! "},
		{addr: chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4", desc: "地址被截短", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b8", desc: "地址尾数被修改", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: chainID + "KG7Y7hWLhjxiBNZ8UBgja40cLcSHKW4b4", desc: "地址中间o转0", errcode: types.ErrInvalidAddress, errmsg: "Base58Addr parse error! "},
		{addr: tempAddr + "a", desc: "地址增加一个字符a", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: tempAddr + "ab4", desc: "增加三个字符ab4", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: tempAddr, desc: "地址为不存在的合约地址：" + tempAddr, errcode: types.CodeOK},
		{addr: t.sdk.Helper().GenesisHelper().Token().Address(), isvalid: true, desc: "本币合约地址", errcode: types.CodeOK},
		{addr: t.sdk.Message().Contract().Address(), isvalid: true, desc: "本合约地址", errcode: types.CodeOK},
	}
	for i, c := range cases {
		printTestCase(i+1, c.desc)
		icontract, err := t.contractOfAddress(c.addr)
		AssertError(err, c.errcode, c.errmsg)
		if c.isvalid == false {
			Assert(icontract == nil)
		} else if c.errcode == types.CodeOK {
			AssertEquals(icontract.Address(), c.addr)
		}
		printPass()
	}
}

func (t *TestContractHelper) runCasesContractOfToken() {
	fmt.Println("\nTest Case: 根据代币地址构造合约对象，runCasesContractOfToken()")

	chainID := t.sdk.Helper().GenesisHelper().ChainID()
	tempAddr := algorithm.CalcContractAddress(chainID, t.sdk.Helper().GenesisHelper().Token().Owner().Address(), "test", "1.0")

	hasToken := false
	if t.sdk.Message().Contract().Token() != "" {
		hasToken = true
	}

	var cases = []struct {
		addr    types.Address
		isvalid bool //是否返回有效对象
		desc    string
		errcode uint32
		errmsg  string
	}{
		{addr: "", desc: "地址为空", errcode: types.ErrInvalidAddress, errmsg: "Address chainID is error! "},
		{addr: "h", desc: "地址为 h", errcode: types.ErrInvalidAddress, errmsg: "Address chainID is error! "},
		{addr: chainID + "h", desc: "地址为chainID+h", errcode: types.ErrInvalidAddress, errmsg: "Base58Addr parse error! "},
		{addr: "hjsd43878@#*9", desc: "地址为 hjsd43878@#*9", errcode: types.ErrInvalidAddress, errmsg: "Address chainID is error! "},
		{addr: chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4", desc: "地址被截短", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b8", desc: "地址尾数被修改", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: chainID + "KG7Y7hWLhjxiBNZ8UBgja40cLcSHKW4b4", desc: "地址中间o转0", errcode: types.ErrInvalidAddress, errmsg: "Base58Addr parse error! "},
		{addr: tempAddr + "a", desc: "地址增加一个字符a", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: tempAddr + "ab4", desc: "增加三个字符ab4", errcode: types.ErrInvalidAddress, errmsg: "Address checksum is error! "},
		{addr: tempAddr, desc: "地址为不存在的合约地址：" + tempAddr, errcode: types.CodeOK},
		{addr: t.sdk.Helper().GenesisHelper().Token().Address(), isvalid: true, desc: "本币合约地址", errcode: types.CodeOK},
		{addr: t.sdk.Message().Contract().Address(), isvalid: hasToken, desc: "本合约地址", errcode: types.CodeOK},
	}
	for i, c := range cases {
		printTestCase(i+1, c.desc)
		icontract, err := t.contractOfToken(c.addr)
		AssertError(err, c.errcode, c.errmsg)
		if c.isvalid == false {
			Assert(icontract == nil)
		} else if c.errcode == types.CodeOK {
			AssertEquals(icontract.Address(), c.addr)
		}
		printPass()
	}
	//合约已注册代币，测试代币地址
	if hasToken {
		printTestCase(len(cases)+1, "本合约代币地址")
		taddr := t.sdk.Message().Contract().Token()
		icontract, err := t.contractOfToken(taddr)
		AssertError(err, types.CodeOK, "")
		AssertEquals(icontract.Address(), taddr)
		printPass()
	}
}
