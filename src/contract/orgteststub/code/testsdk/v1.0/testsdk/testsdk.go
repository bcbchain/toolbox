package testsdk

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
)

//TestSdk test sdk interface
//@:contract:testsdk
//@:version:1.0
//@:organization:orgGyRrMVF7ukfHNwaZhgWMTbQAYz7d7RcBh
//@:author:b37e7627431feb18123b81bcf1f41ffd37efdb90513d48ff2c7f8a0c27a9d06c
type TestSdk struct {
	sdk sdk.ISmartContract
	//@:public:store
	mydata Mystruct
	//@:public:store:cache
	datai16 int16
}

//Mystruct is a test structure
type Mystruct struct {
	Index int64
	Data  string
	Mymap map[int64]string
}

//InitChain init when deployed on the blockchain first time
//@:constructor
func (dt *TestSdk) InitChain() {

}

//SetDatai16 set data
//@:public:method:gas[500]
func (dt *TestSdk) SetDatai16(d int16) {
	dt._setDatai16(d)
}

//GetDatai16 get data
//@:public:method:gas[500]
func (dt *TestSdk) GetDatai16() int16 {
	return dt._datai16()
}

//SetOwner set data
//@:public:method:gas[500]
func (dt *TestSdk) SetOwner(owner types.Address) (err types.Error) {
	dt.sdk.Message().Contract().SetOwner(owner)

	return
}

//SetMydata set data
//@:public:method:gas[500]
func (dt *TestSdk) SetMydata(v Mystruct) types.Error {

	dt._setMydata(v)
	return types.Error{ErrorCode: types.CodeOK}
}

//GetMydata get data
//@:public:method:gas[500]
func (dt *TestSdk) GetMydata() (Mystruct, types.Error) {

	v := dt._mydata()
	return v, types.Error{ErrorCode: types.CodeOK}
}
