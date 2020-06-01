package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	err := Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "genesis",
	Short: "genesis",
	Long:  "generate genesis files for blockchain",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(org) == 0 {
			org = "genesis"
		}
		return genesis(chainID, password, pathOfCharter, pathOfContracts, pathOfOutput, org)
	},
}

var (
	// command line flags
	chainID         string
	password        string
	pathOfCharter   string
	pathOfContracts string
	pathOfOutput    string
	org             string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addFlags() {
	RootCmd.PersistentFlags().StringVarP(&chainID, "chainid", "i", "", "chain id of blockChain")
	RootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password of owner's wallet")
	RootCmd.PersistentFlags().StringVarP(&pathOfCharter, "charter", "c", "", "path of charter")
	RootCmd.PersistentFlags().StringVarP(&pathOfContracts, "contracts", "t", "", "path of contracts")
	RootCmd.PersistentFlags().StringVarP(&pathOfOutput, "output", "o", "", "path of output")
	RootCmd.PersistentFlags().StringVarP(&org, "org", "r", "", "name of org")
}
