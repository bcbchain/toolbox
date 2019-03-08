package main

import (
	"common/wal"
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
	Use:   "addr",
	Short: "Gen address",
	Long:  "Gen address by chainID and wallet name",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return addr(keyStorePath, walName, walPassword, chainID, file)
	},
}

var (
	// wallet flag
	keyStorePath string
	walName      string
	walPassword  string
	chainID      string

	//address file
	file string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addrFlags() {
	RootCmd.PersistentFlags().StringVarP(&keyStorePath, "keystore", "k", "", "path of key store")
	RootCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	RootCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	RootCmd.PersistentFlags().StringVarP(&chainID, "chainid", "i", "", "chainID of blockchain")
	RootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "file of address")
}

func addFlags() {
	addrFlags()
}

func addr(keyStorePath, walName, walPassword, chainID, file string) (err error) {

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if walName == "" {
		return errors.New("Need wallet name")
	}
	if chainID == "" {
		return errors.New("Need chainID")
	}

	acc, err := wal.LoadAccount(keyStorePath, walName, walPassword)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" failed, %v", walName, err.Error()))
		return
	}

	addr := acc.Address(chainID)
	if file == "" {
		fmt.Println("OK")
		fmt.Println("Address: ", addr)
		return nil
	} else {
		fi, e := os.Create(file)
		if e != nil {
			Error(fmt.Sprintf("Create file \"%v\" failed, %v", file, e.Error()))
		}
		addrFileStr := `{"addr":"` + addr + `"}`
		_, e = fi.Write([]byte(addrFileStr))
		if e != nil {
			Error(fmt.Sprintf("Write address to \"%v\" failed, %v", file, e.Error()))
		}
		fmt.Println("OK")
	}

	return nil
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
