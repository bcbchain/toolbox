package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/types"
)

// ProcessFollower - processing the follower :-) nonsense
func ProcessFollower(byzantium, proxyApp, aAddr string, listenPortN int) {
	persistentPeers := ""
	voters := strings.Split(byzantium, ",")
	for i, v := range voters {
		nodeID := getNodeID(v)
		if nodeID != "" {
			if i != 0 {
				persistentPeers += ","
			}
			persistentPeers += nodeID
		}
	}
	var genesisDoc *types.GenesisDoc
	for _, v := range voters {
		genesisDoc = getGenesis(v)
		if genesisDoc != nil {
			break
		}
	}

	conf := cfg.DefaultConfig()
	_ = viper.Unmarshal(conf) // nolint unhandled
	configFilePath := filepath.Join(conf.RootDir, "config", "config.toml")

	home := os.Getenv("TMHOME")
	if strings.HasPrefix(home, "/etc") {
		paths := strings.Split(home, "/")
		var myHome string
		if len(paths) > 2 {
			myHome = "/home/" + paths[2]
		} else {
			myHome = "/home/tmcore"
		}
		conf.DBPath = myHome + "/data"
		conf.LogPath = myHome + "/log"
		conf.Mempool.WalPath = myHome + "/data/mempool.wal"
		conf.Consensus.WalPath = myHome + "/data/cs.wal/wal"
	}
	conf.P2P.PersistentPeers = persistentPeers

	if proxyApp != "" {
		conf.ProxyApp = []string{proxyApp}
	}
	if aAddr != "" {
		conf.P2P.AAddress = aAddr
	}
	if listenPortN != 0 {
		conf.P2P.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", listenPortN)
		conf.RPC.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", listenPortN+1)
	}

	cfg.WriteConfigFile(configFilePath, conf)

	validatorsFile := conf.ValidatorsFile()
	outByte, err := cdc.MarshalJSONIndent(genesisDoc.Validators, "", "  ")
	if err != nil {
		fmt.Printf("last step,marshal genesisDoc err: %v\n", err)
		return
	}
	_ = ioutil.WriteFile(validatorsFile, outByte, 0600) // nolint unhandled
}

type nodeInfo struct {
	ID            string `json:"id"`
	AnnouncedAddr string `json:"listen_addr"`
}
type statusResult struct {
	NodeInfo nodeInfo `json:"node_info"`
}
type statusResponse struct {
	Result statusResult `json:"result"`
}

func getNodeID(node string) string {
	url := "https://" + node + "/status"
	url2 := "http://" + node + "/status"
	nodeID := getNodeIDFromURL(url)
	if nodeID == "" {
		return getNodeIDFromURL(url2)
	}
	return nodeID
}

func getNodeIDFromURL(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("get status from %s cause err1: %v\n", url, err)
		return ""
	}
	client := &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get status from %s cause err2: %v\n", url, err)
		return ""
	}
	defer func() { _ = resp.Body.Close() }() // nolint unhandled

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get status from %s cause err3: %v\n", url, err)
		return ""
	}
	// fmt.Println("response Body:", string(body))
	var nodeStatus statusResponse
	err = json.Unmarshal(body, &nodeStatus)
	if err != nil {
		fmt.Printf("get Status from %s parse err: %v\n", url, err)
		return ""
	}
	if nodeStatus.Result.NodeInfo.ID == "" {
		fmt.Printf("got node Id='' from %s\n", url)
		return ""
	}
	if nodeStatus.Result.NodeInfo.AnnouncedAddr == "" {
		fmt.Printf("bad listen address:(%v) from %s", nodeStatus.Result.NodeInfo.AnnouncedAddr, url)
		return ""
	}
	return nodeStatus.Result.NodeInfo.ID + "@" + nodeStatus.Result.NodeInfo.AnnouncedAddr
}

type genesisResult struct {
	Genesis types.GenesisDoc `json:"genesis"`
}
type genesisResponse struct {
	Result genesisResult `json:"result"`
}

func getGenesisFromURL(url string) *types.GenesisDoc {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("get genesis from %s cause err1: %v\n", url, err)
		return nil
	}
	client := &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get genesis from %s cause err2: %v\n", url, err)
		return nil
	}
	defer func() { _ = resp.Body.Close() }() // nolint unhandled

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get genesis from %s cause err3: %v\n", url, err)
		return nil
	}
	// fmt.Println("response Body:", string(body))
	var genesisResponse genesisResponse
	err = cdc.UnmarshalJSON(body, &genesisResponse)
	if err != nil {
		fmt.Printf("get genesis from %s parse err: %v\n", url, err)
		return nil
	}
	genesis := genesisResponse.Result.Genesis

	return &genesis
}
func getGenesis(node string) *types.GenesisDoc {
	url := "https://" + node + "/genesis"
	url2 := "http://" + node + "/genesis"
	genesisDoc := getGenesisFromURL(url)
	if genesisDoc != nil {
		return genesisDoc
	}
	return getGenesisFromURL(url2)
}
