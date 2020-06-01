package main

import (
	"time"

	"github.com/bcbchain/bclib/tendermint/go-crypto"
)

type BlockSize struct {
	MaxBytes int   `json:"max_bytes"` // NOTE: must not be 0 nor greater than 100MB
	MaxTxs   int   `json:"max_txs"`
	MaxGas   int64 `json:"max_gas"`
}
type TxSize struct {
	MaxBytes int   `json:"max_bytes"`
	MaxGas   int64 `json:"max_gas"`
}

type BlockGossip struct {
	BlockPartSizeBytes int `json:"block_part_size_bytes"` // NOTE: must not be 0
}

type EvidenceParams struct {
	MaxAge int64 `json:"max_age"` // only accept new evidence more recent than this
}

type ConsensusParams struct {
	BlockSize      `json:"block_size_params"`
	TxSize         `json:"tx_size_params"`
	BlockGossip    `json:"block_gossip_params"`
	EvidenceParams `json:"evidence_params"`
}

type GenesisValidator struct {
	Name       string         `json:"name"`
	RewardAddr crypto.Address `json:"reward_addr"`
	Power      int64          `json:"power"`
}

type GenesisAppState struct {
	Organization   string          `json:"organization"`
	Token          tokenInfo       `json:"token"`
	RewardStrategy []GenesisReward `json:"rewardStrategy"`
	Contracts      []contractInfo  `json:"contracts"`
}

type GenesisDoc struct {
	ChainID         string             `json:"chain_id"`
	ChainVersion    string             `json:"chain_version"`
	GenesisTime     time.Time          `json:"genesis_time"`
	AppHash         string             `json:"app_hash"`
	ConsensusParams *ConsensusParams   `json:"consensus_params,omitempty"`
	AppState        GenesisAppState    `json:"app_state,omitempty"`
	Validators      []GenesisValidator `json:"validators,omitempty"`
}

type tokenInfo struct {
	Address          crypto.Address `json:"address"`
	Owner            crypto.Address `json:"owner"`
	Name             string         `json:"name"`
	Symbol           string         `json:"symbol"`
	TotalSupply      uint64         `json:"totalSupply"`
	AddSupplyEnabled bool           `json:"addSupplyEnabled"`
	BurnEnabled      bool           `json:"burnEnabled"`
	GasPrice         uint64         `json:"gasPrice"`
}

type codeSign struct {
	PubKey    string `json:"pubkey"`
	Signature string `json:"signature"`
}

type contractInfo struct {
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	Code       string   `json:"code"`
	CodeHash   string   `json:"codeHash"`
	CodeDevSig codeSign `json:"codeDevSig"`
	CodeOrgSig codeSign `json:"codeOrgSig"`
	code       []byte
}

type GenesisToken struct {
	TokenName        string `json:"tokenName"`
	TokenSymbol      string `json:"tokenSymbol"`
	TotalSupply      uint64 `json:"totalSupply"`
	AddSupplyEnabled bool   `json:"addSupplyEnabled"`
	BurnEnabled      bool   `json:"burnEnabled"`
	GasPrice         uint64 `json:"gasPrice"` //代币交易的gas价格
}

type GenesisReward struct {
	Name    string `json:"name"`
	Reward  string `json:"rewardPercent"`
	Address string `json:"address"`
}

type GenesisCharter struct {
	Token   GenesisToken    `json:"token"`
	Rewards []GenesisReward `json:"rewardStrategy"`
}

type NodeDef struct {
	Name       string         `json:"name"`
	RewardAddr crypto.Address `json:"reward_addr"`
	Power      int            `json:"power"`
	ListenPort int            `json:"listen_port"`
	Announce   string         `json:"announce"`
	Public     string         `json:"public"`
	IPIn       string         `json:"ip_in"`
	IPOut      string         `json:"ip_out"`
	IPPriv     string         `json:"ip_priv"`
	Apps       []string       `json:"apps"`
}

type Account struct {
	Addr crypto.Address `json:"addr"`
	name string
}
