package relay

import (
	"errors"
	"fmt"
	"github.com/bcbchain/bclib/tendermint/abci/types"
	"github.com/bcbchain/sdk/sdk/ibc"
	"github.com/bcbchain/sdk/sdk/jsoniter"
	"github.com/bcbchain/sdk/sdk/std"
	core_types "github.com/bcbchain/tendermint/rpc/core/types"
	"github.com/bcbchain/toolbox/relay/common"
	"github.com/bcbchain/toolbox/relay/relaydb"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func StartScanner() {
	log := common.GetLogger()
	log.Info("Start Scan Block! ")

	chainIDs := relaydb.GetChainIDs()
	log.Info(fmt.Sprintf("GetChainIDs : %v", chainIDs))

	// 遍历chainInfo,在所有需要搬运的侧链的urls上扫块
	for _, chainID := range chainIDs {
		go scanBlock(chainID)
	}
}

func scanBlock(chainID string) {

	log := common.GetLogger()
	log.Debug(chainID + " scanner is running")

	// 等待程序退出
	loop := true
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Printf("captured %v, exiting...\n", sig)
			loop = false
		}
	}()

	// 轮询直到程序退出
	chainInfo := relaydb.GetChainInfo(chainID)
	for loop {
		// 更新LastScanHeight
		chainInfo.LastScanHeight = chainInfo.LastScanHeight + 1

		// 根据数据库中存储的链信息去对应的urls上扫块
	BRS:
		blkResults := getBlockResults(chainInfo.OpenUrls, chainInfo.LastScanHeight)
		if blkResults == nil {
			time.Sleep(5 * time.Second)
			goto BRS
		}

		// updateOpenURLs
		resetOpenURLs(blkResults.Results.DeliverTx)

		mapForPackets := getMapPackets(blkResults.Results.DeliverTx)
		// 如果有ibc交易
		if len(mapForPackets) != 0 {
		PF:
			p, err := getProof(chainInfo.OpenUrls, chainInfo.LastScanHeight)
			if err != nil {
				//return err
				time.Sleep(time.Second * 5)
				goto PF
			}

			for queueID, packets := range mapForPackets {
				log.Debug(fmt.Sprintf("scanner queueID: %v", queueID))
				var pktsProof ibc.PktsProof
				pktsProof.Packets = packets
				pktsProof.Precommits = p.Precommits
				pktsProof.Header = p.Header

				relaydb.SetPktsProof(queueID, chainInfo.LastScanHeight, pktsProof)

				if relaydb.AddChainQueueID(chainID, queueID) {
					relaydb.SetLastHeight(queueID, chainInfo.LastScanHeight-1)

					go carry(chainID, queueID)
				}
			}
		}

		relaydb.SetChainInfo(chainInfo)
	}
}

func getProof(urls []string, scanHeight int64) (pktsProof ibc.PktsProof, err error) {

	//获取header
	headerHeight := scanHeight + 1
	result1, err := block(urls, headerHeight)
	if err != nil {
		return
	}

	precommitHeight := headerHeight + 1
	result2, err := block(urls, precommitHeight)
	if err != nil {
		return
	}

	headerByte, err := jsoniter.Marshal(result1.BlockMeta.Header)
	if err != nil {
		return
	}

	var header ibc.Header
	err = jsoniter.Unmarshal(headerByte, &header)

	pktsProof.Header = header

	preCommitBytes, err := jsoniter.Marshal(result2.Block.LastCommit.Precommits)
	var preCommits []ibc.Precommit
	err = jsoniter.Unmarshal(preCommitBytes, &preCommits)
	if err != nil {
		return
	}
	pktsProof.Precommits = preCommits

	return
}

func getBlockResults(urls []string, height int64) (blkResults *core_types.ResultBlockResults) {

	blkResults = new(core_types.ResultBlockResults)
	err := common.DoHttpRequestAndParse(urls, "block_results", map[string]interface{}{"height": height}, blkResults)
	if err != nil {
		//if err.Error() == "Height must be less than or equal to the current blockchain height" {
		//	return nil
		//}
		return nil
		//panic(err)
	}

	return
}

func getMapPackets(deliverTxs []*types.ResponseDeliverTx) map[string][]ibc.Packet {
	mapForPackets := make(map[string][]ibc.Packet)
	for _, deliverTx := range deliverTxs {
		packets := getPackets(deliverTx)
		for _, packet := range packets {
			var queueIDPackets []ibc.Packet
			var ok bool
			if queueIDPackets, ok = mapForPackets[packet.QueueID]; !ok {
				queueIDPackets = make([]ibc.Packet, 0)
			}
			queueIDPackets = append(queueIDPackets, packet)
			mapForPackets[packet.QueueID] = queueIDPackets
		}
	}

	return mapForPackets
}

// getPackets get ibc::packet receipt from deliver tags.
// deliverTx only have one ibc::packet receipt in default
func getPackets(deliverTx *types.ResponseDeliverTx) []ibc.Packet {
	if deliverTx.Code != types.CodeTypeOK {
		return nil
	}

	packets := make([]ibc.Packet, 0)
	var err error
	for _, tag := range deliverTx.Tags {
		if strings.Contains(string(tag.Key), "/ibc::packet/") {
			var receipt std.Receipt
			err = jsoniter.Unmarshal(tag.Value, &receipt)
			if err != nil {
				panic(err)
			}

			var packet ibc.Packet
			err = jsoniter.Unmarshal(receipt.Bytes, &packet)
			if err != nil {
				panic(err)
			}

			packets = append(packets, packet)
		}
	}

	return packets
}

func resetOpenURLs(deliverTxs []*types.ResponseDeliverTx) {

	for _, deliverTx := range deliverTxs {
		if deliverTx.Code != types.CodeTypeOK {
			return
		}

		var err error
		for _, tag := range deliverTx.Tags {
			if strings.HasSuffix(string(tag.Key), "/netgovernance.setOpenURL") {
				var receipt std.Receipt
				err = jsoniter.Unmarshal(tag.Value, &receipt)
				if err != nil {
					panic(err)
				}

				newOpenURLs := new(relaydb.SetOpenURL)
				err = jsoniter.Unmarshal(receipt.Bytes, newOpenURLs)
				if err != nil {
					panic(err)
				}

				chainIDs := relaydb.GetChainIDs()
				for _, v := range chainIDs {
					if v == newOpenURLs.SideChainID {
						oldChainInfo := relaydb.GetChainInfo(v)
						newChainInfo := new(relaydb.ChainInfo)
						newChainInfo.ChainID = v
						newChainInfo.OpenUrls = newOpenURLs.OpenURLs
						newChainInfo.LastScanHeight = oldChainInfo.LastScanHeight
						relaydb.SetChainInfo(newChainInfo)
					}
				}
			}
		}
	}
}

func block(urls []string, scanHeight int64) (result *core_types.ResultBlock, err error) {

	result = new(core_types.ResultBlock)
	params := map[string]interface{}{"height": scanHeight}
	err = common.DoHttpRequestAndParse(urls, "block", params, result)
	if err != nil {
		return nil, errors.New("GetBlockResults fail, err info : " + err.Error())
	}

	return
}
