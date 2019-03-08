package testaccount

import (
	"blockchain/algorithm"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"encoding/hex"
	"fmt"
)

//nolint unhandled
//transferTestTransferValue tuses sender and contract account to test transfer value
// this contains all of test cases of "value"
func (t *TestAccount) transferTestTransferValue(byType string) {
	fmt.Println("\nTEST CASE: 测试转账金额")
	fmt.Printf(" transferTestTransferValue(%v)\n", byType)

	accAddr := t.sdk.Message().Contract().Account()
	acct := t.sdk.Helper().AccountHelper().AccountOf(accAddr)
	sender := t.sdk.Message().Sender()

	tokenAddr := t.sdk.Helper().TokenHelper().Token().Address()
	tokenSymbol := t.sdk.Helper().TokenHelper().Token().Symbol()
	tokenName := t.sdk.Helper().TokenHelper().Token().Name()

	// 正常转账测试用例列表
	var cases = []struct {
		from   sdk.IAccount
		to     sdk.IAccount
		value  func() bn.Number
		desc   string
		code   uint32
		errmsg string
	}{
		{sender, acct, func() bn.Number { return bn.N(-1E10) }, "转出 -1E10cong", types.ErrInvalidParameter, "Value cannot be negative"},
		{sender, acct, func() bn.Number { return bn.N(-1) }, "转出 -1cong", types.ErrInvalidParameter, "Value cannot be negative"},
		{sender, acct, func() bn.Number { return bn.N(0) }, "转出 0cong", types.CodeOK, ""},
		{sender, acct, func() bn.Number { return bn.N(1) }, "转出 1cong", types.CodeOK, ""},
		{sender, acct, func() bn.Number { return bn.N(2) }, "转出 2cong", types.CodeOK, ""},
		{sender, acct, func() bn.Number { return sender.Balance().SubI(1) }, "剩余1cong", types.CodeOK, ""},
		{sender, acct, func() bn.Number { return sender.Balance() }, "转出剩余全部余额", types.CodeOK, ""},

		{acct, sender, func() bn.Number { return acct.Balance().SubI(2) }, "剩余2cong", types.CodeOK, ""},
		{acct, sender, func() bn.Number { return acct.Balance().DivI(2) }, "转出一半", types.CodeOK, ""},
		{acct, sender, func() bn.Number { return acct.Balance() }, "转出剩余全部余额", types.CodeOK, ""},

		{sender, acct, func() bn.Number { return sender.Balance() }, "转出账户全部余额", types.CodeOK, ""},
		{acct, sender, func() bn.Number { return acct.Balance() }, "转出账户全部余额", types.CodeOK, ""},
		{sender, acct, func() bn.Number { return sender.Balance().AddI(1) }, "转出账户全部余额+1cong", types.ErrInsufficientBalance, "Insufficient balance"},
		{sender, acct, func() bn.Number { return sender.Balance().AddI(2) }, "转出账户全部余额+2cong", types.ErrInsufficientBalance, "Insufficient balance"},
		{sender, acct, func() bn.Number { return sender.Balance().AddI(1E15) }, "转出账户全部余额+1E15cong", types.ErrInsufficientBalance, "Insufficient balance"},
	}
	for i, item := range cases {
		switch byType {
		case byNone:
			t.runTransfer(i, item.from, item.to, "", item.value(), item.desc, item.code, item.errmsg)
		case byToken:
			t.runTransferByToken(i, item.from, item.to, "", tokenAddr, item.value(), item.desc, item.code, item.errmsg)
		case byName:
			t.runTransferByTokenName(i, item.from, item.to, "", tokenName, item.value(), item.desc, item.code, item.errmsg)
		case bySymbol:
			t.runTransferByTokenSymbol(i, item.from, item.to, "", tokenSymbol, item.value(), item.desc, item.code, item.errmsg)
		}
	}
}

//transferSenderNoAuth test transfer functions with a sender with no authority to do
//nolint unhandled
func (t *TestAccount) transferSenderNoAuth(byType string) {
	fmt.Println("\nTEST CASE: 构造普通用户，测试转账权限")
	fmt.Printf(" transferSenderNoAuth(%v)\n", byType)

	accAddr := t.sdk.Message().Contract().Account()
	acct := t.sdk.Helper().AccountHelper().AccountOf(accAddr)

	//Transfer to a normal user
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)
	t.Transfer(recv.Address(), bn.N(1000000000))

	tokenAddr := t.sdk.Helper().TokenHelper().Token().Address()
	tokenSymbol := t.sdk.Helper().TokenHelper().Token().Symbol()
	tokenName := t.sdk.Helper().TokenHelper().Token().Name()

	// 正常转账测试用例列表
	var cases = []struct {
		from   sdk.IAccount
		to     sdk.IAccount
		value  func() bn.Number
		desc   string
		code   uint32
		errmsg string
	}{
		// 拥有者账户为转出账户
		{recv, acct, func() bn.Number { return bn.N(0) }, "转出 0cong", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, acct, func() bn.Number { return bn.N(1) }, "转出 1cong", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, acct, func() bn.Number { return bn.N(2) }, "转出 2cong", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, acct, func() bn.Number { return recv.Balance().SubI(1) }, "剩余1cong", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, acct, func() bn.Number { return recv.Balance() }, "转出账户全部余额", types.ErrNoAuthorization, "No authorization to execute contract"},
	}
	for i, item := range cases {
		switch byType {
		case byNone:
			t.runTransfer(i, item.from, item.to, "", item.value(), item.desc, item.code, item.errmsg)
		case byToken:
			t.runTransferByToken(i, item.from, item.to, "", tokenAddr, item.value(), item.desc, item.code, item.errmsg)
		case byName:
			t.runTransferByTokenName(i, item.from, item.to, "", tokenName, item.value(), item.desc, item.code, item.errmsg)
		case bySymbol:
			t.runTransferByTokenSymbol(i, item.from, item.to, "", tokenSymbol, item.value(), item.desc, item.code, item.errmsg)
		}
	}
}

//transferTestToAddress test all kind of "to" address
//nolint unhandled
func (t *TestAccount) transferTestToAddress(byType string) {
	fmt.Println("\nTEST CASE: 测试接收者地址")
	fmt.Printf(" transferTestToAddress(%v)\n", byType)

	sender := t.sdk.Message().Sender()
	//Transfer to a normal user
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)

	tokenAddr := t.sdk.Helper().TokenHelper().Token().Address()
	tokenSymbol := t.sdk.Helper().TokenHelper().Token().Symbol()
	tokenName := t.sdk.Helper().TokenHelper().Token().Name()

	chainID := t.sdk.Block().ChainID()
	// 正常转账测试用例列表
	var cases = []struct {
		from   sdk.IAccount
		to     sdk.IAccount
		toAddr types.Address
		value  func() bn.Number
		desc   string
		code   uint32
		errmsg string
	}{
		// 拥有者账户为转出账户
		{sender, recv, "h", func() bn.Number { return bn.N(1) }, "地址为 h", types.ErrInvalidAddress, "Address chainID is error! "},
		{sender, recv, chainID + "h", func() bn.Number { return bn.N(1) }, "地址为chainID+h", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{sender, recv, "hjsd43878@#*9", func() bn.Number { return bn.N(1) }, "地址为 hjsd43878@#*9", types.ErrInvalidAddress, "Address chainID is error! "},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4", func() bn.Number { return bn.N(1) }, "地址被截短", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b8", func() bn.Number { return bn.N(2) }, "地址尾数被修改", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ8UBgja40cLcSHKW4b4", func() bn.Number { return recv.Balance().SubI(1) }, "地址中间o转0", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{sender, recv, recv.Address() + "a", func() bn.Number { return bn.N(1E9) }, "地址增加一个字符a", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, recv, recv.Address() + "ab4", func() bn.Number { return bn.N(1E9) }, "地址增加三个字符ab4", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, sender, sender.Address(), func() bn.Number { return bn.N(1E9) }, "转账给自己", types.ErrInvalidParameter, "Cannot transfer to self"},
		{sender, recv, recv.Address(), func() bn.Number { return bn.N(1E9) }, "正常地址", types.CodeOK, ""},
	}
	for i, item := range cases {
		switch byType {
		case byNone:
			t.runTransfer(i, item.from, item.to, item.toAddr, item.value(), item.desc, item.code, item.errmsg)
		case byToken:
			t.runTransferByToken(i, item.from, item.to, item.toAddr, tokenAddr, item.value(), item.desc, item.code, item.errmsg)
		case byName:
			t.runTransferByTokenName(i, item.from, item.to, item.toAddr, tokenName, item.value(), item.desc, item.code, item.errmsg)
		case bySymbol:
			t.runTransferByTokenSymbol(i, item.from, item.to, item.toAddr, tokenSymbol, item.value(), item.desc, item.code, item.errmsg)
		}
	}
}

//transferTestToAddress test token address
//nolint unhandled
func (t *TestAccount) transferTestTokenAddress() {
	fmt.Println("\nTEST CASE: 测试代币地址")
	fmt.Printf(" transferTestTokenAddress()\n")

	sender := t.sdk.Message().Sender()

	//Transfer to a normal user
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)
	chainID := t.sdk.Block().ChainID()
	addr := t.sdk.Helper().TokenHelper().Token().Address()

	newtoken := algorithm.CalcContractAddress(chainID, sender.Address(), "token", "1.0.1")
	// 正常转账测试用例列表
	var cases = []struct {
		from      sdk.IAccount
		to        sdk.IAccount
		tokenAddr types.Address
		value     func() bn.Number
		desc      string
		code      uint32
		errmsg    string
	}{
		// 拥有者账户为转出账户
		{sender, recv, "", func() bn.Number { return bn.N(1) }, "代币地址为空", types.ErrInvalidAddress, "Address chainID is error! "},
		{sender, recv, "h", func() bn.Number { return bn.N(1) }, "代币地址为 h", types.ErrInvalidAddress, "Address chainID is error! "},
		{sender, recv, chainID + "h", func() bn.Number { return bn.N(1) }, "代币地址为chainID+h", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{sender, recv, "hjsd43878@#*9", func() bn.Number { return bn.N(1) }, "代币地址为 hjsd43878@#*9", types.ErrInvalidAddress, "Address chainID is error! "},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4", func() bn.Number { return bn.N(1) }, "代币地址被截短", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b8", func() bn.Number { return bn.N(2) }, "代币地址尾数被修改", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ8UBgja40cLcSHKW4b4", func() bn.Number { return recv.Balance().SubI(1) }, "代币地址中间o转0", types.ErrInvalidAddress, "Base58Addr parse error! "},
		{sender, recv, addr + "a", func() bn.Number { return bn.N(1E9) }, "代币地址增加一个字符a", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, recv, addr + "ab4", func() bn.Number { return bn.N(1E9) }, "代币地址增加三个字符ab4", types.ErrInvalidAddress, "Address checksum is error! "},
		{sender, recv, addr, func() bn.Number { return bn.N(1E9) }, "本合约代币地址", types.CodeOK, ""},
		//其他代币地址, 通过stub测试已存在代币的跨合约转账
		{sender, recv, newtoken, func() bn.Number { return bn.N(1E9) }, "不存在的代币", types.ErrInvalidParameter, "Token not found(address=" + newtoken + ")"},
		//本币地址
		{sender, recv, t.sdk.Helper().GenesisHelper().Token().Address(), func() bn.Number { return bn.N(1E9) }, "Sender转出本币", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, sender, t.sdk.Helper().GenesisHelper().Token().Address(), func() bn.Number { return bn.N(1E9) }, "普通用户转出本币", types.ErrNoAuthorization, "No authorization to execute contract"},
	}
	for i, item := range cases {
		t.runTransferByToken(i, item.from, item.to, "", item.tokenAddr, item.value(), item.desc, item.code, item.errmsg)
	}
}

//transferTestTokenName test transfer by token name
//nolint unhandled
func (t *TestAccount) transferTestTokenName() {
	fmt.Println("\nTEST CASE: 测试代币名称")
	fmt.Printf(" transferTestTokenName()\n")

	sender := t.sdk.Message().Sender()

	//normal user
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)
	chainID := t.sdk.Block().ChainID()
	name := t.sdk.Helper().TokenHelper().Token().Name()

	// 正常转账测试用例列表
	var cases = []struct {
		from      sdk.IAccount
		to        sdk.IAccount
		tokenName string
		value     func() bn.Number
		desc      string
		code      uint32
		errmsg    string
	}{
		// 拥有者账户为转出账户
		{sender, recv, "", func() bn.Number { return bn.N(1) }, "代币名称为空", types.ErrInvalidParameter, "Token not found(name=)"},
		{sender, recv, "h", func() bn.Number { return bn.N(1) }, "代币名称为 h", types.ErrInvalidParameter, "Token not found(name=h)"},
		{sender, recv, chainID, func() bn.Number { return bn.N(1) }, "代币名称为chainID", types.ErrInvalidParameter, "Token not found(name=" + chainID + ")"},
		{sender, recv, "h878@#*9", func() bn.Number { return bn.N(1) }, "代币名称为 h878@#*9", types.ErrInvalidParameter, "Token not found(name=h878@#*9)"},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b4", func() bn.Number { return bn.N(2) }, "代币名称为一个地址", types.ErrInvalidParameter, "Token not found(name=" + chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b4)"},
		{sender, recv, "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b4KG7Y7hWLhjx", func() bn.Number { return bn.N(1E9) }, "代币名称超过40字符", types.ErrInvalidParameter, "Token not found(name=KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b4KG7Y7hWLhjx)"},
		{sender, recv, name + "ab4", func() bn.Number { return bn.N(1E9) }, "代币名称增加三个字符ab4", types.ErrInvalidParameter, "Token not found(name=" + name + "ab4)"},
		{sender, recv, name, func() bn.Number { return bn.N(1E9) }, "本合约代币名称", types.CodeOK, ""},
		//本币地址
		{sender, recv, t.sdk.Helper().GenesisHelper().Token().Name(), func() bn.Number { return bn.N(1E9) }, "Sender转出本币", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, sender, t.sdk.Helper().GenesisHelper().Token().Name(), func() bn.Number { return bn.N(1E9) }, "普通用户转出本币", types.ErrNoAuthorization, "No authorization to execute contract"},
	}
	for i, item := range cases {
		t.runTransferByTokenName(i, item.from, item.to, "", item.tokenName, item.value(), item.desc, item.code, item.errmsg)
	}
}

//transferTestTokenSymbol test transfer by token name
//nolint unhandled
func (t *TestAccount) transferTestTokenSymbol() {
	fmt.Println("\nTEST CASE: 测试代币符号")
	fmt.Printf(" transferTestTokenSymbol()\n")

	sender := t.sdk.Message().Sender()

	//normal user
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)
	chainID := t.sdk.Block().ChainID()
	symbol := t.sdk.Helper().TokenHelper().Token().Symbol()

	// 正常转账测试用例列表
	var cases = []struct {
		from   sdk.IAccount
		to     sdk.IAccount
		symbol string
		value  func() bn.Number
		desc   string
		code   uint32
		errmsg string
	}{
		// 拥有者账户为转出账户
		{sender, recv, "", func() bn.Number { return bn.N(1) }, "代币符号为空", types.ErrInvalidParameter, "Token not found(symbol=)"},
		{sender, recv, "h", func() bn.Number { return bn.N(1) }, "代币符号为 h", types.ErrInvalidParameter, "Token not found(symbol=h)"},
		{sender, recv, chainID, func() bn.Number { return bn.N(1) }, "代币符号为chainID", types.ErrInvalidParameter, "Token not found(symbol=" + chainID + ")"},
		{sender, recv, "h878@#*9", func() bn.Number { return bn.N(1) }, "代币符号为 h878@#*9", types.ErrInvalidParameter, "Token not found(symbol=h878@#*9)"},
		{sender, recv, chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b4", func() bn.Number { return bn.N(2) }, "代币符号为一个地址", types.ErrInvalidParameter, "Token not found(symbol=" + chainID + "KG7Y7hWLhjxiBNZ9UBgja4ocLcSHKW4b4)"},
		{sender, recv, "KG7Y7hWLhjxiBNZ9UBgja4", func() bn.Number { return bn.N(1E9) }, "代币符号超过20字符", types.ErrInvalidParameter, "Token not found(symbol=KG7Y7hWLhjxiBNZ9UBgja4)"},
		{sender, recv, symbol + "ab4", func() bn.Number { return bn.N(1E9) }, "代币符号增加三个字符ab4", types.ErrInvalidParameter, "Token not found(symbol=" + symbol + "ab4)"},
		{sender, recv, symbol, func() bn.Number { return bn.N(1E9) }, "本合约代币符号", types.CodeOK, ""},
		//本币地址
		{sender, recv, t.sdk.Helper().GenesisHelper().Token().Symbol(), func() bn.Number { return bn.N(1E9) }, "Sender转出本币", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, sender, t.sdk.Helper().GenesisHelper().Token().Symbol(), func() bn.Number { return bn.N(1E9) }, "普通用户转出本币", types.ErrNoAuthorization, "No authorization to execute contract"},
	}
	for i, item := range cases {
		t.runTransferByTokenSymbol(i, item.from, item.to, "", item.symbol, item.value(), item.desc, item.code, item.errmsg)
	}
}

//transferTestGenesisToken test transfer genesis from kinds of accounts
//nolint unhandled
func (t *TestAccount) transferTestGenesisToken(byType string) {
	fmt.Println("\nTEST CASE: 测试本币转账-" + byType)
	fmt.Printf(" transferTestGenesisToken(%v)\n", byType)

	sender := t.sdk.Message().Sender()
	contractAcc := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Account())
	//Transfer to a normal user
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)

	tokenAddr := t.sdk.Helper().GenesisHelper().Token().Address()
	tokenSymbol := t.sdk.Helper().GenesisHelper().Token().Symbol()
	tokenName := t.sdk.Helper().GenesisHelper().Token().Name()

	// 正常转账测试用例列表
	var cases = []struct {
		from   sdk.IAccount
		to     sdk.IAccount
		value  func() bn.Number
		desc   string
		code   uint32
		errmsg string
	}{
		// 拥有者账户为转出账户
		{sender, recv, func() bn.Number { return bn.N(1) }, "从Sender转出", types.ErrNoAuthorization, "No authorization to execute contract"},
		{sender, sender, func() bn.Number { return bn.N(1) }, "从Sender转出给自己", types.ErrInvalidParameter, "Cannot transfer to self"},
		{recv, sender, func() bn.Number { return bn.N(1) }, "从普通用户转出", types.ErrNoAuthorization, "No authorization to execute contract"},
		{recv, recv, func() bn.Number { return bn.N(1) }, "从普通用户转出给自己", types.ErrInvalidParameter, "Cannot transfer to self"},
		{contractAcc, recv, func() bn.Number { return bn.N(2) }, "从合约账户转出", types.CodeOK, ""},
		{contractAcc, contractAcc, func() bn.Number { return bn.N(2) }, "从合约账户转出给自己", types.ErrInvalidParameter, "Cannot transfer to self"},
	}
	for i, item := range cases {
		switch byType {
		case byToken:
			t.runTransferGenTokenByToken(i, item.from, item.to, "", tokenAddr, item.value(), item.desc, item.code, item.errmsg)
		case byName:
			t.runTransferGenTokenByTokenName(i, item.from, item.to, "", tokenName, item.value(), item.desc, item.code, item.errmsg)
		case bySymbol:
			t.runTransferGenTokenByTokenSymbol(i, item.from, item.to, "", tokenSymbol, item.value(), item.desc, item.code, item.errmsg)
		}
	}
}

//transferTestGenesisToken test transfer genesis from kinds of accounts
//nolint unhandled
func (t *TestAccount) testAccountPubKey() {
	fmt.Println("\nTEST CASE: 测试账户PubKey  -- testAccountPubKey()")

	// 正常转账测试用例列表
	var cases = []struct {
		hexStr string
		desc   string
		code   uint32
		errmsg string
	}{
		{"", "pubkey 为空", types.ErrInvalidParameter, "Invalid PubKey"},
		{"A", "长度为1", types.ErrInvalidParameter, "Invalid PubKey"},
		{"AA", "长度为2", types.ErrInvalidParameter, "Invalid PubKey"},
		{"AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B10", "长度小1", types.ErrInvalidParameter, "Invalid PubKey"},
		{"AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B108334", "长度大1", types.ErrInvalidParameter, "Invalid PubKey"},
		{"AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B10844444444444", "长度大10", types.ErrInvalidParameter, "Invalid PubKey"},
		{"AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B10GG", "包含非16进制字符", types.ErrInvalidParameter, "Invalid PubKey"},
		{"AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083", "正常用例", types.CodeOK, ""},
		{"BBE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1084", "正常用例", types.CodeOK, ""},
		{"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "32个AA", types.CodeOK, ""},
		{"0000000000000000000000000000000000000000000000000000000000000000", "32个00", types.CodeOK, ""},
		{"1111111111111111111111111111111111111111111111111111111111111111", "32个11", types.CodeOK, ""},
		{"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", "32个FF", types.CodeOK, ""},
	}
	for i, item := range cases {
		printTestCase(i, item.desc)
		pubkey, _ := hex.DecodeString(item.hexStr)
		acc, err := t.createAccoutWithPubKey(pubkey)
		AssertError(err, item.code, item.errmsg)
		if item.code == types.CodeOK {
			AssertEquals(acc.PubKey().String(), item.hexStr)
		}
		printPass()
	}
}

//transferTestGenesisToken test transfer genesis from kinds of accounts
//nolint unhandled
func (t *TestAccount) testAccountBalance(byType string) {
	s := ""
	switch byType {
	case byNone:
		s = "Balance()"
	case byToken:
		s = "BalanceOfToken()"
	case byName:
		s = "BalanceOfName"
	case bySymbol:
		s = "BalanceOfSymbol"
	}
	fmt.Println("\nTEST CASE: 测试账户余额-" + s)
	fmt.Printf(" testAccountBalance(%v)\n", s)

	sender := t.sdk.Message().Sender()
	//Transfer to a normal user
	pubkey, _ := hex.DecodeString("AAE0014B0B08BB79B17B996ECABEDA6BF02534B64917631BB5DE59FB411B1083")
	recv := t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)
	chainID := t.sdk.Helper().GenesisHelper().ChainID()

	totalSupply := t.sdk.Helper().TokenHelper().Token().TotalSupply()
	tokenAddr := t.sdk.Helper().TokenHelper().Token().Address()
	tokenSymbol := t.sdk.Helper().TokenHelper().Token().Symbol()
	tokenName := t.sdk.Helper().TokenHelper().Token().Name()
	AssertEquals(recv.Balance().String(), bn.N(0).String())

	// 正常转账测试用例列表
	var cases = []struct {
		from    sdk.IAccount
		to      sdk.IAccount
		value   func() bn.Number
		balance bn.Number
		desc    string
		code    uint32
		errmsg  string
	}{
		// 拥有者账户为转出账户
		{sender, recv, func() bn.Number { return bn.N(0) }, bn.N(0), "转入0cong", types.CodeOK, ""},
		{sender, recv, func() bn.Number { return bn.N(1) }, bn.N(1), "转入1cong", types.CodeOK, ""},
		{sender, recv, func() bn.Number { return bn.N(2) }, bn.N(3), "转入2cong", types.CodeOK, ""},
		{sender, recv, func() bn.Number { return bn.N(1E9) }, bn.N(1E9 + 3), "转入1E9cong", types.CodeOK, ""},
		{sender, recv, func() bn.Number { return sender.Balance() }, totalSupply, "全部转入", types.CodeOK, ""},
	}
	for i, item := range cases {
		printTestCase(i, item.desc)
		item.from.Transfer(item.to.Address(), item.value())
		switch byType {
		case byNone:
			AssertEquals(item.to.Balance().String(), item.balance.String())
			AssertEquals(item.to.Balance().Add(item.from.Balance()).String(), totalSupply.String())

		case byToken:
			AssertEquals(item.to.BalanceOfToken(tokenAddr).String(), item.balance.String())
			AssertEquals(item.to.BalanceOfToken(tokenAddr).Add(item.from.BalanceOfName(tokenName)).String(), totalSupply.String())
		case byName:
			AssertEquals(item.to.BalanceOfName(tokenName).String(), item.balance.String())
			AssertEquals(item.to.BalanceOfName(tokenName).Add(item.from.BalanceOfSymbol(tokenSymbol)).String(), totalSupply.String())
		case bySymbol:
			AssertEquals(item.to.BalanceOfSymbol(tokenSymbol).String(), item.balance.String())
			AssertEquals(item.to.BalanceOfSymbol(tokenSymbol).Add(item.from.BalanceOfToken(tokenAddr)).String(), totalSupply.String())
		}
		printPass()
	}
	if byType == bySymbol || byType == byName {
		var cases2 = []struct {
			acc    sdk.IAccount
			str    string
			desc   string
			code   uint32
			errmsg string
		}{
			{sender, "", "不存在的代币:名字、符号", types.CodeOK, ""},
			{sender, "a", "不存在的代币:名字、符号a", types.CodeOK, ""},
			{sender, chainID, "不存在的代币:名字、符号" + chainID, types.CodeOK, ""},
			{sender, sender.Address(), "不存在的代币:名字、符号为发送者地址" + sender.Address(), types.CodeOK, ""},
			{sender, "abcdef123456", "不存在的代币:名字、符号为abcdef123456", types.CodeOK, ""},
		}
		z := bn.N(0).String()
		for i, item := range cases2 {
			printTestCase(i+len(cases), item.desc)
			AssertEquals(item.acc.BalanceOfSymbol(item.str).String(), z)
			AssertEquals(item.acc.BalanceOfName(item.str).String(), z)
			printPass()
		}
	} else if byType == byToken {
		var cases3 = []struct {
			acc    sdk.IAccount
			addr   types.Address
			desc   string
			code   uint32
			errmsg string
		}{
			{sender, "", "不存在的代币:地址为空", types.ErrInvalidAddress, "Address chainID is error! "},
			{sender, "a", "不存在的代币：地址为a", types.ErrInvalidAddress, "Address chainID is error! "},
			{sender, chainID, "不存在的代币：地址为" + chainID, types.ErrInvalidAddress, "Base58Addr parse error! "},
			{sender, sender.Address(), "不存在的代币名字、符号为发送者地址" + sender.Address(), types.CodeOK, ""},
		}
		z := bn.N(0).String()
		for i, item := range cases3 {
			printTestCase(i+len(cases), item.desc)
			b, err := t.balanceByToken(item.acc, item.addr)
			AssertError(err, item.code, item.errmsg)
			if item.code == types.CodeOK {
				AssertEquals(b.String(), z)
			}
			printPass()
		}
	}
}
