package main

import (
	"blockchain/smccheck/parsecode"
	"blockchain/smcsdk/sdk/crypto/sha3"
	"blockchain/types"
	"common/fs"
	"common/wal"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tendermint/go-crypto"

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
	Use:   "smcpack",
	Short: "Pack smart contract",
	Long:  "Pack smart contract",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return packsmc(keyStorePath, walName, walPassword, smcPath)
	},
}

var (
	// wallet flag
	keyStorePath string
	walName      string
	walPassword  string

	// app flag
	smcPath string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addFlags() {
	RootCmd.PersistentFlags().StringVarP(&keyStorePath, "keystore", "k", "", "path of key store")
	RootCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	RootCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	RootCmd.PersistentFlags().StringVarP(&smcPath, "path", "s", "", "path of smart contract")
}

func packsmc(keyStorePath, walName, walPassword, smcPath string) (err error) {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if walName == "" {
		return errors.New("Need wallet name")
	}
	if smcPath == "" {
		return errors.New("Need path of smart contract")
	}

	acc, err := wal.LoadAccount(keyStorePath, walName, walPassword)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" failed, %v", walName, err.Error()))
		return
	}

	res, errOfCheck := parsecode.Check(smcPath)
	if errOfCheck.ErrorCode != types.CodeOK {
		Error(fmt.Sprintf("Parse contract code failed in \"%v\", %v", smcPath, errOfCheck.ErrorDesc))
		return
	}

	p := acc.PubKey().(crypto.PubKeyEd25519)
	if strings.ToLower(hex.EncodeToString(p[:])) != strings.ToLower(res.Author) {
		Error(fmt.Sprintf("Author pubKey don't match in \"%v\"", smcPath))
		return
	}

	contractTarName := res.ContractName + "-" + res.Version + ".tar.gz"
	tempDir, err := ioutil.TempDir(".", "temp")
	defer os.RemoveAll(tempDir)

	err = fs.CopyDir(smcPath, tempDir+"/", "(.go)$", "(_autogen_)")
	if err != nil {
		Error(fmt.Sprintf("Copy \"%v\" failed, %v", smcPath, err.Error()))
		return
	}

	err = fs.TarGz(tempDir, contractTarName, 0)
	if err != nil {
		Error(fmt.Sprintf("Tar \"%v\" failed, %v", smcPath, err.Error()))
		return
	}

	tarByte, err := ioutil.ReadFile(contractTarName)
	if err != nil {
		Error(fmt.Sprintf("Read tar.gz file \"%v\" failed, %v", contractTarName, err.Error()))
		return
	}

	err = acc.Sign2File(sha3.Sum256(tarByte), contractTarName+".sig")
	if err != nil {
		Error(fmt.Sprintf("Sign to \"%v\" failed, %v", contractTarName, err.Error()))
		return
	}

	fmt.Println("OK")
	fmt.Println("OutputFile: " + contractTarName)
	fmt.Println("            " + contractTarName + ".sig")
	return
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
