package mycrossmc

//_setStoredData This is a method of MyCrossmc
func (mc *MyCrossmc) _setStoredData(v uint64) {
	mc.sdk.Helper().StateHelper().Set("/storedData", &v)
}

//_storedData This is a method of MyCrossmc
func (mc *MyCrossmc) _storedData() uint64 {

	return *mc.sdk.Helper().StateHelper().GetEx("/storedData", new(uint64)).(*uint64)
}

//_chkStoredData This is a method of MyCrossmc
func (mc *MyCrossmc) _chkStoredData() bool {
	return mc.sdk.Helper().StateHelper().Check("/storedData")
}
