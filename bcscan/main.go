package main

import (
	"errors"
	"fmt"
	"github.com/bcbchain/toolbox/bcscan/conf"
	"github.com/bcbchain/toolbox/bcscan/db"
	"github.com/bcbchain/toolbox/bcscan/log"
	"github.com/bcbchain/toolbox/bcscan/scan"
	"github.com/bcbchain/toolbox/bcscan/utils"
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
		return Scan(from, to, filter, output)
	},
}

var (
	from   int64
	to     int64
	filter string
	output string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func scanFlags() {
	RootCmd.PersistentFlags().Int64VarP(&from, "from", "f", 0, "begin height")
	RootCmd.PersistentFlags().Int64VarP(&to, "to", "t", 0, "end height")
	RootCmd.PersistentFlags().StringVarP(&filter, "filter", "l", "",
		"type of scan information, must be \"tx\", or \"header\" or \"tx,header\" or \"header,tx\"")
	RootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output db name")
}

func addFlags() {
	scanFlags()
}

func Scan(from, to int64, filter, output string) (err error) {
	if len(output) == 0 {
		return errors.New("output can not be empty")
	}

	filterList := make([]string, 0)
	if filterList, err = utils.ValidFilter(filter); err != nil {
		return err
	}

	if err = checkHeight(from, to); err != nil {
		return err
	}

	if err = conf.LoadBCSConfig(); err != nil {
		return err
	}

	c := conf.GetConfig()
	if c == nil {
		return errors.New("can not get config")
	}

	if err = db.InitDB(output, c.DBPort); err != nil {
		return err
	}
	defer db.Close()

	log.InitLogger(c)

	controller := scan.NewController(c, from, to, filterList)
	return controller.Start()
}

func checkHeight(from, to int64) error {
	if from < 0 {
		return errors.New("from must greater than zero")
	}

	if to < 0 {
		return errors.New("to must greater than zero")
	}

	if from > to {
		return errors.New("to must greater than from")
	}

	return nil
}
