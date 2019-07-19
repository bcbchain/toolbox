package runtimeerror

import (
	"blockchain/smcsdk/sdk"
	"fmt"
)

//RuntimeError This is struct of contract
//@:contract:runtimeerror
//@:version:2.0
//@:organization:orgJgaGConUyK81zibntUBjQ33PKctpk1K1G
//@:author:5e8339cb1a5cce65602fd4f57e115905348f7e83bcbe38dd77694dbe1f8903c9
type RuntimeError struct {
	sdk sdk.ISmartContract

	//This is a sample field which is to store in db
	//@:public:store
	sampleStore string
}

//InitChain Constructor of this RuntimeError
//@:constructor
func (r *RuntimeError) InitChain() {
	//TODO
	//This method is automatically selected when the block height reaches the contract effective block height.
}

//UpdateChain Constructor of this RuntimeError
//@:constructor
func (r *RuntimeError) UpdateChain() {
	//TODO
	//This method is automatically selected when the block height reaches the new version contract effective block height.
}

//@:public:method:gas[500]
func (r *RuntimeError) IndexOutOfRange() {
	aaa := make([]string, 0)
	fmt.Println(aaa[10])
}

//@:public:method:gas[500]
func (r *RuntimeError) SliceBoundsOutOfRange() {
	aaa := make([]string, 0)
	fmt.Println(aaa[:10])
}

//@:public:method:gas[500]
func (r *RuntimeError) DividedByZero() {
	i := 10
	j := 0
	fmt.Println(i / j)
}

//@:public:method:gas[500]
func (r *RuntimeError) IntegerOverflow() {

}

//@:public:method:gas[500]
func (r *RuntimeError) FloatingPointError() {

}

//SampleMethod This is a sample method
//@:public:method:gas[500]
func (r *RuntimeError) NilPanic() {
	var b *string
	fmt.Println(*b)
}
