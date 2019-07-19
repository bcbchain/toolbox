package walv1

import (
	common2 "cmd/bcc/common"
	"errors"
	"fmt"
	"io/ioutil"
	conv "strconv"

	"blockchain/algorithm"
	"common/fs"

	"github.com/tendermint/go-amino"
	crypto "github.com/tendermint/go-crypto"
	cmn "github.com/tendermint/tmlibs/common"
)

type Address = string

var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}

type Account struct {
	Name         string         `json:"name"`
	PrivKey      crypto.PrivKey `json:"privKey"`
	PubKey       crypto.PubKey  `json:"pubKey"`
	Address      Address        `json:"address"`
	Nonce        uint64         `json:"nonce"`
	KeystorePath string         `json:"keystore"`
}

func NewAccount(name string, keystoreDir string) (*Account, error) {
	cfg := common2.GetBCCConfig()
	var keystorePath string
	if keystoreDir != "" {
		keystorePath = keystoreDir + "/" + name + ".wal"
		exists, err := fs.PathExists(keystorePath)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("The account of " + name + " is already exist!")
		}
	}

	privKey := crypto.GenPrivKeyEd25519()
	pubKey := privKey.PubKey()
	address := pubKey.Address(cfg.DefaultChainID)

	acct := Account{
		Name:         name,
		PrivKey:      privKey,
		PubKey:       pubKey,
		Address:      address,
		Nonce:        0,
		KeystorePath: keystorePath,
	}
	return &acct, nil
}

func NewAccountEx(prefix string, index int, keystoreDir string) (*Account, error) {
	name := prefix + conv.Itoa(index)
	return NewAccount(name, keystoreDir)
}

func NewAccountExTwo(name string, keystoreDir string) (*Account, error) {
	return NewAccount(name, keystoreDir)
}

func LoadAccount(keystorePath string, password, fingerprint []byte) (*Account, error) {
	acct := Account{}
	err := acct.Load(keystorePath, password, fingerprint)
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

func (acct *Account) Save(password, fingerprint []byte) error {
	if acct.KeystorePath == "" {
		cmn.PanicSanity("Cannot save account because KeystorePath not set")
	}
	jsonBytes, err := cdc.MarshalJSON(acct)
	if err != nil {
		return err
	}
	walBytes := algorithm.EncryptWithPassword(jsonBytes, password, fingerprint)
	err = cmn.WriteFileAtomic(acct.KeystorePath, walBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (acct *Account) Load(keystorePath string, password, fingerprint []byte) error {
	if keystorePath == "" {
		cmn.PanicSanity("Cannot loads account because keystorePath not set")
	}
	walBytes, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		return errors.New("account does not exist")
	}
	jsonBytes, err := algorithm.DecryptWithPassword(walBytes, password, fingerprint)
	if err != nil {
		return fmt.Errorf("the password is wrong err info : %s", err)
	}
	err = cdc.UnmarshalJSON(jsonBytes, acct)
	if err != nil {
		return err
	}
	acct.KeystorePath = keystorePath

	if acct.PrivKey == nil {
		return errors.New("the wal file is corrupted")
	}
	return nil
}
