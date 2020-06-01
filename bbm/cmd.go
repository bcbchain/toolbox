package main

import (
	"errors"
	"fmt"
	"github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/bcbchain/bclib/tx/v2"
	"github.com/bcbchain/bclib/types"
	bn2 "github.com/bcbchain/sdk/sdk/bn"
	"github.com/bcbchain/toolbox/bbm/account"
	"github.com/bcbchain/toolbox/bbm/common"
	"github.com/bcbchain/toolbox/bbm/rpcclient"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	protoTransfer = "Transfer(types.Address,bn.Number)"
	fee           = 1250000
)

func prepare(accounts, value string) error {
	// check param
	if accounts == "" {
		return errors.New("Need number of accounts")
	}
	iAccounts, err := strconv.Atoi(accounts)
	if nil != err {
		return err
	}

	if value == "" {
		return errors.New("Need amount for per account")
	}
	iValue, err := strconv.Atoi(value)
	if nil != err {
		return err
	}
	logger.Info("Get flags", "accounts", iAccounts, "value", iValue)

	// first, revert last accounts before create new accounts
	keys, err := account.GetKeys(0, iAccounts-1)
	if nil != err {
		return err
	}
	if len(keys) > 0 {
		logger.Info("Find accounts, revert balance to genesis", "accounts", len(keys))
		err = doRevert(keys)
		if nil != err {
			return err
		}
	}

	// create account
	err = account.CreateAccounts(iAccounts)
	if nil != err {
		return err
	}
	keys, err = account.GetKeys(0, iAccounts-1)
	if nil != err {
		return err
	}
	logger.Info("Create Accounts", "accounts number", len(keys))
	if len(keys) != iAccounts {
		return errors.New("Create accounts fail")
	}

	// transfer to new account
	err = doPrepare(keys, iValue)
	if nil != err {
		return errors.New("preTransfer to test account error")
	}

	// test get balance
	err = checkBalance(keys, iValue)
	if nil != err {
		return errors.New("preTransfer to test account error")
	}
	logger.Info("prepare is all ok")

	return nil
}

func transferGenesis(account []byte, value int64) error {
	client := rpcclient.InitClient(bbmConfig.NodeAddr, logger)
	token := common.GetGenesisToken(&client)
	chainID := common.GetChainID(&client)
	toAddr := crypto.PrivKeyEd25519FromBytes(account).PubKey().Address(chainID)

	nonce := common.AccountNonce(&client, token.Owner) + 1

	var tx string
	// tx1 and tx2
	if 2 == bbmConfig.txVersion {
		tx2.Init(chainID)
		privKey := "0x" + bbmConfig.GenesisKey

		params := make([]interface{}, 0)
		params = append(params, toAddr)
		params = append(params, int64(value))
		tx = common.GenerateTx(token.Address, protoTransfer, params, nonce, 500000, "", privKey)
	} else if 1 == bbmConfig.txVersion {
		var params []string
		params = append(params, toAddr, common.UintToHex(uint64(value)))
		tx = common.GenerateTx1(nonce, 500000, "", token.Address, "Transfer(smc.Address,big.Int)smc.Error", params, chainID, bbmConfig.GenesisKey)
	}

	response := client.SendTx([]byte(tx))
	if response.Code != 200 {
		logger.Debug("transfer response", "code", response.Code)
		return errors.New("transfer response error")
	}
	return nil
}

func getStep(step, second int64) int64 {
	// minimum step
	var minStep int64 = 5
	var maxStep int64 = 1000
	switch {
	case second > 20:
		// min step
		return minStep
	case second > 10:
		// sub step
		step = step / 2
		if step <= minStep {
			step = minStep
		}
		return step
	case second < 4:
		// add step
		step = step + step/2
		if step >= maxStep {
			step = maxStep
		}
		return step
	default:
		// keep
		return step
	}
	return minStep
}

func doPrepare(accounts map[string][]byte, iValue int) error {
	iNumAccounts := len(accounts)
	// get rounds
	var iRounds int
	for iRounds = 0; ; iRounds++ {
		var iTimes int
		pow := math.Pow(2, float64(iRounds))
		iTimes = iTimes + int(pow)
		if iTimes >= (iNumAccounts - 1) {
			break
		}
	}

	// get total amounts
	totalFee := (iNumAccounts - 1) * fee
	totalBalance := totalFee + (iNumAccounts * iValue)
	logger.Info("Quick Transfer", "totalAmount", totalBalance, "rounds", iRounds)

	// transfer first account(genesis to account[0])
	err := transferGenesis(accounts["0"], int64(totalBalance))
	if nil != err {
		return err
	}

	// trasferMap record all transfer times of every account (map[id]times)
	transferMap := make(map[int]int, iNumAccounts)
	// record base times
	for i := 0; i <= iRounds; i++ {
		subTotal := int(math.Pow(2, float64(i)))
		for ii := 0; ii < subTotal; ii++ {
			if subTotal+ii+1 > iNumAccounts {
				break
			}
			transferMap[ii] += 1
		}
	}

	// add sub account times
	for i := iRounds; i >= 0; i-- {
		subTotal := int(math.Pow(2, float64(i)))
		for ii := 0; ii < subTotal; ii++ {
			current := subTotal + ii
			if current+1 > iNumAccounts {
				break
			}
			transferMap[ii] += transferMap[current]
		}
	}

	// transfer
	logger.Info("Transfer to test accounts,Please wait a moment")
	var wg sync.WaitGroup
	step := 10
	for i := 0; i <= iRounds; i++ {
		subTotal := int(math.Pow(2, float64(i)))
		for ii := 0; ii < subTotal; ii++ {
			current := subTotal + ii
			if current+1 > iNumAccounts {
				// transfer done
				break
			}
			wg.Add(1)
			fromKey := fmt.Sprintf("%d", ii)
			toKey := fmt.Sprintf("%d", current)

			value := transferMap[current]*(fee+iValue) + iValue
			logger.Debug("doPrepare", "From", fromKey, "To", toKey, "Value", value)
			go transferPrePare(accounts[fromKey], accounts[toKey], &wg, value)
			// linux system have 1024 open files limit
			if (ii+1)%step == 0 {
				begin := time.Now().Unix()
				wg.Wait()
				end := time.Now().Unix()
				step = int(getStep(int64(step), end-begin))
				logger.Debug("Dynamic Step", "Step", step, "seconds", end-begin)
			}
		}
		wg.Wait()
		logger.Info("doPrepare Create", "TotalRounds", iRounds, "Current", i+1)
	}

	return nil
}

func transferPrePare(from, to []byte, wg *sync.WaitGroup, iValue int) {
	defer wg.Done()
	client := rpcclient.InitClient(bbmConfig.NodeAddr, logger)
	token := common.GetGenesisToken(&client)
	chainID := common.GetChainID(&client)
	toAddr := crypto.PrivKeyEd25519FromBytes(to).PubKey().Address(chainID)
	fromAddr := crypto.PrivKeyEd25519FromBytes(from).PubKey().Address(chainID)
	nonce := common.AccountNonce(&client, fromAddr) + 1

	var tx string
	if 2 == bbmConfig.txVersion {
		tx2.Init(chainID)
		privKey := "0x" + string(fmt.Sprintf("%x", from))
		params := make([]interface{}, 0)
		params = append(params, toAddr)
		params = append(params, iValue)
		tx = common.GenerateTx(token.Address, protoTransfer, params, nonce, 500000, "", privKey)
	} else if 1 == bbmConfig.txVersion {
		var params []string
		params = append(params, toAddr, common.UintToHex(uint64(iValue)))
		tx = common.GenerateTx1(nonce, 500000, "", token.Address, "Transfer(smc.Address,big.Int)smc.Error", params, chainID, fmt.Sprintf("%x", from))
	}

	response := client.SendTx([]byte(tx))
	if response.Code != 200 {
		logger.Debug("transfer response", "code", response.Code)
	}

	return
}

func checkBalance(accounts map[string][]byte, iValue int) error {
	client := rpcclient.InitClient(bbmConfig.NodeAddr, logger)
	token := common.GetGenesisToken(&client)
	tokenAddr := token.Address
	chainID := common.GetChainID(&client)
	for i := 0; i < len(accounts); i++ {
		key := fmt.Sprintf("%d", i)
		accountAddr := crypto.PrivKeyEd25519FromBytes(accounts[key]).PubKey().Address(chainID)
		item := common.BalanceOf(&client, accountAddr, tokenAddr)
		logger.Debug("check balance", "account", key, "balance", item)
		if 0 != item.CmpI(int64(iValue)) {
			return errors.New("balance of accout is error")
		}
	}

	return nil
}

func syncTest(from, to, round string) error {
	// check param
	froms := strings.Split(from, "-")
	if 2 != len(froms) {
		return errors.New("Flags (-f, --from) example: \"1-100\"")
	}
	tos := strings.Split(to, "-")
	if 2 != len(tos) {
		return errors.New("Flags (-t, --to) example: \"100-200\"")
	}

	fromBegin, err := strconv.Atoi(froms[0])
	if nil != err {
		return errors.New("Flags (-f, --from) must be a int number")
	}

	fromEnd, err := strconv.Atoi(froms[1])
	if nil != err {
		return errors.New("Flags (-f, --from) must be a int number")
	}

	toBegin, err := strconv.Atoi(tos[0])
	if nil != err {
		return errors.New("Flags (-t, --to) must be a int number")
	}

	toEnd, err := strconv.Atoi(tos[1])
	if nil != err {
		return errors.New("Flags (-t, --to) must be a int number")
	}

	logger.Info("Get flags", "fromBegin", fromBegin, "fromEnd", fromEnd, "toBegin", toBegin, "toEnd", toEnd)

	if (fromEnd - fromBegin) != (toEnd - toBegin) {
		return errors.New("from and to account not equal")
	}

	if (fromEnd < fromBegin) || (toEnd < toBegin) || (toBegin <= fromEnd) {
		return errors.New("from and to size error")
	}

	if round == "" {
		return errors.New("Need number of test rounds")
	}
	iRound, err := strconv.Atoi(round)
	if nil != err {
		return errors.New("Flags (-r, --round) must be a int number")
	}
	logger.Info("Get flags", "fromBegin", fromBegin, "fromEnd", fromEnd, "toBegin", toBegin, "toEnd", toEnd, "round", iRound)

	// get accounts
	fromKeys, _ := account.GetKeys(fromBegin-1, fromEnd-1)
	if (fromEnd - fromBegin + 1) != len(fromKeys) {
		return errors.New("fromKeys number error")
	}

	toKeys, _ := account.GetKeys(toBegin-1, toEnd-1)
	if (toEnd - toBegin + 1) != len(toKeys) {
		return errors.New("toKeys number error")
	}

	logger.Info("Get Accounts successful")

	// create waitgroup
	var wg sync.WaitGroup
	goNumbers := fromEnd - fromBegin + 1
	wg.Add(goNumbers)

	// get begin block height
	client := rpcclient.InitClient(bbmConfig.NodeAddr, logger)
	bcBeginHeight := common.GetCurrentBlockHeight(&client)
	if bcBeginHeight == 0 {
		return errors.New("Get Block Height error")
	}

	// print tps dynamically
	quit := make(chan int)
	go printDynamicTPS(&client, bcBeginHeight, quit)

	token := common.GetGenesisToken(&client)
	chainID := common.GetChainID(&client)

	// go testing
	for i := 0; i < goNumbers; i++ {
		from := fmt.Sprintf("%d", fromBegin+i-1)
		to := fmt.Sprintf("%d", toBegin+i-1)
		go doSyncTest(&client, &token.Address, chainID, fromKeys[from], toKeys[to], iRound, &wg)
	}
	wg.Wait()

	// send quit single to print thread
	quit <- 0

	// get test end block height
	bcEndHeight := common.GetCurrentBlockHeight(&client)
	if bcEndHeight == 0 {
		return errors.New("Get Block Height error")
	}

	if bcEndHeight == bcBeginHeight {
		logger.Info("SyncTest Done, no transaction found")
		return nil
	}

	// compute tps
	tps, err := getTPS(&client, bcBeginHeight, bcEndHeight)
	if nil != err {
		logger.Info("get txs and time error")
		return nil
	}
	logger.Info("SyncTest Done", "Total TPS", tps)

	return nil
}

func printDynamicTPS(client *rpcclient.RPCClient, beginBlock int64, quit chan int) {
	for {
		// wait for quit single
		select {
		case <-quit:
			return
		default:
		}

		// print tps per 10 blocks
		currentHeight := common.GetCurrentBlockHeight(client)
		if (currentHeight - beginBlock) > int64(bbmConfig.PrintTPSFrequency) {
			tps, _ := getDynamicTPS(client, beginBlock, beginBlock+int64(bbmConfig.PrintTPSFrequency))
			logger.Info("testing", "block height", beginBlock, "TPS", tps)
			beginBlock += int64(bbmConfig.PrintTPSFrequency)
		} else {
			time.Sleep(time.Second * 1)
		}
	}
}

func doSyncTest(client *rpcclient.RPCClient, contractAddr *types.Address, chainID string, from, to []byte, round int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < round; i++ {
		toAddr := crypto.PrivKeyEd25519FromBytes(to).PubKey().Address(chainID)
		fromAddr := crypto.PrivKeyEd25519FromBytes(from).PubKey().Address(chainID)
		nonce := common.AccountNonce(client, fromAddr) + 1

		var tx string
		// tx1 and tx2
		if 2 == bbmConfig.txVersion {
			tx2.Init(chainID)
			privKey := "0x" + string(fmt.Sprintf("%x", from))
			params := make([]interface{}, 0)
			params = append(params, toAddr)
			params = append(params, 1)
			tx = common.GenerateTx(*contractAddr, protoTransfer, params, nonce, 500000, "", privKey)
		} else if 1 == bbmConfig.txVersion {
			var params []string
			params = append(params, toAddr, common.UintToHex(uint64(1)))
			tx = common.GenerateTx1(nonce, 500000, "", *contractAddr, "Transfer(smc.Address,big.Int)smc.Error", params, chainID, fmt.Sprintf("%x", from))
		}

		response := client.SendTx([]byte(tx))

		if response.Code != 200 {
			logger.Info("transfer response", "round", i, "code", response.Code)
			break
		}

		// swap from and to
		from, to = to, from
	}

	return
}

func asyncTest(from, to string) error {
	// check param
	froms := strings.Split(from, "-")
	if 2 != len(froms) {
		return errors.New("Flags (-f, --from) example: \"1-100\"")
	}

	fromBegin, err := strconv.Atoi(froms[0])
	if nil != err {
		return errors.New("Flags (-f, --from) must be a int number")
	}

	fromEnd, err := strconv.Atoi(froms[1])
	if nil != err {
		return errors.New("Flags (-f, --from) must be a int number")
	}

	if "" == to {
		return errors.New("Need number of -t,--to")
	}

	toBegin, err := strconv.Atoi(to)
	if nil != err {
		return errors.New("Flags (-t, --to) must be a int number")
	}

	logger.Info("Get flags", "fromBegin", fromBegin, "fromEnd", fromEnd, "toBegin", toBegin)

	if (fromEnd < fromBegin) || (toBegin <= fromEnd) {
		return errors.New("from or to size error")
	}

	// get accounts
	fromKeys, _ := account.GetKeys(fromBegin-1, fromEnd-1)
	if (fromEnd - fromBegin + 1) != len(fromKeys) {
		return errors.New("fromKeys number error")
	}

	toKeys, _ := account.GetKeys(toBegin-1, toBegin-1)
	if 1 != len(toKeys) {
		return errors.New("toKeys number error")
	}

	logger.Info("Get Accounts successful")

	// transfer to test account
	var wg sync.WaitGroup
	goNumbers := fromEnd - fromBegin + 1
	wg.Add(goNumbers)

	client := rpcclient.InitClient(bbmConfig.NodeAddr, logger)

	// get begin block height
	bcBeginHeight := common.GetCurrentBlockHeight(&client)
	if bcBeginHeight == 0 {
		return errors.New("Get Block Height error")
	}
	logger.Info("async test", "begin block height", bcBeginHeight)

	// go testing
	chainID := common.GetChainID(&client)
	token := common.GetGenesisToken(&client)
	toKey := fmt.Sprintf("%d", toBegin-1)
	toAddr := crypto.PrivKeyEd25519FromBytes([]byte(toKey)).PubKey().Address(chainID)
	for i := 0; i < goNumbers; i++ {
		fromKey := fmt.Sprintf("%d", fromBegin+i-1)
		//logger.Info("Start test", "account from", fromKey, "account to ", toKey)
		go doAsyncTest(fromKeys[fromKey], &client, &token.Address, &toAddr, chainID, &wg)
	}
	wg.Wait()

	// wait for all txs are commit
	err = waitAsynRequestDone(&client)
	if err != nil {
		logger.Info("server not work")
		return nil
	}

	// get test end block height
	bcEndHeight := common.GetCurrentBlockHeight(&client)
	if bcEndHeight == 0 {
		return errors.New("Get Block Height error")
	}
	logger.Info("async test", "end block height", bcEndHeight)

	// get tps
	tps, err := getTPS(&client, bcBeginHeight, bcEndHeight)
	if nil != err {
		logger.Info("get txs and time error")
		return nil
	}
	logger.Info("Test End", "TPS", tps)
	logger.Info("syncTest Done")

	return nil
}

func doAsyncTest(from []byte, client *rpcclient.RPCClient, contractAddr, toAddr *types.Address, chainID string, wg *sync.WaitGroup) {
	defer wg.Done()

	fromAddr := crypto.PrivKeyEd25519FromBytes(from).PubKey().Address(chainID)
	nonce := common.AccountNonce(client, fromAddr) + 1

	var tx string
	// tx1 and tx2
	if 2 == bbmConfig.txVersion {
		tx2.Init(chainID)
		privKey := "0x" + string(fmt.Sprintf("%x", from))
		params := make([]interface{}, 0)
		params = append(params, *toAddr)
		params = append(params, 1)
		tx = common.GenerateTx(*contractAddr, protoTransfer, params, nonce, 500000, "", privKey)
	} else if 1 == bbmConfig.txVersion {
		var params []string
		params = append(params, *toAddr, common.UintToHex(uint64(1)))
		tx = common.GenerateTx1(nonce, 500000, "", *contractAddr, "Transfer(smc.Address,big.Int)smc.Error", params, chainID, fmt.Sprintf("%x", from))
	}
	client.SendTxAsync([]byte(tx))
	return
}

func waitAsynRequestDone(client *rpcclient.RPCClient) error {
	// waif server receive first tx
	time.Sleep(time.Second * 2)
	for {
		r, err := client.NumUnconfirmedTxs()
		if nil != err {
			logger.Info("get unconfirmed tx err")
			return err
		}
		// done
		if r.N == 0 {
			break
		}

		time.Sleep(time.Second * 2)
	}
	return nil
}

func getDynamicTPS(client *rpcclient.RPCClient, beginBlock, endBlock int64) (float64, error) {

	// get start block
	r, err := client.BlockWithHeight(beginBlock)
	if nil != err {
		logger.Info("get get block error", "height", beginBlock)
		return 0, err
	}
	beginTime := r.BlockMeta.Header.Time
	beginTxs := r.BlockMeta.Header.TotalTxs

	// find end block
	r, err = client.BlockWithHeight(endBlock)
	if nil != err {
		logger.Info("get get block error", "height", endBlock)
		return 0, err
	}
	endTime := r.BlockMeta.Header.Time
	endTxs := r.BlockMeta.Header.TotalTxs

	tps := float64(endTxs-beginTxs) / endTime.Sub(beginTime).Seconds()
	tps, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tps), 64)
	return tps, nil
}

func getTPS(client *rpcclient.RPCClient, beginBlock, endBlock int64) (float64, error) {
	var beginTime time.Time
	var beginTxs int64
	var beginHeight int64

	// find start block
	for i := beginBlock; i <= endBlock; i++ {
		r, err := client.BlockWithHeight(i)
		if nil != err {
			logger.Info("get get block error", "height", i)
			return 0, err
		}

		if (i == beginBlock) && (r.BlockMeta.Header.NumTxs != 0) {
			logger.Info("getTPS error, block is dirty, please don't call blockchain when testing")
			return 0, errors.New("getTPS error, block is dirty, please don't call block when testing")
		}

		// got it
		if r.BlockMeta.Header.NumTxs != 0 {
			break
		}

		// record last block's time and txs
		beginHeight = i
		beginTime = r.BlockMeta.Header.Time
		beginTxs = r.BlockMeta.Header.TotalTxs
	}

	logger.Info("Test summary", "begin block", beginHeight, "time", beginTime, "total txs", beginTxs)

	// find end block
	var endTime time.Time
	var endTxs int64
	var endHeight int64
	for i := endBlock; i >= beginBlock; i-- {
		r, err := client.BlockWithHeight(i)
		if nil != err {
			logger.Info("get get block error", "height", i)
			return 0, err
		}

		// got it
		if r.BlockMeta.Header.NumTxs != 0 {
			endHeight = i
			endTime = r.BlockMeta.Header.Time
			endTxs = r.BlockMeta.Header.TotalTxs
			break
		}
	}

	if 0 == endTxs {
		logger.Info("Test summary, not transaction found")
		return 0, nil
	}

	logger.Info("Test summary", "end block", endHeight, "time", endTime, "total txs", endTxs)
	if endTxs == beginTxs {
		logger.Info("beginTxs equal endTxs, test fails")
		return 0, errors.New("beginTxs equal endTxs, test fails")
	}
	logger.Info("Test summary", "txs", float64(endTxs-beginTxs), "second", endTime.Sub(beginTime).Seconds())
	tps := float64(endTxs-beginTxs) / endTime.Sub(beginTime).Seconds()
	tps, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tps), 64)
	return tps, nil
}

func revert() error {
	// get all account
	keys, _ := account.GetAllKeys()
	if len(keys) == 0 {
		return errors.New("accounts is nil")
	}
	logger.Info("Revert", "total keys", len(keys))

	// revert
	doRevert(keys)

	logger.Info("revert Done")
	return nil
}

func doRevert(keys map[string][]byte) error {
	// get chaincode
	client := rpcclient.InitClient(bbmConfig.NodeAddr, logger)
	token := common.GetGenesisToken(&client)
	chainID := common.GetChainID(&client)

	var wg sync.WaitGroup
	var totalBalance bn2.Number = bn2.N(0)
	var mu sync.Mutex
	var i int
	var result []string
	step := 10
	for id, key := range keys {
		wg.Add(1)
		go transferRevert(&client, chainID, key, token.Address, token.Owner, &wg, &mu, &totalBalance, id, &result)
		// linux system have 1024 open files limit
		if (i+1)%step == 0 {
			begin := time.Now().Unix()
			wg.Wait()
			end := time.Now().Unix()
			step = int(getStep(int64(step), end-begin))
			logger.Info("Revert ", "totalAccounts", len(keys), "current", i)
		}
		i++
	}
	wg.Wait()

	// clear db
	for _, v := range result {
		err := account.ClearKey(v)
		if err != nil {
			logger.Debug("DB clear key error", err)
			return err
		}
	}
	time.Sleep(time.Second)
	logger.Info("Revert ", "total balance", totalBalance)
	return nil
}

func transferRevert(client *rpcclient.RPCClient, chainID string, from []byte, tokenAddr, genesisAddr types.Address, wg *sync.WaitGroup, mu *sync.Mutex, totalBalance *bn2.Number, id string, result *[]string) {
	defer wg.Done()

	fromAddr := crypto.PrivKeyEd25519FromBytes(from).PubKey().Address(chainID)
	balance := common.BalanceOf(client, fromAddr, tokenAddr)

	if !balance.IsGreaterThanI(fee) {
		*result = append(*result, id)
		logger.Debug("account balance not enough for pay fee, continue", "balance", balance)
		return
	}

	privKey := "0x" + string(fmt.Sprintf("%x", from))
	nonce := common.AccountNonce(client, fromAddr) + 1

	var tx string
	// tx1 and tx2
	if 2 == bbmConfig.txVersion {
		tx2.Init(chainID)
		params := make([]interface{}, 0)
		params = append(params, genesisAddr)
		params = append(params, balance.SubI(fee))
		tx = common.GenerateTx(tokenAddr, protoTransfer, params, nonce, 500000, "", privKey)
	} else if 1 == bbmConfig.txVersion {
		var params []string

		params = append(params, genesisAddr, common.UintToHex(balance.SubI(fee).V.Uint64()))
		tx = common.GenerateTx1(nonce, 500000, "", tokenAddr, "Transfer(smc.Address,big.Int)smc.Error", params, chainID, fmt.Sprintf("%x", from))
	}

	response := client.SendTx([]byte(tx))
	if response.Code != 200 {
		logger.Info("transfer response", "code", response.Code)
		return
	}

	*result = append(*result, id)
	mu.Lock()
	*totalBalance = totalBalance.Add(balance.SubI(fee))
	mu.Unlock()
}
