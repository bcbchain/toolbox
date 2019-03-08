package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"cmd/walv1tov2/walv1"
	"common/wal"

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
	Use:   "walv1towalv2",
	Short: "walv1towalv2",
	Long:  "convert wallet file from version 1 to version 2",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateWal(keyStorePath, walName, walV1Password, walV2Password)
	},
}

var (
	// wallet flag
	keyStorePath  string
	walName       string
	walV1Password string
	walV2Password string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addFlags() {
	RootCmd.PersistentFlags().StringVarP(&keyStorePath, "keystore", "k", "", "path of key store")
	RootCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	RootCmd.PersistentFlags().StringVarP(&walV1Password, "password", "p", "", "password of wallet version 1")
	RootCmd.PersistentFlags().StringVarP(&walV2Password, "password_new", "q", "", "password of wallet version 2")
}

func updateWal(keyStorePath, walName, walV1Password, walV2Password string) error {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if walName == "" {
		return errors.New("Need wallet name")
	}
	if walV1Password == "" {
		return errors.New("Need old password to decrypto wallet")
	}
	if walV2Password == "" {
		return errors.New("Need new password to encrypto new version wallet")
	}

	//加载老板钱包文件
	acct1, err := walv1.LoadAccount(
		filepath.Join(keyStorePath, walName+".wal"),
		[]byte(walV1Password),
		nil,
	)
	if err != nil {
		Error(err.Error())
	}

	_, err = wal.ImportAccount(keyStorePath, walName, walV2Password, acct1.PrivKey)
	if err != nil {
		Error(err.Error())
	}

	fmt.Println("OK")
	return nil
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
