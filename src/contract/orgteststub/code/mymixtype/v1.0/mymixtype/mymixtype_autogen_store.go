package mymixtype

//_setBasic This is a method of Mymixtype
func (m *Mymixtype) _setBasic(v BasicTypes) {
	m.sdk.Helper().StateHelper().McSet("/basic", &v)
}

//_basic This is a method of Mymixtype
func (m *Mymixtype) _basic() BasicTypes {

	return *m.sdk.Helper().StateHelper().McGetEx("/basic", new(BasicTypes)).(*BasicTypes)
}

//_clrBasic This is a method of Mymixtype
func (m *Mymixtype) _clrBasic() {
	m.sdk.Helper().StateHelper().McClear("/basic")
}

//_chkBasic This is a method of Mymixtype
func (m *Mymixtype) _chkBasic() bool {
	return m.sdk.Helper().StateHelper().Check("/basic")
}

//_McChkBasic This is a method of Mymixtype
func (m *Mymixtype) _McChkBasic() bool {
	return m.sdk.Helper().StateHelper().McCheck("/basic")
}

//_setSlice This is a method of Mymixtype
func (m *Mymixtype) _setSlice(v SliceTypes) {
	m.sdk.Helper().StateHelper().Set("/slice", &v)
}

//_slice This is a method of Mymixtype
func (m *Mymixtype) _slice() SliceTypes {

	return *m.sdk.Helper().StateHelper().GetEx("/slice", new(SliceTypes)).(*SliceTypes)
}

//_chkSlice This is a method of Mymixtype
func (m *Mymixtype) _chkSlice() bool {
	return m.sdk.Helper().StateHelper().Check("/slice")
}

//_setMap1 This is a method of Mymixtype
func (m *Mymixtype) _setMap1(v MapStringTypes) {
	m.sdk.Helper().StateHelper().McSet("/map1", &v)
}

//_map1 This is a method of Mymixtype
func (m *Mymixtype) _map1() MapStringTypes {

	return *m.sdk.Helper().StateHelper().McGetEx("/map1", new(MapStringTypes)).(*MapStringTypes)
}

//_clrMap1 This is a method of Mymixtype
func (m *Mymixtype) _clrMap1() {
	m.sdk.Helper().StateHelper().McClear("/map1")
}

//_chkMap1 This is a method of Mymixtype
func (m *Mymixtype) _chkMap1() bool {
	return m.sdk.Helper().StateHelper().Check("/map1")
}

//_McChkMap1 This is a method of Mymixtype
func (m *Mymixtype) _McChkMap1() bool {
	return m.sdk.Helper().StateHelper().McCheck("/map1")
}

//_setMap2 This is a method of Mymixtype
func (m *Mymixtype) _setMap2(v MapOtherTypes) {
	m.sdk.Helper().StateHelper().McSet("/map2", &v)
}

//_map2 This is a method of Mymixtype
func (m *Mymixtype) _map2() MapOtherTypes {

	return *m.sdk.Helper().StateHelper().McGetEx("/map2", new(MapOtherTypes)).(*MapOtherTypes)
}

//_clrMap2 This is a method of Mymixtype
func (m *Mymixtype) _clrMap2() {
	m.sdk.Helper().StateHelper().McClear("/map2")
}

//_chkMap2 This is a method of Mymixtype
func (m *Mymixtype) _chkMap2() bool {
	return m.sdk.Helper().StateHelper().Check("/map2")
}

//_McChkMap2 This is a method of Mymixtype
func (m *Mymixtype) _McChkMap2() bool {
	return m.sdk.Helper().StateHelper().McCheck("/map2")
}

//_setMap3 This is a method of Mymixtype
func (m *Mymixtype) _setMap3(v MapSliceTypes) {
	m.sdk.Helper().StateHelper().McSet("/map3", &v)
}

//_map3 This is a method of Mymixtype
func (m *Mymixtype) _map3() MapSliceTypes {

	return *m.sdk.Helper().StateHelper().McGetEx("/map3", new(MapSliceTypes)).(*MapSliceTypes)
}

//_clrMap3 This is a method of Mymixtype
func (m *Mymixtype) _clrMap3() {
	m.sdk.Helper().StateHelper().McClear("/map3")
}

//_chkMap3 This is a method of Mymixtype
func (m *Mymixtype) _chkMap3() bool {
	return m.sdk.Helper().StateHelper().Check("/map3")
}

//_McChkMap3 This is a method of Mymixtype
func (m *Mymixtype) _McChkMap3() bool {
	return m.sdk.Helper().StateHelper().McCheck("/map3")
}

//_setComplex This is a method of Mymixtype
func (m *Mymixtype) _setComplex(v ComplexDefine) {
	m.sdk.Helper().StateHelper().Set("/complex", &v)
}

//_complex This is a method of Mymixtype
func (m *Mymixtype) _complex() ComplexDefine {

	return *m.sdk.Helper().StateHelper().GetEx("/complex", new(ComplexDefine)).(*ComplexDefine)
}

//_chkComplex This is a method of Mymixtype
func (m *Mymixtype) _chkComplex() bool {
	return m.sdk.Helper().StateHelper().Check("/complex")
}

//_setLong This is a method of Mymixtype
func (m *Mymixtype) _setLong(v LongType) {
	m.sdk.Helper().StateHelper().Set("/long", &v)
}

//_long This is a method of Mymixtype
func (m *Mymixtype) _long() LongType {

	return *m.sdk.Helper().StateHelper().GetEx("/long", new(LongType)).(*LongType)
}

//_chkLong This is a method of Mymixtype
func (m *Mymixtype) _chkLong() bool {
	return m.sdk.Helper().StateHelper().Check("/long")
}
