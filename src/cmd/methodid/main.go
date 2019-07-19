package main

import (
	"blockchain/algorithm"
	"errors"
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
	Use:   "methodid",
	Short: "Calc method ID",
	Long:  "Method ID calculation tool",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return MethodId(methodProtoType)
	},
}

var methodProtoType string

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func methodIdFlags() {
	RootCmd.PersistentFlags().StringVarP(&methodProtoType, "method", "m", "", "prototype of method")
}

func addFlags() {
	methodIdFlags()
}

func MethodId(methodProtoType string) error {
	if methodProtoType == "" {
		return errors.New("Need method prototype")
	}
	methodID := algorithm.ConvertMethodID(algorithm.CalcMethodId(methodProtoType))

	fmt.Println("OK")
	fmt.Println("methodId: ", methodID)

	return nil
}
