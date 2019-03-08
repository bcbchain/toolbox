package testsdk

func (ts *TestSdk) _setMydata(v Mystruct) {
	ts.sdk.Helper().StateHelper().Set("/mydata", &v)
}
func (ts *TestSdk) _mydata() Mystruct {
	return *ts.sdk.Helper().StateHelper().GetEx("/mydata", new(Mystruct)).(*Mystruct)
}
func (ts *TestSdk) _chkMydata() bool {
	return ts.sdk.Helper().StateHelper().Check("/mydata")
}

func (ts *TestSdk) _setDatai16(v int16) {
	ts.sdk.Helper().StateHelper().McSet("/datai16", &v)
}
func (ts *TestSdk) _datai16() int16 {
	return *ts.sdk.Helper().StateHelper().McGetEx("/datai16", new(int16)).(*int16)
}
func (ts *TestSdk) _clrDatai16() {
	ts.sdk.Helper().StateHelper().McClear("/datai16")
}
func (ts *TestSdk) _chkDatai16() bool {
	return ts.sdk.Helper().StateHelper().Check("/datai16")
}
