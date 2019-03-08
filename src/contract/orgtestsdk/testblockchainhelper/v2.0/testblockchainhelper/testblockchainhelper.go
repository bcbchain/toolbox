package testblockchainhelper

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"strings"
)

//TestblockChainHelper This is struct of contract
//@:contract:testblockchainhelper
//@:version:2.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111
//@:author:011b8728b2e5fd42c769bffe46cc92348f3f14f24a42a860d3d7760784f97869
type TestblockChainHelper struct {
	sdk sdk.ISmartContract

	//This is a sample field which is to store in db
	//@:public:store
	sampleStore string
}

//TestCalcAccountFromPubKey test
//@:public:method:gas[500]
func (t *TestblockChainHelper) TestCalcAccountFromPubKey() {
	t.testCalcAccountFromPubKey()
}

//TestCalcAccountFromName test
//@:public:method:gas[500]
func (t *TestblockChainHelper) TestCalcAccountFromName() {
}

//TestCalcContractAddress test
//@:public:method:gas[500]
func (t *TestblockChainHelper) TestCalcContractAddress() {
}

//TestCalcOrgID test
//@:public:method:gas[500]
func (t *TestblockChainHelper) TestCalcOrgID() {
}

//TestCheckAddress test
//@:public:method:gas[500]
func (t *TestblockChainHelper) TestCheckAddress() {
}

//TestGetBlock test
//@:public:method:gas[500]
func (t *TestblockChainHelper) TestGetBlock() {
}

func (t *TestblockChainHelper) runCalcAccountFromPubKey(index int, pubKey types.PubKey, desc string, code uint32, errmsg string) {
	printTestCase(index, desc)

	addr, err := t.sdkCalcAccountFromPubKey(pubKey)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		Assert(len(addr) > 0)
		Assert(strings.HasPrefix(addr, t.sdk.Block().ChainID()))
	}
	printPass()
}

func (t *TestblockChainHelper) sdkCalcAccountFromPubKey(pubKey types.PubKey) (addr types.Address, err types.Error) {
	defer funcRecover(&err)
	addr = t.sdk.Helper().BlockChainHelper().CalcAccountFromPubKey(pubKey)
	return
}

func (t *TestblockChainHelper) runCalcAccountFromName(index int, name string, orgID string, desc string, code uint32, errmsg string) {
	printTestCase(index, desc)
	addr, err := t.sdkCalcAccountFromName(name, orgID)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		Assert(len(addr) > 0)
		Assert(strings.HasPrefix(addr, t.sdk.Block().ChainID()))
	}

	printPass()
}

func (t *TestblockChainHelper) sdkCalcAccountFromName(name string, orgID string) (addr types.Address, err types.Error) {
	defer funcRecover(&err)
	addr = t.sdk.Helper().BlockChainHelper().CalcAccountFromName(name, orgID)
	return
}

func (t *TestblockChainHelper) runCalcContractAddress(index int, name, version string, owner types.Address, desc string, code uint32, errmsg string) {
	printTestCase(index, desc)
	addr, err := t.sdkCalcContractAddress(name, version, owner)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		Assert(len(addr) > 0)
		Assert(strings.HasPrefix(addr, t.sdk.Block().ChainID()))
	}

	printPass()
}

func (t *TestblockChainHelper) sdkCalcContractAddress(name, version string, owner types.Address) (addr types.Address, err types.Error) {
	defer funcRecover(&err)
	addr = t.sdk.Helper().BlockChainHelper().CalcContractAddress(name, version, owner)
	return
}

func (t *TestblockChainHelper) runCalcOrgID(index int, name string, desc string, code uint32, errmsg string) {
	printTestCase(index, desc)
	addr, err := t.sdkCalcOrgID(name)
	AssertError(err, code, errmsg)
	if code == types.CodeOK {
		Assert(len(addr) > 0)
		Assert(strings.HasPrefix(addr, t.sdk.Block().ChainID()))
	}

	printPass()
}

func (t *TestblockChainHelper) sdkCalcOrgID(name string) (addr types.Address, err types.Error) {
	defer funcRecover(&err)
	addr = t.sdk.Helper().BlockChainHelper().CalcOrgID(name)
	return
}
