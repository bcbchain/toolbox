package core

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/types"
)

const transferMethodID = "44d8ca60"

// ----- param struct ----
type CallParam struct {
	OrgName      string `json:"orgName"`
	Contract     string `json:"contract"`
	Method       string `json:"method"`
	GasLimit     string `json:"gasLimit"`
	SplitBy      string `json:"splitBy"`
	Params       string `json:"params"`
	ParamsFile   string `json:"paramsFile"`
	Note         string `json:"note"`
	Pay          string `json:"pay"`
	ChainID      string `json:"chainID"`
	KeyStorePath string `json:"keyStorePath"`
}

type DeployContractParam struct {
	ContractName string `json:"contractName"`
	Version      string `json:"version"`
	OrgName      string `json:"orgName"`
	CodeFile     string `json:"codeFile"`
	EffectHeight string `json:"effectHeight"`
	Owner        string `json:"owner"`
	ChainID      string `json:"chainID"`
	KeyStorePath string `json:"keyStorePath"`
	GasLimit     string `json:"gasLimit"`
	Note         string `json:"note"`
}

type RegisterTokenParam struct {
	TokenName        string `json:"tokenName"`
	TokenSymbol      string `json:"tokenSymbol"`
	TotalSupply      string `json:"totalSupply"`
	AddSupplyEnabled string `json:"addSupplyEnabled"`
	BurnEnabled      string `json:"burnEnabled"`
	GasPrice         string `json:"gasPrice"`
	ChainID          string `json:"chainID"`
	KeyStorePath     string `json:"keyStorePath"`
	GasLimit         string `json:"gasLimit"`
	Note             string `json:"note"`
}

type RegisterOrgParam struct {
	OrgName      types.Address `json:"OrgName"`
	ChainID      string        `json:"chainID"`
	KeyStorePath string        `json:"keyStorePath"`
	GasLimit     string        `json:"gasLimit"`
	Note         string        `json:"note"`
}

type TransferParam struct {
	Token        types.Address `json:"token"`
	GasLimit     string        `json:"gasLimit"`
	Note         string        `json:"note"`
	To           types.Address `json:"to"`
	Value        string        `json:"value"`
	ChainID      string        `json:"chainID"`
	KeyStorePath string        `json:"keyStorePath"`
}

// TransferResult - transfer result
type TransferResult struct {
	Code   uint32 `json:"code"`
	Log    string `json:"log"`
	Fee    uint64 `json:"fee"`
	TxHash string `json:"txHash"`
	Height int64  `json:"height"`
}

// BlockHeightResult - block height result
type BlockHeightResult struct {
	LastBlock int64 `json:"lastBlock"`
}

// Message - message struct
type Message struct {
	SmcAddress types.Address `json:"smcAddress"`
	SmcName    string        `json:"smcName"`
	Method     string        `json:"method"`
	To         string        `json:"to"`
	Value      string        `json:"value"`
}

// TxResult - transaction struct
type TxResult struct {
	TxHash      string        `json:"txHash"`
	TxTime      string        `json:"txTime"`
	Code        uint32        `json:"code"`
	Log         string        `json:"log"`
	BlockHash   string        `json:"blockHash"`
	BlockHeight int64         `json:"blockHeight"`
	From        types.Address `json:"from"`
	Nonce       uint64        `json:"nonce"`
	GasLimit    uint64        `json:"gasLimit"`
	Fee         uint64        `json:"fee"`
	Note        string        `json:"note"`
	Messages    []Message     `json:"messages"`
}

// BlockResult - block struct
type BlockResult struct {
	BlockHeight     int64         `json:"blockHeight"`
	BlockHash       string        `json:"blockHash"`
	ParentHash      string        `json:"parentHash"`
	ChainID         string        `json:"chainID"`
	ValidatorHash   string        `json:"validatorHash"`
	ConsensusHash   string        `json:"consensusHash"`
	BlockTime       string        `json:"blockTime"`
	BlockSize       int           `json:"blockSize"`
	ProposerAddress types.Address `json:"proposerAddress"`
	Txs             []TxResult    `json:"txs"`
}

// BalanceItemResult - item of all balance struct
type BalanceItemResult struct {
	TokenAddress types.Address `json:"tokenAddress"`
	TokenName    string        `json:"tokenName"`
	Balance      string        `json:"balance"`
}

// TokenBalance - balance of token with account address
type TokenBalance struct {
	Address types.Address `json:"address"` //代币的合约账户地址
	Balance bn.Number     `json:"balance"` //代币的余额
}

// NonceResult - nonce struct
type NonceResult struct {
	Nonce uint64 `json:"nonce"`
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

// VersionResult - version struct
type VersionResult struct {
	Version string `json:"version"`
}

//-------------------------------------
// 定义交易数据结构

type MethodInfo struct {
	MethodID  uint32
	ParamData []byte
}
