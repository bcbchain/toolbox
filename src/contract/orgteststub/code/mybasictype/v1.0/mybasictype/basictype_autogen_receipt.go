package mybasictype

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

var _ receipt = (*BasicType)(nil)

//emitAddressT This is a method of BasicType
func (bt *BasicType) emitAddressT(address_ types.Address) {
	type addressT struct {
		Address_ types.Address `json:"address_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		addressT{
			Address_: address_,
		},
	)
}

//emitHashT This is a method of BasicType
func (bt *BasicType) emitHashT(hash_ types.Hash) {
	type hashT struct {
		Hash_ types.Hash `json:"hash_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		hashT{
			Hash_: hash_,
		},
	)
}

//emitHexBytesT This is a method of BasicType
func (bt *BasicType) emitHexBytesT(hexBytes_ types.HexBytes) {
	type hexBytesT struct {
		HexBytes_ types.HexBytes `json:"hexBytes_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		hexBytesT{
			HexBytes_: hexBytes_,
		},
	)
}

//emitPubKeyT This is a method of BasicType
func (bt *BasicType) emitPubKeyT(pubKey_ types.PubKey) {
	type pubKeyT struct {
		PubKey_ types.PubKey `json:"pubKey_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		pubKeyT{
			PubKey_: pubKey_,
		},
	)
}

//emitNumberT This is a method of BasicType
func (bt *BasicType) emitNumberT(number_ bn.Number) {
	type numberT struct {
		Number_ bn.Number `json:"number_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		numberT{
			Number_: number_,
		},
	)
}

//emitIntT This is a method of BasicType
func (bt *BasicType) emitIntT(int_ int) {
	type intT struct {
		Int_ int `json:"int_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		intT{
			Int_: int_,
		},
	)
}

//emitInt8T This is a method of BasicType
func (bt *BasicType) emitInt8T(int8_ int8) {
	type int8T struct {
		Int8_ int8 `json:"int8_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		int8T{
			Int8_: int8_,
		},
	)
}

//emitInt16T This is a method of BasicType
func (bt *BasicType) emitInt16T(int16_ int16) {
	type int16T struct {
		Int16_ int16 `json:"int16_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		int16T{
			Int16_: int16_,
		},
	)
}

//emitInt32T This is a method of BasicType
func (bt *BasicType) emitInt32T(int32_ int32) {
	type int32T struct {
		Int32_ int32 `json:"int32_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		int32T{
			Int32_: int32_,
		},
	)
}

//emitInt64T This is a method of BasicType
func (bt *BasicType) emitInt64T(int64_ int64) {
	type int64T struct {
		Int64_ int64 `json:"int64_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		int64T{
			Int64_: int64_,
		},
	)
}

//emitUintT This is a method of BasicType
func (bt *BasicType) emitUintT(uint_ uint) {
	type uintT struct {
		Uint_ uint `json:"uint_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		uintT{
			Uint_: uint_,
		},
	)
}

//emitUint8T This is a method of BasicType
func (bt *BasicType) emitUint8T(uint8_ uint8) {
	type uint8T struct {
		Uint8_ uint8 `json:"uint8_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		uint8T{
			Uint8_: uint8_,
		},
	)
}

//emitUint16T This is a method of BasicType
func (bt *BasicType) emitUint16T(uint16_ uint16) {
	type uint16T struct {
		Uint16_ uint16 `json:"uint16_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		uint16T{
			Uint16_: uint16_,
		},
	)
}

//emitUint32T This is a method of BasicType
func (bt *BasicType) emitUint32T(uint32_ uint32) {
	type uint32T struct {
		Uint32_ uint32 `json:"uint32_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		uint32T{
			Uint32_: uint32_,
		},
	)
}

//emitUint64T This is a method of BasicType
func (bt *BasicType) emitUint64T(uint64_ uint64) {
	type uint64T struct {
		Uint64_ uint64 `json:"uint64_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		uint64T{
			Uint64_: uint64_,
		},
	)
}

//emitBoolT This is a method of BasicType
func (bt *BasicType) emitBoolT(bool_ bool) {
	type boolT struct {
		Bool_ bool `json:"bool_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		boolT{
			Bool_: bool_,
		},
	)
}

//emitByteT This is a method of BasicType
func (bt *BasicType) emitByteT(byte_ byte) {
	type byteT struct {
		Byte_ byte `json:"byte_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		byteT{
			Byte_: byte_,
		},
	)
}

//emitSliceT This is a method of BasicType
func (bt *BasicType) emitSliceT(slice_ []byte) {
	type sliceT struct {
		Slice_ []byte `json:"slice_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		sliceT{
			Slice_: slice_,
		},
	)
}

//emitMapT This is a method of BasicType
func (bt *BasicType) emitMapT(map_ map[string]int) {
	type mapT struct {
		Map_ map[string]int `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		mapT{
			Map_: map_,
		},
	)
}

//emitMap1T This is a method of BasicType
func (bt *BasicType) emitMap1T(map_ map[uint]map[string]types.Address) {
	type map1T struct {
		Map_ map[uint]map[string]types.Address `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map1T{
			Map_: map_,
		},
	)
}

//emitMap2T This is a method of BasicType
func (bt *BasicType) emitMap2T(map_ map[int]map[int8]types.Hash) {
	type map2T struct {
		Map_ map[int]map[int8]types.Hash `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map2T{
			Map_: map_,
		},
	)
}

//emitMap3T This is a method of BasicType
func (bt *BasicType) emitMap3T(map_ map[int]map[uint64]types.HexBytes) {
	type map3T struct {
		Map_ map[int]map[uint64]types.HexBytes `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map3T{
			Map_: map_,
		},
	)
}

//emitMap4T This is a method of BasicType
func (bt *BasicType) emitMap4T(map_ map[bool]map[byte]types.PubKey) {
	type map4T struct {
		Map_ map[bool]map[byte]types.PubKey `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map4T{
			Map_: map_,
		},
	)
}

//emitMap5T This is a method of BasicType
func (bt *BasicType) emitMap5T(map_ map[bool]map[bn.Number]bool) {
	type map5T struct {
		Map_ map[bool]map[bn.Number]bool `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map5T{
			Map_: map_,
		},
	)
}

//emitMap6T This is a method of BasicType
func (bt *BasicType) emitMap6T(map_ map[byte]map[string]bn.Number) {
	type map6T struct {
		Map_ map[byte]map[string]bn.Number `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map6T{
			Map_: map_,
		},
	)
}

//emitMap7T This is a method of BasicType
func (bt *BasicType) emitMap7T(map_ map[string]map[types.Address]byte) {
	type map7T struct {
		Map_ map[string]map[types.Address]byte `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map7T{
			Map_: map_,
		},
	)
}

//emitMap8T This is a method of BasicType
func (bt *BasicType) emitMap8T(map_ map[types.Address]map[bn.Number]string) {
	type map8T struct {
		Map_ map[types.Address]map[bn.Number]string `json:"map_"`
	}

	bt.sdk.Helper().ReceiptHelper().Emit(
		map8T{
			Map_: map_,
		},
	)
}
