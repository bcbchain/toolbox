package testtokenhelper

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"fmt"
)

func (t *TestTokenHelper) testRegisterTokenName() {
	fmt.Println("\nTest Case: testRegisterTokenName()")

	cases := []struct {
		name             string
		symbol           string
		totalSupply      bn.Number
		addSupplyEnabled bool
		burnEnabled      bool
		desc             string
		code             uint32
		errMsg           string
	}{
		{"", "test1", bn.N(1E15), true, true, "name 为空", types.ErrInvalidParameter, "Token name cannot be less than 3 characters"},
		{"a", "test2", bn.N(1E15), true, true, "name 长度小于最小长度", types.ErrInvalidParameter, "Token name cannot be less than 3 characters"},
		{"testNameTooLongMoreThanFortyLengthTestCase", "test3", bn.N(1E15), true, true, "name 长度过长", types.ErrInvalidParameter, "The token name's length cannot great than 40"},
		{"test", "test", bn.N(1E15), true, true, "正常用例", types.CodeOK, ""},
		{"test", "test", bn.N(1E15), true, true, "重复注册 token", types.ErrInvalidParameter, "The contract has registered token already"},
	}

	for i, c := range cases {
		t.runRegisterToken(i, c.desc, c.code, c.errMsg, c.name, c.symbol, c.totalSupply, c.addSupplyEnabled, c.burnEnabled)
	}
}

func (t *TestTokenHelper) testRegisterTokenSymbol() {
	fmt.Println("\nTest Case: testRegisterTokenSymbol()")

	cases := []struct {
		name             string
		symbol           string
		totalSupply      bn.Number
		addSupplyEnabled bool
		burnEnabled      bool
		desc             string
		code             uint32
		errMsg           string
	}{
		{"testName", "", bn.N(1E15), true, true, "name 为空", types.ErrInvalidParameter, "Token symbol cannot be less than 3 characters"},
		{"testName1", "a", bn.N(1E15), true, true, "name 长度小于最小长度", types.ErrInvalidParameter, "Token symbol cannot be less than 3 characters"},
		{"testName2", "testSymbolTooLongTestCase", bn.N(1E15), true, true, "name 长度过长", types.ErrInvalidParameter, "The token symbol's length cannot great than 20"},
		{"testName", "symbol", bn.N(1E15), true, true, "正常用例", types.CodeOK, ""},
		{"testName", "symbol", bn.N(1E15), true, true, "重复注册 token", types.ErrInvalidParameter, "The contract has registered token already"},
	}

	for i, c := range cases {
		t.runRegisterToken(i, c.desc, c.code, c.errMsg, c.name, c.symbol, c.totalSupply, c.addSupplyEnabled, c.burnEnabled)
	}
}

func (t *TestTokenHelper) testRegisterTokenTotalSupply() {
	fmt.Println("\nTest Case: testRegisterTokenTotalSupply()")

	cases := []struct {
		name             string
		symbol           string
		totalSupply      bn.Number
		addSupplyEnabled bool
		burnEnabled      bool
		desc             string
		code             uint32
		errMsg           string
	}{
		{"testName", "symbol", bn.N(0).Sub(bn.N(100)), true, true, "总量小于0", types.ErrInvalidParameter, "The totalSupply cannot be less than 1000000000"},
		{"testName", "symbol", bn.N(0), true, true, "总量为0", types.ErrInvalidParameter, "The totalSupply cannot be less than 1000000000"},
		{"testName", "symbol", bn.N(1E15), true, true, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		t.runRegisterToken(i, c.desc, c.code, c.errMsg, c.name, c.symbol, c.totalSupply, c.addSupplyEnabled, c.burnEnabled)
	}
}

func (t *TestTokenHelper) testRegisterTokenAddSupplyEnabled() {
	fmt.Println("\nTest Case: testRegisterTokenAddSupplyEnabled()")

	cases := []struct {
		name             string
		symbol           string
		totalSupply      bn.Number
		addSupplyEnabled bool
		burnEnabled      bool
		desc             string
		code             uint32
		errMsg           string
	}{
		{"testName1", "symbol1", bn.N(1E15), true, true, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		t.runRegisterToken(i, c.desc, c.code, c.errMsg, c.name, c.symbol, c.totalSupply, c.addSupplyEnabled, c.burnEnabled)
	}
}

func (t *TestTokenHelper) testRegisterTokenBurnEnabled() {
	fmt.Println("\nTest Case: testRegisterTokenBurnEnabled()")

	cases := []struct {
		name             string
		symbol           string
		totalSupply      bn.Number
		addSupplyEnabled bool
		burnEnabled      bool
		desc             string
		code             uint32
		errMsg           string
	}{
		{"testName1", "symbol1", bn.N(1E15), true, true, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		t.runRegisterToken(i, c.desc, c.code, c.errMsg, c.name, c.symbol, c.totalSupply, c.addSupplyEnabled, c.burnEnabled)
	}
}

func (t *TestTokenHelper) testRegisterTokenDuplicate() {
	fmt.Println("\nTest Case: testRegisterTokenBurnEnabled()")

	cases := []struct {
		name             string
		symbol           string
		totalSupply      bn.Number
		addSupplyEnabled bool
		burnEnabled      bool
		desc             string
		code             uint32
		errMsg           string
	}{
		{"testName", "symbol", bn.N(1E15), true, true, "正常用例", types.CodeOK, ""},
		{"testName", "symbol", bn.N(1E15), true, true, "异常用例-相同参数的token重复注册", types.ErrInvalidParameter, "The contract has registered token already"},
		{"testName", "symbol1", bn.N(1E16), false, false, "异常用例-重复注册，名字相同", types.ErrInvalidParameter, "The contract has registered token already"},
		{"testName1", "symbol", bn.N(1E16), false, false, "异常用例-重复注册，符号相同", types.ErrInvalidParameter, "The contract has registered token already"},
		{"testName1", "symbol1", bn.N(1E15), false, false, "异常用例-重复注册，总量相同", types.ErrInvalidParameter, "The contract has registered token already"},
		{"testName1", "symbol1", bn.N(1E16), true, false, "异常用例-重复注册，是否允许增发相同", types.ErrInvalidParameter, "The contract has registered token already"},
		{"testName1", "symbol1", bn.N(1E16), false, true, "异常用例-重复注册，是否允许燃烧相同", types.ErrInvalidParameter, "The contract has registered token already"},
	}

	for i, c := range cases {
		t.runRegisterToken(i, c.desc, c.code, c.errMsg, c.name, c.symbol, c.totalSupply, c.addSupplyEnabled, c.burnEnabled)
	}
}

func (t *TestTokenHelper) testToken() {
	fmt.Println("\nTest Case: TestToken()")
	printTestCase(0, "正常用例")

	name := "test-token1"
	symbol := "test1"
	totalSupply := bn.N(1E15)
	addSupplyEnabled := true
	burnEnabled := true
	t.sdk.Helper().TokenHelper().RegisterToken(name, symbol, totalSupply, addSupplyEnabled, burnEnabled)
	token := t.sdk.Helper().TokenHelper().Token()

	Assert(token.Name() == name)
	Assert(token.Symbol() == symbol)
	Assert(token.TotalSupply() == totalSupply)
	Assert(token.AddSupplyEnabled() == addSupplyEnabled)
	Assert(token.BurnEnabled() == burnEnabled)
	Assert(token.Address() == t.sdk.Message().Contract().Address())
	printPass()
}

func (t *TestTokenHelper) testTokenOfAddress() {
	fmt.Println("\nTest Case: TestTokenOfAddress()")
	printTestCase(0, "正常用例")

	name := "test-token2"
	symbol := "test2"
	totalSupply := bn.N(1E15)
	addSupplyEnabled := true
	burnEnabled := true
	t.sdk.Helper().TokenHelper().RegisterToken(name, symbol, totalSupply, addSupplyEnabled, burnEnabled)
	token := t.sdk.Helper().TokenHelper().TokenOfAddress(t.sdk.Message().Contract().Address())

	Assert(token.Name() == name)
	Assert(token.Symbol() == symbol)
	Assert(token.TotalSupply() == totalSupply)
	Assert(token.AddSupplyEnabled() == addSupplyEnabled)
	Assert(token.BurnEnabled() == burnEnabled)
	printPass()
}

func (t *TestTokenHelper) testTokenOfName() {
	fmt.Println("\nTest Case: TestTokenOfName()")
	printTestCase(0, "正常用例")

	name := "test-token3"
	symbol := "test3"
	totalSupply := bn.N(1E15)
	addSupplyEnabled := true
	burnEnabled := true
	t.sdk.Helper().TokenHelper().RegisterToken(name, symbol, totalSupply, addSupplyEnabled, burnEnabled)
	token := t.sdk.Helper().TokenHelper().TokenOfName(name)

	Assert(token.Name() == name)
	Assert(token.Symbol() == symbol)
	Assert(token.TotalSupply() == totalSupply)
	Assert(token.AddSupplyEnabled() == addSupplyEnabled)
	Assert(token.BurnEnabled() == burnEnabled)
	Assert(token.Address() == t.sdk.Message().Contract().Address())
	printPass()
}

func (t *TestTokenHelper) testTokenOfSymbol() {
	fmt.Println("\nTest Case: TestTokenOfSymbol()")
	printTestCase(0, "正常用例")

	name := "test-token4"
	symbol := "test4"
	totalSupply := bn.N(1E15)
	addSupplyEnabled := true
	burnEnabled := true
	t.sdk.Helper().TokenHelper().RegisterToken(name, symbol, totalSupply, addSupplyEnabled, burnEnabled)
	token := t.sdk.Helper().TokenHelper().TokenOfSymbol(symbol)

	Assert(token.Name() == name)
	Assert(token.Symbol() == symbol)
	Assert(token.TotalSupply() == totalSupply)
	Assert(token.AddSupplyEnabled() == addSupplyEnabled)
	Assert(token.BurnEnabled() == burnEnabled)
	Assert(token.Address() == t.sdk.Message().Contract().Address())
	printPass()
}

func (t *TestTokenHelper) testTokenOfContract() {
	fmt.Println("\nTest Case: TestTokenOfContract()")
	printTestCase(0, "正常用例")

	name := "test-token5"
	symbol := "test5"
	totalSupply := bn.N(1E15)
	addSupplyEnabled := true
	burnEnabled := true
	t.sdk.Helper().TokenHelper().RegisterToken(name, symbol, totalSupply, addSupplyEnabled, burnEnabled)
	token := t.sdk.Helper().TokenHelper().TokenOfContract(t.sdk.Message().Contract().Address())

	Assert(token.Name() == name)
	Assert(token.Symbol() == symbol)
	Assert(token.TotalSupply() == totalSupply)
	Assert(token.AddSupplyEnabled() == addSupplyEnabled)
	Assert(token.BurnEnabled() == burnEnabled)
	Assert(token.Address() == t.sdk.Message().Contract().Address())
	printPass()
}

func (t *TestTokenHelper) testBaseGasPrice() {
	fmt.Println("\nTest Case: TestBaseGasPrice()")
	printTestCase(0, "正常用例")

	baseGasPrice := t.sdk.Helper().TokenHelper().BaseGasPrice()
	Assert(baseGasPrice == 2500)
	printPass()
}
