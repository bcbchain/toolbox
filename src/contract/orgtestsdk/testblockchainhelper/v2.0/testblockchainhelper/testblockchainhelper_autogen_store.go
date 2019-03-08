package testblockchainhelper

//_setSampleStore This is a method of TestblockChainHelper
func (t *TestblockChainHelper) _setSampleStore(v string) {
	t.sdk.Helper().StateHelper().Set("/sampleStore", &v)
}

//_sampleStore This is a method of TestblockChainHelper
func (t *TestblockChainHelper) _sampleStore() string {

	return *t.sdk.Helper().StateHelper().GetEx("/sampleStore", new(string)).(*string)
}

//_chkSampleStore This is a method of TestblockChainHelper
func (t *TestblockChainHelper) _chkSampleStore() bool {
	return t.sdk.Helper().StateHelper().Check("/sampleStore")
}
