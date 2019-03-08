package testblock

import (
	"blockchain/smcsdk/sdk"
)

//TestBlock This is struct of contract
//@:contract:testblock
//@:version:2.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111
//@:author:011b8728b2e5fd42c769bffe46cc92348f3f14f24a42a860d3d7760784f97869
type TestBlock struct {
	sdk sdk.ISmartContract

	//This is a sample field which is to store in db
	//@:public:store
	sampleStore string
}

//TestChainID test sdl IBlock ChainID
//@:public:method:gas[500]
func (t *TestBlock) TestChainID() {
	t.testChainID()
}

// TestBlockHash test
//@:public:method:gas[500]
func (t *TestBlock) TestBlockHash() {
	t.testBlockHash()
}

// TestHeight test
//@:public:method:gas[500]
func (t *TestBlock) TestHeight() {
	t.testHeight()
}

// TestTime test
//@:public:method:gas[500]
func (t *TestBlock) TestTime() {
	t.testTime()
}

// TestNow test
//@:public:method:gas[500]
func (t *TestBlock) TestNow() {
	t.testNow()
}

// TestNumTxs test
//@:public:method:gas[500]
func (t *TestBlock) TestNumTxs() {
	t.testNumTxs()
}

// TestDataHash test
//@:public:method:gas[500]
func (t *TestBlock) TestDataHash() {
	t.testDataHash()
}

// TestProposerAddress test
//@:public:method:gas[500]
func (t *TestBlock) TestProposerAddress() {
	t.testProposerAddress()
}

// TestRewardAddress test
//@:public:method:gas[500]
func (t *TestBlock) TestRewardAddress() {
	t.testRewardAddress()
}

// TestRandomNumber test
//@:public:method:gas[500]
func (t *TestBlock) TestRandomNumber() {
	t.testRandomNumber()
}

// TestVersion test
//@:public:method:gas[500]
func (t *TestBlock) TestVersion() {
	t.testVersion()
}

// TestLastBlockHash test
//@:public:method:gas[500]
func (t *TestBlock) TestLastBlockHash() {
	t.testLastBlockHash()
}

// TestLastCommitHash test
//@:public:method:gas[500]
func (t *TestBlock) TestLastCommitHash() {
	t.testLastCommitHash()
}

// TestAppHash test
//@:public:method:gas[500]
func (t *TestBlock) TestAppHash() {
	t.testLastAppHash()
}

// TestLastFee test
//@:public:method:gas[500]
func (t *TestBlock) TestLastFee() {
	t.testLastFee()
}
