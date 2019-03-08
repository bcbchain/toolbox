package testblock

//_setSampleStore This is a method of TestBlock
func (tb *TestBlock) _setSampleStore(v string) {
	tb.sdk.Helper().StateHelper().Set("/sampleStore", &v)
}

//_sampleStore This is a method of TestBlock
func (tb *TestBlock) _sampleStore() string {

	return *tb.sdk.Helper().StateHelper().GetEx("/sampleStore", new(string)).(*string)
}

//_chkSampleStore This is a method of TestBlock
func (tb *TestBlock) _chkSampleStore() bool {
	return tb.sdk.Helper().StateHelper().Check("/sampleStore")
}
