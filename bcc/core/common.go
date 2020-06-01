package core

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bcbchain/bclib/tx/v2"
	"github.com/bcbchain/bclib/wal"
	"github.com/bcbchain/sdk/sdk/bn"
	"github.com/bcbchain/sdk/sdk/std"
	"github.com/bcbchain/toolbox/bcc/common"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/bcbchain/bclib/tendermint/go-crypto"
)

const (
	nonceErrDesc = "Invalid nonce"
	smcErrDesc   = "The contract has expired"
)

func nodeAddrSlice(chainID string) []string {
	if len(chainID) == 0 {
		chainID = common.GetBCCConfig().DefaultChainID
	}

	if _, ok := common.GetBCCConfig().Urls[chainID]; ok {
		return common.GetBCCConfig().Urls[chainID]
	}

	return []string{}
}

func getAccountPriKey(keyStorePath, name, password string) (priKeyHex string, err error) {

	acct, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		return
	}

	priKey := acct.PrivateKey.(crypto.PrivKeyEd25519)

	return "0x" + hex.EncodeToString(priKey[:]), nil
}

func FuncRecover(errPtr *error) {
	if err := recover(); err != nil {
		msg := ""
		if errInfo, ok := err.(error); ok {
			msg = errInfo.Error()
		}

		if errInfo, ok := err.(string); ok {
			msg = errInfo
		}

		*errPtr = errors.New(msg)
	}
}

func prepare(splitBy, keyStorePath, chainID string) (string, string, string) {
	if splitBy == "" {
		splitBy = "@"
	}

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
	}
	crypto.SetChainId(chainID)
	tx2.Init(chainID)

	return splitBy, keyStorePath, chainID
}

func requireNotEmpty(key, data string) {

	if len(data) == 0 {
		panic(errors.New(fmt.Sprintf("%s cannot be emtpy", key)))
	}
}

func requireUint64(key, valueStr string, base int) (uint64, error) {
	value, err := strconv.ParseUint(valueStr, base, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("%s error=%s", key, err.Error()))
	}

	return value, nil
}

func checkPay(pay string) (value bn.Number, token string, err error) {
	token = ""
	value = bn.N(0)
	if len(pay) > 0 {

		// step 1. check format
		firstIndex := strings.Index(pay, "(")
		lastIndex := strings.Index(pay, ")")
		if firstIndex == -1 || lastIndex < firstIndex {
			err = errors.New("pay option's format error, right format example: 1.02(bcb)")
			return
		}

		// step 2. check token
		token = pay[firstIndex+1 : lastIndex]
		if len(token) <= 0 {
			err = errors.New("pay option's format error, token cannot be empty")
			return
		}

		// step 3. check value
		valueStr := pay[:firstIndex]
		potIndex := strings.Index(valueStr, ".")
		if potIndex != -1 && len(strings.TrimRight(valueStr[potIndex+1:], "0")) > 9 {
			err = errors.New("pay option's format error, value's decimals cannot great than 9 chars")
			return
		}

		zeroCount := 0
		valueStr = strings.Replace(valueStr, ".", "", -1)
		if potIndex == -1 {
			zeroCount = 9
		} else {
			zeroCount = 9 - (len(valueStr) - potIndex)
		}
		for zeroCount > 0 {
			valueStr += "0"
			zeroCount--
		}
		value = bn.NewNumberStringBase(valueStr, 10)
		if value.IsLEI(0) {
			err = errors.New("pay option's format error, value must be number and greater than zero")
		}
	}

	return
}

func checkTransfer(value string) (validValue string, err error) {
	potIndex := strings.Index(value, ".")
	if potIndex != -1 && len(strings.TrimRight(value[potIndex+1:], "0")) > 9 {
		err = errors.New("pay option's format error, value's decimals cannot great than 9 chars")
		return "", err
	}

	zeroCount := 0
	validValue = strings.Replace(value, ".", "", -1)
	if potIndex == -1 {
		zeroCount = 9
	} else {
		zeroCount = 9 - (len(validValue) - potIndex)
	}
	for zeroCount > 0 {
		validValue += "0"
		zeroCount--
	}
	return
}

func checkVersion(version string) (err error) {
	if len(version) < 3 {
		return errors.New("invalid version")
	}

	if len(strings.Trim(version, ".")) != len(version) {
		return errors.New("invalid version")
	}

	verStr := strings.Replace(version, ".", "", -1)
	verN := bn.NewNumberStringBase(verStr, 10)
	if verN.IsLessThanI(0) {
		return errors.New("invalid version")
	}

	return
}

// CheckUTF8 check format
func CheckUTF8(buf []byte) bool {
	nBytes := 0
	for i := 0; i < len(buf); i++ {
		if nBytes == 0 {
			if (buf[i] & 0x80) != 0 {
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1
					nBytes++
				}

				if nBytes < 2 || nBytes > 6 {
					return false
				}

				nBytes--
			}
		} else {
			if buf[i]&0xc0 != 0x80 {
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}

//  查询方法ID
func QueryMethodID(orgName, contractName, method, chainID, keyStorePath string, bSmcErr bool) (uint32, error) {

	contract, err := getContract(orgName, contractName, chainID, bSmcErr, keyStorePath)
	if err != nil {
		return 0, errors.New("getContract error: " + err.Error())
	}
	var item std.Method
	for _, methodItem := range contract.Methods {
		if strings.HasPrefix(methodItem.ProtoType, method) {
			item = methodItem
			break
		}
	}

	if len(item.MethodID) == 0 {
		return 0, errors.New("invalid method")
	}

	methodid, _ := requireUint64("methodID", item.MethodID, 16)

	return uint32(methodid), nil
}

func GetAbiObject(abiFilePath, addr string) (abiStr string, err error) {

	// read the sol contract abiFile
	if abiFilePath != "" {
		codeData, err := ioutil.ReadFile(abiFilePath)
		if err != nil {
			return "", err
		}
		abiStr = string(codeData)

	} else {
		if addr != "" {
			addrs := nodeAddrSlice(common.GetBCCConfig().DefaultChainID)
			res := new(std.BvmContract)
			err = common.DoHttpQueryAndParse(addrs, "/bvm/contract/"+addr, &res)
			if err != nil {
				return
			}

			abiStr = res.BvmAbi
		}
	}

	return
}
