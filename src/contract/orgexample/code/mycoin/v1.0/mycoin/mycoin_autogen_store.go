package mycoin

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"fmt"
)

//_setTotalSupply This is a method of Mycoin
func (m *Mycoin) _setTotalSupply(v bn.Number) {
	m.sdk.Helper().StateHelper().McSet("/totalSupply", &v)
}

//_totalSupply This is a method of Mycoin
func (m *Mycoin) _totalSupply() bn.Number {
	temp := bn.N(0)
	return *m.sdk.Helper().StateHelper().McGetEx("/totalSupply", &temp).(*bn.Number)
}

//_clrTotalSupply This is a method of Mycoin
func (m *Mycoin) _clrTotalSupply() {
	m.sdk.Helper().StateHelper().McClear("/totalSupply")
}

//_chkTotalSupply This is a method of Mycoin
func (m *Mycoin) _chkTotalSupply() bool {
	return m.sdk.Helper().StateHelper().Check("/totalSupply")
}

//_McChkTotalSupply This is a method of Mycoin
func (m *Mycoin) _McChkTotalSupply() bool {
	return m.sdk.Helper().StateHelper().McCheck("/totalSupply")
}

//_delTotalSupply This is a method of Mycoin
func (m *Mycoin) _delTotalSupply() {
	m.sdk.Helper().StateHelper().Delete("/totalSupply")
}

//_McDelTotalSupply This is a method of Mycoin
func (m *Mycoin) _McDelTotalSupply() {
	m.sdk.Helper().StateHelper().McDelete("/totalSupply")
}

//_setBalanceOf This is a method of Mycoin
func (m *Mycoin) _setBalanceOf(k types.Address, v bn.Number) {
	m.sdk.Helper().StateHelper().Set(fmt.Sprintf("/balanceOf/%v", k), &v)
}

//_balanceOf This is a method of Mycoin
func (m *Mycoin) _balanceOf(k types.Address) bn.Number {
	temp := bn.N(0)
	return *m.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/balanceOf/%v", k), &temp).(*bn.Number)
}

//_chkBalanceOf This is a method of Mycoin
func (m *Mycoin) _chkBalanceOf(k types.Address) bool {
	return m.sdk.Helper().StateHelper().Check(fmt.Sprintf("/balanceOf/%v", k))
}

//_delBalanceOf This is a method of Mycoin
func (m *Mycoin) _delBalanceOf(k types.Address) {
	m.sdk.Helper().StateHelper().Delete(fmt.Sprintf("/balanceOf/%v", k))
}
