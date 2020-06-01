package main

import (
	cmn "github.com/bcbchain/bclib/tendermint/tmlibs/common"
	"github.com/bcbchain/toolbox/relay/common"
	"github.com/bcbchain/toolbox/relay/relay"
	"github.com/bcbchain/toolbox/relay/relaydb"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	err := Execute()
	if err != nil {
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "relay",
	Short: "Relay",
	Long:  "Relay",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return start(walName, walPassword, keyStorePath)
	},
}

var (
	// wallet flag
	keyStorePath string
	walName      string
	walPassword  string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addFlags() {
	RootCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	RootCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	RootCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", ".keystore", "path of key store")
}

func start(walName, walPassword, keyStorePath string) (err error) {
	err = common.InitAll(walName, walPassword, keyStorePath)
	if err != nil {
		panic(err)
	}

	//初始化数据库
	err = relaydb.InitDB()
	if err != nil {
		panic(err)
	}

	// 启动中继
	relay.Start()

	// Wait forever
	cmn.TrapSignal(func(signal os.Signal) {
	})
	return
}
