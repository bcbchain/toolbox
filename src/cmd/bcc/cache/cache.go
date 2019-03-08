package cache

import (
	"blockchain/smcsdk/sdk/std"
	"common/jsoniter"
	"os"
	"path/filepath"
)

type accountCache struct {
	Nonce uint64 `json:"nonce"`
}

var (
	cacheFilePath = ""
)

func Init(path string) {
	cacheFilePath = path
}

func pathOfContract(orgID, contractName, keyStorePath string) string {
	return filepath.Join(keyStorePath, orgID+"_"+contractName+".toolCache")
}

func pathOfAccount(accountName, keyStorePath string) string {
	return filepath.Join(keyStorePath, accountName+".toolCache")
}

// Contract contract toolCache
func Contract(orgID, contractName, keyStorePath string) (contract *std.Contract, err error) {
	if keyStorePath == "" {
		keyStorePath = cacheFilePath
	}

	contractCachePath := pathOfContract(orgID, contractName, keyStorePath)

	f, err := os.Open(contractCachePath)
	if err != nil {
		return
	}
	defer f.Close()

	var b []byte
	_, err = f.Read(b)
	if err != nil {
		return
	}

	contract = new(std.Contract)
	err = jsoniter.Unmarshal(b, contract)

	return
}

// SetContract
func SetContract(contract *std.Contract, keyStorePath string) (err error) {

	if keyStorePath == "" {
		keyStorePath = cacheFilePath
	}

	contractCachePath := pathOfContract(contract.OrgID, contract.Name, keyStorePath)

	f, err := os.OpenFile(contractCachePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm|os.ModeAppend)
	if err != nil {
		return
	}
	defer f.Close()

	resBytes, err := jsoniter.Marshal(contract)
	if err != nil {
		return
	}

	_, err = f.Write(resBytes)

	return
}

// Nonce account toolCache
func Nonce(name, keyStorePath string) (nonce uint64, err error) {

	if keyStorePath == "" {
		keyStorePath = cacheFilePath
	}

	accountCachePath := pathOfAccount(name, keyStorePath)

	f, err := os.Open(accountCachePath)
	if err != nil {
		return
	}
	defer f.Close()

	var b []byte
	_, err = f.Read(b)
	if err != nil {
		return
	}

	var ac accountCache
	err = jsoniter.Unmarshal(b, &ac)
	if err != nil {
		return
	}

	return ac.Nonce, nil
}

// SetNonce
func SetNonce(name string, nonce uint64, keyStorePath string) (err error) {

	if keyStorePath == "" {
		keyStorePath = cacheFilePath
	}

	accountCachePath := pathOfAccount(name, keyStorePath)

	f, err := os.OpenFile(accountCachePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}

	var b []byte
	_, err = f.Read(b)
	if err != nil {
		return
	}

	var ac accountCache
	if len(b) > 0 {
		err = jsoniter.Unmarshal(b, &ac)
		if err != nil {
			return
		}
	}

	ac.Nonce = nonce
	resBytes, err := jsoniter.Marshal(ac)
	if err != nil {
		return
	}

	_, err = f.Write(resBytes)

	return
}
