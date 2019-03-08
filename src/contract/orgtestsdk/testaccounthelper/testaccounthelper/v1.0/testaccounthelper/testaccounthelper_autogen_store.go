package testaccounthelper

//_setSampleStore This is a method of TestAccountHelper
func (tah *TestAccountHelper) _setSampleStore(v string) {
	tah.sdk.Helper().StateHelper().Set("/sampleStore", &v)
}

//_sampleStore This is a method of TestAccountHelper
func (tah *TestAccountHelper) _sampleStore() string {

	return *tah.sdk.Helper().StateHelper().GetEx("/sampleStore", new(string)).(*string)
}

//_chkSampleStore This is a method of TestAccountHelper
func (tah *TestAccountHelper) _chkSampleStore() bool {
	return tah.sdk.Helper().StateHelper().Check("/sampleStore")
}
