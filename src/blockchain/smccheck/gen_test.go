package smccheck

import (
	"blockchain/smccheck/gen"
	"testing"
)

func TestGen(t *testing.T) {
	path := "/Users/zerppen/GIBlockChain/trunk/code/v2.0/bcsmc-sdk/src/test/contract"

	var contractInfoList []gen.ContractInfo
	contractInfoList = append(contractInfoList, gen.ContractInfo{
		Name:         "myplayerbook",
		Version:      "1.0",
		EffectHeight: 50,
		LoseHeight:   1000,
	})

	contractInfoList = append(contractInfoList, gen.ContractInfo{
		Name:         "myplayerbook",
		Version:      "2.0",
		EffectHeight: 1000,
		LoseHeight:   0,
	})

	Gen(path, contractInfoList)
}
