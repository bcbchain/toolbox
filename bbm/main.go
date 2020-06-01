package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/bcbchain/bclib/tendermint/tmlibs/log"
	"os"
)

var (
	// bbm cmd flag//
	accounts string
	from     string
	to       string
	value    string
	round    string
)

const (
	configFilePath = "./.config/bbm.yaml"
)

func main() {
	// init evn
	err := InitEvn()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// exec cmd
	err = Execute()
	if err != nil {
		os.Exit(1)
	}
}

func InitEvn() error {
	//load config
	err := LoadBBMConfig(configFilePath)
	if err != nil {
		return err
	}

	//set log
	logger = log.NewTMLogger(bbmConfig.LoggerFilePath, "bbm")
	logger.SetOutputToScreen(bbmConfig.LoggerScreen)
	logger.SetOutputToFile(bbmConfig.LoggerFile)
	logger.AllowLevel(bbmConfig.LoggerLevel)
	//logger.Info("Start bbm program")

	//get bcbchain version
	err = InitTXVersion()
	if err != nil {
		return err
	}
	return nil
}

var RootCmd = &cobra.Command{
	Use:   "bbm",
	Short: "Blockchain benchmark tool",
	Long:  "Blockchain benchmark tool",
}

func Execute() error {
	addFlags()
	addCommand()
	return RootCmd.Execute()
}

var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Create accounts for test",
	Long:  "Create accounts for test",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := prepare(accounts, value)
		if nil != err {
			logger.Error("prepare", "Error", err)
		}
		return err
	},
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Bcbchain performance test in an sync manner",
	Long:  "Bcbchain performance test in an sync manner",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := syncTest(from, to, round)
		if nil != err {
			logger.Error("sync", "Error", err)
		}
		return err
	},
}

var asyncCmd = &cobra.Command{
	Use:   "async",
	Short: "Bcbchain performance test in an async manner",
	Long:  "Bcbchain performance test in an async manner",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := asyncTest(from, to)
		if nil != err {
			logger.Error("async", "Error", err)
		}
		return err
	},
}

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Recover funds from the test account",
	Long:  "Recover funds from the test account",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := revert()
		if nil != err {
			logger.Error("revert", "Error", err)
		}
		return err
	},
}

func prepareFlags() {
	prepareCmd.PersistentFlags().StringVarP(&accounts, "accounts", "a", "", "number of accounts created for testing")
	prepareCmd.PersistentFlags().StringVarP(&value, "value", "v", "", "amount per account")
}

func syncFlags() {
	syncCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "number of accounts transferred out")
	syncCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "number of accounts transferred to")
	syncCmd.PersistentFlags().StringVarP(&round, "round", "r", "", "number of test rounds")
}

func asyncFlags() {
	asyncCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "number of accounts transferred out")
	asyncCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "number of accounts transferred to")
}

func addFlags() {
	prepareFlags()
	syncFlags()
	asyncFlags()
}

func addCommand() {
	RootCmd.AddCommand(prepareCmd)
	RootCmd.AddCommand(syncCmd)
	RootCmd.AddCommand(asyncCmd)
	RootCmd.AddCommand(revertCmd)
}
