package mymixtype

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
)

//@:contract:mymixtype
//@:version:1.0
//@:organization:orgGyRrMVF7ukfHNwaZhgWMTbQAYz7d7RcBh
//@:author:b37e7627431feb18123b81bcf1f41ffd37efdb90513d48ff2c7f8a0c27a9d06c
type Mymixtype struct {
	sdk sdk.ISmartContract

	//@:public:store:cache
	basic BasicTypes

	//@:public:store
	slice SliceTypes

	//@:public:store:cache
	map1 MapStringTypes

	//@:public:store:cache
	map2 MapOtherTypes

	//@:public:store:cache
	map3 MapSliceTypes

	//@:public:store
	complex ComplexDefine

	//@:public:store
	long LongType
}

//@:public:receipt
type receipt interface {
	emitBasic(int2 int, int82 int8, int162 int16, int322 int32, int642 int64, uint2 uint, uint82 uint8, uint162 uint16, uint322 uint32, uint642 uint64, string2 string, bool2 bool, byte2 byte)
	emitSlice(intS []int, int8S []int8, int16S []int16, int32S []int32, int64S []int64, uintS []uint, uint8S []uint8, uint16S []uint16, uint32S []uint32, uint64S []uint64, stringS []string, boolS []bool, byteS []byte)
	emitMapString(int_ map[string]int, int8_ map[string]int8, int16_ map[string]int16, int32_ map[string]int32, int64_ map[string]int64, uint_ map[string]uint, uint8_ map[string]uint8, uint16_ map[string]uint16, uint32_ map[string]uint32, uint64_ map[string]uint64, bool_ map[string]bool, byte_ map[string]byte, string_ map[string]string)
	emitMapOther(int_ map[int]string, int8_ map[int8]string, int16_ map[int16]string, int32_ map[int32]string, int64_ map[int64]string, uint_ map[uint]string, uint8_ map[uint8]string, uint16_ map[uint16]string, uint32_ map[uint32]string, uint64_ map[uint64]string, bool_ map[bool]string, byte_ map[byte]string, string_ map[string]string)
	emitMapSlice(int_ map[string][]int, int8_ map[string][]int8, int16_ map[string][]int16, int32_ map[string][]int32, int64_ map[string][]int64, uint_ map[string][]uint, uint8_ map[string][]uint8, uint16_ map[string][]uint16, uint32_ map[string][]uint32, uint64_ map[string][]uint64, bool_ map[string][]bool, byte_ map[string][]byte, string_ map[string][]string)
	emitComplexT(address types.Address, hash types.Hash, hexBytes types.HexBytes, pubKey types.PubKey, number bn.Number)
	emitLongT(basic BasicTypes, slice SliceTypes, map1 MapStringTypes, map2 MapOtherTypes, complex ComplexDefine)
}

//@:constructor
func (m *Mymixtype) InitChain() {

}

//@:public:method:gas[500]
func (m *Mymixtype) Basic(basic_ *BasicTypes) (data []byte) {
	data, _ = jsoniter.Marshal(basic_)

	m._setBasic(*basic_)

	m.emitBasic(
		basic_.Int_,
		basic_.Int8_,
		basic_.Int16_,
		basic_.Int32_,
		basic_.Int64_,
		basic_.Uint_,
		basic_.Uint8_,
		basic_.Uint16_,
		basic_.Uint32_,
		basic_.Uint64_,
		basic_.String_,
		basic_.Bool_,
		basic_.Byte_,
	)

	return
}

//@:public:method:gas[500]
func (m *Mymixtype) Slice(slice_ *SliceTypes) (data []byte) {
	data, _ = jsoniter.Marshal(slice_)

	m._setSlice(*slice_)

	m.emitSlice(
		slice_.IntS,
		slice_.Int8S,
		slice_.Int16S,
		slice_.Int32S,
		slice_.Int64S,
		slice_.UintS,
		slice_.Uint8S,
		slice_.Uint16S,
		slice_.Uint32S,
		slice_.Uint64S,
		slice_.StringS_,
		slice_.BoolS,
		slice_.ByteS,
	)

	return
}

//@:public:method:gas[500]
func (m *Mymixtype) MapString(map_ *MapStringTypes) (data []byte) {
	data, _ = jsoniter.Marshal(map_)

	m._setMap1(*map_)

	m.emitMapString(
		map_.Mint,
		map_.Mint8,
		map_.Mint16,
		map_.Mint32,
		map_.Mint64,
		map_.Muint,
		map_.Muint8,
		map_.Muint16,
		map_.Muint32,
		map_.Muint64,
		map_.Mbool,
		map_.Mbyte,
		map_.Mstring,
	)

	return
}

//@:public:method:gas[500]
func (m *Mymixtype) MapOther(map_ *MapOtherTypes) (data []byte) {
	data, _ = jsoniter.Marshal(map_)

	m._setMap2(*map_)

	m.emitMapOther(
		map_.Mint,
		map_.Mint8,
		map_.Mint16,
		map_.Mint32,
		map_.Mint64,
		map_.Muint,
		map_.Muint8,
		map_.Muint16,
		map_.Muint32,
		map_.Muint64,
		map_.Mbool,
		map_.Mbyte,
		map_.Mstring,
	)

	return
}

//@:public:method:gas[500]
func (m *Mymixtype) MapSlice(map_ *MapSliceTypes) (data []byte) {
	data, _ = jsoniter.Marshal(map_)

	m._setMap3(*map_)

	m.emitMapSlice(
		map_.Mint,
		map_.Mint8,
		map_.Mint16,
		map_.Mint32,
		map_.Mint64,
		map_.Muint,
		map_.Muint8,
		map_.Muint16,
		map_.Muint32,
		map_.Muint64,
		map_.Mbool,
		map_.Mbyte,
		map_.Mstring,
	)

	return
}

//@:public:method:gas[500]
func (m *Mymixtype) Complex(complex_ *ComplexDefine) (data []byte) {
	data, _ = jsoniter.Marshal(complex_)

	m._setComplex(*complex_)

	m.emitComplexT(
		complex_.Address_,
		complex_.Hash_,
		complex_.HexBytes_,
		complex_.PubKey_,
		complex_.Number_,
	)

	return
}

//@:public:method:gas[500]
func (m *Mymixtype) Long(long_ *LongType) (data []byte) {
	data, _ = jsoniter.Marshal(long_)

	m._setLong(*long_)

	m.emitLongT(
		long_.Basic_,
		long_.Slice_,
		long_.Map1_,
		long_.Map2_,
		long_.Complex_,
	)

	return
}
