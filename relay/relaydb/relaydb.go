package relaydb

import (
	"github.com/bcbchain/bclib/bcdb"
	"github.com/bcbchain/sdk/sdk/ibc"
	"github.com/bcbchain/sdk/sdk/jsoniter"
	"github.com/bcbchain/toolbox/relay/blockchain"
	"github.com/bcbchain/toolbox/relay/common"
)

var (
	sdb *bcdb.GILevelDB
)

func InitDB() error {
	config := common.GetConfig()
	var err error
	sdb, err = bcdb.OpenDB(config.DBName, "", "")
	if err != nil {
		panic(err)
	}

	return updateOpenURLs()
}

func updateOpenURLs() error {

	chainIDs := GetChainIDs()
	if len(chainIDs) == 0 {
		config := common.GetConfig()
		chainIDs := config.SideChainIDs

		for _, chainID := range chainIDs {
			openUrls, err := blockchain.QueryOpenURLs(chainID)
			if err != nil {
				return err
			}

			chainInfo := &ChainInfo{
				ChainID:        chainID,
				OpenUrls:       openUrls,
				LastScanHeight: 1,
			}
			SetChainInfo(chainInfo)
		}

		SetChainIDs(chainIDs)
	}

	return nil
}

func GetChainIDs() []string {
	chainIDsByte, err := sdb.Get(keyOfChainIDs())
	if err != nil {
		panic(err)
	}

	if len(chainIDsByte) == 0 {
		return nil
	}

	var chainIDs []string
	err = jsoniter.Unmarshal(chainIDsByte, &chainIDs)
	if err != nil {
		panic(err)
	}

	return chainIDs
}

func SetChainIDs(chainIDs []string) {
	chainIDsByte, err := jsoniter.Marshal(chainIDs)
	if err != nil {
		panic(err)
	}

	err = sdb.Set(keyOfChainIDs(), chainIDsByte)
	if err != nil {
		panic(err)
	}
}

func GetChainInfo(chainID string) *ChainInfo {
	chainInfoByte, err := sdb.Get(keyOfChainInfo(chainID))
	if err != nil {
		panic(err)
	}

	if len(chainInfoByte) == 0 {
		return nil
	}

	var chainInfo ChainInfo
	err = jsoniter.Unmarshal(chainInfoByte, &chainInfo)
	if err != nil {
		panic(err)
	}

	return &chainInfo
}

func SetChainInfo(chainInfo *ChainInfo) {
	chainInfoByte, err := jsoniter.Marshal(chainInfo)
	if err != nil {
		panic(err)
	}

	err = sdb.Set(keyOfChainInfo(chainInfo.ChainID), chainInfoByte)
	if err != nil {
		panic(err)
	}
}

func GetLastHeight(queueID string) int64 {
	seqBytes, err := sdb.Get(keyOfLastHeight(queueID))
	if err != nil {
		panic(err)
	}

	if len(seqBytes) == 0 {
		return 0
	}

	var height int64
	err = jsoniter.Unmarshal(seqBytes, &height)
	if err != nil {
		panic(err)
	}

	return height
}

func SetLastHeight(queueID string, height int64) {
	seqBytes, err := jsoniter.Marshal(height)
	if err != nil {
		panic(err)
	}

	err = sdb.Set(keyOfLastHeight(queueID), seqBytes)
	if err != nil {
		panic(err)
	}
}

func GetPktsProof(queueID string, height int64) *ibc.PktsProof {
	pktsProofBytes, err := sdb.Get(keyOfPktsProof(queueID, height))
	if err != nil {
		panic(err)
	}

	if len(pktsProofBytes) == 0 {
		return nil
	}

	pktsProof := new(ibc.PktsProof)
	err = jsoniter.Unmarshal(pktsProofBytes, pktsProof)
	if err != nil {
		panic(err)
	}

	return pktsProof
}

func SetPktsProof(queueID string, height int64, pktsProof ibc.PktsProof) {
	pktsProofBytes, err := jsoniter.Marshal(pktsProof)
	if err != nil {
		panic(err)
	}

	err = sdb.Set(keyOfPktsProof(queueID, height), pktsProofBytes)
	if err != nil {
		panic(err)
	}
}

func GetChainQueueIDs(chainID string) []string {
	queueIDsBytes, err := sdb.Get(keyOfChainQueueIDs(chainID))
	if err != nil {
		panic(err)
	}

	if len(queueIDsBytes) == 0 {
		return nil
	}

	var queueIDs []string
	err = jsoniter.Unmarshal(queueIDsBytes, &queueIDs)
	if err != nil {
		panic(err)
	}

	return queueIDs
}

func AddChainQueueID(chainID string, queueID string) (bNew bool) {
	queueIDs := GetChainQueueIDs(chainID)

	bNew = true
	for _, item := range queueIDs {
		if item == queueID {
			bNew = false
			break
		}
	}

	if bNew {

		queueIDs = append(queueIDs, queueID)

		queueIDsBytes, err := jsoniter.Marshal(queueIDs)
		if err != nil {
			panic(err)
		}

		err = sdb.Set(keyOfChainQueueIDs(chainID), queueIDsBytes)
		if err != nil {
			panic(err)
		}
	}

	return
}
