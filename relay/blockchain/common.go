package blockchain

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/bcbchain/bclib/wal"
	"github.com/bcbchain/sdk/sdk/std"
	"github.com/bcbchain/sdk/sdk/types"
	"github.com/bcbchain/toolbox/bcc/common"
	"github.com/bcbchain/toolbox/bcc/core"
	"strconv"
	"strings"
)

func getAccountPriKey(keyStorePath, name, password string) (priKeyHex string, err error) {

	acct, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		return
	}

	priKey := acct.PrivateKey.(crypto.PrivKeyEd25519)

	return "0x" + hex.EncodeToString(priKey[:]), nil
}

func getAccountNonce(urls []string, exAddress types.Address) (nonce uint64, err error) {

	defer core.FuncRecover(&err)

	type account struct {
		Nonce uint64 `json:"nonce"`
	}
	an := new(account)

	err = common.DoHttpQueryAndParse(urls, "/account/ex/"+exAddress+"/account", an)
	nonce = an.Nonce + 1

	return
}

func getContract(urls []string, orgID, contractName string) (contract *std.Contract, err error) {

	defer core.FuncRecover(&err)

	var contractVersionList std.ContractVersionList
	err = common.DoHttpQueryAndParse(urls, "/contract/"+orgID+"/"+contractName, &contractVersionList)
	if err != nil {
		return
	}

	contractAddrListLen := len(contractVersionList.ContractAddrList)
	err = common.DoHttpQueryAndParse(urls, "/contract/"+contractVersionList.ContractAddrList[contractAddrListLen-1], &contract)

	return
}

func requireUint64(key, valueStr string, base int) (uint64, error) {
	value, err := strconv.ParseUint(valueStr, base, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("%s error=%s", key, err.Error()))
	}

	return value, nil
}

func QueryMethodID(orgID, contractName, method string, urls []string) (methodID uint32, err error) {

	contract, err := getContract(urls, orgID, contractName)
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
