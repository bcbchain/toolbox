package main

import (
	"fmt"
	"github.com/bcbchain/toolbox/bcparser/db"
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
	Use:   "bcscan",
	Short: "Scan block chain",
	Long:  "Scan the block chain and save header and transaction information",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return Parse(input)
	},
}

var (
	input string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func parseFlags() {
	RootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "scan db file")
}

func addFlags() {
	parseFlags()
}

func Parse(input string) error {
	if err := db.InitDB(input, ""); err != nil {
		return err
	}

	//fi, err := os.Create("aaabbb")
	//if err != nil {
	//	panic(err)
	//}
	//defer fi.Close()

	lastHeight := db.GetLastHeight()
	for i := int64(1); i <= lastHeight; i++ {
		header := db.GetHeader(i)
		fmt.Printf("Height: %d \nHeader: %v\n", i, header)
		//fi.WriteString(fmt.Sprintf("Height: %d \nHeader: %v\n", i, header))

		txs := db.GetTx(i)
		for _, v := range txs {
			fmt.Printf("Height: %d \nTx: %s\n", i, v)
			//fi.WriteString(fmt.Sprintf("Height: %d \nTx: %s\n", i, v))
		}
	}

	return nil
}
