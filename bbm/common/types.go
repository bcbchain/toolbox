package common

import (
	"github.com/bcbchain/bcbchain/abciapp_v1.0/keys"
	"github.com/bcbchain/bclib/types"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
)

// 定义交易数据结构
type Transaction struct {
	Nonce    uint64       // 交易发起者发起交易的计数值，从1开始，必须单调增长，增长步长为1。
	GasLimit uint64       // 交易发起者愿意为执行此次交易支付的GAS数量的最大值。
	Note     string       // UTF-8编码的备注信息，要求小于256个字符。
	To       keys.Address // 合约地址
	Data     []byte       // 调用智能合约所需要的参数，RLP编码格式。
}

type MethodInfo struct {
	MethodID  uint32
	ParamData []byte
}

//Client client interface
type Client interface {
	Query(key string) []byte
	Status() core_types.ResultStatus
	SendTx(tx []byte) types.Response
	InitChain()
	EndAndCommit()
	SendTxAsync(tx []byte) core_types.ResultBroadcastTxCommit
	BlockWithHeight(height int64) (*core_types.ResultBlock, error)
	Validators(height int64) *core_types.ResultValidators
}
