package account

import (
	"fmt"
	"github.com/bcbchain/bclib/bcdb"
	"github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/pkg/errors"
	"os"
	"strings"
)

const (
	dbPath = "/bbm"
)

func CreateAccounts(accountNum int) error {
	pwd, _ := os.Getwd()
	db, err := bcdb.OpenDB(pwd+dbPath, "", "")
	if err != nil {
		return errors.New("openDB error:" + err.Error())
	}
	defer db.Close()

	// generate ecc key
	for i := 0; i < accountNum; i++ {
		privKey := crypto.GenPrivKeyEd25519()
		key := fmt.Sprintf("%d", i)
		err = db.Set([]byte(key), privKey[:])
		if err != nil {
			return errors.New("DB Set error:" + err.Error())
		}
	}
	return nil
}

func GetKeys(begin, end int) (map[string][]byte, error) {
	pwd, _ := os.Getwd()
	db, err := bcdb.OpenDB(pwd+dbPath, "", "")
	if err != nil {
		return nil, errors.New("openDB error:" + err.Error())
	}
	defer db.Close()

	result := make(map[string][]byte)
	for i := begin; i <= end; i++ {
		key := fmt.Sprintf("%d", i)
		v, err := db.Get([]byte(key))
		if err != nil {
			if len(result) > 0 {
				return result, nil
			}
			return nil, errors.New("DB Get error:" + err.Error())
		}
		if 0 != len(v) {
			result[key] = v
		}
	}

	return result, nil
}

func GetAllKeys() (map[string][]byte, error) {
	pwd, _ := os.Getwd()
	db, err := bcdb.OpenDB(pwd+dbPath, "", "")
	if err != nil {
		return nil, errors.New("openDB error:" + err.Error())
	}
	defer db.Close()

	allKey := strings.Split(string(db.GetAllKey()), ";")
	result := make(map[string][]byte)
	for _, v := range allKey {
		data, err := db.Get([]byte(v))
		if err != nil {
			if len(result) > 0 {
				return result, nil
			}
			return nil, errors.New("DB Get error:" + err.Error())
		}
		if 0 != len(data) {
			result[v] = data
		}
	}

	return result, nil
}

func ClearDB() error {
	pwd, _ := os.Getwd()
	db, err := bcdb.OpenDB(pwd+dbPath, "", "")
	if err != nil {
		return errors.New("openDB error:" + err.Error())
	}
	defer db.Close()

	allKey := strings.Split(string(db.GetAllKey()), ";")
	for _, v := range allKey {
		err := db.Delete([]byte(v))
		if err != nil {
			return errors.New("DB Delete error:" + err.Error())
		}
	}

	return nil
}

func ClearKey(key string) error {
	pwd, _ := os.Getwd()
	db, err := bcdb.OpenDB(pwd+dbPath, "", "")
	if err != nil {
		return errors.New("openDB error:" + err.Error())
	}
	defer db.Close()

	err = db.Delete([]byte(key))
	if err != nil {
		return errors.New("DB Delete error:" + err.Error())
	}

	return nil
}
