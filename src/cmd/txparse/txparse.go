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
		return TxParse(txStr, chainID)
	},
}

var (
	txStr   string
	chainID string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func txParseFlags() {
	RootCmd.PersistentFlags().StringVarP(&txStr, "tx", "t", "", "transaction string")
	RootCmd.PersistentFlags().StringVarP(&chainID, "chainID", "c", "", "block chain id")
}

func addFlags() {
	txParseFlags()
}

func TxParse(txStr, chainID string) error {
	if txStr == "" {
		return errors.New("Need transaction string")
	}
	tx2.Init(chainID)

	tx, pubKey, err := tx2.TxParse(txStr)
	if err != nil {
		Error(fmt.Sprintf("Parse transaction \"%v\" failed, %v", txStr, err.Error()))
		return err
	}

	fmt.Println("OK")
	fmt.Printf("Tx: %#v \n", tx)
	fmt.Println("PubKey: ", hex.EncodeToString(pubKey[:]))
	fmt.Println("Contract: ", tx.Messages[0].Contract)
	fmt.Printf("MethodID: 0x%x\n", tx.Messages[0].MethodID)
	fmt.Println("Items count: ", len(tx.Messages[0].Items))

	return nil
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
