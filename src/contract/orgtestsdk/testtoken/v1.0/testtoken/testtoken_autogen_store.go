package testtoken

//_setSampleStore This is a method of TestToken
func (tt *TestToken) _setSampleStore(v string) {
	tt.sdk.Helper().StateHelper().Set("/sampleStore", &v)
}

//_sampleStore This is a method of TestToken
func (tt *TestToken) _sampleStore() string {

	return *tt.sdk.Helper().StateHelper().GetEx("/sampleStore", new(string)).(*string)
}

//_chkSampleStore This is a method of TestToken
func (tt *TestToken) _chkSampleStore() bool {
	return tt.sdk.Helper().StateHelper().Check("/sampleStore")
}
