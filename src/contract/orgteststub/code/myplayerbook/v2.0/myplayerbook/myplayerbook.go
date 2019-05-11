package myplayerbook

import (
	"bytes"
	"fmt"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//MyPlayerBook a demo of players management
//@:contract:myplayerbook
//@:version:2.0
//@:organization:orgGyRrMVF7ukfHNwaZhgWMTbQAYz7d7RcBh
//@:author:b37e7627431feb18123b81bcf1f41ffd37efdb90513d48ff2c7f8a0c27a9d06c
type MyPlayerBook struct {
	sdk sdk.ISmartContract

	//@:public:store:cache
	//@:comment:(address => data) player data
	plyr map[types.Address]Player
}

//Player player struct
type Player struct {
	Address types.Address
	Name    string
}

type myType string

const registrationFee = 1000000 // unit in "cong"

//InitChain init when deployed on the blockchain first time
//@:constructor
func (pb *MyPlayerBook) InitChain() {
}

//GetPlayer gets player's data
//@:public:method:gas[100]
//@:public:interface:gas[100]
func (pb *MyPlayerBook) GetPlayer(addr types.Address) string {
	plyr := pb._plyr(addr)
	return plyr.Name
}

//RegisterName register a player with name
//@:public:method:gas[500]
//@:public:interface:gas[500]
func (pb *MyPlayerBook) RegisterName(index int64, plyr Player) {

	var a myType

	// 检查收据，确认已支付注册费
	sdk.Require(a.bRegistrationFee(pb.sdk.Message()) == true,
		types.ErrUserDefined, "")

	name, _err := nameFilter(plyr.Name)
	sdk.Require(
		_err == "" && name != "",
		types.ErrInvalidParameter, "")

	pb._setPlyr(plyr.Address, plyr)
}

//MultiTypesParam test parameters
//@:public:method:gas[500]
//@:public:interface:gas[500]
func (pb *MyPlayerBook) MultiTypesParam(index uint64, flt float64, bl bool, bt byte, hash types.Hash, hb types.HexBytes, bi bn.Number, mp map[int]string) {
	fmt.Println(index, flt, bl, bt, hash, hb, bi, mp)
}

func (m *myType) bRegistrationFee(smcAPI sdk.IMessage) bool {
	feeReceipts := smcAPI.GetTransferToMe()
	for _, feeReceipt := range feeReceipts {
		if feeReceipt != nil && feeReceipt.Value.CmpI(registrationFee) >= 0 {
			return true
		}
	}
	return false
}

/**
* @dev filters name strings
* -converts uppercase to lower case.
* -makes sure it does not start/end with a space
* -makes sure it does not contain multiple spaces in a row
* -cannot be only numbers
* -cannot start with 0x
* -restricts characters to A-Z, a-z, 0-9, and space.
* @return reprocessed string in bytes32 format
 */
func nameFilter(_input string) (_name string, _err string) { // nolint gocyclo
	_temp := []byte(_input)
	_length := len(_temp)

	//sorry limited to 32 characters
	if !(_length <= 32 && _length > 0) {
		_err = "string must be between 1 and 32 characters"
		return
	}
	// make sure it doesnt start with or end with space
	if !(_temp[0] != 0x20 && _temp[_length-1] != 0x20) {
		_err = "string cannot start or end with space"
		return
	}
	// make sure first two characters are not 0x
	if len(_temp) >= 2 {
		if _temp[0] == 0x30 {
			if !(_temp[1] != 0x78) {
				_err = "string cannot start with 0x"
				return
			}
			if !(_temp[1] != 0x58) {
				_err = "string cannot start with 0X"
				return
			}
		}
	}
	if _temp[0] >= '0' && _temp[0] <= '9' {
		_err = "string cannot start with digit"
		return
	}

	if len(_temp) >= 3 {
		_tempHeader := bytes.ToLower(_temp)[:3]
		if _tempHeader[0] == 'b' && _tempHeader[1] == 'c' && _tempHeader[2] == 'b' {
			_err = "string cannot start with bcb address prefix"
			return
		}
	}

	// create a bool to track if we have a non number character
	_hasFullNumber := true

	// convert & check
	for i := 0; i < _length; i++ {
		// if its uppercase A-Z
		if _temp[i] >= 'A' && _temp[i] <= 'Z' {
			// convert to lower case a-z
			_temp[i] = byte(uint(_temp[i]) + 32)

			// we have a non number
			_hasFullNumber = false
		} else {
			// require character is a space
			// OR lowercase a-z
			// or 0-9
			if !(_temp[i] == ' ' ||
				(_temp[i] >= 'a' && _temp[i] <= 'z') ||
				(_temp[i] >= '0' && _temp[i] <= '9')) {
				_err = "string contains invalid characters"
				return
			}

			// make sure theres not 2x spaces in a row
			if _temp[i] == ' ' {
				if _temp[i+1] == ' ' {
					_err = "string cannot contain consecutive spaces"
					return
				}
			}

			// see if we have a character other than a number
			if _hasFullNumber == true && (_temp[i] < '0' || _temp[i] > '9') {
				_hasFullNumber = false
			}
		}
	}

	if _hasFullNumber == true {
		_err = "string cannot be only numbers"
		return
	}

	_name = string(_temp)
	return
}
