package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bcbchain/bcbchain/hyperledger/burrow/crypto"
	crypto2 "github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/bcbchain/bclib/wal"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func main() {
	err := Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "addr",
	Short: "Gen address",
	Long:  "Gen address by chainID and wallet name",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return addr(keyStorePath, walName, walPassword, chainID, file, bvmAddr, bvmConversion, bcbConversion)
	},
}

var (
	// wallet flag
	keyStorePath  string
	walName       string
	walPassword   string
	chainID       string
	bvmAddr       string
	bvmConversion string
	bcbConversion string

	//address file
	file string
)

func Execute() error {
	addFlags()
	return RootCmd.Execute()
}

func addrFlags() {
	RootCmd.PersistentFlags().StringVarP(&keyStorePath, "keystorepath", "k", "", "path of key store")
	RootCmd.PersistentFlags().StringVarP(&walName, "name", "n", "", "name of wallet")
	RootCmd.PersistentFlags().StringVarP(&walPassword, "password", "p", "", "password of wallet")
	RootCmd.PersistentFlags().StringVarP(&chainID, "chainid", "i", "", "chainID of blockchain")
	RootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "file of address")
	RootCmd.PersistentFlags().StringVarP(&bvmAddr, "bvmaddr", "d", "false", "whether to display bvm addr")
	RootCmd.PersistentFlags().StringVarP(&bvmConversion, "bvm2bcb", "b", "", "bvm addr to bcb")
	RootCmd.PersistentFlags().StringVarP(&bcbConversion, "bcb2bvm", "a", "", "bcb addr to bvm")
}

func addFlags() {
	addrFlags()
}

func addr(keyStorePath, walName, walPassword, chainID, file, bvmAddr, bvmConversion, bcbConversion string) (err error) {

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}

	if chainID == "" {
		return errors.New("Need chainID")
	}
	crypto2.SetChainId(chainID)

	if bvmConversion != "" {
		newBvmAddr := crypto.BVMAddress{}
		binCode, err := hex.DecodeString(bvmConversion)
		if err != nil {
			return err
		}
		copy(newBvmAddr[:], binCode[:])
		fmt.Println("bcbAddr: ", crypto.ToAddr(newBvmAddr))

		return nil
	}

	if bcbConversion != "" {
		fmt.Println("bvmAddr: ", strings.ToLower(crypto.ToBVM(bcbConversion).String()))
		return nil
	}

	acc, err := wal.LoadAccount(keyStorePath, walName, walPassword)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" failed, %v", walName, err.Error()))
		return
	}

	addr := acc.Address(chainID)
	if file == "" {
		if bvmAddr == "true" {
			fmt.Println("OK")
			fmt.Println("Address:    ", addr)
			fmt.Println("BvmAddress: ", crypto.ToBVM(addr))
			return nil
		}

		fmt.Println("OK")
		fmt.Println("Address: ", addr)
		return nil
	} else {
		fi, e := os.Create(file)
		if e != nil {
			Error(fmt.Sprintf("Create file \"%v\" failed, %v", file, e.Error()))
		}
		addrFileStr := `{"addr":"` + addr + `"}`
		_, e = fi.Write([]byte(addrFileStr))
		if e != nil {
			Error(fmt.Sprintf("Write address to \"%v\" failed, %v", file, e.Error()))
		}
		fmt.Println("OK")
	}

	return nil
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
