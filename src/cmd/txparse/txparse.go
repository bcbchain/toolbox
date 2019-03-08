package main

import (
	"blockchain/tx2"
	"encoding/hex"
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
	Use:   "txparse",
	Short: "Parsing a transaction",
	Long:  "Parsing a transaction",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return TxParse(txStr)
	},
}

var txStr string

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func txParseFlags() {
	RootCmd.PersistentFlags().StringVarP(&txStr, "tx", "t", "", "transaction string")
}

func addFlags() {
	txParseFlags()
}

func TxParse(txStr string) error {
	if txStr == "" {
		return errors.New("Need transaction string")
	}

	tx, pubKey, err := tx2.TxParse(txStr)
	if err != nil {
		Error(fmt.Sprintf("Parse transaction \"%v\" failed, %v", txStr, err.Error()))
		return err
	}

	fmt.Println("OK")
	fmt.Printf("Tx: %#v \n", tx)
	fmt.Println("PubKey: ", hex.EncodeToString(pubKey[:]))

	return nil
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
