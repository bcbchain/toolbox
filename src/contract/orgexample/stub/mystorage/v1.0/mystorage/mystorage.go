package mystorage

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/rlp"
	sdktypes "blockchain/smcsdk/sdk/types"
	"blockchain/types"
	stubTypes "contract/stubcommon/types"
	"strconv"

	"contract/orgexample/code/mystorage/v1.0/mystorage"
	"github.com/tendermint/tmlibs/log"
)

type StorageStub struct {
	logger log.Logger
}

var methodIdToFunc = map[string]func([]sdktypes.HexBytes, sdk.ISmartContract) types.Response{
	"45676666": set,
	"45678888": get,
}

var _ stubTypes.IContractStub = (*StorageStub)(nil)

func New(logger log.Logger) stubTypes.IContractStub {

	var stub StorageStub
	stub.logger = logger

	return &stub
}

func (stub *StorageStub) Invoke(smcApi sdk.ISmartContract) (response types.Response) {
	// TODO 手续费
	if smcApi.Message() == nil {
		response.Code = 4001
		response.Log = "Message can not be nil."

		return response
	}

	var itemsBytes = smcApi.Message().Items()

	f, ok := methodIdToFunc[smcApi.Message().MethodID()]
	if !ok {

		response.Code = sdktypes.ErrStubDefined
		response.Log = "Invalid method id."

		return response
	}

	response = f(itemsBytes, smcApi)

	return response
}

func set(itemsBytes []sdktypes.HexBytes, smcApi sdk.ISmartContract) types.Response {

	response := types.Response{}
	if len(itemsBytes) != 1 {

		response.Code = sdktypes.ErrStubDefined
		response.Log = "Parameters invalid."

		return response
	}

	var date uint64
	if err := rlp.DecodeBytes(itemsBytes[0], &date); err != nil {

		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	storage := new(mystorage.MyStorage)
	storage.SetSdk(smcApi)

	response.Code = 200

	return response
}

func get(itemsBytes []sdktypes.HexBytes, smcApi sdk.ISmartContract) types.Response {

	response := types.Response{}
	if len(itemsBytes) != 0 {

		response.Code = sdktypes.ErrStubDefined
		response.Log = "Parameters invalid."

		return response
	}

	storage := new(mystorage.MyStorage)
	storage.SetSdk(smcApi)

	data := storage.Get()
	response.Code = 200
	response.Data = strconv.FormatUint(data, 10)

	return response
}
