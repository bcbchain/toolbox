package testmessage

import (
	"blockchain/smcsdk/sdk"
)

//TestMessage This is struct of contract
//@:contract:testmessage
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111
//@:author:8d9d958020a77c9d3e04467d1e4095e43916a948a36168078bd28d9eeae7cfb4
type TestMessage struct {
	sdk sdk.ISmartContract
}

//InitChain Constructor of this TestMessage
//@:constructor
func (t *TestMessage) InitChain() {

}

//TestContract test Contract() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestContract() {
	t.testContract()
}

//TestMethodID test MethodID() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestMethodID() {
	t.testMethodID()
}

//TestItems test Data() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestItems() {
	t.testItems()
}

//TestGasPrice test GasPrice() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestGasPrice() {
	t.testGasPrice()
}

//TestSender test Sender() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestSender() {
	t.testSender()
}

//TestOrigins test Origins() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestOrigins() {
	t.testOrigins()
}

//TestInputReceipts test InputReceipts() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestInputReceipts() {
	t.testInputReceipts()
}

//TestGetTransferToMe test GetTransferToMe() interface of sdk IMessage
//@:public:method:gas[200]
func (t *TestMessage) TestGetTransferToMe() {
	t.testGetTransferToMe()
}
