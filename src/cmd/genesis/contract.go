package main

import (
	"common/fs"
	"common/jsoniter"
	"common/sig"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"blockchain/smcsdk/sdk/crypto/sha3"
)

func findContract(path string) (contractInfos []contractInfo, releaseVersion string) {

	contractInfos = make([]contractInfo, 0)

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		Error(err.Error())
	}
	for _, fi := range fis {
		contractInfo := contractInfo{}
		if strings.HasSuffix(fi.Name(), ".tar.gz") {

			contractInfo.Code = fi.Name()
			nameAndVersion := strings.Replace(fi.Name(), ".tar.gz", "", 1)
			l := strings.Split(nameAndVersion, "-")
			version := l[len(l)-1]
			codeBytes, err := ioutil.ReadFile(filepath.Join(path, fi.Name()))
			if err != nil {
				Error(err.Error())
			}

			codeHash := sha3.Sum256(codeBytes)
			contractInfo.CodeHash = hex.EncodeToString(codeHash)
			contractInfo.Name = strings.Replace(nameAndVersion, "-"+version, "", 1)
			contractInfo.Version = version
			contractInfo.code = codeBytes

			devSigFile := filepath.Join(path, fi.Name()+".sig")
			orgSigFile := filepath.Join(path, fi.Name()+".sig.sig")
			if ok, _ := fs.PathExists(devSigFile); !ok {
				if strings.HasPrefix(fi.Name(), "genesis-smart-contract_") {
					if releaseVersion != "" {
						Error("Found too many genesis smart contract release package")
					}
					releaseVersion = strings.TrimPrefix(fi.Name(), "genesis-smart-contract_")
					releaseVersion = strings.TrimSuffix(releaseVersion, ".tar.gz")
				}
				continue
			}

			if !verifyDevSign(devSigFile, codeHash) {
				Error("verifySign failed for " + devSigFile)
			}

			if !verifyOrgSign(devSigFile, orgSigFile) {
				Error("verifySign failed for " + orgSigFile)
			}

			contractInfo.CodeDevSig = getSignature(devSigFile)
			contractInfo.CodeOrgSig = getSignature(orgSigFile)
			contractInfos = append(contractInfos, contractInfo)
		}
	}

	if len(contractInfos) == 0 {
		Error("Can not find any genesis smart contract")
	}
	if releaseVersion == "" {
		Error("Can not find any genesis smart contract release package")
	}
	return
}

func verifyDevSign(devSigPath string, data []byte) bool {
	sigBytes, err := ioutil.ReadFile(devSigPath)
	if err != nil {
		Error(err.Error())
	}

	codeSigna := new(codeSign)
	err = jsoniter.Unmarshal(sigBytes, codeSigna)
	if err != nil {
		Error(err.Error())
	}

	pk, err := hex.DecodeString(codeSigna.PubKey)
	if err != nil {
		Error(err.Error())
	}

	s, err := hex.DecodeString(codeSigna.Signature)
	if err != nil {
		Error(err.Error())
	}

	ok, err := sig.Verify(pk, data, s)
	if err != nil {
		Error(fmt.Sprintf("Verify signature for %v failed: %v", devSigPath, err.Error()))
	}

	return ok
}

func verifyOrgSign(devSigPath, orgSigPath string) bool {
	devSigBytes, err := ioutil.ReadFile(devSigPath)
	if err != nil {
		Error(err.Error())
	}

	codeDevSigna := new(codeSign)
	err = jsoniter.Unmarshal(devSigBytes, codeDevSigna)
	if err != nil {
		Error(err.Error())
	}

	devSig, err := hex.DecodeString(codeDevSigna.Signature)
	if err != nil {
		Error(err.Error())
	}

	orgSigBytes, err := ioutil.ReadFile(orgSigPath)
	if err != nil {
		Error(err.Error())
	}

	codeOrgSigna := new(codeSign)
	err = jsoniter.Unmarshal(orgSigBytes, codeOrgSigna)
	if err != nil {
		Error(err.Error())
	}

	orgPk, err := hex.DecodeString(codeOrgSigna.PubKey)
	if err != nil {
		Error(err.Error())
	}

	orgSig, err := hex.DecodeString(codeOrgSigna.Signature)
	if err != nil {
		Error(err.Error())
	}

	ok, err := sig.Verify(orgPk, devSig, orgSig)
	if err != nil {
		Error(fmt.Sprintf("Verify signature for %v failed: %v", orgSigPath, err.Error()))
	}

	return ok
}

func getSignature(sigPath string) codeSign {
	sigBytes, err := ioutil.ReadFile(sigPath)
	if err != nil {
		Error(err.Error())
	}

	cs := codeSign{}
	err = jsoniter.Unmarshal(sigBytes, &cs)
	if err != nil {
		Error(err.Error())
	}

	return cs
}
