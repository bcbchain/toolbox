package core

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bcbchain/bcbchain/abciapp_v1.0/keys"
	tx1 "github.com/bcbchain/bcbchain/abciapp_v1.0/tx/tx"
	types2 "github.com/bcbchain/bcbchain/abciapp_v1.0/types"
	types3 "github.com/bcbchain/bclib/tendermint/abci/types"
	"github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/bcbchain/bclib/tx/v2"
	"github.com/bcbchain/bclib/types"
	"github.com/bcbchain/bclib/wal"
	"github.com/bcbchain/sdk/sdk"
	"github.com/bcbchain/sdk/sdk/bn"
	"github.com/bcbchain/sdk/sdk/rlp"
	"github.com/bcbchain/sdk/sdk/std"
	"github.com/bcbchain/sdk/sdkimpl/helper"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
	"github.com/bcbchain/toolbox/bcc/cache"
	"github.com/bcbchain/toolbox/bcc/common"
	"io/ioutil"
	"math/big"
	"strings"
	"time"
)

//Query query state db with path
func Query(path string, data []byte, height int64, trusted bool, chainID string) (query *types2.ResultABCIQuery, err error) {
	return
}

//BlockHeight blockHeight query
func BlockHeight(chainID string) (blkHeight *BlockHeightResult, err error) {

	defer FuncRecover(&err)

	addrS := nodeAddrSlice(chainID)

	result := new(core_types.ResultABCIInfo)
	params := map[string]interface{}{}
	err = common.DoHttpRequestAndParseExBlock(addrS, "abci_info", params, result)
	if err != nil {
		return
	}
	blkHeight = new(BlockHeightResult)
	blkHeight.LastBlock = result.Response.LastBlockHeight

	return
}

//Block block information query
func Block(height *int64, bTime string, num *int64, chainID string) (blk *BlockResult, err error) {

	_, _, chainID = prepare("", "", chainID)

	if height == nil && bTime == "" {
		blkResult, err := BlockHeight(chainID)
		if err != nil {
			return nil, err
		}
		height = &blkResult.LastBlock
	}

	if height != nil {
		return block(*height, chainID)
	} else {
		return blockEx(bTime, num, chainID)
	}
}

func BlockForRpc(height int64, bTime string, num int64, chainID string) (blk *BlockResult, err error) {

	return Block(&height, bTime, &num, chainID)
}

//block block information query
func block(height int64, chainID string) (blk *BlockResult, err error) {

	defer FuncRecover(&err)

	addrS := nodeAddrSlice(chainID)

	result := new(core_types.ResultBlock)
	params := map[string]interface{}{"height": height}
	err = common.DoHttpRequestAndParseExBlock(addrS, "block", params, result)
	if err != nil {
		return
	}

	blk = new(BlockResult)
	blk.BlockHeight = result.BlockMeta.Header.Height
	blk.BlockHash = "0x" + hex.EncodeToString(result.BlockMeta.BlockID.Hash)
	blk.ParentHash = "0x" + hex.EncodeToString(result.BlockMeta.Header.LastBlockID.Hash)
	blk.ChainID = result.BlockMeta.Header.ChainID
	blk.ValidatorHash = "0x" + hex.EncodeToString(result.BlockMeta.Header.ValidatorsHash)
	blk.ConsensusHash = "0x" + hex.EncodeToString(result.BlockMeta.Header.ConsensusHash)
	blk.BlockTime = result.BlockMeta.Header.Time.String()
	blk.BlockSize = result.BlockSize
	blk.ProposerAddress = result.BlockMeta.Header.ProposerAddress

	blk.Txs = make([]TxResult, 0)
	var blkResults *core_types.ResultBlockResults
	if blkResults, err = blockResults(chainID, height); err != nil {
		return blk, nil
	}

	blk.Txs = make([]TxResult, 0)
	for k, ResDeliver := range blkResults.Results.DeliverTx {
		var tx *TxResult
		if tx, err = transactionBlock(k, ResDeliver, result); err != nil {
			return
		}
		blk.Txs = append(blk.Txs, *tx)
	}

	return
}

//blockEx block information query
func blockEx(bTime string, num *int64, chainID string) (blk *BlockResult, err error) {

	addrS := nodeAddrSlice(chainID)
	cache.SetAddrList(addrS)

	t, err := time.ParseInLocation("2006-01-02 15:04:05", bTime, time.UTC)
	if err != nil {
		return
	}

	nearlyHeight := cache.BinarySearchEx(0, 0, 128, t)
	if nearlyHeight <= 0 {
		return nil, errors.New("cannot find nearly block about time=" + bTime)
	}
	// if it less than require time then nearlyHeight plus one
	if cache.CompareWithTime(nearlyHeight, t, false) == -1 {
		nearlyHeight += 1
	}

	if num != nil && *num > 1 {
		return blockSimpleResult(nearlyHeight, *num, addrS)
	} else {
		return block(nearlyHeight, chainID)
	}
}

// blockSimpleResult get simple block result from h to h+num
func blockSimpleResult(h, num int64, addrS []string) (blk *BlockResult, err error) {
	blk = new(BlockResult)
	blk.Result = make([]simpleBlockResult, 0)

	result := new(core_types.ResultBlock)
	index := int64(0)
	for index < num {
		err = common.DoHttpRequestAndParse(addrS, "block", map[string]interface{}{"height": h + index}, result)
		if err != nil {
			return
		}

		blk.Result = append(blk.Result, simpleBlockResult{
			BlockHeight: h + index,
			BlockHash:   "0x" + result.BlockMeta.BlockID.Hash.String(),
			BlockTime:   result.BlockMeta.Header.Time.String()})

		index++
	}

	return
}

//Transaction transaction information query
func transactionBlock(k int, ResDeliver *types3.ResponseDeliverTx, resultBlock *core_types.ResultBlock) (tx *TxResult, err error) {

	//ParseTX
	var (
		transaction tx1.Transaction
		fromAddr    string
		msg         Message
		GasLimit    uint64
		Nonce       uint64
		Note        string
	)

	messages := make([]Message, 0)
	txStr := string(resultBlock.Block.Txs[k])
	splitTx := strings.Split(txStr, ".")

	if splitTx[1] == "v1" {
		// parse transaction V1
		fromAddr, _, err = transaction.TxParse(crypto.GetChainId(), txStr)
		if err != nil {
			return
		}
		msg, err = messageV1Parse(transaction)
		if err != nil {
			return
		}
		messages = append(messages, msg)
		GasLimit = transaction.GasLimit
		Nonce = transaction.Nonce
		Note = transaction.Note

	} else if splitTx[1] == "v2" {
		// parse transaction V2
		var txv2 types.Transaction
		var pubKey crypto.PubKeyEd25519
		txv2, pubKey, err := tx2.TxParse(txStr)
		if err != nil {
			return nil, err
		}

		fromAddr = pubKey.Address(crypto.GetChainId())

		var msg Message
		for i := 0; i < len(txv2.Messages); i++ {
			msg, err = messageV2Parse(txv2.Messages[i])
			if err != nil {
				return nil, err
			}
			messages = append(messages, msg)
		}
		GasLimit = uint64(txv2.GasLimit)
		Nonce = txv2.Nonce
		Note = txv2.Note
	} else if splitTx[1] == "v2" || splitTx[1] == "v3" {
		// parse transaction V2
		var txv2 types.Transaction
		var pubKey crypto.PubKeyEd25519
		txv2, pubKey, err := tx2.TxParse(txStr)
		if err != nil {
			return nil, err
		}

		fromAddr = pubKey.Address(crypto.GetChainId())

		var msg Message
		for i := 0; i < len(txv2.Messages); i++ {
			msg, err = messageV2Parse(txv2.Messages[i])
			if err != nil {
				return nil, err
			}
			messages = append(messages, msg)
		}
		GasLimit = uint64(txv2.GasLimit)
		Nonce = txv2.Nonce
		Note = txv2.Note
	} else {
		err = errors.New("unsupported tx=" + txStr)
		return
	}

	tx = new(TxResult)
	tx.TxHash = "0x" + strings.ToLower(hex.EncodeToString(ResDeliver.TxHash))
	tx.TxTime = resultBlock.BlockMeta.Header.Time.String()
	tx.Code = ResDeliver.Code
	tx.Log = ResDeliver.Log
	tx.BlockHash = "0x" + hex.EncodeToString(resultBlock.BlockMeta.BlockID.Hash)
	tx.BlockHeight = resultBlock.BlockMeta.Header.Height
	tx.From = fromAddr
	tx.Nonce = Nonce
	tx.GasLimit = GasLimit
	tx.Fee = ResDeliver.Fee
	tx.Note = Note
	tx.Messages = messages
	tx.Tags = make(map[string]*Tag)

	for _, item := range ResDeliver.Tags {
		Tag := new(Tag)
		Tag.Receipt = make(map[string]interface{})
		Receipt := make(map[string]interface{})

		err = json.Unmarshal(item.Value, &Tag)
		if err != nil {
			return nil, err
		}

		aDec, err := base64.StdEncoding.DecodeString(Tag.ReceiptBytes)
		err = json.Unmarshal(aDec, &Receipt)
		if err != nil {
			return nil, err
		}

		Tag.Receipt = Receipt
		Tag.ReceiptHash = "0x" + Tag.ReceiptHash

		tx.Tags[string(item.Key)] = Tag
	}

	return
}

func messageV1Parse(transation tx1.Transaction) (msg Message, err error) {

	var methodInfo tx1.MethodInfo
	if err = rlp.DecodeBytes(transation.Data, &methodInfo); err != nil {
		return
	}
	methodID := fmt.Sprintf("%x", methodInfo.MethodID)

	msg.SmcAddress = transation.To
	if msg.SmcName, msg.Method, err = contractNameAndMethod2(transation.To, methodID); err != nil {
		return
	}

	if methodID == transferMethodIDV1 {
		var itemsBytes = make([][]byte, 0)
		if err = rlp.DecodeBytes(methodInfo.ParamData, &itemsBytes); err != nil {
			return
		}
		msg.To = string(itemsBytes[0])
		msg.Value = new(big.Int).SetBytes(itemsBytes[1][:]).String()
	}

	return
}

func messageV2Parse(message types.Message) (msg Message, err error) {

	methodID := fmt.Sprintf("%x", message.MethodID)

	msg.SmcAddress = message.Contract
	if msg.SmcName, msg.Method, err = contractNameAndMethod2(message.Contract, methodID); err != nil {
		return
	}

	if methodID == transferMethodIDV2 {
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
		msg.Method = "Transfer(types.Address,bn.Number)"
	}

	return
}

func Transaction(chainID, txHash string, resultBlock *core_types.ResultBlock) (tx *TxResult, err error) {

	defer FuncRecover(&err)

	requireNotEmpty("txHash", txHash)

	if txHash[:2] == "0x" {
		txHash = txHash[2:]
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
	}
	tx2.Init(chainID)
	crypto.SetChainId(chainID)

	addrS := nodeAddrSlice(chainID)

	result := new(core_types.ResultTx)
	params := map[string]interface{}{"hash": txHash}
	err = common.DoHttpRequestAndParse(addrS, "tx", params, result)
	if err != nil {
		return
	}

	if resultBlock == nil {
		resultBlock = new(core_types.ResultBlock)
		params = map[string]interface{}{"height": result.Height}
		err = common.DoHttpRequestAndParse(addrS, "block", params, resultBlock)
		if err != nil {
			return
		}
	}

	var txStr string
	var blkResults *core_types.ResultBlockResults
	if blkResults, err = blockResults(chainID, result.Height); err != nil {
		return
	}

	for k, v := range blkResults.Results.DeliverTx {
		hash := hex.EncodeToString(v.TxHash)
		if hash[:2] == "0x" {
			txHash = txHash[2:]
		}
		if strings.ToLower(txHash) == hash {
			txStr = string(resultBlock.Block.Txs[k])
		}
	}

	nonce, gasLimit, fromAddr, note, messages, err := parseTx(chainID, txStr, resultBlock.Block.Height, resultBlock.Block.ChainVersion)
	if err != nil {
		return
	}

	tx = new(TxResult)
	tx.TxHash = "0x" + txHash
	tx.TxTime = resultBlock.BlockMeta.Header.Time.String()
	tx.Code = result.DeliverResult.Code
	tx.Log = result.DeliverResult.Log
	tx.BlockHash = "0x" + hex.EncodeToString(resultBlock.BlockMeta.BlockID.Hash)
	tx.BlockHeight = resultBlock.BlockMeta.Header.Height
	tx.From = fromAddr
	tx.Nonce = nonce
	tx.GasLimit = gasLimit
	tx.Fee = result.DeliverResult.Fee
	tx.Note = note
	tx.Messages = messages
	tx.Tags = make(map[string]*Tag)

	for _, item := range result.DeliverResult.Tags {
		Tag := new(Tag)
		Tag.Receipt = make(map[string]interface{})
		Receipt := make(map[string]interface{})

		err = json.Unmarshal(item.Value, &Tag)
		if err != nil {
			return nil, err
		}

		aDec, err := base64.StdEncoding.DecodeString(Tag.ReceiptBytes)
		err = json.Unmarshal(aDec, &Receipt)
		if err != nil {
			return nil, err
		}

		Tag.Receipt = Receipt
		Tag.ReceiptHash = "0x" + Tag.ReceiptHash

		tx.Tags[string(item.Key)] = Tag
	}

	return
}

func parseTx(chainID, txStr string, height int64, chainVersion *int64) (nonce, gasLimit uint64, fromAddr, note string, messages []Message, err error) {

	messages = make([]Message, 0)

	splitTx := strings.Split(txStr, ".")
	if splitTx[1] == "v1" {
		var txv1 tx1.Transaction
		fromAddr, _, err = txv1.TxParse(chainID, txStr)
		if err != nil {
			return
		}
		nonce = txv1.Nonce
		note = txv1.Note
		gasLimit = txv1.GasLimit

		var msg Message
		msg, err = messageV1(chainID, txv1, height, chainVersion)
		if err != nil {
			return
		}
		messages = append(messages, msg)
	} else if splitTx[1] == "v2" || splitTx[1] == "v3" {
		var txv2 types.Transaction
		var pubKey crypto.PubKeyEd25519
		txv2, pubKey, err = tx2.TxParse(txStr)
		if err != nil {
			return
		}
		fromAddr = pubKey.Address(chainID)
		nonce = txv2.Nonce
		note = txv2.Note
		gasLimit = uint64(txv2.GasLimit)

		var msg Message
		for i := 0; i < len(txv2.Messages); i++ {
			msg, err = message(chainID, txv2.Messages[i], height, chainVersion)
			if err != nil {
				return
			}
			messages = append(messages, msg)
		}
	} else {
		err = errors.New("unsupported tx=" + txStr)
		return
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
		err = common.DoHttpQueryAndParse(addrS, std.KeyOfGenesisToken(), tokenResult)
		if err != nil {
			return nil, errors.New("get genesis token error: " + err.Error())
		}
		tokenAddress = tokenResult.Address
		tokenName = tokenResult.Name
	} else {
		if value, err = common.DoHttpQuery(addrS, std.KeyOfTokenWithName(tokenName)); err != nil {
			return
		}
		if len(value) == 0 {
			return nil, errors.New("invalid tokenName")
		}

		if err = json.Unmarshal(value, &tokenAddress); err != nil {
			return
		}
	}

	if value, err = common.DoHttpQuery(addrS, std.KeyOfAccountToken(accAddress, tokenAddress)); err != nil {
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
		if tokenBalance.BVMBalanceMod != nil {
			result.BVMBalance = tokenBalance.BVMBalanceMod.String()
		}
	}

	tokenInfo := new(std.Token)
	err = common.DoHttpQueryAndParse(addrS, std.KeyOfGenesisToken(), tokenInfo)
	if err != nil {
		return
	}

	result.TokenAddress = tokenAddress
	result.TokenName = tokenInfo.Name

	if tokenName != "" {
		result.TokenName = tokenName
	}

	return
}

func allBalance(chainID string, address types.Address) (items *[]BalanceItemResult, err error) {

	addrS := nodeAddrSlice(chainID)

	tokens := make([]string, 0)
	if err = common.DoHttpQueryAndParse(addrS, std.KeyOfAccount(address), &tokens); err != nil {
		return
	}

	balanceItems := make([]BalanceItemResult, 0)
	for _, token := range tokens {
		splitToken := strings.Split(token, "/")
		if splitToken[4] != "token" || len(splitToken) != 6 {
			continue
		}
		tokenBalance := new(TokenBalance)
		if err = common.DoHttpQueryAndParse(addrS, token, &tokenBalance); err != nil {
			return
		}

		var name string
		if name, err = tokenName(chainID, tokenBalance.Address); err != nil {
			return
		}
		BalanceItemRes := BalanceItemResult{
			TokenAddress: tokenBalance.Address,
			TokenName:    name,
			Balance:      tokenBalance.Balance.String()}

		if tokenBalance.BVMBalanceMod != nil {
			BalanceItemRes.BVMBalance = tokenBalance.BVMBalanceMod.String()
		}

		balanceItems = append(balanceItems, BalanceItemRes)
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
	value, err := common.DoHttpQuery(addrS, std.KeyOfAccountNonce(address))
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

// All info of contract
func ContractInfo(chainID, orgID, contractName string) (contracts map[string]std.Contract, err error) {

	defer FuncRecover(&err)

	ContractList := new(std.ContractVersionList)

	addrS := nodeAddrSlice(chainID)

	err = common.DoHttpQueryAndParse(addrS, std.KeyOfContractsWithName(orgID, contractName), &ContractList)

	contracts = make(map[string]std.Contract)
	for k, v := range ContractList.ContractAddrList {

		contract := new(std.Contract)

		err = common.DoHttpQueryAndParse(addrS, std.KeyOfContract(v), &contract)

		contracts[contractName+string(k)] = *contract
	}
	return
}

// Query contract information with address
func ContractInfoWithAddr(chainID, contractAddr string) (contract *std.Contract, err error) {

	addrS := nodeAddrSlice(chainID)

	err = common.DoHttpQueryAndParse(addrS, std.KeyOfContract(contractAddr), &contract)

	return
}

// Query all contract information
func AllContractInfo(chainID string) (contractAddrList []string, contractList []interface{}, err error) {

	contractAddrList = make([]string, 0)

	addrS := nodeAddrSlice(chainID)
	err = common.DoHttpQueryAndParse(addrS, std.KeyOfAllContracts(), &contractAddrList)

	key := "["
	for _, v := range contractAddrList {
		key = key + v + ","
	}
	key = key[:len(key)-1] + "]"

	contract := new(std.Contract)
	contractList, err = common.DoBulkHttpQueryAndParse(addrS, std.KeyOfContract(key), contract)
	return
}

func TokenInfo(tokenName, chainId string) (err error) {
	tokenAddr := ""
	addrS := nodeAddrSlice(chainId)
	err = common.DoHttpQueryAndParse(addrS, std.KeyOfTokenWithName(tokenName), &tokenAddr)
	if err != nil {
		return err
	}

	token := new(std.Token)
	err = common.DoHttpQueryAndParse(addrS, std.KeyOfToken(tokenAddr), &token)
	if err != nil {
		return err
	}

	fmt.Println("OK")
	jsIndent, _ := json.MarshalIndent(&token, "", "  ")
	fmt.Printf("Response: %s\n", string(jsIndent))

	return
}

// Query all token  information
func AllTokenInfo(chainID string) (err error) {

	tokenAddrList := make([]string, 0)

	addrS := nodeAddrSlice(chainID)
	err = common.DoHttpQueryAndParse(addrS, std.KeyOfAllToken(), &tokenAddrList)

	key := "["
	for _, v := range tokenAddrList {
		key = key + v + ","
	}
	key = key[:len(key)-1] + "]"

	token := std.Token{}
	tokenList, err := common.DoBulkHttpQueryAndParse(addrS, std.KeyOfToken(key), &token)
	fmt.Println("OK")
	fmt.Println("Response: ")
	for _, token := range tokenList {
		fmt.Printf("   token name: %s   token symbol: %s\n   token addr: %s\n\n", token.(std.Token).Name, token.(std.Token).Symbol, token.(std.Token).Address)
	}
	return
}

// Query organization information
func QueryOrgInfo(orgID, chainID string) (OrgInfo *std.Organization, err error) {

	addrS := nodeAddrSlice(chainID)

	OrgInfo = new(std.Organization)

	err = common.DoHttpQueryAndParse(addrS, std.GetOrganizaitionInfo(orgID), &OrgInfo)

	return
}

// CommitTx commit transaction information
func CommitTx(chainID, tx string) (commit *CommitTxResult, err error) {

	defer FuncRecover(&err)

	requireNotEmpty("tx", tx)

	addrS := nodeAddrSlice(chainID)

	var result *core_types.ResultTx
	result, err = common.DoHttpCommitTxAndParseAsync(addrS, tx)
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
	if result.CheckResult.Code == types.CodeBVMQueryOK {
		commit.Data = result.CheckResult.Data
	}

	return
}

//Version version information for the current block
func Version() (result *VersionResult, err error) {

	defer FuncRecover(&err)

	result = new(VersionResult)
	var version []byte
	version, err = ioutil.ReadFile("./version")
	if err != nil {
		err = nil
		result.Version = "0.0.0.1"
		return
	}
	result.Version = string(version)
	result.Version = strings.Replace(result.Version, "\r\n", "", -1)
	result.Version = strings.Replace(result.Version, "\n", "", -1)
	return
}

func blockResults(chainID string, height int64) (blkResults *core_types.ResultBlockResults, err error) {

	addrS := nodeAddrSlice(chainID)

	blkResults = new(core_types.ResultBlockResults)
	params := map[string]interface{}{"height": height}
	err = common.DoHttpRequestAndParseExBlock(addrS, "block_results", params, blkResults)
	if err != nil {
		return
	}

	return
}

func message(chainID string, message types.Message, height int64, chainVersion *int64) (msg Message, err error) {

	methodID := fmt.Sprintf("%x", message.MethodID)

	msg.SmcAddress = message.Contract
	if msg.SmcName, msg.Method, err = contractNameAndMethod(message.Contract, chainID, methodID, height, chainVersion); err != nil {
		return
	}

	if methodID == transferMethodIDV2 {
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

func messageV1(chainID string, tx tx1.Transaction, height int64, chainVersion *int64) (msg Message, err error) {

	var methodInfo tx1.MethodInfo
	if err = rlp.DecodeBytes(tx.Data, &methodInfo); err != nil {
		return
	}
	methodID := fmt.Sprintf("%x", methodInfo.MethodID)

	msg.SmcAddress = tx.To
	if msg.SmcName, msg.Method, err = contractNameAndMethod(tx.To, chainID, methodID, height, chainVersion); err != nil {
		return
	}

	if methodID == transferMethodIDV1 {
		var itemsBytes = make([][]byte, 0)
		if err = rlp.DecodeBytes(methodInfo.ParamData, &itemsBytes); err != nil {
			return
		}
		msg.To = string(itemsBytes[0])
		msg.Value = new(big.Int).SetBytes(itemsBytes[1][:]).String()
	}

	return
}

func contractNameAndMethod2(contractAddress keys.Address, methodID string) (contractName string, method string, err error) {
	contract := new(types2.Contract)
	common.RWLock.RLock()
	v, ok := common.ContractMap[contractAddress]
	common.RWLock.RUnlock()
	if ok == true {
		contract = v
	} else {
		param := map[string]interface{}{"path": std.KeyOfContract(contractAddress)}
		result := new(types2.ResultABCIQuery)
		if err = common.DoHttpRequestAndParseExBlock(common.GetBCCConfig().Urls[common.GetBCCConfig().DefaultChainID], "abci_query", param, result); err != nil {
			return
		}
		err = json.Unmarshal(result.Response.Value, contract)
		if err != nil {
			return
		}
		common.RWLock.Lock()
		common.ContractMap[contractAddress] = contract
		common.RWLock.Unlock()
	}

	for _, methodItem := range contract.Methods {
		if methodItem.MethodId == methodID {
			method = methodItem.Prototype
			break
		}
	}

	return contract.Name, method, nil
}

func contractNameAndMethod(contractAddress types.Address, chainID, methodID string, height int64, chainVersion *int64) (contractName string, method string, err error) {

	addrS := nodeAddrSlice(chainID)

	contract := new(std.Contract)
	if err = common.DoHttpQueryAndParse(addrS, std.KeyOfContract(contractAddress), contract); err != nil {
		return
	}

	if chainVersion != nil && contract.LoseHeight != 0 && contract.LoseHeight < height {
		conVer := new(std.ContractVersionList)
		if err = common.DoHttpQueryAndParse(addrS, std.KeyOfContractsWithName(contract.OrgID, contract.Name), conVer); err == nil {
			for index, eh := range conVer.EffectHeights {
				if eh <= height {
					tmp := new(std.Contract)
					if err = common.DoHttpQueryAndParse(addrS, std.KeyOfContract(conVer.ContractAddrList[index]), tmp); err == nil {
						if tmp.LoseHeight == 0 || (tmp.LoseHeight != 0 && tmp.LoseHeight > height) {
							contract = tmp
							break
						}
					} else {
						return
					}
				}
			}
		} else {
			return
		}
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

	if err = common.DoHttpQueryAndParse(addrS, std.KeyOfToken(tokenAddress), token); err != nil {
		return
	}

	return token.Name, err
}

func contractOfName(chainID, orgID, contractName string) (contract *std.Contract, err error) {

	addrS := nodeAddrSlice(chainID)

	key := std.KeyOfContractsWithName(orgID, contractName)

	contractList := new(std.ContractVersionList)
	err = common.DoHttpQueryAndParse(addrS, key, contractList)
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
				if index < heightLen-1 && contractList.EffectHeights[index+1] > blkHeightResult.LastBlock {
					break
				}
			}
		}

		return contractList.ContractAddrList[effectIndex]
	}

	key = std.KeyOfContract(effectiveAddr())
	contract = new(std.Contract)

	err = common.DoHttpQueryAndParse(addrS, key, contract)

	return
}

func tokenAddressFromName(chainID, tokenName string) (tokenAddr types.Address, err error) {

	addrS := nodeAddrSlice(chainID)

	key := std.KeyOfTokenWithName(tokenName)

	err = common.DoHttpQueryAndParse(addrS, key, &tokenAddr)
	if err != nil {
		return
	}

	return
}

func QueryOfRpc(key, chainID string) (result *core_types.ResultABCIQuery, err error) {
	addrS := nodeAddrSlice(chainID)

	if strings.HasPrefix(key, "/org") && !strings.HasPrefix(key, "/orgJgaGConUyK81zibntUBjQ33PKctpk1K1G") {
		return nil, errors.New("no permission")
	}

	result = new(core_types.ResultABCIQuery)
	result, err = common.DoHttpQueryForRpc(addrS, key)
	if err != nil {
		return nil, errors.New("rpc query failed , error: " + err.Error())
	}

	return result, nil
}

// Query the contract information based on the parameters for Rpc
func ContractInfoForRPC(orgName, contractName, orgID, contractAddr, chainID string) (contract *std.Contract, err error) {

	if orgID != "" && contractName != "" && contractAddr == "" {
		contractList, err := ContractInfo(chainID, orgID, contractName)
		if err != nil {
			return nil, err
		}

		for _, v := range contractList {

			// 校验其他输入参数
			if orgName != "" {
				OrgInfo, err := QueryOrgInfo(orgID, chainID)
				if err != nil {
					return nil, err
				}

				if orgName != OrgInfo.Name {
					fmt.Println("Error: Input orgName is wrong.")
					return nil, err
				}
			}

			return &v, nil
		}

	} else if orgName != "" && contractName != "" && contractAddr == "" {
		contract, err := QueryContractInfoForRpc(orgName, contractName, chainID)
		if err != nil {
			return nil, err
		}

		// 校验其他输入参数
		if orgID != "" && orgID != contract.OrgID {
			fmt.Println("Error: Input orgID is wrong.")
			return nil, err
		}

		return contract, nil

	} else if contractAddr != "" {
		contract, err := ContractInfoWithAddr(chainID, contractAddr)
		if err != nil {
			return nil, err
		}

		// 校验其他输入参数
		if orgName != "" && orgID != "" {
			OrgInfo, err := QueryOrgInfo(orgID, chainID)
			if err != nil {
				return nil, err
			}

			if orgName != OrgInfo.Name {
				fmt.Println("Error: Input orgName is wrong.")
				return nil, nil
			}
		}
		if orgID != "" && orgID != contract.OrgID {
			fmt.Println("Error: orgID orgName is wrong.")
			return nil, nil
		}

		return contract, nil

	} else if orgName == "" && contractName == "" && orgID == "" && contractAddr == "" {
		_, contractList, err := AllContractInfo(chainID)
		if err != nil {
			fmt.Println("Query ContractInfo failed")
			return nil, err
		}
		fmt.Println("OK")
		fmt.Println("Response: ")
		for _, contract := range contractList {
			fmt.Printf("   contract name: %s\n   contract addr: %s\n\n", contract.(std.Contract).Name, contract.(std.Contract).Address)
		}
	} else {
		fmt.Println("Insufficient input parameters")
		return nil, nil
	}

	return
}

type BlockChainHelper struct {
	smc sdk.ISmartContract
}

func QueryContractInfoForRpc(orgName, contractName, chainID string) (contract *std.Contract, err error) {
	bh := helper.BlockChainHelper{}
	orgID := bh.CalcOrgID(orgName)

	addrS := nodeAddrSlice(chainID)

	List := new(std.ContractVersionList)
	err = common.DoHttpQueryAndParse(addrS, std.KeyOfContractsWithName(orgID, contractName), &List)

	for _, v := range List.ContractAddrList {
		contract, err = ContractInfoWithAddr(chainID, v)
		if err != nil {
			return nil, err
		}
		return contract, nil
	}
	return
}
