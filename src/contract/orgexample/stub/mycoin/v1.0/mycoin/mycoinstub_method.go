package mycoin

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	sdktypes "blockchain/smcsdk/sdk/types"
	"blockchain/types"
	stubTypes "contract/stubcommon/types"

	"github.com/tendermint/tmlibs/log"
)

type CoinStub struct {
	logger log.Logger
}

var _ stubTypes.IContractStub = (*CoinStub)(nil)

func New(logger log.Logger) stubTypes.IContractStub {

	var stub CoinStub
	stub.logger = logger

	return &stub
}

func (mc *CoinStub) Invoke(smcApi sdk.ISmartContract) types.Response {

	// TODO 手续费
	switch smcApi.Message().MethodID() {
	case "23445656":
		return transfer(smcApi)
	}
	return types.Response{}
}

func transfer(smcApi sdk.ISmartContract) (response types.Response) {

	itemsBytes := smcApi.Message().Items()

	if len(itemsBytes) != 2 {
		response.Code = sdktypes.ErrStubDefined
		response.Log = "Message can not be nil."

		return response
	}

	type paramTransfer struct {
		to    types.Address
		value bn.Number
	}
	param := paramTransfer{}
	if err := rlp.DecodeBytes(itemsBytes[0], param.to); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}

	if err := rlp.DecodeBytes(itemsBytes[1], param.value); err != nil {
		response.Code = sdktypes.ErrStubDefined
		response.Log = err.Error()

		return response
	}
	inter := IntfcCoinStub{smcApi}
	return inter.core_transfer(&param)
}
