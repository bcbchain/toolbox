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
	Use:   "bcw",
	Short: "Blockchain wallet console",
	Long:  "Blockchain wallet console",
}

func Execute() error {
	addFlags()
	addCommand()
	return RootCmd.Execute()
}

var walletCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new wallet",
	Long:  "Create a new wallet",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return Create(keyStorePath, walName, walPassword)
	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the private key of wallet",
	Long:  "Export the private key of wallet",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return Export(keyStorePath, walName, walPassword)
	},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import the private key to a new wallet",
	Long:  "Import the private key to a new wallet",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return Import(keyStorePath, walName, walPassword, PrivateKey)
	},
}

var signFileCmd = &cobra.Command{
	Use:   "signfile",
	Short: "Sign a file",
	Long:  "Sign a file",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return SignFile(keyStorePath, walName, walPassword, File, Mode)
	},
}

var signDataCmd = &cobra.Command{
	Use:   "signdata",
	Short: "Sign raw data",
	Long:  "Sign raw data",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return SignData(keyStorePath, walName, walPassword, Data, File)
	},
}

var (
	// wallet flag//
	keyStorePath string
	walName      string
	walPassword  string
	File         string
	PrivateKey   string
	Data         string
	Mode         string
	Path         string
)

func addCreateFlags() {
	walletCreateCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	walletCreateCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	walletCreateCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
}

func addExportFlags() {
	exportCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	exportCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	exportCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
}

func addImportFlags() {
	importCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	importCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	importCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	importCmd.PersistentFlags().StringVarP(&PrivateKey, "privkey", "y", "", "hex plaintext of private key")
}

func addSignFileFlags() {
	signFileCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	signFileCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	signFileCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	signFileCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "name of signature file")
	signFileCmd.PersistentFlags().StringVarP(&Mode, "mode", "m", "", "mode of Signature,\"b\" is binary file,\"t\" is text file")

}

func addSignDataFlags() {
	signDataCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	signDataCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	signDataCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	signDataCmd.PersistentFlags().StringVarP(&Data, "data", "d", "", "hex data to be signed")
	signDataCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "name of signature file")

}

func addFlags() {
	addCreateFlags()
	addExportFlags()
	addImportFlags()
	addSignFileFlags()
	addSignDataFlags()
}

func addCommand() {
	RootCmd.AddCommand(walletCreateCmd)
	RootCmd.AddCommand(exportCmd)
	RootCmd.AddCommand(importCmd)
	RootCmd.AddCommand(signFileCmd)
	RootCmd.AddCommand(signDataCmd)
}
