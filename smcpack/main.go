package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bcbchain/bcbchain/smccheck/parsecode"
	"github.com/bcbchain/bclib/fs"
	"github.com/bcbchain/bclib/types"
	"github.com/bcbchain/bclib/wal"
	"github.com/bcbchain/sdk/sdk/crypto/sha3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bcbchain/bclib/tendermint/go-crypto"

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
		return packsmc(keyStorePath, walName, walPassword, smcPath, output)
	},
}

var (
	// wallet flag
	keyStorePath string
	walName      string
	walPassword  string

	// app flag
	smcPath string
	output  string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addFlags() {
	RootCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	RootCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	RootCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	RootCmd.PersistentFlags().StringVarP(&smcPath, "path", "s", "", "path of smart contract")
	RootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "path of output")
}

func packsmc(keyStorePath, walName, walPassword, smcPath, output string) (err error) {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if walName == "" {
		return errors.New("Need wallet name")
	}
	if smcPath == "" {
		return errors.New("Need path of smart contract")
	}
	if output == "" {
		errors.New("Need path of output ")
	}

	acc, err := wal.LoadAccount(keyStorePath, walName, walPassword)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" failed, %v", walName, err.Error()))
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("No such file or directory", err)
		}
		return
	}()

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

	contractTarName := filepath.Join(output, res.ContractName+"-"+res.Version+".tar.gz")
	tempDir, err := ioutil.TempDir(".", "temp")
	defer os.RemoveAll(tempDir)

	smcPath, err = filepath.Abs(smcPath)
	if err != nil {
		Error(fmt.Sprintf("Invalid smcPath \"%v\", %v", smcPath, err.Error()))
		return
	}

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

	codeHash := sha3.Sum256(tarByte)
	err = acc.Sign2File(codeHash, contractTarName+".sig")
	if err != nil {
		Error(fmt.Sprintf("Sign to \"%v\" failed, %v", contractTarName, err.Error()))
		return
	}

	contractHashFile := filepath.Join(output, res.ContractName+"-"+res.Version+".hash")
	fi, err := os.Create(contractHashFile)
	defer fi.Close()
	if err != nil {
		Error(fmt.Sprintf("Create file \"%v\" failed, %v", contractHashFile, err.Error()))
		return
	}

	_, err = fi.Write([]byte(hex.EncodeToString(codeHash)))
	if err != nil {
		Error(fmt.Sprintf("Write code hash to \"%v\" failed, %v", contractHashFile, err.Error()))
		return
	}

	fmt.Println("OK")
	fmt.Println("OutputFile: " + contractTarName)
	fmt.Println("            " + contractTarName + ".sig")
	fmt.Println("            " + contractHashFile)
	return
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
