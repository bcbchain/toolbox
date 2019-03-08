package stub

import (
	"blockchain/smcsdk/sdk"
	"contract/stubcommon/common"
	"contract/stubcommon/types"

	"contract/orgexample/stub/mystorage/v1.0/mystorage"

	"contract/orgexample/stub/mycoin/v1.0/mycoin"
	"github.com/tendermint/tmlibs/log"
)

func NewStub(smc sdk.ISmartContract, logger log.Logger) types.IContractStub {

	switch common.CalcKey(smc.Message().Contract().Name(), smc.Message().Contract().Version()) {
	case "mycoin_1_0":
		return mycoin.New(logger)
	case "mystorage":
		return mystorage.New(logger)
	}
	return nil
}
