package main

import (
	"common/wal"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"github.com/tendermint/go-crypto"
)

func Create(keyStorePath, name, password string) error {

	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if name == "" {
		return errors.New("Need wallet name ")
	}

	acct, err := wal.NewAccount(keyStorePath, name, password)
	if err != nil {
		Error(fmt.Sprintf("New account \"%v\" failed, %v", name, err.Error()))
		return err
	}

	PubK := acct.PubKey().(crypto.PubKeyEd25519)

	fmt.Println("OK")
	fmt.Println("PubKey: ", hex.EncodeToString(PubK[:]))

	return nil
}

func Export(keyStorePath, name, password string) error {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if name == "" {
		return errors.New("Need wallet name ")
	}

	acct, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" failed, %v", name, err.Error()))
		return err
	}

	PriK := acct.PrivateKey.(crypto.PrivKeyEd25519)
	PubK := acct.PubKey().(crypto.PubKeyEd25519)

	fmt.Println("OK")
	fmt.Println("PrivateKey: ", hex.EncodeToString(PriK[:]))
	fmt.Println("PubKey:     ", hex.EncodeToString(PubK[:]))

	return nil
}

func Import(keyStorePath, name, password, privateKey string) error {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if name == "" {
		return errors.New("Need wallet name ")
	}
	if privateKey == "" {
		return errors.New("Need private Key of account ")
	}

	newPrivateKey, err := hex.DecodeString(privateKey)
	if err != nil {
		Error(fmt.Sprintf("Private Key conversion \"%v\" failed, %v", name, err.Error()))
	}
	if len(newPrivateKey) != 64 {
		return errors.New(fmt.Sprintf("Private key \"%v\" length incorrect, %v", privateKey, err.Error()))
	}

	acct, err := wal.ImportAccount(keyStorePath, name, password, crypto.PrivKeyEd25519FromBytes(newPrivateKey))
	if err != nil {
		Error(fmt.Sprintf("Import account \"%v\" failed, %v", name, err.Error()))
	}

	PubK := acct.PubKey().(crypto.PubKeyEd25519)

	fmt.Println("OK")
	fmt.Println("PubKey: ", hex.EncodeToString(PubK[:]))

	return nil
}

func SignFile(keyStorePath, name, password, file, mode string) error {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if name == "" {
		return errors.New("Need wallet name ")
	}
	if file == "" {
		return errors.New("Need signed file ")
	}
	if mode == "" {
		return errors.New("Need mode of file ")
	}

	acct, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" failed, %v", name, err.Error()))
		return err
	}

	var SignatureFile string
	if mode == "b" {
		acct.SignBinFile(file, SignatureFile)

	} else if mode == "t" {
		acct.SignTextFile(file, SignatureFile)

	} else {
		return errors.New(fmt.Sprintf("Invalid mode to \"%v\", %v", name, err.Error()))
	}

	fmt.Println("OK")
	fmt.Println("SignatureFile: ", SignatureFile)

	return err
}

func SignData(keyStorePath, name, password, data string, file string) error {
	if keyStorePath == "" {
		keyStorePath = ".keystore"
	}
	if name == "" {
		return errors.New("Need wallet name ")
	}
	if data == "" {
		return errors.New("Need signed data ")
	}
	if file == "" {
		return errors.New("Need file to be output after signing ")
	}

	newData, err := hex.DecodeString(data)
	if err != nil {
		Error(fmt.Sprintf("Decode hex data failed, %v", err.Error()))
		return err
	}

	acct, err := wal.LoadAccount(keyStorePath, name, password)
	if err != nil {
		Error(fmt.Sprintf("Load account \"%v\" in %v failed, %v", name, keyStorePath, err.Error()))
		return err
	}

	err = acct.Sign2File(newData, file)
	if err != nil {
		Error(fmt.Sprintf("Sign to \"%v\" failed, %v", name, err.Error()))
		return err
	}

	fmt.Println("OK")
	fmt.Println("SignatureFile: ", file)

	return nil
}

func Error(s string) {
	fmt.Printf("ERROR! -- %v\n", s)
	os.Exit(1)
}
