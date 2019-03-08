package testtx

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl"
	"fmt"
	"reflect"

	"gopkg.in/check.v1"
)

var c *check.C

//SetChecker set checker for this module
func SetChecker(checker *check.C) {
	c = checker
}

//Assert assert true
func Assert(b bool) {
	if !b {
		printFail("Assert(b bool)", false, true)
	}
}

//AssertEquals assert a equals b
func AssertEquals(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		printFail("reflect.DeepEqual(a, b)", "a != b", "a == b")
	}
}

//AssertOK assert errcode is CodeOK
func AssertOK(err types.Error) {
	if !reflect.DeepEqual(err.ErrorCode, types.CodeOK) {
		printFail("reflect.DeepEqual(err.ErrorCode, types.CodeOK)", err.ErrorCode, types.CodeOK)
	}
}

//AssertSDB assert key's value in SDB
func AssertSDB(key string, expect interface{}, defaultValue interface{}, s sdk.ISmartContract) {
	s.(*sdkimpl.SmartContract).LlState().Get(key, defaultValue)
	if !reflect.DeepEqual(defaultValue, expect) {
		printFail("reflect.DeepEqual(defaultValue, expect)", defaultValue, expect)
	}
}

// AssertBalance assert account balance
func AssertBalance(acc sdk.IAccount, tokenName string, value bn.Number, s sdk.ISmartContract) {
	_token := s.Helper().TokenHelper().TokenOfName(tokenName)
	key := std.KeyOfAccountToken(acc.Address(), _token.Address())
	accountInfo := std.AccountInfo{}
	s.(*sdkimpl.SmartContract).LlState().Get(key, &accountInfo)
	if !reflect.DeepEqual(_token.Address(), accountInfo.Address) {
		printFail("reflect.DeepEqual(_token.Address(), accountInfo.Address)", _token.Address(), accountInfo.Address)
	}
	if !reflect.DeepEqual(value.String(), accountInfo.Balance.String()) {
		printFail("reflect.DeepEqual(value, accountInfo.Balance)", value.String(), accountInfo.Balance.String())
	}
}

//AssertError assert error code and message
func AssertError(err types.Error, errcode uint32, errmsg string) {
	if !reflect.DeepEqual(err.Error(), errmsg) {
		printFail("reflect.DeepEqual(err.Error(), errmsg)", err.Error(), errmsg)
	}
	if !reflect.DeepEqual(err.ErrorCode, errcode) {
		printFail("reflect.DeepEqual(err.ErrorCode, errcode)", err.ErrorCode, errcode)
	}
}

func printTestCase(index int, desc string) {
	fmt.Printf("Case:%3d (%s) \t\t", index, desc)
}
func printPass() {
	fmt.Printf("---- PASS\n")
}
func printFail(call string, obtained, expected interface{}) {
	fmt.Printf("---- FAIL: %v: obtained: %v, expected: %v\n", call, obtained, expected)
	c.Assert(obtained, check.Equals, expected)
}

//funcRecover recover panic by Assert
func funcRecover(err *types.Error) {
	err.ErrorCode = types.CodeOK
	if rerr := recover(); rerr != nil {
		if _, ok := rerr.(types.Error); ok {
			err.ErrorCode = rerr.(types.Error).ErrorCode
			err.ErrorDesc = rerr.(types.Error).ErrorDesc
		} else {
			panic(rerr)
		}
	}
}
