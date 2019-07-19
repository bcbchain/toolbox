package main

import (
	"blockchain/smccheck/parsecode"
	"blockchain/types"
	"errors"
	"fmt"
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
	Use:   "smccheck",
	Short: "check smart contract format and content",
	Long:  "check smart contract format and content",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return smcCheck(smcPath)
	},
}

var (
	// app flag
	smcPath string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addFlags() {
	RootCmd.PersistentFlags().StringVarP(&smcPath, "path", "s", "", "path of smart contract")
}

func smcCheck(smcPath string) (err error) {
	if smcPath == "" {
		return errors.New("Need path of smart contract")
	}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		return
	}()
	_, errOfCheck := parsecode.Check(smcPath)
	if errOfCheck.ErrorCode != types.CodeOK {
		//Error(fmt.Sprintf("Check contract code failed in \"%v\".\n%v", smcPath, errOfCheck.ErrorDesc))
		return
	} else {
		fmt.Printf("")
		return
	}

}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
