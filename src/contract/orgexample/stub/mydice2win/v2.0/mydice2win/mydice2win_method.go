package mydice2win

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	sdktypes "blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/types"
	stubTypes "contract/stubcommon/types"

	"contract/orgexample/code/mydice2win/v1.0/mydice2win"
	common2 "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"
)

const msg = "Message can not be nil."

// MyBasicTypeStub stub
type MyBasicTypeStub struct {
	logger log.Logger
}

var _ stubTypes.IContractStub = (*MyBasicTypeStub)(nil)

// New new stub
func New(logger log.Logger) stubTypes.IContractStub {

	var stub MyBasicTypeStub
	stub.logger = logger

	return &stub
}

// Invoke invoke method
func (mc *MyBasicTypeStub) Invoke(smc sdk.ISmartContract) (resp types.Response) {
	//emitFee(smc)

	switch smc.Message().MethodID() {
	case "1":
		return setSecretSigner(smc)
	case "2":
		return placeBet(smc)
	}

	return
}

func setSecretSigner(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()

	if len(itemsBytes) != 1 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v1 []byte
	if err := rlp.DecodeBytes(itemsBytes[0], &v1); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	dw := new(mydice2win.Dice2Win)
	dw.SetSdk(smc)

	dw.SetSecretSigner(v1)
	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}

func placeBet(smc sdk.ISmartContract) (response types.Response) {
	itemsBytes := smc.Message().Items()

	if len(itemsBytes) != 6 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = msg

		return response
	}

	var v1 []byte
	if err := rlp.DecodeBytes(itemsBytes[0], &v1); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	var v2 []byte
	if err := rlp.DecodeBytes(itemsBytes[1], &v2); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	var v3 []byte
	if err := rlp.DecodeBytes(itemsBytes[2], &v3); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	var v4 []byte
	if err := rlp.DecodeBytes(itemsBytes[3], &v4); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	var v5 []byte
	if err := rlp.DecodeBytes(itemsBytes[4], &v5); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	var v6 []byte
	if err := rlp.DecodeBytes(itemsBytes[5], &v6); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	dw := new(mydice2win.Dice2Win)
	dw.SetSdk(smc)

	dw.PlaceBet(
		new(bn.Number).SetBytes(v1),
		new(bn.Number).SetBytes(v2).V.Int64(),
		new(bn.Number).SetBytes(v3).V.Int64(),
		v4,
		v5,
		string(v6),
	)

	response.Code = sdktypes.CodeOK
	response.Log = "success"
	response.Tags = make([]common2.KVPair, 0)
	for _, kv := range smc.Message().(*object.Message).OutputReceipts() {
		temp := common2.KVPair{Key: kv.Key, Value: kv.Value}
		response.Tags = append(response.Tags, temp)
	}

	return
}
