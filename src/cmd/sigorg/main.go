package main

import (
	"blockchain/smcsdk/sdk/crypto/sha3"
	"common/sig"
	"common/wal"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

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
	Use:   "sigorg",
	Short: "Sign a smc by organization",
	Long:  "Sign a smart contract by organization",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return sigOrg(keyStorePath, walName, walPassword, smcPkgFile)
	},
}

var (
	// wallet flag
	keyStorePath string
	walName      string
	walPassword  string

	//
	smcPkgFile string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addFlags() {
	RootCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	RootCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	RootCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	RootCmd.PersistentFlags().StringVarP(&smcPkgFile, "smc", "s", "", "file name of smart contract package")
}

func sigOrg(keyStorePath, walName, walPassword, smcPkgFile string) (err error) {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if walName == "" {
		return errors.New("Need wallet name")
	}
	if smcPkgFile == "" {
		return errors.New("Need file name of smart contract package")
	}

	if !strings.HasSuffix(smcPkgFile, ".tar.gz") {
		Error(fmt.Sprintf("Read contract tar.gz file \"%v\" failed", smcPkgFile))
		return
	}

	codeByte, err := ioutil.ReadFile(smcPkgFile)
	if err != nil {
		Error(fmt.Sprintf("Read file \"%v\" failed, %v", smcPkgFile, err.Error()))
		return
	}

	codeSigFile := path.Base(smcPkgFile) + ".sig"
	codeSigByte, err := ioutil.ReadFile(codeSigFile)
	if err != nil {
		Error(fmt.Sprintf("Read signature file \"%v\" failed, %v", codeSigFile, err.Error()))
		return
	}

	type signature struct {
		PubKey    string `json:"pubkey"`
		Signature string `json:"signature"`
	}
	// parse developer signature
	s := new(signature)
	err = json.Unmarshal(codeSigByte, s)
	if err != nil {
		Error(fmt.Sprintf("Unmarshal \"%v\" failed, %v", codeSigFile, err.Error()))
		return
	}

	pk, err := hex.DecodeString(s.PubKey)
	if err != nil {
		Error(fmt.Sprintf("Invalid pubKey in \"%v\" failed, %v", codeSigFile, err.Error()))
		return
	}

	sigByte, err := hex.DecodeString(s.Signature)
	if err != nil {
		Error(fmt.Sprintf("Invalid signature in \"%v\" failed, %v", codeSigFile, err.Error()))
		return
	}

	// verify signature for contract package's hash
	ok, err := sig.Verify(pk, sha3.Sum256(codeByte), sigByte)
	if err != nil {
		Error(fmt.Sprintf("Verify sign \"%v\" failed, %v", codeSigFile, err.Error()))
		return
	}

	if !ok {
		Error(fmt.Sprintf("Verify sign \"%v\" failed, %v", codeSigFile, errors.New("")))
		return
	}

	// load organization account
	acc, err := wal.LoadAccount(keyStorePath, walName, walPassword)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" failed, %v", walName, err.Error()))
		return
	}

	// organization account sign for developer's signature data, not file.
	err = acc.Sign2File(sigByte, codeSigFile+".sig")
	if err != nil {
		Error(fmt.Sprintf("Sign for \"%v\" failed, %v", smcPkgFile, err.Error()))
		return
	}

	fmt.Println("OK")
	fmt.Println("SignatureFile: " + codeSigFile + ".sig")
	return
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
