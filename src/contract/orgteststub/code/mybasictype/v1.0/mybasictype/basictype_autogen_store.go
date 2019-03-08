package mybasictype

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"fmt"
)

//_setAddress_ This is a method of BasicType
func (bt *BasicType) _setAddress_(v types.Address) {
	bt.sdk.Helper().StateHelper().McSet("/address_", &v)
}

//_address_ This is a method of BasicType
func (bt *BasicType) _address_() types.Address {

	return *bt.sdk.Helper().StateHelper().McGetEx("/address_", new(types.Address)).(*types.Address)
}

//_clrAddress_ This is a method of BasicType
func (bt *BasicType) _clrAddress_() {
	bt.sdk.Helper().StateHelper().McClear("/address_")
}

//_chkAddress_ This is a method of BasicType
func (bt *BasicType) _chkAddress_() bool {
	return bt.sdk.Helper().StateHelper().Check("/address_")
}

//_McChkAddress_ This is a method of BasicType
func (bt *BasicType) _McChkAddress_() bool {
	return bt.sdk.Helper().StateHelper().McCheck("/address_")
}

//_setHash_ This is a method of BasicType
func (bt *BasicType) _setHash_(v types.Hash) {
	bt.sdk.Helper().StateHelper().McSet("/hash_", &v)
}

//_hash_ This is a method of BasicType
func (bt *BasicType) _hash_() types.Hash {

	return *bt.sdk.Helper().StateHelper().McGetEx("/hash_", new(types.Hash)).(*types.Hash)
}

//_clrHash_ This is a method of BasicType
func (bt *BasicType) _clrHash_() {
	bt.sdk.Helper().StateHelper().McClear("/hash_")
}

//_chkHash_ This is a method of BasicType
func (bt *BasicType) _chkHash_() bool {
	return bt.sdk.Helper().StateHelper().Check("/hash_")
}

//_McChkHash_ This is a method of BasicType
func (bt *BasicType) _McChkHash_() bool {
	return bt.sdk.Helper().StateHelper().McCheck("/hash_")
}

//_setHexBytes_ This is a method of BasicType
func (bt *BasicType) _setHexBytes_(v types.HexBytes) {
	bt.sdk.Helper().StateHelper().McSet("/hexBytes_", &v)
}

//_hexBytes_ This is a method of BasicType
func (bt *BasicType) _hexBytes_() types.HexBytes {

	return *bt.sdk.Helper().StateHelper().McGetEx("/hexBytes_", new(types.HexBytes)).(*types.HexBytes)
}

//_clrHexBytes_ This is a method of BasicType
func (bt *BasicType) _clrHexBytes_() {
	bt.sdk.Helper().StateHelper().McClear("/hexBytes_")
}

//_chkHexBytes_ This is a method of BasicType
func (bt *BasicType) _chkHexBytes_() bool {
	return bt.sdk.Helper().StateHelper().Check("/hexBytes_")
}

//_McChkHexBytes_ This is a method of BasicType
func (bt *BasicType) _McChkHexBytes_() bool {
	return bt.sdk.Helper().StateHelper().McCheck("/hexBytes_")
}

//_setPubkey_ This is a method of BasicType
func (bt *BasicType) _setPubkey_(v types.PubKey) {
	bt.sdk.Helper().StateHelper().McSet("/pubkey_", &v)
}

//_pubkey_ This is a method of BasicType
func (bt *BasicType) _pubkey_() types.PubKey {

	return *bt.sdk.Helper().StateHelper().McGetEx("/pubkey_", new(types.PubKey)).(*types.PubKey)
}

//_clrPubkey_ This is a method of BasicType
func (bt *BasicType) _clrPubkey_() {
	bt.sdk.Helper().StateHelper().McClear("/pubkey_")
}

//_chkPubkey_ This is a method of BasicType
func (bt *BasicType) _chkPubkey_() bool {
	return bt.sdk.Helper().StateHelper().Check("/pubkey_")
}

//_McChkPubkey_ This is a method of BasicType
func (bt *BasicType) _McChkPubkey_() bool {
	return bt.sdk.Helper().StateHelper().McCheck("/pubkey_")
}

//_setNumber_ This is a method of BasicType
func (bt *BasicType) _setNumber_(v bn.Number) {
	bt.sdk.Helper().StateHelper().McSet("/number_", &v)
}

//_number_ This is a method of BasicType
func (bt *BasicType) _number_() bn.Number {
	temp := bn.N(0)
	return *bt.sdk.Helper().StateHelper().McGetEx("/number_", &temp).(*bn.Number)
}

//_clrNumber_ This is a method of BasicType
func (bt *BasicType) _clrNumber_() {
	bt.sdk.Helper().StateHelper().McClear("/number_")
}

//_chkNumber_ This is a method of BasicType
func (bt *BasicType) _chkNumber_() bool {
	return bt.sdk.Helper().StateHelper().Check("/number_")
}

//_McChkNumber_ This is a method of BasicType
func (bt *BasicType) _McChkNumber_() bool {
	return bt.sdk.Helper().StateHelper().McCheck("/number_")
}

//_setInt_ This is a method of BasicType
func (bt *BasicType) _setInt_(v int) {
	bt.sdk.Helper().StateHelper().Set("/int_", &v)
}

//_int_ This is a method of BasicType
func (bt *BasicType) _int_() int {

	return *bt.sdk.Helper().StateHelper().GetEx("/int_", new(int)).(*int)
}

//_chkInt_ This is a method of BasicType
func (bt *BasicType) _chkInt_() bool {
	return bt.sdk.Helper().StateHelper().Check("/int_")
}

//_setInt8_ This is a method of BasicType
func (bt *BasicType) _setInt8_(v int8) {
	bt.sdk.Helper().StateHelper().Set("/int8_", &v)
}

//_int8_ This is a method of BasicType
func (bt *BasicType) _int8_() int8 {

	return *bt.sdk.Helper().StateHelper().GetEx("/int8_", new(int8)).(*int8)
}

//_chkInt8_ This is a method of BasicType
func (bt *BasicType) _chkInt8_() bool {
	return bt.sdk.Helper().StateHelper().Check("/int8_")
}

//_setInt16_ This is a method of BasicType
func (bt *BasicType) _setInt16_(v int16) {
	bt.sdk.Helper().StateHelper().Set("/int16_", &v)
}

//_int16_ This is a method of BasicType
func (bt *BasicType) _int16_() int16 {

	return *bt.sdk.Helper().StateHelper().GetEx("/int16_", new(int16)).(*int16)
}

//_chkInt16_ This is a method of BasicType
func (bt *BasicType) _chkInt16_() bool {
	return bt.sdk.Helper().StateHelper().Check("/int16_")
}

//_setInt32_ This is a method of BasicType
func (bt *BasicType) _setInt32_(v int32) {
	bt.sdk.Helper().StateHelper().Set("/int32_", &v)
}

//_int32_ This is a method of BasicType
func (bt *BasicType) _int32_() int32 {

	return *bt.sdk.Helper().StateHelper().GetEx("/int32_", new(int32)).(*int32)
}

//_chkInt32_ This is a method of BasicType
func (bt *BasicType) _chkInt32_() bool {
	return bt.sdk.Helper().StateHelper().Check("/int32_")
}

//_setInt64_ This is a method of BasicType
func (bt *BasicType) _setInt64_(v int64) {
	bt.sdk.Helper().StateHelper().Set("/int64_", &v)
}

//_int64_ This is a method of BasicType
func (bt *BasicType) _int64_() int64 {

	return *bt.sdk.Helper().StateHelper().GetEx("/int64_", new(int64)).(*int64)
}

//_chkInt64_ This is a method of BasicType
func (bt *BasicType) _chkInt64_() bool {
	return bt.sdk.Helper().StateHelper().Check("/int64_")
}

//_setUint_ This is a method of BasicType
func (bt *BasicType) _setUint_(v uint) {
	bt.sdk.Helper().StateHelper().Set("/uint_", &v)
}

//_uint_ This is a method of BasicType
func (bt *BasicType) _uint_() uint {

	return *bt.sdk.Helper().StateHelper().GetEx("/uint_", new(uint)).(*uint)
}

//_chkUint_ This is a method of BasicType
func (bt *BasicType) _chkUint_() bool {
	return bt.sdk.Helper().StateHelper().Check("/uint_")
}

//_setUint8_ This is a method of BasicType
func (bt *BasicType) _setUint8_(v uint8) {
	bt.sdk.Helper().StateHelper().Set("/uint8_", &v)
}

//_uint8_ This is a method of BasicType
func (bt *BasicType) _uint8_() uint8 {

	return *bt.sdk.Helper().StateHelper().GetEx("/uint8_", new(uint8)).(*uint8)
}

//_chkUint8_ This is a method of BasicType
func (bt *BasicType) _chkUint8_() bool {
	return bt.sdk.Helper().StateHelper().Check("/uint8_")
}

//_setUint16_ This is a method of BasicType
func (bt *BasicType) _setUint16_(v uint16) {
	bt.sdk.Helper().StateHelper().Set("/uint16_", &v)
}

//_uint16_ This is a method of BasicType
func (bt *BasicType) _uint16_() uint16 {

	return *bt.sdk.Helper().StateHelper().GetEx("/uint16_", new(uint16)).(*uint16)
}

//_chkUint16_ This is a method of BasicType
func (bt *BasicType) _chkUint16_() bool {
	return bt.sdk.Helper().StateHelper().Check("/uint16_")
}

//_setUint32_ This is a method of BasicType
func (bt *BasicType) _setUint32_(v uint32) {
	bt.sdk.Helper().StateHelper().Set("/uint32_", &v)
}

//_uint32_ This is a method of BasicType
func (bt *BasicType) _uint32_() uint32 {

	return *bt.sdk.Helper().StateHelper().GetEx("/uint32_", new(uint32)).(*uint32)
}

//_chkUint32_ This is a method of BasicType
func (bt *BasicType) _chkUint32_() bool {
	return bt.sdk.Helper().StateHelper().Check("/uint32_")
}

//_setUint64_ This is a method of BasicType
func (bt *BasicType) _setUint64_(v uint64) {
	bt.sdk.Helper().StateHelper().Set("/uint64_", &v)
}

//_uint64_ This is a method of BasicType
func (bt *BasicType) _uint64_() uint64 {

	return *bt.sdk.Helper().StateHelper().GetEx("/uint64_", new(uint64)).(*uint64)
}

//_chkUint64_ This is a method of BasicType
func (bt *BasicType) _chkUint64_() bool {
	return bt.sdk.Helper().StateHelper().Check("/uint64_")
}

//_setString_ This is a method of BasicType
func (bt *BasicType) _setString_(v string) {
	bt.sdk.Helper().StateHelper().Set("/string_", &v)
}

//_string_ This is a method of BasicType
func (bt *BasicType) _string_() string {

	return *bt.sdk.Helper().StateHelper().GetEx("/string_", new(string)).(*string)
}

//_chkString_ This is a method of BasicType
func (bt *BasicType) _chkString_() bool {
	return bt.sdk.Helper().StateHelper().Check("/string_")
}

//_setByte_ This is a method of BasicType
func (bt *BasicType) _setByte_(v byte) {
	bt.sdk.Helper().StateHelper().Set("/byte_", &v)
}

//_byte_ This is a method of BasicType
func (bt *BasicType) _byte_() byte {

	return *bt.sdk.Helper().StateHelper().GetEx("/byte_", new(byte)).(*byte)
}

//_chkByte_ This is a method of BasicType
func (bt *BasicType) _chkByte_() bool {
	return bt.sdk.Helper().StateHelper().Check("/byte_")
}

//_setBool_ This is a method of BasicType
func (bt *BasicType) _setBool_(v bool) {
	bt.sdk.Helper().StateHelper().Set("/bool_", &v)
}

//_bool_ This is a method of BasicType
func (bt *BasicType) _bool_() bool {

	return *bt.sdk.Helper().StateHelper().GetEx("/bool_", new(bool)).(*bool)
}

//_chkBool_ This is a method of BasicType
func (bt *BasicType) _chkBool_() bool {
	return bt.sdk.Helper().StateHelper().Check("/bool_")
}

//_setMap_ This is a method of BasicType
func (bt *BasicType) _setMap_(k string, v int) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map_/%v", k), &v)
}

//_map_ This is a method of BasicType
func (bt *BasicType) _map_(k string) int {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map_/%v", k), new(int)).(*int)
}

//_chkMap_ This is a method of BasicType
func (bt *BasicType) _chkMap_(k string) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map_/%v", k))
}

//_setMap1_ This is a method of BasicType
func (bt *BasicType) _setMap1_(k1 uint, k2 string, v types.Address) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map1_/%v/%v", k1, k2), &v)
}

//_map1_ This is a method of BasicType
func (bt *BasicType) _map1_(k1 uint, k2 string) types.Address {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map1_/%v/%v", k1, k2), new(types.Address)).(*types.Address)
}

//_chkMap1_ This is a method of BasicType
func (bt *BasicType) _chkMap1_(k1 uint, k2 string) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map1_/%v/%v", k1, k2))
}

//_setMap2_ This is a method of BasicType
func (bt *BasicType) _setMap2_(k1 int, k2 int8, v types.Hash) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map2_/%v/%v", k1, k2), &v)
}

//_map2_ This is a method of BasicType
func (bt *BasicType) _map2_(k1 int, k2 int8) types.Hash {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map2_/%v/%v", k1, k2), new(types.Hash)).(*types.Hash)
}

//_chkMap2_ This is a method of BasicType
func (bt *BasicType) _chkMap2_(k1 int, k2 int8) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map2_/%v/%v", k1, k2))
}

//_setMap3_ This is a method of BasicType
func (bt *BasicType) _setMap3_(k1 int, k2 uint64, v types.HexBytes) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map3_/%v/%v", k1, k2), &v)
}

//_map3_ This is a method of BasicType
func (bt *BasicType) _map3_(k1 int, k2 uint64) types.HexBytes {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map3_/%v/%v", k1, k2), new(types.HexBytes)).(*types.HexBytes)
}

//_chkMap3_ This is a method of BasicType
func (bt *BasicType) _chkMap3_(k1 int, k2 uint64) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map3_/%v/%v", k1, k2))
}

//_setMap4_ This is a method of BasicType
func (bt *BasicType) _setMap4_(k1 bool, k2 byte, v types.PubKey) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map4_/%v/%v", k1, k2), &v)
}

//_map4_ This is a method of BasicType
func (bt *BasicType) _map4_(k1 bool, k2 byte) types.PubKey {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map4_/%v/%v", k1, k2), new(types.PubKey)).(*types.PubKey)
}

//_chkMap4_ This is a method of BasicType
func (bt *BasicType) _chkMap4_(k1 bool, k2 byte) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map4_/%v/%v", k1, k2))
}

//_setMap5_ This is a method of BasicType
func (bt *BasicType) _setMap5_(k1 bool, k2 bn.Number, v bool) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map5_/%v/%v", k1, k2), &v)
}

//_map5_ This is a method of BasicType
func (bt *BasicType) _map5_(k1 bool, k2 bn.Number) bool {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map5_/%v/%v", k1, k2), new(bool)).(*bool)
}

//_chkMap5_ This is a method of BasicType
func (bt *BasicType) _chkMap5_(k1 bool, k2 bn.Number) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map5_/%v/%v", k1, k2))
}

//_setMap6_ This is a method of BasicType
func (bt *BasicType) _setMap6_(k1 byte, k2 string, v bn.Number) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map6_/%v/%v", k1, k2), &v)
}

//_map6_ This is a method of BasicType
func (bt *BasicType) _map6_(k1 byte, k2 string) bn.Number {
	temp := bn.N(0)
	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map6_/%v/%v", k1, k2), &temp).(*bn.Number)
}

//_chkMap6_ This is a method of BasicType
func (bt *BasicType) _chkMap6_(k1 byte, k2 string) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map6_/%v/%v", k1, k2))
}

//_setMap7_ This is a method of BasicType
func (bt *BasicType) _setMap7_(k1 string, k2 types.Address, v byte) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map7_/%v/%v", k1, k2), &v)
}

//_map7_ This is a method of BasicType
func (bt *BasicType) _map7_(k1 string, k2 types.Address) byte {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map7_/%v/%v", k1, k2), new(byte)).(*byte)
}

//_chkMap7_ This is a method of BasicType
func (bt *BasicType) _chkMap7_(k1 string, k2 types.Address) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map7_/%v/%v", k1, k2))
}

//_setMap8_ This is a method of BasicType
func (bt *BasicType) _setMap8_(k1 types.Address, k2 bn.Number, v string) {
	bt.sdk.Helper().StateHelper().Set(fmt.Sprintf("/map8_/%v/%v", k1, k2), &v)
}

//_map8_ This is a method of BasicType
func (bt *BasicType) _map8_(k1 types.Address, k2 bn.Number) string {

	return *bt.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/map8_/%v/%v", k1, k2), new(string)).(*string)
}

//_chkMap8_ This is a method of BasicType
func (bt *BasicType) _chkMap8_(k1 types.Address, k2 bn.Number) bool {
	return bt.sdk.Helper().StateHelper().Check(fmt.Sprintf("/map8_/%v/%v", k1, k2))
}

//_setSlice_ This is a method of BasicType
func (bt *BasicType) _setSlice_(v []byte) {
	bt.sdk.Helper().StateHelper().Set("/slice_", &v)
}

//_slice_ This is a method of BasicType
func (bt *BasicType) _slice_() []byte {

	return *bt.sdk.Helper().StateHelper().GetEx("/slice_", new([]byte)).(*[]byte)
}

//_chkSlice_ This is a method of BasicType
func (bt *BasicType) _chkSlice_() bool {
	return bt.sdk.Helper().StateHelper().Check("/slice_")
}
