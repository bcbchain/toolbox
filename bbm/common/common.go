package common

import (
	"github.com/bcbchain/sdk/sdk/bn"
	"github.com/bcbchain/sdk/sdk/jsoniter"
	"github.com/bcbchain/sdk/sdk/std"
	"github.com/bcbchain/bclib/types"
	"errors"
	"strings"
	"time"

	types2 "github.com/bcbchain/bclib/tendermint/abci/types"
)

//GetGenesisToken get genesis token name
func GetGenesisToken(client Client) std.Token {
	b := client.Query(std.KeyOfGenesisToken())
	var token std.Token
	if err := jsoniter.Unmarshal(b, &token); err != nil {
		panic(err.Error())
	}
	return token

}

func keyOfGenesisChainID() string {
	return "/genesis/chainid"
}

//GetChainID get chainID
func GetChainID(client Client) string {
	b := client.Query(keyOfGenesisChainID())
	if b == nil {
		panic("chainID must not be nil")
	}

	var chainID string
	err := jsoniter.Unmarshal(b, &chainID) //nolint
	if err != nil {
		//panic("Unmarshal error:" + err.Error())
		// 1.0 和 2.0 的chainId 方式不一样，这里需要强制转换成 string
		return string(b)
	}
	return chainID
}

func keyOfAccountNonce(exAddress types.Address) string {
	return "/account/ex/" + exAddress + "/account"
}

//AccountNonce get account's nonce
func AccountNonce(client Client, exAddr types.Address) uint64 {
	//time.Sleep(time.Second * 1)
	b := client.Query(keyOfAccountNonce(exAddr))
	type AccountInfo struct {
		Nonce uint64
	}
	var nonce AccountInfo
	if e := jsoniter.Unmarshal(b, &nonce); e != nil {
		time.Sleep(time.Second * 1)
		bn := client.Query(keyOfAccountNonce(exAddr))
		if en := jsoniter.Unmarshal(bn, &nonce); en != nil {
			time.Sleep(time.Second * 6)
			bno := client.Query(keyOfAccountNonce(exAddr))
			if eno := jsoniter.Unmarshal(bno, &nonce); eno != nil {
				return 0
			}
		}
	}
	return nonce.Nonce
}

//BalanceOf get external account balance of token
func BalanceOf(client Client, exAddr types.Address, tokenAddr types.Address) bn.Number {
	b := client.Query(std.KeyOfAccountToken(exAddr, tokenAddr))
	if b == nil {
		return bn.N(0)
	}
	var acc std.AccountInfo
	if e := jsoniter.Unmarshal(b, &acc); e != nil {
		return bn.N(0)
	}
	return acc.Balance
}

//GetContractByName get contract data of given name
func GetContractByName(client Client, orgID, name string) (types.Address, error) {
	b := client.Query(std.KeyOfContractsWithName(orgID, name))
	if b == nil {
		return "", errors.New("query contract failed")
	}

	var contractList std.ContractVersionList
	err := jsoniter.Unmarshal(b, &contractList)
	if err != nil {
		return "", err
	}

	//blockHeight := GetCurrentBlockHeight(client) + 1
	//从最后一个开始对比，如果生效高度小于当前区块高度，则此合约有效
	for i := len(contractList.EffectHeights) - 1; i >= 0; {
		blockHeight := GetCurrentBlockHeight(client) + 1
		if blockHeight > 0 && blockHeight >= contractList.EffectHeights[i] {
			return contractList.ContractAddrList[i], nil
		} else {
			time.Sleep(2 * time.Second)
		}
	}

	return "", errors.New("not found")
}

func GetCurrentBlockHeight(client Client) int64 {
	b := client.Query(std.KeyOfAppState())

	var appState types2.AppState
	err := jsoniter.Unmarshal(b, &appState)
	if err != nil {
		return -1
	}
	return appState.BlockHeight
}

//GetTokenByName get contract data of given name
func GetTokenByName(client Client, name string) (types.Address, error) {
	if strings.ToLower(name) == "bcb" {
		t := GetGenesisToken(client)
		return t.Address, nil
	}
	b := client.Query(std.KeyOfTokenWithName(name))
	if b == nil {
		return "", errors.New("Query token failed,key:" + std.KeyOfTokenWithName(name))
	}

	var tokenAddr types.Address
	err := jsoniter.Unmarshal(b, &tokenAddr)
	return tokenAddr, err
}

func GetV1ContractByName(client Client, name string) (string, error) {
	key := std.KeyOfAllContracts()
	b := client.Query(key)
	if b == nil {
		return "", errors.New("cat not get all contract")
	}
	var cs []string
	err := jsoniter.Unmarshal(b, &cs)
	if err != nil {
		return "", err
	}

	for _, v := range cs {
		bb := client.Query(std.KeyOfContract(v))
		con := std.Contract{}
		err := jsoniter.Unmarshal(bb, &con)
		if err != nil {
			return "", err
		}
		if con.Name == name && con.ChainVersion == 0 {
			return con.Address, nil
		}
	}
	return "", errors.New("can not get contract : " + name)
}

func GetGenesisOrgID() string {
	return "orgJgaGConUyK81zibntUBjQ33PKctpk1K1G"
}
