package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bcbchain/bclib/tx/v2"
	types2 "github.com/bcbchain/bclib/types"
	"strings"

	"github.com/bcbchain/bclib/algorithm"
	"github.com/bcbchain/bclib/tendermint/go-crypto"
	"github.com/bcbchain/sdk/sdk/bn"
	"github.com/bcbchain/sdk/sdk/rlp"
	"github.com/bcbchain/sdk/sdk/types"
	"github.com/btcsuite/btcutil/base58"
)

type Ed25519Sig struct {
	SigType  string
	PubKey   crypto.PubKeyEd25519
	SigValue crypto.SignatureEd25519
}

type Transaction struct {
	Nonce    uint64
	GasLimit int64
	Note     string
	Msgs     []Message
}

type Message struct {
	To       types.Address
	MethodID uint32
	Items    []Item
}

type Item []byte

func privKeyEd25519FromBytes(data []byte) crypto.PrivKeyEd25519 {
	var privkey crypto.PrivKeyEd25519
	copy(privkey[:], data)
	return privkey
}

//调用 blockchain/bcchainlib/tx2 进行的示例代码
func demo() {
	var toContract1, toContract2 types.Address
	var methodID1, methodID2 uint32

	toContract1 = "bcbMWedWqzzW8jkt5tntTomQQEN7fSwWFhw6"
	methodID1 = algorithm.BytesToUint32(algorithm.CalcMethodId("Transfer(types.Address,bn.Number)"))

	toContract2 = "bcbCpeczqoSoxLxx1x3UyuKsaS4J8yamzWzz"
	methodID2 = algorithm.BytesToUint32(algorithm.CalcMethodId("RegisterName(string,types.Address)"))

	tx2.Init("bcb")

	// step 1 打包方法调用需要的参数
	var toAccount types.Address = "bcbCpeczqoSoxLxx1x3UyuKsaS4J8yamzWzz"
	var value bn.Number = bn.N(1000000000)
	params1 := tx2.WrapInvokeParams(toAccount, value)

	var nickName string = "mmmmhhhh"
	var referee types.Address = "bcb5rzgE1tSJbJuegEj4vbAkotmwRkxwiSyV"
	params2 := tx2.WrapInvokeParams(nickName, referee)

	msg1 := types2.Message{toContract1, methodID1, params1}
	msg2 := types2.Message{toContract2, methodID2, params2}

	// step 2 打包交易净荷
	var nonce uint64 = 1
	var gasLimit int64 = 500
	var note string = "Example for cascade invoke smart contract."
	payload := tx2.WrapPayload(nonce, gasLimit, note, msg1, msg2)

	// step 3 对交易净荷签名并组成最终交易数据
	privKey := "0x4a2c14697282e658b3ed7dd5324de1a102d216d6fa50d5937ffe89f35cbc12aa68eb9a09813bdf7c0869bf34a244cc545711509fe70f978d121afd3a4ae610e6"
	tx := tx2.WrapTx(payload, privKey)

	fmt.Printf("TX:%v\n\n", tx)
}

func main() {
	crypto.SetChainId("bcb")
	//privKey := crypto.PrivKeyEd25519(crypto.GenPrivKeyEd25519())
	privKBytes, _ := hex.DecodeString("4a2c14697282e658b3ed7dd5324de1a102d216d6fa50d5937ffe89f35cbc12aa68eb9a09813bdf7c0869bf34a244cc545711509fe70f978d121afd3a4ae610e6")
	privKey := crypto.PrivKeyEd25519(privKeyEd25519FromBytes(privKBytes))
	pubKey := privKey.PubKey()
	address := pubKey.Address("bcb")

	pkBytes := make([]byte, 64)
	copy(pkBytes[:], privKey[:])

	fmt.Println("<==== sender ===========================================================>")
	fmt.Println("privKey:", hex.EncodeToString(pkBytes))
	fmt.Println("pubKey: ", pubKey)
	fmt.Println("address:", address)
	fmt.Println("")

	var nonce uint64
	nonce = 1
	nonceRlp, err := rlp.EncodeToBytes(nonce)
	if err != nil {
		panic(err)
	}

	var gasLimit int64
	gasLimit = 500
	gasLimitRlp, err := rlp.EncodeToBytes(gasLimit)
	if err != nil {
		panic(err)
	}

	var note string
	note = "Example for cascade invoke smart contract."
	noteRlp, err := rlp.EncodeToBytes(note)
	if err != nil {
		panic(err)
	}

	//fisrt contract
	var toContract1 types.Address
	toContract1 = "bcbMWedWqzzW8jkt5tntTomQQEN7fSwWFhw6"
	toContract1Rlp, err := rlp.EncodeToBytes(toContract1)
	if err != nil {
		panic(err)
	}

	var methodID1 uint32
	methodID1 = algorithm.BytesToUint32(algorithm.CalcMethodId("Transfer(types.Address,bn.Number)"))
	methodID1Rlp, err := rlp.EncodeToBytes(methodID1)
	if err != nil {
		panic(err)
	}

	var toAccount types.Address
	toAccount = "bcbCpeczqoSoxLxx1x3UyuKsaS4J8yamzWzz"
	toAccountRlp, err := rlp.EncodeToBytes(toAccount)
	if err != nil {
		panic(err)
	}

	var value bn.Number
	value = bn.N(1000000000)
	valBytes := value.Bytes()
	valueRlp, err := rlp.EncodeToBytes(valBytes)
	if err != nil {
		panic(err)
	}

	//second contract
	var toContract2 types.Address
	toContract2 = "bcbCpeczqoSoxLxx1x3UyuKsaS4J8yamzWzz"
	toContract2Rlp, err := rlp.EncodeToBytes(toContract2)
	if err != nil {
		panic(err)
	}

	var methodID2 uint32
	methodID2 = algorithm.BytesToUint32(algorithm.CalcMethodId("RegisterName(string,types.Address)"))
	methodID2Rlp, err := rlp.EncodeToBytes(methodID2)
	if err != nil {
		panic(err)
	}

	var nickName string
	nickName = "mmmmhhhh"
	nickNameRlp, err := rlp.EncodeToBytes(nickName)
	if err != nil {
		panic(err)
	}

	var referee types.Address
	referee = "bcb5rzgE1tSJbJuegEj4vbAkotmwRkxwiSyV"
	refereeRlp, err := rlp.EncodeToBytes(referee)
	if err != nil {
		panic(err)
	}

	fmt.Printf("原始参数：\n")
	fmt.Printf("<==== raw params ========================================================>\n")
	fmt.Printf("nonce:               %v\n", nonce)
	fmt.Printf("gasLimit:            %v\n", gasLimit)
	fmt.Printf("note:                %v\n", note)
	fmt.Printf("msgs: [\n")
	fmt.Printf("  {\n")
	fmt.Printf("    toContract:      %v\n", toContract1)
	fmt.Printf("    methodID:        0x%v\n", methodID1)
	fmt.Printf("    items{\n")
	fmt.Printf("      toAccount:     %v\n", toAccount)
	fmt.Printf("      value:         %v\n", value)
	fmt.Printf("    }\n")
	fmt.Printf("  }\n")
	fmt.Printf("  {\n")
	fmt.Printf("    toContract:      %v\n", toContract2)
	fmt.Printf("    methodID:        0x%v\n", methodID2)
	fmt.Printf("    items{\n")
	fmt.Printf("      nickName:      %v\n", nickName)
	fmt.Printf("      referee:       %v\n", referee)
	fmt.Printf("    }\n")
	fmt.Printf("  }\n")
	fmt.Printf("]\n")
	fmt.Printf("\n")

	fmt.Printf("对原始数据中调用合约方法的每个参数单独进行RLP编码：\n")
	fmt.Printf("<==== raw params ========================================================>\n")
	fmt.Printf("nonce:               %v\n", nonce)
	fmt.Printf("gasLimit:            %v\n", gasLimit)
	fmt.Printf("note:                %v\n", note)
	fmt.Printf("msgs: [\n")
	fmt.Printf("  {\n")
	fmt.Printf("    toContract:      %v\n", toContract1)
	fmt.Printf("    methodID:        0x%v\n", methodID1)
	fmt.Printf("    items{\n")
	fmt.Printf("      toAccount(RLP):0x%v\n", hex.EncodeToString(toAccountRlp))
	fmt.Printf("      value(RLP):    0x%v\n", hex.EncodeToString(valueRlp))
	fmt.Printf("    }\n")
	fmt.Printf("  }\n")
	fmt.Printf("  {\n")
	fmt.Printf("    toContract:      %v\n", toContract2)
	fmt.Printf("    methodID:        0x%v\n", methodID2)
	fmt.Printf("    items{\n")
	fmt.Printf("      nickName(RLP): 0x%v\n", hex.EncodeToString(nickNameRlp))
	fmt.Printf("      referee(RLP):  0x%v\n", hex.EncodeToString(refereeRlp))
	fmt.Printf("    }\n")
	fmt.Printf("  }\n")
	fmt.Printf("]\n")
	fmt.Printf("\n")

	//items for first contract
	itemInfo1 := make([]Item, 2)
	itemInfo1[0] = toAccountRlp
	itemInfo1[1] = valueRlp
	itemInfo1Rlp, err := rlp.EncodeToBytes(itemInfo1)
	if err != nil {
		panic(err)
	}

	message1 := Message{
		To:       toContract1,
		MethodID: methodID1,
		Items:    itemInfo1,
	}
	message1Rlp, err := rlp.EncodeToBytes(message1)
	if err != nil {
		panic(err)
	}

	//items for second contract
	itemInfo2 := make([]Item, 2)
	itemInfo2[0] = nickNameRlp
	itemInfo2[1] = refereeRlp
	itemInfo2Rlp, err := rlp.EncodeToBytes(itemInfo2)
	if err != nil {
		panic(err)
	}

	message2 := Message{
		To:       toContract2,
		MethodID: methodID2,
		Items:    itemInfo2,
	}
	message2Rlp, err := rlp.EncodeToBytes(message2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("对原始数据中调用合约方法的参数表进行RLP编码：\n")
	fmt.Printf("<==== raw params (partial rlp encoded) ==================================>\n")
	fmt.Printf("nonce:               %v\n", nonce)
	fmt.Printf("gasLimit:            %v\n", gasLimit)
	fmt.Printf("note:                %v\n", note)
	fmt.Printf("msgs: [\n")
	fmt.Printf("  {\n")
	fmt.Printf("    toContract(RLP): 0x%v\n", hex.EncodeToString(toContract1Rlp))
	fmt.Printf("    methodID(RLP):   0x%v\n", hex.EncodeToString(methodID1Rlp))
	fmt.Printf("    items(RLP):      0x%v\n", hex.EncodeToString(itemInfo1Rlp))
	fmt.Printf("  },\n")
	fmt.Printf("  {\n")
	fmt.Printf("    toContract(RLP): 0x%v\n", hex.EncodeToString(toContract2Rlp))
	fmt.Printf("    methodID(RLP):   0x%v\n", hex.EncodeToString(methodID2Rlp))
	fmt.Printf("    items(RLP):      0x%v\n", hex.EncodeToString(itemInfo2Rlp))
	fmt.Printf("  }\n")
	fmt.Printf("]\n")
	fmt.Printf("\n")

	fmt.Printf("对原始数据中调用合约方法的单个消息进行RLP编码：\n")
	fmt.Printf("<==== raw params (partial rlp encoded) ==================================>\n")
	fmt.Printf("nonce:               %v\n", nonce)
	fmt.Printf("gasLimit:            %v\n", gasLimit)
	fmt.Printf("note:                %v\n", note)
	fmt.Printf("msgs: [\n")
	fmt.Printf("  message1(RLP):     0x%v\n", hex.EncodeToString(message1Rlp))
	fmt.Printf("  message2(RLP):     0x%v\n", hex.EncodeToString(message2Rlp))
	fmt.Printf("]\n")
	fmt.Printf("\n")

	txInfo := Transaction{
		Nonce:    nonce,
		GasLimit: gasLimit,
		Note:     note,
		Msgs:     make([]Message, 2),
	}
	txInfo.Msgs[0] = message1
	txInfo.Msgs[1] = message2
	msgsBytesRlp, err := rlp.EncodeToBytes(txInfo.Msgs)
	txPayloadBytesRlp, err := rlp.EncodeToBytes(txInfo)

	fmt.Printf("对原始数据中调用合约方法的消息表及交易参数进行RLP编码：\n")
	fmt.Printf("<==== raw params (partial rlp encoded) ==================================>\n")
	fmt.Printf("nonce:               0x%v\n", hex.EncodeToString(nonceRlp))
	fmt.Printf("gasLimit:            0x%v\n", hex.EncodeToString(gasLimitRlp))
	fmt.Printf("note:                0x%v\n", hex.EncodeToString(noteRlp))
	fmt.Printf("msgs:                0x%v\n", hex.EncodeToString(msgsBytesRlp))
	fmt.Printf("\n")

	fmt.Printf("对对交易数据进行RLP编码：得到待签名的净荷数据\n")
	fmt.Printf("<==== rlp for tx data (payload for sign) ================================>\n")
	fmt.Printf("payload:             0x%v\n", hex.EncodeToString(txPayloadBytesRlp))
	fmt.Printf("\n")

	sigInfo := Ed25519Sig{
		"ed25519",
		pubKey.(crypto.PubKeyEd25519),
		privKey.Sign(txPayloadBytesRlp).(crypto.SignatureEd25519),
	}
	sigBytes := make([]byte, 64)
	copy(sigBytes[:], sigInfo.SigValue[:])

	fmt.Printf("对净荷数据进行签名：\n")
	fmt.Printf("<==== raw sign data for payload =========================================>\n")
	fmt.Printf("sig:type:            %v\n", sigInfo.SigType)
	fmt.Printf("sig:pubkey:          0x%v\n", sigInfo.PubKey)
	fmt.Printf("sig:signdata:        0x%v\n", hex.EncodeToString(sigBytes))
	fmt.Printf("\n")

	typeRlp, err := rlp.EncodeToBytes(sigInfo.SigType)
	if err != nil {
		panic(err)
	}
	pubKeyRlp, err := rlp.EncodeToBytes(sigInfo.PubKey)
	if err != nil {
		panic(err)
	}
	sigDataRlp, err := rlp.EncodeToBytes(sigBytes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("对签名数据各字段进行RLP编码：\n")
	fmt.Printf("<==== rlp for raw sign data =============================================>\n")
	fmt.Printf("sig:type:            0x%v\n", hex.EncodeToString(typeRlp))
	fmt.Printf("sig:pubkey:          0x%v\n", hex.EncodeToString(pubKeyRlp))
	fmt.Printf("sig:signdata:        0x%v\n", hex.EncodeToString(sigDataRlp))
	fmt.Printf("\n")

	mrlpsigs := make([]byte, 0)
	mrlpsigs = append(mrlpsigs, typeRlp...)
	mrlpsigs = append(mrlpsigs, pubKeyRlp...)
	mrlpsigs = append(mrlpsigs, sigDataRlp...)

	fmt.Printf("对上述RLP编码的签名数据按顺序合并：\n")
	fmt.Printf("<==== merged sign data list =============================================>\n")
	fmt.Printf("sig:signdata:        0x%v\n", hex.EncodeToString(mrlpsigs))
	fmt.Printf("\n")

	mrlpSigsRlp, err := rlp.EncodeToBytes(sigInfo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("对合并后的编码签名数据再次进行RLP编码：得到最终的签名数据\n")
	fmt.Printf("<==== rlp for merged sign data list =====================================>\n")
	fmt.Printf("rlp:signdata:        0x%v\n", hex.EncodeToString(mrlpSigsRlp))
	fmt.Printf("\n")

	txString := base58.Encode(txPayloadBytesRlp)

	fmt.Printf("对净荷数据进行Base58编码：\n")
	fmt.Printf("<==== Base58 for payload ================================================>\n")
	fmt.Printf("data:  %v\n", txString)
	fmt.Printf("\n")

	sigString := base58.Encode(mrlpSigsRlp)

	fmt.Printf("对签名数据进行Base58编码：\n")
	fmt.Printf("<==== Base58 for sign data ==============================================>\n")
	fmt.Printf("data:  %v\n", sigString)
	fmt.Printf("\n")

	MAC := string("bcb") + "<tx>"
	Version := "v2"
	SignerNumber := "<1>"

	finalTx := MAC + "." + Version + "." + txString + "." + SignerNumber + "." + sigString

	fmt.Printf("将经过Base58编码的净荷与签名数据按规范构造最终的交易数据：\n")
	fmt.Printf("<==== final tx data =====================================================>\n")
	fmt.Printf("%v\n", finalTx)
	fmt.Printf("\n")

	hash := algorithm.CalcCodeHash(finalTx)

	fmt.Printf("计算交易数据的哈希：\n")
	fmt.Printf("<==== final tx hash =====================================================>\n")
	fmt.Printf("%v\n", strings.ToUpper(hex.EncodeToString(hash)))
	fmt.Printf("\n")

	/*
		var jsonBuf bytes.Buffer
		txp.InitUnWrapper("bcb")
		parsedTx := txp.UnpackAndParseTx(finalTx)
		json.Indent(&jsonBuf, []byte(parsedTx), "", "  ")
		fmt.Printf("parse final tx:\n")
		fmt.Printf("%v\n", jsonBuf.String())
		fmt.Printf("\n")
	*/
}
