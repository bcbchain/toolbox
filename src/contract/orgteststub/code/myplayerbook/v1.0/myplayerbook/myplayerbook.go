package myplayerbook

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"bytes"
)

//MyPlayerBook a simple contract for demo
//@:contract:myplayerbook
//@:comment:玩家信息管理合约
//@:version:1.0
//@:organization:orgGyRrMVF7ukfHNwaZhgWMTbQAYz7d7RcBh
//@:author:b37e7627431feb18123b81bcf1f41ffd37efdb90513d48ff2c7f8a0c27a9d06c
type MyPlayerBook struct {
	sdk sdk.ISmartContract

	//@:public:store:cache
	//@:comment: total number of players
	pID int64
	//@:public:store:cache
	//@:comment:(addr => pID) returns player id by address
	pIDxAddr map[types.Address]int64
	//@:public:store:cache
	//@:comment:(name => pID) returns player id by name
	pIDxName map[string]int64
	//@:public:store:cache
	//@:comment:(pID => data) player data
	plyr map[int64]Player
	//@:public:store:cache
	//@:comment:(pID => names) list of names a player owns.  (used so you can change your display name amongst any name you own)
	plyrNames map[int64][]string
}

//Player player structure
type Player struct {
	Addr types.Address // player address
	Name string        // player name
	Win  bn.Number     // winnings vault
	Gen  bn.Number     // general vault
	Aff  bn.Number     // affiliate vault
	Lrnd int64         // last round played
	Laff int64         // last affiliate id used
}

//InitChain init while
func (pb *MyPlayerBook) InitChain() {
	pb.pIDxAddr = make(map[types.Address]int64)
	pb.pIDxName = make(map[string]int64)
	pb.plyr = make(map[int64]Player)
	pb.plyrNames = make(map[int64][]string)
}

//GetPlayer get player information
//@:public:method:gas[100]
//@:export:interface:gas[100]
func (pb *MyPlayerBook) GetPlayer(_pID int64) *Player {
	return pb._player(_pID)
}

//GetPlayerID get player id
//@:public:method:gas[100]
//@:export:interface:gas[100]
func (pb *MyPlayerBook) GetPlayerID(addr types.Address) int64 {
	return pb._pIDxAddr(addr)
}

//GetPlayerName get player name
//@:public:method:gas[100]
//@:export:interface:gas[100]
func (pb *MyPlayerBook) GetPlayerName(_pID int64) string {
	return pb._player(_pID).Name
}

//GetPlayerLAff get player
//@:public:method:gas[100]
//@:export:interface:gas[100]
func (pb *MyPlayerBook) GetPlayerLAff(_pID int64) int64 {
	return pb._player(_pID).Laff
}

//GetPlayerAddr get player address
//@:public:method:gas[100]
//@:export:interface:gas[100]
func (pb *MyPlayerBook) GetPlayerAddr(_pID int64) types.Address {
	return pb._player(_pID).Addr
}

//RegisterNameXid register name
//@:public:method:gas[500]
//@:export:interface:gas[500]
func (pb *MyPlayerBook) RegisterNameXid(_nameString string, _affCode int64) (err types.Error) {

	//TODO: 检查收据，确认已支付注册费

	_name, _err := nameFilter(_nameString)
	sdk.Require(
		_err == "" && _name != "",
		types.ErrInvalidParameter, "")

	sdk.Require(
		pb._pIDxName(_name) == 0,
		types.ErrUserDefined, "Dup name with any other player")

	_addr := pb.sdk.Message().Sender().Address()
	//_paid := _bcb
	//_now := pb.sdk.Block().Now()
	_isNewPlayer := pb.determinePID(_addr)

	// manage affiliate residuals
	// if no affiliate code was given, no new affiliate code was given, or the
	// player tried to use their own pID as an affiliate code, lolz
	_pID := pb.pIDxAddr[_addr]
	_plyr := pb.plyr[_pID]
	if _affCode != 0 && _affCode != _plyr.Laff && _affCode != _pID {
		// update last affiliate
		_plyr.Laff = _affCode
	} else if _affCode == _pID {
		_affCode = 0
	}
	return pb.registerNameCore(_pID, _addr, _affCode, _name, _isNewPlayer, true)
}

func (pb *MyPlayerBook) registerNameCore(_pID int64, _addr types.Address, _affID int64, _name string, _isNewPlayer, _all bool) (err types.Error) {
	// if names already has been used, require that current msg sender owns the name
	sdk.Require(pb.determinePName(_pID, _name) == false,
		types.ErrInvalidParameter, "")
	// add name to player profile, registry, and name book
	pb._setPID(_pID)
	pb._setPIDxAddr(_addr, _pID)

	_plyr := pb.plyr[_pID]
	_plyr.Name = _name
	pb._setPlayer(_pID, &_plyr)

	pb.pIDxName[_name] = _pID
	pb._setPIDxName(_name, _pID)

	pb.plyrNames[_pID] = append(pb.plyrNames[_pID], _name)
	pb._setPlyrNames(_pID, pb.plyrNames[_pID])

	return
}

func (pb *MyPlayerBook) determinePID(_addr types.Address) bool {
	if pb.pIDxAddr[_addr] == 0 {
		pb.pID++
		pb.pIDxAddr[_addr] = pb.pID
		pb.plyr[pb.pID] = Player{Addr: _addr}

		// set the new player bool to true
		return true
	}

	return false
}

func (pb *MyPlayerBook) determinePName(k int64, name string) bool {
	for _, v := range pb._plyrNames(k) {
		if v == name {
			// set the new player bool to true
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
func nameFilter(_input string) (_name string, _err string) {
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
