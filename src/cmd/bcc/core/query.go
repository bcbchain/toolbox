package core

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	"blockchain/smcsdk/sdk/std"
	"blockchain/tx2"
	"blockchain/types"
	"cmd/bcc/common"
	"common/wal"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tendermint/go-crypto"
	"io/ioutil"
	"strings"

	"github.com/tendermint/tendermint/rpc/core/types"
)

//BlockHeight blockHeight query
func BlockHeight(chainID string) (blkHeight *BlockHeightResult, err error) {

	defer FuncRecover(&err)

	addrS := nodeAddrSlice(chainID)

	result := new(core_types.ResultABCIInfo)
	params := map[string]interface{}{}
	err = DoHttpRequestAndParse(addrS, "abci_info", params, result)
	if err != nil {
		return
	}
	blkHeight = new(BlockHeightResult)
	blkHeight.LastBlock = result.Response.LastBlockHeight

	return
}

//Block block information query
func Block(height int64, chainID string) (blk *BlockResult, err error) {

	defer FuncRecover(&err)

	_, _, chainID = prepare("", "", chainID)

	addrS := nodeAddrSlice(chainID)

	result := new(core_types.ResultBlock)
	params := map[string]interface{}{"height": height}
	err = DoHttpRequestAndParse(addrS, "block", params, result)
	if err != nil {
		return
	}

	blk = new(BlockResult)
	blk.BlockHeight = result.BlockMeta.Header.Height
	blk.BlockHash = hex.EncodeToString(result.BlockMeta.BlockID.Hash)
	blk.ParentHash = hex.EncodeToString(result.BlockMeta.Header.LastBlockID.Hash)
	blk.ChainID = result.BlockMeta.Header.ChainID
	blk.ValidatorHash = hex.EncodeToString(result.BlockMeta.Header.ValidatorsHash)
	blk.ConsensusHash = hex.EncodeToString(result.BlockMeta.Header.ConsensusHash)
	blk.BlockTime = result.BlockMeta.Header.Time.String()
	blk.BlockSize = result.BlockSize
	blk.ProposerAddress = result.BlockMeta.Header.ProposerAddress

	blk.Txs = make([]TxResult, 0)
	var blkResults *core_types.ResultBlockResults
	if blkResults, err = blockResults(chainID, height); err != nil {
		return
	}

	for _, blkResult := range blkResults.Results.DeliverTx {
		var tx *TxResult
		if tx, err = Transaction(chainID, hex.EncodeToString(blkResult.TxHash), result); err != nil {
			return
		}
		blk.Txs = append(blk.Txs, *tx)
	}

	return
}

//Transaction transaction information query
func Transaction(chainID, txHash string, resultBlock *core_types.ResultBlock) (tx *TxResult, err error) {

	defer FuncRecover(&err)

	requireNotEmpty("txHash", txHash)

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
		tx2.Init(chainID)
		crypto.SetChainId(chainID)
	}

	addrS := nodeAddrSlice(chainID)

	result := new(core_types.ResultTx)
	params := map[string]interface{}{"hash": txHash}
	err = DoHttpRequestAndParse(addrS, "tx", params, result)
	if err != nil {
		return
	}

	if resultBlock == nil {
		resultBlock = new(core_types.ResultBlock)
		params = map[string]interface{}{"height": result.Height}
		err = DoHttpRequestAndParse(addrS, "block", params, resultBlock)
		if err != nil {
			return
		}
	}

	// parse transaction //todo 兼容v1
	transaction, fromAddr, err := tx2.TxParse(string(resultBlock.Block.Txs[result.Index]))
	if err != nil {
		return
	}

	tx = new(TxResult)
	tx.TxHash = txHash
	tx.TxTime = resultBlock.BlockMeta.Header.Time.String()
	tx.Code = result.DeliverResult.Code
	tx.Log = result.DeliverResult.Log
	tx.BlockHash = hex.EncodeToString(resultBlock.BlockMeta.BlockID.Hash)
	tx.BlockHeight = resultBlock.BlockMeta.Header.Height
	tx.From = fromAddr.Address()
	tx.Nonce = transaction.Nonce
	tx.GasLimit = uint64(transaction.GasLimit)
	tx.Fee = result.DeliverResult.Fee
	tx.Note = transaction.Note
	tx.Messages = make([]Message, 0)

	var msg Message
	for i := 0; i < len(transaction.Messages); i++ {
		msg, err = message(chainID, transaction.Messages[i])
		if err != nil {
			return
		}
		tx.Messages = append(tx.Messages, msg)
	}

	return
}

// Balance balance information query
func Balance(address types.Address, name, password, tokenName string, all bool, chainID, keyStorePath string) (result *[]BalanceItemResult, err error) {

	defer FuncRecover(&err)

	_, keyStorePath, chainID = prepare("", keyStorePath, chainID)

	// if account address is empty, then load account with name and password
	if address == "" {
		acct, err := wal.LoadAccount(keyStorePath, name, password)
		if err != nil {
			return nil, err
		}

		address = acct.Address(chainID)
	}

	if all == false {
		items := make([]BalanceItemResult, 0)
		item, err := balanceOfToken(address, tokenName, chainID)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
		result = &items
	} else {
		temp, err := allBalance(chainID, address)
		if err != nil {
			return nil, err
		}
		result = temp
	}

	return
}

func balanceOfToken(accAddress types.Address, tokenName, chainID string) (result *BalanceItemResult, err error) {

	addrS := nodeAddrSlice(chainID)

	var value []byte
	var tokenAddress types.Address
	if tokenName == "" {
		tokenResult := new(std.Token)
		err = DoHttpQueryAndParse(addrS, std.KeyOfGenesisToken(), tokenResult)
		if err != nil {
			return nil, errors.New("get genesis token error: " + err.Error())
		}
		tokenAddress = tokenResult.Address
		tokenName = tokenResult.Name
	} else {
		if value, err = DoHttpQuery(addrS, std.KeyOfTokenWithName(tokenName)); err != nil {
			return
		}
		if len(value) == 0 {
			return nil, errors.New("invalid tokenName")
		}

		if err = json.Unmarshal(value, &tokenAddress); err != nil {
			return
		}
	}

	if value, err = DoHttpQuery(addrS, std.KeyOfAccountToken(accAddress, tokenAddress)); err != nil {
		return
	}

	result = new(BalanceItemResult)
	if len(value) == 0 {
		result.Balance = "0"
	} else {
		var tokenBalance TokenBalance
		if err = json.Unmarshal(value, &tokenBalance); err != nil {
			return
		}
		result.Balance = tokenBalance.Balance.String()
	}
	result.TokenAddress = tokenAddress
	result.TokenName = tokenName

	return
}

func allBalance(chainID string, address types.Address) (items *[]BalanceItemResult, err error) {

	addrS := nodeAddrSlice(chainID)

	tokens := make([]string, 0)
	if err = DoHttpQueryAndParse(addrS, std.KeyOfAccount(address), &tokens); err != nil {
		return
	}

	balanceItems := make([]BalanceItemResult, 0)
	for _, token := range tokens {
		splitToken := strings.Split(token, "/")
		if splitToken[4] != "token" || len(splitToken) != 6 {
			continue
		}
		tokenBalance := new(TokenBalance)
		if err = DoHttpQueryAndParse(addrS, token, tokenBalance); err != nil {
			return
		}

		var name string
		if name, err = tokenName(chainID, tokenBalance.Address); err != nil {
			return
		}

		balanceItems = append(balanceItems,
			BalanceItemResult{
				TokenAddress: tokenBalance.Address,
				TokenName:    name,
				Balance:      tokenBalance.Balance.String()})
	}
	items = &balanceItems

	return
}

// Nonce nonce information query
func Nonce(address types.Address, name, password, chainID, keyStorePath string) (result *NonceResult, err error) {

	defer FuncRecover(&err)

	addrS := nodeAddrSlice(chainID)

	_, keyStorePath, chainID = prepare("", keyStorePath, chainID)

	if address == "" {
		acct, err := wal.LoadAccount(keyStorePath, name, password)
		if err != nil {
			return nil, err
		}

		address = acct.Address(chainID)
	}

	type account struct {
		Nonce uint64 `json:"nonce"`
	}

	a := new(account)
	value, err := DoHttpQuery(addrS, std.KeyOfAccountNonce(address))
	if err != nil {
		return
	}

	result = new(NonceResult)
	if len(value) == 0 {
		result.Nonce = 1
	} else {
		err = json.Unmarshal(value, a)
		if err != nil {
			return
		}

		result.Nonce = a.Nonce + 1
	}

	return
}

// CommitTx commit transaction information
func CommitTx(chainID, tx string) (commit *CommitTxResult, err error) {

	defer FuncRecover(&err)

	requireNotEmpty("tx", tx)

	addrS := nodeAddrSlice(chainID)

	var result *core_types.ResultBroadcastTxCommit
	result, err = DoHttpCommitTxAndParse(addrS, tx)
	if err != nil {
		return
	}

	commit = new(CommitTxResult)
	if result.CheckTx.Code != types.CodeOK {
		commit.Code = result.CheckTx.Code
		commit.Log = result.CheckTx.Log
	} else {
		commit.Code = result.DeliverTx.Code
		commit.Log = result.DeliverTx.Log
	}

	commit.Fee = result.DeliverTx.Fee
	commit.TxHash = hex.EncodeToString(result.Hash)
	commit.Height = result.Height
	commit.Data = result.DeliverTx.Data

	return
}

//Version version information for the current block
func Version() (result *VersionResult, err error) {

	defer FuncRecover(&err)

	result = new(VersionResult)
	var version []byte
	version, err = ioutil.ReadFile("./version")
	if err != nil {
		common.GetLogger().Error("Read version file error", "error", err)
		err = nil
		result.Version = "0.0.0.0"
		return
	}
	result.Version = string(version)

	return
}

func blockResults(chainID string, height int64) (blkResults *core_types.ResultBlockResults, err error) {

	addrS := nodeAddrSlice(chainID)

	blkResults = new(core_types.ResultBlockResults)
	params := map[string]interface{}{"height": height}
	err = DoHttpRequestAndParse(addrS, "block_results", params, blkResults)
	if err != nil {
		return
	}

	return
}

func message(chainID string, message types.Message) (msg Message, err error) {

	//var methodInfo MethodInfo
	//for i := 0; i < len(message.Items); i++ {
	//	if err = rlp.DecodeBytes(message.Items[i], &methodInfo); err != nil {
	//		return
	//	}
	//}
	methodID := fmt.Sprintf("%x", message.MethodID)

	msg.SmcAddress = message.Contract
	if msg.SmcName, msg.Method, err = contractNameAndMethod(message.Contract, chainID, methodID); err != nil {
		return
	}

	if methodID == transferMethodID {
		if len(message.Items) != 2 {
			return msg, errors.New("items count error")
		}

		var to types.Address
		if err = rlp.DecodeBytes(message.Items[0], &to); err != nil {
			return
		}

		var value bn.Number
		if err = rlp.DecodeBytes(message.Items[1], &value); err != nil {
			return
		}
		msg.To = to
		msg.Value = value.String()
	}

	return
}

func contractNameAndMethod(contractAddress types.Address, chainID, methodID string) (contractName string, method string, err error) {

	addrS := nodeAddrSlice(chainID)

	contract := new(std.Contract)
	if err = DoHttpQueryAndParse(addrS, std.KeyOfContract(contractAddress), contract); err != nil {
		return
	}

	for _, methodItem := range contract.Methods {
		if methodItem.MethodID == methodID {
			method = methodItem.ProtoType
			break
		}
	}

	return contract.Name, method, nil
}

func tokenName(chainID string, tokenAddress types.Address) (name string, err error) {

	addrS := nodeAddrSlice(chainID)

	token := new(std.Token)

	if err = DoHttpQueryAndParse(addrS, std.KeyOfToken(tokenAddress), token); err != nil {
		return
	}

	return token.Name, err
}

func contractOfName(chainID, orgID, contractName string) (contract *std.Contract, err error) {

	addrS := nodeAddrSlice(chainID)

	key := std.KeyOfContractsWithName(orgID, contractName)

	contractList := new(std.ContractVersionList)
	err = DoHttpQueryAndParse(addrS, key, contractList)
	if err != nil {
		return nil, errors.New("Is orgName or contract's name right? error: " + err.Error())
	}

	var blkHeightResult *BlockHeightResult
	blkHeightResult, err = BlockHeight(chainID)
	if err != nil {
		return
	}

	effectiveAddr := func() types.Address {
		heightLen := len(contractList.EffectHeights)
		var effectIndex int
		for index, item := range contractList.EffectHeights {
			effectIndex = index
			if item < blkHeightResult.LastBlock {
				if index < heightLen-1 && contractList.EffectHeights[index] > item {
					break
				}
			}
		}

		return contractList.ContractAddrList[effectIndex]
	}

	key = std.KeyOfContract(effectiveAddr())
	contract = new(std.Contract)

	err = DoHttpQueryAndParse(addrS, key, contract)

	return
}

func contractOfTokenName(chainID, tokenName string) (contract *std.Contract, err error) {

	addrS := nodeAddrSlice(chainID)

	key := std.KeyOfTokenWithName(tokenName)
	tokenAddr := new(string)

	err = DoHttpQueryAndParse(addrS, key, tokenAddr)
	if err != nil {
		return
	}

	key = std.KeyOfContract(*tokenAddr)
	contract = new(std.Contract)

	err = DoHttpQueryAndParse(addrS, key, contract)
	if err != nil {
		return
	}

	return contractOfName(chainID, contract.OrgID, contract.Name)
}
