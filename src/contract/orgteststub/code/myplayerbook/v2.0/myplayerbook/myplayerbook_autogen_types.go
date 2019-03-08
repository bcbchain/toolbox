package myplayerbook

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//RegisterNameParam structure of parameters of RegisterName() of v2.0
type RegisterNameParam struct {
	Index int64
	Plyr  Player
}

//GetPlayerParam structure of parameters of RegisterName() of v2.0
type GetPlayerParam struct {
	Addr types.Address
}

//MultiTypesParam structure of parameters of RegisterName() of v2.0
type MultiTypesParam struct {
	Index uint64
	Flt   float64
	Bl    bool
	Bt    byte
	Hash  types.Hash
	Hb    types.HexBytes
	Bi    bn.Number
	Mp    map[int]string
}
