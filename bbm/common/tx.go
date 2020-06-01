package common

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/bcbchain/bcbchain/abciapp_v1.0/keys"
	"github.com/bcbchain/bclib/algorithm"
	"github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/bcbchain/bclib/tx/v2"
	"github.com/bcbchain/bclib/types"
	"github.com/bcbchain/bclib/utils"
	"github.com/bcbchain/sdk/sdk/rlp"
	"strconv"
	"strings"

	"github.com/bcbchain/bcbchain/abciapp_v1.0/kms"
	"github.com/bcbchain/bcbchain/abciapp_v1.0/tx"
	"github.com/btcsuite/btcutil/base58"
)

//GenerateTx generate tx with one contract method request
func GenerateTx(contract types.Address, method string, params []interface{}, nonce uint64, gaslimit int64, note string, privKey string) string {
	items := tx2.WrapInvokeParams(params...)
	message := types.Message{
		Contract: contract,
		MethodID: algorithm.BytesToUint32(algorithm.CalcMethodId(method)),
		Items:    items,
	}
	payload := tx2.WrapPayload(nonce, gaslimit, note, message)
	return tx2.WrapTx(payload, privKey)
}

// privateKey 可以是 名字:密码 也可以是 私钥的十六进制string
func GenerateTx1(nonce, gasLimit uint64, note, to, methodProtoType string, items []string, chainID, privateKey string) string {
	methodID := utils.BytesToHex(algorithm.CalcMethodId(methodProtoType))
	if !strings.Contains(privateKey, ":") {
		s := packAndSignTxV1(ToHex(nonce), ToHex(gasLimit), note, to, methodID, items, chainID, privateKey)
		return s
	}
	kms.InitKMS("./.config", "local_mode", "", "")

	//crypto.SetChainId("local")
	kms.SigMode = "local_mode"

	s := tx.PackAndSignTx(ToHex(nonce+1), ToHex(gasLimit), note, to,
		methodID, items, privateKey)
	return s
}

func packAndSignTxV1(nonce, gasLimit, note, to, methodId string, items []string, chainID, privateKey string) string {
	//parse nonce
	nonceInt, err := utils.ParseHexUint64(nonce, "nonce")
	if err != nil {
		return err.Error()
	}

	//parse gasLimit
	gasLimitInt, err := utils.ParseHexUint64(gasLimit, "gasLimit")
	if err != nil {
		return err.Error()
	}

	toAddress := to

	//parse methodId & items => data
	var mi MethodInfo

	//parse methodId
	_, err = utils.ParseHexUint32(methodId, "methodId")
	if err != nil {
		return err.Error()
	}
	dataBytes, _ := hex.DecodeString(string([]byte(methodId[2:])))
	mi.MethodID = binary.BigEndian.Uint32(dataBytes)

	var itemsBytes = make([]([]byte), 0)
	for i, item := range items {
		var itemBytes []byte
		if strings.HasPrefix(item, "0x") {

			if strings.Contains(item, ",") {
				addrs := strings.Split(item, ",")
				var addrStr string
				for _, value := range addrs {
					if strings.HasPrefix(value, "0x") {
						addrStr += strings.TrimPrefix(value, "0x")
					}
				}
				itemBytes, err = hex.DecodeString(addrStr)
			} else {
				itemBytes, err = utils.ParseHexString(item, string("item[")+strconv.Itoa(i)+"]", 0) //??
			}

			if err != nil {
				return err.Error()
			}

		} else {
			itemBytes = []byte(item)
		}
		itemsBytes = append(itemsBytes, itemBytes)
	}
	mi.ParamData, err = rlp.EncodeToBytes(itemsBytes)
	if err != nil {
		return err.Error()
	}

	data, err := rlp.EncodeToBytes(mi)
	if err != nil {
		return err.Error()
	}

	tx1 := NewTransaction(nonceInt, gasLimitInt, note, toAddress, data)

	p, e := hex.DecodeString(privateKey)

	if e != nil {
		panic("can not gen tv1 tx")
	}
	txStr, err := tx1.GenTxV1ByPrivKey(chainID, p)
	if err != nil {
		errInfo := string("{\"code\":-2, \"message\":\"tx.Transaction.TxGen failed(") + err.Error() + ")\",\"data\":\"\"}"
		return errInfo
	}
	return txStr
}

func NewTransaction(nonce uint64, gaslimit uint64, note string, to keys.Address, data []byte) Transaction {
	tx := Transaction{
		Nonce:    nonce,
		GasLimit: gaslimit,
		Note:     note,
		To:       to,
		Data:     data,
	}
	return tx
}

func ToHex(val uint64) string {
	valBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(valBytes, val)
	return string("0x") + hex.EncodeToString(valBytes)
}

func UintToHex(val uint64) string {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	return string("0x") + hex.EncodeToString(buf)
}

func (tx *Transaction) GenTxV1ByPrivKey(chainID string, passphrase []byte) (string, error) {
	//RLP编码tx
	size, r, err := rlp.EncodeToReader(tx)
	if err != nil {
		return "", err
	}
	txBytes := make([]byte, size)
	r.Read(txBytes)

	sigInfo, err := LocalSignData(passphrase, txBytes)
	if err != nil {
		return "", err
	}

	//RLP编码签名信息
	size, r, err = rlp.EncodeToReader(sigInfo)
	if err != nil {
		return "", err
	}
	sigBytes := make([]byte, size)
	r.Read(sigBytes) //转换为字节流

	txString := base58.Encode(txBytes)
	sigString := base58.Encode(sigBytes)

	MAC := string(chainID) + "<tx>"
	Version := "v1"
	SignerNumber := "<1>"

	return MAC + "." + Version + "." + txString + "." + SignerNumber + "." + sigString, nil
}

func LocalSignData(privKey []byte, data []byte) (*types.Ed25519Sig, error) {

	if len(data) <= 0 {
		return nil, errors.New("user data which wants be signed length needs more than 0")
	}

	p := crypto.PrivKeyEd25519FromBytes(privKey)
	sigInfo := types.Ed25519Sig{
		SigType:  "ed25519",
		PubKey:   p.PubKey().(crypto.PubKeyEd25519),
		SigValue: p.Sign(data).(crypto.SignatureEd25519),
	}

	return &sigInfo, nil
}
