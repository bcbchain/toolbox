package mybasictype

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
)

// BasicType test contract
//@:contract:mybasictype
//@:version:1.0
//@:organization:orgGyRrMVF7ukfHNwaZhgWMTbQAYz7d7RcBh
//@:author:b37e7627431feb18123b81bcf1f41ffd37efdb90513d48ff2c7f8a0c27a9d06c
type BasicType struct {
	sdk sdk.ISmartContract

	// type define by self
	//@:public:store:cache
	address_ types.Address
	//@:public:store:cache
	hash_ types.Hash
	//@:public:store:cache
	hexBytes_ types.HexBytes
	//@:public:store:cache
	pubkey_ types.PubKey
	//@:public:store:cache
	number_ bn.Number

	// basic types in golang
	//@:public:store
	int_ int
	//@:public:store
	int8_ int8
	//@:public:store
	int16_ int16
	//@:public:store
	int32_ int32
	//@:public:store
	int64_ int64
	//@:public:store
	uint_ uint
	//@:public:store
	uint8_ uint8
	//@:public:store
	uint16_ uint16
	//@:public:store
	uint32_ uint32
	//@:public:store
	uint64_ uint64
	//@:public:store
	string_ string
	//@:public:store
	byte_ byte
	//@:public:store
	bool_ bool

	// advance type in golang
	//@:public:store
	map_ map[string]int
	//@:public:store
	map1_ map[uint]map[string]types.Address
	//@:public:store
	map2_ map[int]map[int8]types.Hash
	//@:public:store
	map3_ map[int]map[uint64]types.HexBytes
	//@:public:store
	map4_ map[bool]map[byte]types.PubKey
	//@:public:store
	map5_ map[bool]map[bn.Number]bool
	//@:public:store
	map6_ map[byte]map[string]bn.Number
	//@:public:store
	map7_ map[string]map[types.Address]byte
	//@:public:store
	map8_ map[types.Address]map[bn.Number]string
	//@:public:store
	slice_ []byte
}

//@:public:receipt
type receipt interface {
	emitAddressT(address_ types.Address)
	emitHashT(hash_ types.Hash)
	emitHexBytesT(hexBytes_ types.HexBytes)
	emitPubKeyT(pubKey_ types.PubKey)
	emitNumberT(number_ bn.Number)
	emitIntT(int_ int)
	emitInt8T(int8_ int8)
	emitInt16T(int16_ int16)
	emitInt32T(int32_ int32)
	emitInt64T(int64_ int64)
	emitUintT(uint_ uint)
	emitUint8T(uint8_ uint8)
	emitUint16T(uint16_ uint16)
	emitUint32T(uint32_ uint32)
	emitUint64T(uint64_ uint64)
	emitBoolT(bool_ bool)
	emitByteT(byte_ byte)
	emitSliceT(slice_ []byte)
	emitMapT(map_ map[string]int)
	emitMap1T(map_ map[uint]map[string]types.Address)
	emitMap2T(map_ map[int]map[int8]types.Hash)
	emitMap3T(map_ map[int]map[uint64]types.HexBytes)
	emitMap4T(map_ map[bool]map[byte]types.PubKey)
	emitMap5T(map_ map[bool]map[bn.Number]bool)
	emitMap6T(map_ map[byte]map[string]bn.Number)
	emitMap7T(map_ map[string]map[types.Address]byte)
	emitMap8T(map_ map[types.Address]map[bn.Number]string)
}

// InitChain init function
//@:constructor
func (b *BasicType) InitChain() {
}

// EchoAddress echo address
//@:public:method:gas[500]
func (b *BasicType) EchoAddress(v types.Address) (value types.Address) {
	b._setAddress_(v)

	value = v

	// fire event
	b.emitAddressT(v)

	return
}

// EchoHash echo hash
//@:public:method:gas[500]
func (b *BasicType) EchoHash(v types.Hash) (value types.Hash) {
	b._setHash_(v)

	value = v

	// fire event
	b.emitHashT(v)

	return
}

// EchoHexBytes echo hexBytes
//@:public:method:gas[500]
func (b *BasicType) EchoHexBytes(v types.HexBytes) (value types.HexBytes) {
	b._setHexBytes_(v)

	value = v

	// fire event
	b.emitHexBytesT(v)

	return
}

// EchoPubKey echo pubKey
//@:public:method:gas[500]
func (b *BasicType) EchoPubKey(v types.PubKey) (value types.PubKey) {
	b._setHexBytes_(v)

	value = v

	// fire event
	b.emitHexBytesT(v)

	return
}

// EchoNumber echo number
//@:public:method:gas[500]
func (b *BasicType) EchoNumber(v bn.Number) (value bn.Number) {
	b._setNumber_(v)

	value = v

	// fire event
	b.emitNumberT(v)

	return
}

// EchoInt echo int
//@:public:method:gas[500]
func (b *BasicType) EchoInt(v int) (value int) {
	b._setInt_(v)

	value = v

	// fire event
	b.emitIntT(v)

	return
}

// EchoInt8 echo int8
//@:public:method:gas[500]
func (b *BasicType) EchoInt8(v int8) (value int8) {
	b._setInt8_(v)

	value = v

	// fire event
	b.emitInt8T(v)

	return
}

// EchoInt16 echo int16
//@:public:method:gas[500]
func (b *BasicType) EchoInt16(v int16) (value int16) {
	b._setInt16_(v)

	value = v

	// fire event
	b.emitInt16T(v)

	return
}

// EchoInt32 echo int32
//@:public:method:gas[500]
func (b *BasicType) EchoInt32(v int32) (value int32) {
	b._setInt32_(v)

	value = v

	// fire event
	b.emitInt32T(v)

	return
}

// EchoInt64 echo int64
//@:public:method:gas[500]
func (b *BasicType) EchoInt64(v int64) (value int64) {
	b._setInt64_(v)

	value = v

	// fire event
	b.emitInt64T(v)

	return
}

// EchoUint echo uint
//@:public:method:gas[500]
func (b *BasicType) EchoUint(v uint) (value uint) {
	b._setUint_(v)

	value = v

	// fire event
	b.emitUintT(v)

	return
}

// EchoUint8 echo uint8
//@:public:method:gas[500]
func (b *BasicType) EchoUint8(v uint8) (value uint8) {
	b._setUint8_(v)

	value = v

	// fire event
	b.emitUint8T(v)

	return
}

// EchoUint16 echo uint16
//@:public:method:gas[500]
func (b *BasicType) EchoUint16(v uint16) (value uint16) {
	b._setUint16_(v)

	value = v

	// fire event
	b.emitUint16T(v)

	return
}

// EchoUint32 echo uint32
//@:public:method:gas[500]
func (b *BasicType) EchoUint32(v uint32) (value uint32) {
	b._setUint32_(v)

	value = v

	// fire event
	b.emitUint32T(v)

	return
}

// EchoUint64 echo uint64
//@:public:method:gas[500]
func (b *BasicType) EchoUint64(v uint64) (value uint64) {
	b._setUint64_(v)

	value = v

	// fire event
	b.emitUint64T(v)

	return
}

// EchoBool echo bool
//@:public:method:gas[500]
func (b *BasicType) EchoBool(v bool) (value bool) {
	b._setBool_(v)

	value = v

	// fire event
	b.emitBoolT(v)

	return
}

// EchoByte echo byte
//@:public:method:gas[500]
func (b *BasicType) EchoByte(v byte) (value byte) {
	b._setByte_(v)

	value = v

	// fire event
	b.emitByteT(v)

	return
}

// EchoBytes echo bytes
//@:public:method:gas[500]
func (b *BasicType) EchoBytes(v []byte) (value []byte) {
	b._setSlice_(v)

	value = v

	// fire event
	b.emitSliceT(v)

	return
}

// EchoMap echo map
//@:public:method:gas[500]
func (b *BasicType) EchoMap(v map[string]int) (value []byte) {
	for k, v := range v {
		b._setMap_(k, v)
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMapT(v)

	return
}

// EchoMap1 echo map1
//@:public:method:gas[500]
func (b *BasicType) EchoMap1(v map[uint]map[string]types.Address) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap1_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap1T(v)

	return
}

// EchoMap2 echo map2
//@:public:method:gas[500]
func (b *BasicType) EchoMap2(v map[int]map[int8]types.Hash) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap2_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap2T(v)

	return
}

// EchoMap3 echo map3
//@:public:method:gas[500]
func (b *BasicType) EchoMap3(v map[int]map[uint64]types.HexBytes) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap3_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap3T(v)

	return
}

// EchoMap4 echo map4
//@:public:method:gas[500]
func (b *BasicType) EchoMap4(v map[bool]map[byte]types.PubKey) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap4_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap4T(v)

	return
}

// EchoMap5 echo map5
//@:public:method:gas[500]
func (b *BasicType) EchoMap5(v map[bool]map[bn.Number]bool) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap5_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap5T(v)

	return
}

// EchoMap6 echo map6
//@:public:method:gas[500]
func (b *BasicType) EchoMap6(v map[byte]map[string]bn.Number) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap6_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap6T(v)

	return
}

// EchoMap7 echo map7
//@:public:method:gas[500]
func (b *BasicType) EchoMap7(v map[string]map[types.Address]byte) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap7_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap7T(v)

	return
}

// EchoMap8 echo map8
//@:public:method:gas[500]
func (b *BasicType) EchoMap8(v map[types.Address]map[bn.Number]string) (value []byte) {
	for k, v := range v {
		for k1, v1 := range v {
			b._setMap8_(k, k1, v1)
		}
	}

	resBytes, _ := jsoniter.Marshal(v)
	value = resBytes

	// fire event
	b.emitMap8T(v)

	return
}
