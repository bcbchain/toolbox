package relay

import (
	"fmt"
	"github.com/bcbchain/sdk/sdk/ibc"
	"github.com/bcbchain/toolbox/relay/blockchain"
	"github.com/bcbchain/toolbox/relay/common"
	"github.com/bcbchain/toolbox/relay/relaydb"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var genesisOrgID = "orgJgaGConUyK81zibntUBjQ33PKctpk1K1G"

func StartCarrier() {
	chainIDs := relaydb.GetChainIDs()

	for _, chainID := range chainIDs {
		go Carry(chainID)
	}
}

func Carry(chainID string) {

	queueIDs := relaydb.GetChainQueueIDs(chainID)
	for _, queueID := range queueIDs {
		go carry(chainID, queueID)
	}
}

func carry(srcChainID, queueID string) {
	log := common.GetLogger()
	log.Debug(queueID + " carrier is running")

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

	// 开始搬运直到程序退出
	toChainID := chainIDFromQueueID(queueID)
	for loop {
		toChainInfo := relaydb.GetChainInfo(toChainID)
		if toChainInfo == nil {
			break
		}

		lastCarryHeight := relaydb.GetLastHeight(queueID)
		srcChainInfo := relaydb.GetChainInfo(srcChainID)
		pktsProofs := make([]ibc.PktsProof, 0)
		for lastCarryHeight < srcChainInfo.LastScanHeight {
			lastCarryHeight += 1
			pktsProof := relaydb.GetPktsProof(queueID, lastCarryHeight)
			if pktsProof == nil {
				continue
			}

			pktsProofs = append(pktsProofs, *pktsProof)
		}

		if len(pktsProofs) > 0 {
			//log.Debug(fmt.Sprintf("carrier pktsProofs: %v", pktsProofs))

			// 如果网络问题则重试!
			param := blockchain.IbcInputParam{
				PktsProofs: pktsProofs,
			}

			orgID := getOrgID(pktsProofs)
			result, err := blockchain.IbcInput(toChainID, orgID, param, toChainInfo.OpenUrls)
			log.Debug(fmt.Sprintf("IbcInput Result: %v", result))
			if err != nil {
				time.Sleep(5 * time.Second)
				continue
			}
		}

		relaydb.SetLastHeight(queueID, lastCarryHeight)

		// 1 秒轮询一次，查看是否有需要搬运的数据
		time.Sleep(1 * time.Second)
	}
}

func chainIDFromQueueID(queueID string) string {
	queueIDSplit := strings.Split(queueID, "->")

	return queueIDSplit[1]
}

func getOrgID(pktsProofs []ibc.PktsProof) string {
	for _, pktsProof := range pktsProofs {
		for _, packet := range pktsProof.Packets {
			if packet.OrgID != genesisOrgID {
				return packet.OrgID
			}
		}
	}

	return genesisOrgID
}
