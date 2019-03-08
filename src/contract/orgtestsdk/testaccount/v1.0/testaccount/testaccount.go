package testaccount

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

const (
	byNone   = "Transfer"
	byToken  = "TransferByToken"
	byName   = "TransferByName"
	bySymbol = "TransferBySymbol"

	tokenName   = "mytoken"
	tokenSymbol = "mytn"
)

//TestAccount This is struct of contract
//@:contract:testaccount
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzf
//@:author:8d9d958020a77c9d3e04467d1e4095e43916a948a36168078bd28d9eeae7cfb4
type TestAccount struct {
	sdk sdk.ISmartContract
}

//InitChain Constructor of this TestAccount
//@:constructor
func (t *TestAccount) InitChain() {

}

//TestTransfer test sdk Transfer() interface
//@:public:method:gas[500]
func (t *TestAccount) TestTransfer() {

	//test transfer value
	t.transferTestTransferValue(byNone)
	//test "to" address
	t.transferTestToAddress(byNone)
	//transfer by a sender with no authority to do
	t.transferSenderNoAuth(byNone)
}

//TestTransferByToken test sdk TransferByToken() interface
//@:public:method:gas[500]
func (t *TestAccount) TestTransferByToken() {

	//test transfer genesis from kinds of accounts
	t.transferTestGenesisToken(byToken)
	//test transfer value
	t.transferTestTransferValue(byToken)
	//test "to" address
	t.transferTestToAddress(byToken)
	//transfer by a sender with no authority to do
	t.transferSenderNoAuth(byToken)
	//test transfer by token address
	t.transferTestTokenAddress()
}

//TestTransferByName test sdk TransferByName() interface
//@:public:method:gas[500]
func (t *TestAccount) TestTransferByName() {

	//test transfer genesis from kinds of accounts
	t.transferTestGenesisToken(byName)
	//test transfer value
	t.transferTestTransferValue(byName)
	//test "to" address
	t.transferTestToAddress(byName)
	//transfer by a sender with no authority to do
	t.transferSenderNoAuth(byName)
	//test transfer by token name
	t.transferTestTokenName()
}

//TestTransferBySymbol test sdk TransferBySymbol() interface
//@:public:method:gas[500]
func (t *TestAccount) TestTransferBySymbol() {

	//test transfer genesis from kinds of accounts
	t.transferTestGenesisToken(bySymbol)
	//test transfer value
	t.transferTestTransferValue(bySymbol)
	// test "to" address
	t.transferTestToAddress(bySymbol)
	//transfer by a sender with no authority to do
	t.transferSenderNoAuth(bySymbol)
	//test transfer by token symbol
	t.transferTestTokenSymbol()
}

//TestAccountPubKey test sdk PubKey() interface
//@:public:method:gas[500]
func (t *TestAccount) TestAccountPubKey() {
	t.testAccountPubKey()
}

//TestBalance test sdk Balance() interface
//@:public:method:gas[500]
func (t *TestAccount) TestBalance() {
	t.testAccountBalance(byNone)
}

//TestBalanceOfToken test sdk BalanceOfToken() interface
//@:public:method:gas[500]
func (t *TestAccount) TestBalanceOfToken() {
	t.testAccountBalance(byToken)
}

//TestBalanceOfName test sdk BalanceOfName() interface
//@:public:method:gas[500]
func (t *TestAccount) TestBalanceOfName() {
	t.testAccountBalance(byName)
}

//TestBalanceOfSymbol test sdk BalanceOfSymbol() interface
//@:public:method:gas[500]
func (t *TestAccount) TestBalanceOfSymbol() {
	t.testAccountBalance(bySymbol)
}

//Transfer declare transfer interface of token of the contract owned
//@:public:method:gas[500]
//@:public:interface:gas[500]
func (t *TestAccount) Transfer(to types.Address, value bn.Number) {
	t.sdk.Message().Sender().Transfer(to, value)
}

func (t *TestAccount) runTransferByToken(index int, from, to sdk.IAccount, toAddr types.Address, tokenAddr types.Address, value bn.Number, desc string, code uint32, errmsg string) {

	printTestCase(index, desc)
	fromBalance := from.Balance()
	toBalance := to.Balance()

	var addr types.Address
	if toAddr != "" {
		addr = toAddr
	} else {
		addr = to.Address()
	}
	err := t.transferByToken(from, tokenAddr, addr, value)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		esb := fromBalance.Sub(value)
		tob := toBalance.Add(value)
		token := t.sdk.Helper().TokenHelper().TokenOfAddress(tokenAddr)
		AssertBalance(from, token.Name(), esb, t.sdk)
		AssertBalance(to, token.Name(), tob, t.sdk)

		AssertEquals(from.BalanceOfToken(tokenAddr).String(), esb.String())
		AssertEquals(to.BalanceOfToken(tokenAddr).String(), tob.String())
	}

	printPass()
}

func (t *TestAccount) runTransferByTokenName(index int, from, to sdk.IAccount, toAddr types.Address, tokenName string, value bn.Number, desc string, code uint32, errmsg string) {

	printTestCase(index, desc)
	fromBalance := from.Balance()
	toBalance := to.Balance()

	var addr types.Address
	if toAddr != "" {
		addr = toAddr
	} else {
		addr = to.Address()
	}
	err := t.transferByTokenName(from, tokenName, addr, value)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		esb := fromBalance.Sub(value)
		tob := toBalance.Add(value)
		token := t.sdk.Helper().TokenHelper().TokenOfName(tokenName)
		AssertBalance(from, token.Name(), esb, t.sdk)
		AssertBalance(to, token.Name(), tob, t.sdk)

		AssertEquals(from.BalanceOfName(tokenName).String(), esb.String())
		AssertEquals(to.BalanceOfName(tokenName).String(), tob.String())
	}

	printPass()
}

func (t *TestAccount) runTransferByTokenSymbol(index int, from, to sdk.IAccount, toAddr types.Address, symbol string, value bn.Number, desc string, code uint32, errmsg string) {

	printTestCase(index, desc)
	fromBalance := from.Balance()
	toBalance := to.Balance()

	var addr types.Address
	if toAddr != "" {
		addr = toAddr
	} else {
		addr = to.Address()
	}
	err := t.transferByTokenSymbol(from, symbol, addr, value)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		esb := fromBalance.Sub(value)
		tob := toBalance.Add(value)
		token := t.sdk.Helper().TokenHelper().TokenOfSymbol(symbol)
		AssertBalance(from, token.Name(), esb, t.sdk)
		AssertBalance(to, token.Name(), tob, t.sdk)

		AssertEquals(from.BalanceOfSymbol(symbol).String(), esb.String())
		AssertEquals(to.BalanceOfSymbol(symbol).String(), tob.String())
	}

	printPass()
}

func (t *TestAccount) runTransfer(index int, from, to sdk.IAccount, toAddr types.Address, value bn.Number, desc string, code uint32, errmsg string) {

	printTestCase(index, desc)
	fromBalance := from.Balance()
	toBalance := to.Balance()

	var addr types.Address
	if toAddr != "" {
		addr = toAddr
	} else {
		addr = to.Address()
	}
	err := t.sdkTransfer(from, addr, value)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		esb := fromBalance.Sub(value)
		tob := toBalance.Add(value)
		AssertBalance(from, t.sdk.Helper().TokenHelper().Token().Name(), esb, t.sdk)
		AssertBalance(to, t.sdk.Helper().TokenHelper().Token().Name(), tob, t.sdk)

		AssertEquals(from.Balance().String(), esb.String())
		AssertEquals(to.Balance().String(), tob.String())
	}

	printPass()
}

func (t *TestAccount) transferByToken(sender sdk.IAccount, tokenAddr, to types.Address, value bn.Number) (err types.Error) {
	defer funcRecover(&err)
	sender.TransferByToken(tokenAddr, to, value)
	return
}
func (t *TestAccount) transferByTokenName(sender sdk.IAccount, tokenName string, to types.Address, value bn.Number) (err types.Error) {
	defer funcRecover(&err)
	sender.TransferByName(tokenName, to, value)
	return
}
func (t *TestAccount) transferByTokenSymbol(sender sdk.IAccount, symbol, to types.Address, value bn.Number) (err types.Error) {
	defer funcRecover(&err)
	sender.TransferBySymbol(symbol, to, value)
	return
}

func (t *TestAccount) sdkTransfer(sender sdk.IAccount, to types.Address, value bn.Number) (err types.Error) {
	defer funcRecover(&err)
	sender.Transfer(to, value)
	return
}

func (t *TestAccount) runTransferGenTokenByToken(index int, from, to sdk.IAccount, toAddr types.Address, tokenAddr types.Address, value bn.Number, desc string, code uint32, errmsg string) {

	printTestCase(index, desc)
	fromBalance := from.BalanceOfToken(tokenAddr)
	toBalance := to.BalanceOfToken(tokenAddr)

	var addr types.Address
	if toAddr != "" {
		addr = toAddr
	} else {
		addr = to.Address()
	}
	err := t.transferByToken(from, tokenAddr, addr, value)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		esb := fromBalance.Sub(value)
		tob := toBalance.Add(value)
		token := t.sdk.Helper().TokenHelper().TokenOfAddress(tokenAddr)
		AssertBalance(from, token.Name(), esb, t.sdk)
		AssertBalance(to, token.Name(), tob, t.sdk)

		AssertEquals(from.BalanceOfToken(tokenAddr).String(), esb.String())
		AssertEquals(to.BalanceOfToken(tokenAddr).String(), tob.String())
	}

	printPass()
}

func (t *TestAccount) runTransferGenTokenByTokenName(index int, from, to sdk.IAccount, toAddr types.Address, tokenName string, value bn.Number, desc string, code uint32, errmsg string) {

	printTestCase(index, desc)
	fromBalance := from.BalanceOfName(tokenName)
	toBalance := to.BalanceOfName(tokenName)

	var addr types.Address
	if toAddr != "" {
		addr = toAddr
	} else {
		addr = to.Address()
	}
	err := t.transferByTokenName(from, tokenName, addr, value)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		esb := fromBalance.Sub(value)
		tob := toBalance.Add(value)
		AssertBalance(from, tokenName, esb, t.sdk)
		AssertBalance(to, tokenName, tob, t.sdk)

		AssertEquals(from.BalanceOfName(tokenName).String(), esb.String())
		AssertEquals(to.BalanceOfName(tokenName).String(), tob.String())
	}

	printPass()
}

func (t *TestAccount) runTransferGenTokenByTokenSymbol(index int, from, to sdk.IAccount, toAddr types.Address, symbol string, value bn.Number, desc string, code uint32, errmsg string) {

	printTestCase(index, desc)
	fromBalance := from.BalanceOfSymbol(symbol)
	toBalance := to.BalanceOfSymbol(symbol)

	var addr types.Address
	if toAddr != "" {
		addr = toAddr
	} else {
		addr = to.Address()
	}
	err := t.transferByTokenSymbol(from, symbol, addr, value)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		esb := fromBalance.Sub(value)
		tob := toBalance.Add(value)
		token := t.sdk.Helper().TokenHelper().TokenOfSymbol(symbol)
		AssertBalance(from, token.Name(), esb, t.sdk)
		AssertBalance(to, token.Name(), tob, t.sdk)

		AssertEquals(from.BalanceOfSymbol(symbol).String(), esb.String())
		AssertEquals(to.BalanceOfSymbol(symbol).String(), tob.String())
	}

	printPass()
}

func (t *TestAccount) createAccoutWithPubKey(pubkey []byte) (acc sdk.IAccount, err types.Error) {
	defer funcRecover(&err)
	acc = t.sdk.Helper().AccountHelper().AccountOfPubKey(pubkey)
	return
}

func (t *TestAccount) balanceByToken(acc sdk.IAccount, addr types.Address) (b bn.Number, err types.Error) {
	defer funcRecover(&err)
	b = acc.BalanceOfToken(addr)
	return
}
