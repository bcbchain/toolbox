package testaccounthelper

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
)

//TestAccountHelper This is struct of contract
//@:contract:testAccountHelper
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzf
//@:author:011b8728b2e5fd42c769bffe46cc92348f3f14f24a42a860d3d7760784f97869
type TestAccountHelper struct {
	sdk sdk.ISmartContract
}

//InitChain Constructor of this TestAccountHelper
//@:constructor
func (tah *TestAccountHelper) InitChain() {

}

//TestAccountOf test accountOf address
//@:public:method:gas[500]
func (tah *TestAccountHelper) TestAccountOf() {
	tah.testAccountOf()
}

//TestAccountOfPubKey test make an account by pubKey
//@:public:method:gas[500]
func (tah *TestAccountHelper) TestAccountOfPubKey() {
	tah.testAccountOfPubKey()
}

func (tah *TestAccountHelper) runTestAccountOf(index int, addr types.Address, desc string, code uint32, errmsg string) {
	printTestCase(index, desc)
	acc, err := tah.sdkAccountOf(addr)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		Assert(acc.Address() == addr)
	}
	printPass()
}

func (tah TestAccountHelper) sdkAccountOf(addr types.Address) (acc sdk.IAccount, err types.Error) {
	defer funcRecover(&err)
	acc = tah.sdk.Helper().AccountHelper().AccountOf(addr)
	return
}

func (tah *TestAccountHelper) runTestAccountOfPubKey(index int, pk types.PubKey, desc string, code uint32, errmsg string) {
	printTestCase(index, desc)
	acc, err := tah.sdkAccountOfPubKey(pk)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		Assert(len(acc.PubKey()) == len(pk))
		for i, v := range pk {
			Assert(v == acc.PubKey()[i])
		}
	}
	printPass()
}

func (tah TestAccountHelper) sdkAccountOfPubKey(pk types.PubKey) (acc sdk.IAccount, err types.Error) {
	defer funcRecover(&err)
	acc = tah.sdk.Helper().AccountHelper().AccountOfPubKey(pk)
	return
}
