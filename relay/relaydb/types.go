package relaydb

import (
	"fmt"
)

type ChainInfo struct {
	ChainID        string   `json:"chainID"`
	OpenUrls       []string `json:"openUrls"`
	LastScanHeight int64    `json:"lastScanHeight"`
}

type SetOpenURL struct {
	SideChainID string   `json:"sideChainID"`
	OpenURLs    []string `json:"openURLs"`
}

func keyOfChainInfo(chainID string) []byte {
	return []byte("/relay/" + chainID + "/chaininfo")
}

func keyOfPktsProof(queueID string, height int64) []byte {
	return []byte(fmt.Sprintf("/relay/%s/%d/pktsproof", queueID, height))
}

func keyOfChainIDs() []byte {
	return []byte("/relay/chainIDs")
}

func keyOfLastHeight(queueID string) []byte {
	return []byte("/relay/" + queueID + "/lastheight")
}

func keyOfChainQueueIDs(chainID string) []byte {
	return []byte("/relay/" + chainID + "/queueIDs")
}
