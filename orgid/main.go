package main

import (
	"errors"
	"fmt"
	"github.com/bcbchain/sdk/sdkimpl/helper"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	err := Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "orgid",
	Short: "Calc organization ID",
	Long:  "Organization ID calculation tool",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return OrgId(orgName)
	},
}

var orgName string

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func orgIdFlags() {
	RootCmd.PersistentFlags().StringVarP(&orgName, "name", "n", "", "name of organization")
}

func addFlags() {
	orgIdFlags()
}

func OrgId(orgName string) error {
	if orgName == "" {
		return errors.New("Need organization name")
	}

	bcHelper := helper.NewBlockChainHelper(nil)

	fmt.Println("OK")
	fmt.Println("orgId: ", bcHelper.CalcOrgID(orgName))

	return nil
}
