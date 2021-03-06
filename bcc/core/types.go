package core

import (
	"github.com/bcbchain/sdk/sdk/bn"
	"github.com/bcbchain/bclib/types"
)

const transferMethodIDV1 = "af0228bc"
const transferMethodIDV2 = "44d8ca60"

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
	OrgName      string `json:"orgName"`
	ChainID      string `json:"chainID"`
	KeyStorePath string `json:"keyStorePath"`
	GasLimit     string `json:"gasLimit"`
	Note         string `json:"note"`
}

type SetOrgSignersParam struct {
	OrgName      string `json:"orgName"`
	PubKeys      string `json:"pubKeys"`
	ChainID      string `json:"chainID"`
	KeyStorePath string `json:"keyStorePath"`
	GasLimit     string `json:"gasLimit"`
	Note         string `json:"note"`
}

type SetOrgDeployerParam struct {
	OrgName      string        `json:"orgName"`
	Deployer     types.Address `json:"deployer"`
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

type BVMDeployParam struct {
	TokenAddr    string   `json:"tokenAddr"`
	TokenName    string   `json:"tokenName"`
	SourceFile   string   `json:"sourceFile"`
	Library      string   `json:"library"`
	ContractName string   `json:"contractName"`
	BinFile      string   `json:"codeFile"`
	AbiFile      string   `json:"abiFile"`
	ParamsArray  []string `json:"params"`
	GasLimit     string   `json:"gasLimit"`
	Note         string   `json:"note"`
	ChainID      string   `json:"chainID"`
	KeyStorePath string   `json:"keyStorePath"`
}

type BVMCallParam struct {
	AbiFile      string   `json:"abiFile"`
	ContractAddr string   `json:"contractAddr"`
	Value        string   `json:"value"`
	Method       string   `json:"method"`
	ParamsArray  []string `json:"params"`
	GasLimit     string   `json:"gasLimit"`
	Note         string   `json:"note"`
	ChainID      string   `json:"chainID"`
	KeyStorePath string   `json:"keyStorePath"`
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
	TxHash      string          `json:"txHash"`
	TxTime      string          `json:"txTime"`
	Code        uint32          `json:"code"`
	Log         string          `json:"log"`
	BlockHash   string          `json:"blockHash"`
	BlockHeight int64           `json:"blockHeight"`
	From        types.Address   `json:"from"`
	Nonce       uint64          `json:"nonce"`
	GasLimit    uint64          `json:"gasLimit"`
	Fee         uint64          `json:"fee"`
	Note        string          `json:"note"`
	Messages    []Message       `json:"messages"`
	Tags        map[string]*Tag `json:"tags"`
}

// BlockResult - block struct
type BlockResult struct {
	// detail result about height
	BlockHeight     int64         `json:"blockHeight,omitempty"`
	BlockHash       string        `json:"blockHash,omitempty"`
	ParentHash      string        `json:"parentHash,omitempty"`
	ChainID         string        `json:"chainID,omitempty"`
	ValidatorHash   string        `json:"validatorHash,omitempty"`
	ConsensusHash   string        `json:"consensusHash,omitempty"`
	BlockTime       string        `json:"blockTime,omitempty"`
	BlockSize       int           `json:"blockSize,omitempty"`
	ProposerAddress types.Address `json:"proposerAddress,omitempty"`
	Txs             []TxResult    `json:"txs,omitempty"`

	// simple result contain several blocks
	Result []simpleBlockResult `json:"result,omitempty"`
}

// simpleBlockResult simple block information contain height,hash and time
type simpleBlockResult struct {
	BlockHeight int64  `json:"blockHeight"`
	BlockHash   string `json:"blockHash"`
	BlockTime   string `json:"blockTime"`
}

// BalanceItemResult - item of all balance struct
type BalanceItemResult struct {
	TokenAddress types.Address `json:"tokenAddress"`
	TokenName    string        `json:"tokenName"`
	Balance      string        `json:"balance"`
	BVMBalance   string        `json:"bvmBalance,omitempty"`
}

// TokenBalance - balance of token with account address
type TokenBalance struct {
	Address       types.Address `json:"address"`         //代币的合约账户地址
	Balance       bn.Number     `json:"balance"`         //代币的余额
	BVMBalanceMod *bn.Number    `json:"bbMod,omitempty"` // 十八位精度的后九位
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

// BVMCallResult - commit tx of bvm call result
type BVMCallResult struct {
	Code   uint32 `json:"code"`
	Log    string `json:"log"`
	Fee    uint64 `json:"fee,omitempty"`
	TxHash string `json:"txHash"`
	Height int64  `json:"height,omitempty"`
	Data   []byte `json:"data,omitempty"`
}

// VersionResult - version struct
type VersionResult struct {
	Version string `json:"version"`
}

//-------------------------------------
// 定义交易数据结构

type MethodInfo struct {
	MethodID  uint32 `json:"MethodID"`
	ParamData []byte `json:"ParamData"`
}

// Tag - Tag struct
type Tag struct {
	Name         string                 `json:"Name"`
	ContractAddr string                 `json:"ContractAddr"`
	ReceiptBytes string                 `json:"ReceiptBytes"`
	ReceiptHash  string                 `json:"ReceiptHash"`
	Receipt      map[string]interface{} `json:"Receipt"`
}
