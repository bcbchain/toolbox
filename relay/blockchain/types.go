package blockchain

import (
	"github.com/bcbchain/sdk/sdk/ibc"
)

type IbcInputParam struct {
	PktsProofs []ibc.PktsProof
}

// CommitTxResult - commit tx result
type CommitTxResult struct {
	Code         uint32 `json:"code"`
	Log          string `json:"log"`
	Fee          uint64 `json:"fee"`
	TxHash       string `json:"txHash"`
	Height       int64  `json:"height"`
	Data         string `json:"data,omitempty"`
	SmcAddress   string `json:"smcAddress,omitempty"`
	TokenAddress string `json:"tokenAddress,omitempty"`
	OrgID        string `json:"orgID,omitempty"`
}
