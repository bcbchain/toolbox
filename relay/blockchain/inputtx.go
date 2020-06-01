package blockchain

import (
	"github.com/bcbchain/bclib/tx/v2"
	"github.com/bcbchain/bclib/types"
	"github.com/bcbchain/bclib/wal"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
	"github.com/bcbchain/toolbox/relay/common"
	"strconv"
)

var genesisOrgID = "orgJgaGConUyK81zibntUBjQ33PKctpk1K1G"

func IbcInput(toChainID string, orgID string, ibcInputParams IbcInputParam, urls []string) (result *CommitTxResult, err error) {
	defer common.FuncRecover(&err)

	walName, walPassword, keyStorePath := common.GetWalletInfo()

	contractName := "ibc"
	params := make([]interface{}, 0)
	params = append(params, ibcInputParams.PktsProofs)

	//查询方法ID(可以直接定义成一个常量)
	methodID, err := QueryMethodID(genesisOrgID, contractName, "Input", urls)
	if err != nil {
		return
	}

	result, err = packAndCommitTx(toChainID, walName, walPassword, orgID, contractName, "10000000", "", keyStorePath, toChainID, methodID, params, urls)

	return
}

func packAndCommitTx(toChainID, name, password, orgID, contractName, gasLimit, note, keyStorePath, chainID string,
	methodID uint32, values []interface{}, urls []string) (result *CommitTxResult, err error) {
	defer common.FuncRecover(&err)

	acc, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		return
	}

	addr := acc.Address(chainID)
	nonce, err := getAccountNonce(urls, addr)
	if err != nil {
		nonce = 1
	}

	contract, err := getContract(urls, genesisOrgID, contractName)
	if err != nil {
		return
	}

	uGasLimit, err := strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		return
	}

	privStr, err := getAccountPriKey(keyStorePath, name, password)
	if err != nil {
		return
	}
	txStr := GenerateTx(toChainID, orgID+"."+contract.Address, methodID, values, nonce, int64(uGasLimit), note, privStr)

	common.GetLogger().Debug(txStr)
	result, err = CommitTx(urls, txStr)

	return
}

//GenerateTx generate tx with one contract method request
func GenerateTx(toChainID string, contract types.Address, method uint32, params []interface{}, nonce uint64, gaslimit int64, note string, privKey string) string {
	items := tx2.WrapInvokeParams(params...)
	message := types.Message{
		Contract: contract,
		MethodID: method,
		Items:    items,
	}
	payload := tx2.WrapPayload(nonce, gaslimit, note, message)

	return tx2.WrapTxEx(toChainID, payload, privKey)
}

// CommitTx commit transaction information
func CommitTx(urls []string, tx string) (commit *CommitTxResult, err error) {

	defer common.FuncRecover(&err)

	var result *core_types.ResultTx
	result, err = common.DoHttpCommitTxAndParseAsync(urls, tx)
	if err != nil {
		return
	}

	commit = new(CommitTxResult)
	if result.CheckResult.Code != types.CodeOK {
		commit.Code = result.CheckResult.Code
		commit.Log = result.CheckResult.Log
	} else {
		commit.Code = result.DeliverResult.Code
		commit.Log = result.DeliverResult.Log
	}

	commit.Fee = result.DeliverResult.Fee
	commit.TxHash = "0x" + result.Hash
	commit.Height = result.Height
	commit.Data = result.DeliverResult.Data

	return
}
