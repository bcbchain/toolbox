package testtoken

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

// 可以增发可以燃烧
func (t *TestToken) setTotalSupplyBurnAndAddEnabled() {

	pk := []byte("f1edf8f50848b8fa121a24e2a3a83cc5")
	tmp := t.sdk.Helper().AccountHelper().AccountOfPubKey(pk).Address()
	account := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Owner().Address())
	transferCount := account.Balance().DivI(2)
	account.Transfer(tmp, transferCount)

	token := t.sdk.Helper().TokenHelper().TokenOfAddress(t.sdk.Message().Contract().Token())
	cases := []struct {
		totalSupply func() bn.Number
		token       sdk.IToken
		desc        string
		code        uint32
		errMsg      string
	}{
		{func() bn.Number { return transferCount.DivI(2) }, token, "异常用例-合约拥有者在该代币账户上的余额更新后小于0", types.ErrInvalidParameter, "The owner's balance not enough to burn"},
		{func() bn.Number { return token.TotalSupply().Sub(bn.N(1E2)) }, token, "正常用例-新的总来小于原来的总量", types.CodeOK, ""},
		{func() bn.Number { return token.TotalSupply().Add(bn.N(1E2)) }, token, "正常用例-新的总来大于原来的总量", types.CodeOK, ""},
		{func() bn.Number { return bn.N(100).Sub(bn.N(200)) }, token, "异常用例-新的总量为负数", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
		{func() bn.Number { return bn.N(100) }, token, "异常用例-新的总量小于 1000000000 cong", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
	}

	for i, c := range cases {
		t.runSetTotalSupply(i, c.token, c.totalSupply(), c.code, c.desc, c.errMsg)
	}
}

// 可以增发不可以燃烧
func (t *TestToken) setTotalSupplyBurnDisableAndAddEnabled() {

	//pk := []byte("f1edf8f50848b8fa121a24e2a3a83cc5")
	//tmp := t.sdk.Helper().AccountHelper().AccountOfPubKey(pk).Address()
	//account := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Owner())
	//transferCount := account.Balance().DivI(2)
	//account.Transfer(tmp, transferCount)

	token := t.sdk.Helper().TokenHelper().TokenOfAddress(t.sdk.Message().Contract().Token())
	cases := []struct {
		totalSupply func() bn.Number
		token       sdk.IToken
		desc        string
		code        uint32
		errMsg      string
	}{
		{func() bn.Number { return token.TotalSupply().DivI(2) }, token, "异常用例-不允许燃烧", types.ErrBurnNotEnabled, "Burn supply is not enabled"},
		{func() bn.Number { return token.TotalSupply().Add(bn.N(1E2)) }, token, "正常用例-新的总来大于原来的总量", types.CodeOK, ""},
		{func() bn.Number { return bn.N(100).Sub(bn.N(200)) }, token, "异常用例-新的总量为负数", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
		{func() bn.Number { return bn.N(100) }, token, "异常用例-新的总量小于 1000000000 cong", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
	}

	for i, c := range cases {
		t.runSetTotalSupply(i, c.token, c.totalSupply(), c.code, c.desc, c.errMsg)
	}
}

// 可以燃烧不可以增发
func (t *TestToken) setTotalSupplyBurnEnabledAndAddDisabled() {

	pk := []byte("f1edf8f50848b8fa121a24e2a3a83cc5")
	tmp := t.sdk.Helper().AccountHelper().AccountOfPubKey(pk).Address()
	account := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Owner().Address())
	transferCount := account.Balance().DivI(2)
	account.Transfer(tmp, transferCount)

	token := t.sdk.Helper().TokenHelper().TokenOfAddress(t.sdk.Message().Contract().Token())
	cases := []struct {
		totalSupply func() bn.Number
		token       sdk.IToken
		desc        string
		code        uint32
		errMsg      string
	}{
		{func() bn.Number { return transferCount.DivI(2) }, token, "异常用例-合约拥有者在该代币账户上的余额更新后小于0", types.ErrInvalidParameter, "The owner's balance not enough to burn"},
		{func() bn.Number { return token.TotalSupply().DivI(2) }, token, "正常用例", types.CodeOK, ""},
		{func() bn.Number { return token.TotalSupply().Add(bn.N(1E2)) }, token, "异常用例-不可以增发", types.ErrAddSupplyNotEnabled, "Add supply is not enabled"},
		{func() bn.Number { return bn.N(100).Sub(bn.N(200)) }, token, "异常用例-新的总量为负数", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
		{func() bn.Number { return bn.N(100) }, token, "异常用例-新的总量小于 1000000000 cong", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
	}

	for i, c := range cases {
		t.runSetTotalSupply(i, c.token, c.totalSupply(), c.code, c.desc, c.errMsg)
	}
}

// 都不可以
func (t *TestToken) setTotalSupplyBurnAndAddDisabled() {

	//pk := []byte("f1edf8f50848b8fa121a24e2a3a83cc5")
	//tmp := t.sdk.Helper().AccountHelper().AccountOfPubKey(pk).Address()
	//account := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Owner())
	//transferCount := account.Balance().DivI(2)
	//account.Transfer(tmp, transferCount)

	token := t.sdk.Helper().TokenHelper().TokenOfAddress(t.sdk.Message().Contract().Token())
	cases := []struct {
		totalSupply func() bn.Number
		token       sdk.IToken
		desc        string
		code        uint32
		errMsg      string
	}{
		{func() bn.Number { return token.TotalSupply().DivI(2) }, token, "异常用例-不可以燃烧", types.ErrBurnNotEnabled, "Burn supply is not enabled"},
		{func() bn.Number { return token.TotalSupply().Add(bn.N(1E2)) }, token, "异常用例-不可以增发", types.ErrAddSupplyNotEnabled, "Add supply is not enabled"},
		{func() bn.Number { return bn.N(100).Sub(bn.N(200)) }, token, "异常用例-新的总量为负数", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
		{func() bn.Number { return bn.N(100) }, token, "异常用例-新的总量小于 1000000000 cong", types.ErrInvalidParameter, "TotalSupply must great than or equal 1000000000 cong"},
	}

	for i, c := range cases {
		t.runSetTotalSupply(i, c.token, c.totalSupply(), c.code, c.desc, c.errMsg)
	}
}

// token set gas price
func (t *TestToken) setGasPriceNormal() {

	token := t.sdk.Helper().TokenHelper().TokenOfAddress(t.sdk.Message().Contract().Token())
	cases := []struct {
		gasPrice int64
		token    sdk.IToken
		desc     string
		code     uint32
		errMsg   string
	}{
		{-10, token, "异常用例-gasPrice小于0", types.ErrInvalidParameter, "New gasPrice cannot less than baseGasPrice=2500"},
		{0, token, "异常用例-gasPrice等于0", types.ErrInvalidParameter, "New gasPrice cannot less than baseGasPrice=2500"},
		{10, token, "异常用例-gasPrice小于基础燃料价格", types.ErrInvalidParameter, "New gasPrice cannot less than baseGasPrice=2500"},
		{10000000000, token, "异常用例-gasPrice小于基础燃料价格", types.ErrInvalidParameter, "New gasPrice cannot great than maxGasPrice=1000000000"},
		{1000000000, token, "正常用例-gasPrice等于最大燃料价格", types.CodeOK, ""},
		{t.sdk.Helper().TokenHelper().BaseGasPrice(), token, "正常用例-gasPrice等于最底燃料价格", types.CodeOK, ""},
		{10000000, token, "正常用例", types.CodeOK, ""},
		{10000000, token, "正常用例-重复设置也可以", types.CodeOK, ""},
	}

	for i, c := range cases {
		t.runSetGasPrice(i, c.token, c.gasPrice, c.code, c.desc, c.errMsg)
	}
}

func (t *TestToken) testAddress(address types.Address) {
	Assert(address == t.sdk.Helper().TokenHelper().Token().Address())
}

func (t *TestToken) testOwner(owner types.Address) {
	Assert(owner == t.sdk.Helper().TokenHelper().Token().Owner().Address())
}

func (t *TestToken) testName(name string) {
	Assert(name == t.sdk.Helper().TokenHelper().Token().Name())
}

func (t *TestToken) testSymbol(symbol string) {
	Assert(symbol == t.sdk.Helper().TokenHelper().Token().Symbol())
}

func (t *TestToken) testTotalSupply(total bn.Number) {
	Assert(total.IsEqual(t.sdk.Helper().TokenHelper().Token().TotalSupply()))
}

func (t *TestToken) testAddSupplyEnabled(enabled bool) {
	Assert(enabled == t.sdk.Helper().TokenHelper().Token().AddSupplyEnabled())
}

func (t *TestToken) testBurnEnabled(enabled bool) {
	Assert(enabled == t.sdk.Helper().TokenHelper().Token().BurnEnabled())
}

func (t *TestToken) testGasPrice(gasPrice int64) {
	Assert(gasPrice == t.sdk.Helper().TokenHelper().Token().GasPrice())
}
