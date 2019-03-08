package mymixtype

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

type BasicTypes struct {
	Int_    int
	Int8_   int8
	Int16_  int16
	Int32_  int32
	Int64_  int64
	Uint_   uint
	Uint8_  uint8
	Uint16_ uint16
	Uint32_ uint32
	Uint64_ uint64
	Bool_   bool
	Byte_   byte
	String_ string
}

type SliceTypes struct {
	IntS     []int
	Int8S    []int8
	Int16S   []int16
	Int32S   []int32
	Int64S   []int64
	UintS    []uint
	Uint8S   []uint8
	Uint16S  []uint16
	Uint32S  []uint32
	Uint64S  []uint64
	BoolS    []bool
	ByteS    []byte
	StringS_ []string
}

type MapStringTypes struct {
	Mint    map[string]int
	Mint8   map[string]int8
	Mint16  map[string]int16
	Mint32  map[string]int32
	Mint64  map[string]int64
	Muint   map[string]uint
	Muint8  map[string]uint8
	Muint16 map[string]uint16
	Muint32 map[string]uint32
	Muint64 map[string]uint64
	Mbool   map[string]bool
	Mbyte   map[string]byte
	Mstring map[string]string
}

func (m *MapStringTypes) Init() {
	m.Mint = make(map[string]int)
	m.Mint8 = make(map[string]int8)
	m.Mint16 = make(map[string]int16)
	m.Mint32 = make(map[string]int32)
	m.Mint64 = make(map[string]int64)
	m.Muint = make(map[string]uint)
	m.Muint8 = make(map[string]uint8)
	m.Muint16 = make(map[string]uint16)
	m.Muint32 = make(map[string]uint32)
	m.Muint64 = make(map[string]uint64)
	m.Mbool = make(map[string]bool)
	m.Mbyte = make(map[string]byte)
	m.Mstring = make(map[string]string)
}

type MapOtherTypes struct {
	Mint    map[int]string
	Mint8   map[int8]string
	Mint16  map[int16]string
	Mint32  map[int32]string
	Mint64  map[int64]string
	Muint   map[uint]string
	Muint8  map[uint8]string
	Muint16 map[uint16]string
	Muint32 map[uint32]string
	Muint64 map[uint64]string
	Mbool   map[bool]string
	Mbyte   map[byte]string
	Mstring map[string]string
}

func (m *MapOtherTypes) Init() {
	m.Mint = make(map[int]string)
	m.Mint8 = make(map[int8]string)
	m.Mint16 = make(map[int16]string)
	m.Mint32 = make(map[int32]string)
	m.Mint64 = make(map[int64]string)
	m.Muint = make(map[uint]string)
	m.Muint8 = make(map[uint8]string)
	m.Muint16 = make(map[uint16]string)
	m.Muint32 = make(map[uint32]string)
	m.Muint64 = make(map[uint64]string)
	m.Mbool = make(map[bool]string)
	m.Mbyte = make(map[byte]string)
	m.Mstring = make(map[string]string)
}

type MapSliceTypes struct {
	Mint    map[string][]int
	Mint8   map[string][]int8
	Mint16  map[string][]int16
	Mint32  map[string][]int32
	Mint64  map[string][]int64
	Muint   map[string][]uint
	Muint8  map[string][]uint8
	Muint16 map[string][]uint16
	Muint32 map[string][]uint32
	Muint64 map[string][]uint64
	Mbool   map[string][]bool
	Mbyte   map[string][]byte
	Mstring map[string][]string
}

func (m *MapSliceTypes) Init() {
	m.Mint = make(map[string][]int)
	m.Mint8 = make(map[string][]int8)
	m.Mint16 = make(map[string][]int16)
	m.Mint32 = make(map[string][]int32)
	m.Mint64 = make(map[string][]int64)
	m.Muint = make(map[string][]uint)
	m.Muint8 = make(map[string][]uint8)
	m.Muint16 = make(map[string][]uint16)
	m.Muint32 = make(map[string][]uint32)
	m.Muint64 = make(map[string][]uint64)
	m.Mbool = make(map[string][]bool)
	m.Mbyte = make(map[string][]byte)
	m.Mstring = make(map[string][]string)
}

type ComplexDefine struct {
	Address_  types.Address
	Hash_     types.Hash
	HexBytes_ types.HexBytes
	PubKey_   types.PubKey
	Number_   bn.Number
}

type LongType struct {
	Basic_   BasicTypes
	Slice_   SliceTypes
	Map1_    MapStringTypes
	Map2_    MapOtherTypes
	Complex_ ComplexDefine
}
