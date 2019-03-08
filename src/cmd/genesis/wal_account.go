package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func findWalAccount(path string) (accounts []Account) {

	accounts = make([]Account, 0)

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		Error(err.Error())
	}
	for _, fi := range fis {
		if strings.HasSuffix(fi.Name(), ".json") {
			account := loadAccountFile(filepath.Join(path, fi.Name()))
			if account != nil {
				account.name = strings.TrimSuffix(fi.Name(), ".json")
				accounts = append(accounts, *account)
			}
		}
	}

	if len(accounts) == 0 {
		Error(fmt.Sprintf("Can not find any wal account in %v", path))
	}
	return
}

func loadAccountFile(file string) *Account {
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	acct := &Account{}
	err = cdc.UnmarshalJSON(jsonBytes, acct)
	if err != nil {
		return nil
	}

	return acct
}
