package mymixtype

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

var _ receipt = (*Mymixtype)(nil)

//emitBasic This is a method of Mymixtype
func (m *Mymixtype) emitBasic(int2 int, int82 int8, int162 int16, int322 int32, int642 int64, uint2 uint, uint82 uint8, uint162 uint16, uint322 uint32, uint642 uint64, string2 string, bool2 bool, byte2 byte) {
	type basic struct {
		Int2    int    `json:"int2"`
		Int82   int8   `json:"int82"`
		Int162  int16  `json:"int162"`
		Int322  int32  `json:"int322"`
		Int642  int64  `json:"int642"`
		Uint2   uint   `json:"uint2"`
		Uint82  uint8  `json:"uint82"`
		Uint162 uint16 `json:"uint162"`
		Uint322 uint32 `json:"uint322"`
		Uint642 uint64 `json:"uint642"`
		String2 string `json:"string2"`
		Bool2   bool   `json:"bool2"`
		Byte2   byte   `json:"byte2"`
	}

	m.sdk.Helper().ReceiptHelper().Emit(
		basic{
			Int2:    int2,
			Int82:   int82,
			Int162:  int162,
			Int322:  int322,
			Int642:  int642,
			Uint2:   uint2,
			Uint82:  uint82,
			Uint162: uint162,
			Uint322: uint322,
			Uint642: uint642,
			String2: string2,
			Bool2:   bool2,
			Byte2:   byte2,
		},
	)
}

//emitSlice This is a method of Mymixtype
func (m *Mymixtype) emitSlice(intS []int, int8S []int8, int16S []int16, int32S []int32, int64S []int64, uintS []uint, uint8S []uint8, uint16S []uint16, uint32S []uint32, uint64S []uint64, stringS []string, boolS []bool, byteS []byte) {
	type slice struct {
		IntS    []int    `json:"intS"`
		Int8S   []int8   `json:"int8S"`
		Int16S  []int16  `json:"int16S"`
		Int32S  []int32  `json:"int32S"`
		Int64S  []int64  `json:"int64S"`
		UintS   []uint   `json:"uintS"`
		Uint8S  []uint8  `json:"uint8S"`
		Uint16S []uint16 `json:"uint16S"`
		Uint32S []uint32 `json:"uint32S"`
		Uint64S []uint64 `json:"uint64S"`
		StringS []string `json:"stringS"`
		BoolS   []bool   `json:"boolS"`
		ByteS   []byte   `json:"byteS"`
	}

	m.sdk.Helper().ReceiptHelper().Emit(
		slice{
			IntS:    intS,
			Int8S:   int8S,
			Int16S:  int16S,
			Int32S:  int32S,
			Int64S:  int64S,
			UintS:   uintS,
			Uint8S:  uint8S,
			Uint16S: uint16S,
			Uint32S: uint32S,
			Uint64S: uint64S,
			StringS: stringS,
			BoolS:   boolS,
			ByteS:   byteS,
		},
	)
}

//emitMapString This is a method of Mymixtype
func (m *Mymixtype) emitMapString(int_ map[string]int, int8_ map[string]int8, int16_ map[string]int16, int32_ map[string]int32, int64_ map[string]int64, uint_ map[string]uint, uint8_ map[string]uint8, uint16_ map[string]uint16, uint32_ map[string]uint32, uint64_ map[string]uint64, bool_ map[string]bool, byte_ map[string]byte, string_ map[string]string) {
	type mapString struct {
		Int_    map[string]int    `json:"int_"`
		Int8_   map[string]int8   `json:"int8_"`
		Int16_  map[string]int16  `json:"int16_"`
		Int32_  map[string]int32  `json:"int32_"`
		Int64_  map[string]int64  `json:"int64_"`
		Uint_   map[string]uint   `json:"uint_"`
		Uint8_  map[string]uint8  `json:"uint8_"`
		Uint16_ map[string]uint16 `json:"uint16_"`
		Uint32_ map[string]uint32 `json:"uint32_"`
		Uint64_ map[string]uint64 `json:"uint64_"`
		Bool_   map[string]bool   `json:"bool_"`
		Byte_   map[string]byte   `json:"byte_"`
		String_ map[string]string `json:"string_"`
	}

	m.sdk.Helper().ReceiptHelper().Emit(
		mapString{
			Int_:    int_,
			Int8_:   int8_,
			Int16_:  int16_,
			Int32_:  int32_,
			Int64_:  int64_,
			Uint_:   uint_,
			Uint8_:  uint8_,
			Uint16_: uint16_,
			Uint32_: uint32_,
			Uint64_: uint64_,
			Bool_:   bool_,
			Byte_:   byte_,
			String_: string_,
		},
	)
}

//emitMapOther This is a method of Mymixtype
func (m *Mymixtype) emitMapOther(int_ map[int]string, int8_ map[int8]string, int16_ map[int16]string, int32_ map[int32]string, int64_ map[int64]string, uint_ map[uint]string, uint8_ map[uint8]string, uint16_ map[uint16]string, uint32_ map[uint32]string, uint64_ map[uint64]string, bool_ map[bool]string, byte_ map[byte]string, string_ map[string]string) {
	type mapOther struct {
		Int_    map[int]string    `json:"int_"`
		Int8_   map[int8]string   `json:"int8_"`
		Int16_  map[int16]string  `json:"int16_"`
		Int32_  map[int32]string  `json:"int32_"`
		Int64_  map[int64]string  `json:"int64_"`
		Uint_   map[uint]string   `json:"uint_"`
		Uint8_  map[uint8]string  `json:"uint8_"`
		Uint16_ map[uint16]string `json:"uint16_"`
		Uint32_ map[uint32]string `json:"uint32_"`
		Uint64_ map[uint64]string `json:"uint64_"`
		Bool_   map[bool]string   `json:"bool_"`
		Byte_   map[byte]string   `json:"byte_"`
		String_ map[string]string `json:"string_"`
	}

	m.sdk.Helper().ReceiptHelper().Emit(
		mapOther{
			Int_:    int_,
			Int8_:   int8_,
			Int16_:  int16_,
			Int32_:  int32_,
			Int64_:  int64_,
			Uint_:   uint_,
			Uint8_:  uint8_,
			Uint16_: uint16_,
			Uint32_: uint32_,
			Uint64_: uint64_,
			Bool_:   bool_,
			Byte_:   byte_,
			String_: string_,
		},
	)
}

//emitMapSlice This is a method of Mymixtype
func (m *Mymixtype) emitMapSlice(int_ map[string][]int, int8_ map[string][]int8, int16_ map[string][]int16, int32_ map[string][]int32, int64_ map[string][]int64, uint_ map[string][]uint, uint8_ map[string][]uint8, uint16_ map[string][]uint16, uint32_ map[string][]uint32, uint64_ map[string][]uint64, bool_ map[string][]bool, byte_ map[string][]byte, string_ map[string][]string) {
	type mapSlice struct {
		Int_    map[string][]int    `json:"int_"`
		Int8_   map[string][]int8   `json:"int8_"`
		Int16_  map[string][]int16  `json:"int16_"`
		Int32_  map[string][]int32  `json:"int32_"`
		Int64_  map[string][]int64  `json:"int64_"`
		Uint_   map[string][]uint   `json:"uint_"`
		Uint8_  map[string][]uint8  `json:"uint8_"`
		Uint16_ map[string][]uint16 `json:"uint16_"`
		Uint32_ map[string][]uint32 `json:"uint32_"`
		Uint64_ map[string][]uint64 `json:"uint64_"`
		Bool_   map[string][]bool   `json:"bool_"`
		Byte_   map[string][]byte   `json:"byte_"`
		String_ map[string][]string `json:"string_"`
	}

	m.sdk.Helper().ReceiptHelper().Emit(
		mapSlice{
			Int_:    int_,
			Int8_:   int8_,
			Int16_:  int16_,
			Int32_:  int32_,
			Int64_:  int64_,
			Uint_:   uint_,
			Uint8_:  uint8_,
			Uint16_: uint16_,
			Uint32_: uint32_,
			Uint64_: uint64_,
			Bool_:   bool_,
			Byte_:   byte_,
			String_: string_,
		},
	)
}

//emitComplexT This is a method of Mymixtype
func (m *Mymixtype) emitComplexT(address types.Address, hash types.Hash, hexBytes types.HexBytes, pubKey types.PubKey, number bn.Number) {
	type complexT struct {
		Address  types.Address  `json:"address"`
		Hash     types.Hash     `json:"hash"`
		HexBytes types.HexBytes `json:"hexBytes"`
		PubKey   types.PubKey   `json:"pubKey"`
		Number   bn.Number      `json:"number"`
	}

	m.sdk.Helper().ReceiptHelper().Emit(
		complexT{
			Address:  address,
			Hash:     hash,
			HexBytes: hexBytes,
			PubKey:   pubKey,
			Number:   number,
		},
	)
}

//emitLongT This is a method of Mymixtype
func (m *Mymixtype) emitLongT(basic BasicTypes, slice SliceTypes, map1 MapStringTypes, map2 MapOtherTypes, complex ComplexDefine) {
	type longT struct {
		Basic   BasicTypes     `json:"basic"`
		Slice   SliceTypes     `json:"slice"`
		Map1    MapStringTypes `json:"map1"`
		Map2    MapOtherTypes  `json:"map2"`
		Complex ComplexDefine  `json:"complex"`
	}

	m.sdk.Helper().ReceiptHelper().Emit(
		longT{
			Basic:   basic,
			Slice:   slice,
			Map1:    map1,
			Map2:    map2,
			Complex: complex,
		},
	)
}
