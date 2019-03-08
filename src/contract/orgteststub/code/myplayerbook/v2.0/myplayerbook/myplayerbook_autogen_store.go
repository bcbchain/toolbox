package myplayerbook

import (
	"blockchain/smcsdk/sdk/types"
	"fmt"
)

//_setPlyr This is a method of MyPlayerBook
func (mpb *MyPlayerBook) _setPlyr(k types.Address, v Player) {
	mpb.sdk.Helper().StateHelper().McSet(fmt.Sprintf("/plyr/%v", k), &v)
}

//_plyr This is a method of MyPlayerBook
func (mpb *MyPlayerBook) _plyr(k types.Address) Player {

	return *mpb.sdk.Helper().StateHelper().McGetEx(fmt.Sprintf("/plyr/%v", k), new(Player)).(*Player)
}

//_clrPlyr This is a method of MyPlayerBook
func (mpb *MyPlayerBook) _clrPlyr(k types.Address) {
	mpb.sdk.Helper().StateHelper().McClear(fmt.Sprintf("/plyr/%v", k))
}

//_chkPlyr This is a method of MyPlayerBook
func (mpb *MyPlayerBook) _chkPlyr(k types.Address) bool {
	return mpb.sdk.Helper().StateHelper().Check(fmt.Sprintf("/plyr/%v", k))
}

//_McChkPlyr This is a method of MyPlayerBook
func (mpb *MyPlayerBook) _McChkPlyr(k types.Address) bool {
	return mpb.sdk.Helper().StateHelper().McCheck(fmt.Sprintf("/plyr/%v", k))
}
