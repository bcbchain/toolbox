package myplayerbook

import (
	"fmt"

	"blockchain/smcsdk/sdk/types"
)

func (pb *MyPlayerBook) _pID() int64 {
	return pb.sdk.Helper().StateHelper().McGetInt64("/lastplayerid")
}

func (pb *MyPlayerBook) _setPID(v int64) {
	pb.sdk.Helper().StateHelper().McSetInt64("/lastplayerid", v)
}

func (pb *MyPlayerBook) _pIDxAddr(k types.Address) int64 {
	return pb.sdk.Helper().StateHelper().McGetInt64("/pidxaddress/" + k)
}

func (pb *MyPlayerBook) _setPIDxAddr(k types.Address, v int64) {
	pb.sdk.Helper().StateHelper().McSetInt64("/pidxaddress/"+k, v)
}

func (pb *MyPlayerBook) _player(k int64) *Player {
	return pb.sdk.Helper().StateHelper().McGet(fmt.Sprintf("/players/%v", k), &Player{}).(*Player)
}

func (pb *MyPlayerBook) _setPlayer(k int64, v *Player) {
	pb.sdk.Helper().StateHelper().McSet(fmt.Sprintf("/players/%v", k), v)
}

//Variable：pIDxName_
// (name => pID) returns player id by name
func (pb *MyPlayerBook) _pIDxName(k string) int64 {
	return pb.sdk.Helper().StateHelper().McGetInt64("/pidxname/" + k)
}

func (pb *MyPlayerBook) _setPIDxName(k string, v int64) {
	pb.sdk.Helper().StateHelper().McSetInt64("/pidxname/"+k, v)
}

func (pb *MyPlayerBook) _plyrNames(k int64) []string {
	return pb.sdk.Helper().StateHelper().McGetStrings(fmt.Sprintf("/playername/%d", k))
}

func (pb *MyPlayerBook) _setPlyrNames(k int64, v []string) {
	pb.sdk.Helper().StateHelper().McSetStrings(fmt.Sprintf("/playername/%d", k), v)
}
