package core

import (
	tx1 "blockchain/abciapp_v1.0/tx/tx"
	types2 "blockchain/abciapp_v1.0/types"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdkimpl/helper"
	"blockchain/tx2"
	"blockchain/types"
	"cmd/bcc/cache"
	"cmd/bcc/common"
	"common/wal"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tendermint/go-crypto"
	"github.com/tendermint/tendermint/rpc/core/types"
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
	err = common.DoHttpRequestAndParse(addrS, "abci_info", params, result)
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
	err = common.DoHttpRequestAndParse(addrS, "block", params, result)
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

	for _, blkResult := range blkResults.Results.DeliverTx {
		var tx *TxResult
		if tx, err = Transaction(chainID, hex.EncodeToString(blkResult.TxHash), result); err != nil {
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
func Transaction(chainID, txHash string, resultBlock *core_types.ResultBlock) (tx *TxResult, err error) {

	defer FuncRecover(&err)

	requireNotEmpty("txHash", txHash)

	if txHash[:2] == "0x" {
		txHash = txHash[2:]
	}

	if chainID == "" {
		chainID = common.GetBCCConfig().DefaultChainID
		tx2.Init(chainID)
		crypto.SetChainId(chainID)
	}

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
		if txHash == hash {
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
	} else if splitTx[1] == "v2" {
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
	}

	tokenInfo := new(std.Token)
	err = common.DoHttpQueryAndParse(addrS, std.KeyOfGenesisToken(), tokenInfo)
	if err != nil {
		return
	}

	result.TokenAddress = tokenAddress
	result.TokenName = tokenInfo.Name

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
		if err = common.DoHttpQueryAndParse(addrS, token, tokenBalance); err != nil {
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
func AllContractInfo(chainID string) (ContractAddrList []string, err error) {

	ContractAddrList = make([]string, 0)

	addrS := nodeAddrSlice(chainID)

	err = common.DoHttpQueryAndParse(addrS, std.KeyOfAllContracts(), &ContractAddrList)

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
		result.Version = "0.0.0.0"
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
	err = common.DoHttpRequestAndParse(addrS, "block_results", params, blkResults)
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
		ContractAddrList, err := AllContractInfo(chainID)
		if err != nil {
			return nil, err
		}
		fmt.Println("OK")
		fmt.Println("Response:")
		for _, v := range ContractAddrList {
			jsIndent, _ := json.MarshalIndent(&v, "", "  ")
			fmt.Printf("  %s\n", string(jsIndent))
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
