package testtx

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//TestTx This is struct of contract
//@:contract:testtx
//@:version:2.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111
//@:author:011b8728b2e5fd42c769bffe46cc92348f3f14f24a42a860d3d7760784f97869
type TestTx struct {
	sdk sdk.ISmartContract
}

//InitChain Constructor of this TestTx
//@:constructor
func (t *TestTx) InitChain() {

}

//Transfer This is a sample method
//@:public:method:gas[500]
//@:public:interface:gas[500]
func (t *TestTx) Transfer(to types.Address, value bn.Number) {

}

// TestNote test tx note
//@:public:method:gas[500]
func (t *TestTx) TestNote() {
	t.testNote()
}

// TestGasLimit test tx gas limit
//@:public:method:gas[500]
func (t *TestTx) TestGasLimit() {
	t.testGasLimit()
}

// TestGasLeft test tx gas left
//@:public:method:gas[500]
func (t *TestTx) TestGasLeft() {
	t.testGasLeft()
}

// TestSigner test tx signer
//@:public:method:gas[500]
func (t *TestTx) TestSigner() {
	t.testSigner()
}
