package stub

import (
	"blockchain/smcsdk/sdk"
	everycolor30 "contract/orgexample/stub/everycolor/v3.0/everycolor"
	myballot10 "contract/orgexample/stub/myballot/v1.0/myballot"
	mycoin10 "contract/orgexample/stub/mycoin/v1.0/mycoin"
	mydice2win10 "contract/orgexample/stub/mydice2win/v1.0/mydice2win"
	mystorage10 "contract/orgexample/stub/mystorage/v1.0/mystorage"
	"contract/stubcommon/common"
	"contract/stubcommon/types"
	"fmt"
	"github.com/tendermint/tmlibs/log"
)

func NewStub(smc sdk.ISmartContract, logger log.Logger) types.IContractStub {

	logger.Debug(fmt.Sprintf("NewStub error, contract=%s,version=%s", smc.Message().Contract().Name(), smc.Message().Contract().Version()))
	switch common.CalcKey(smc.Message().Contract().Name(), smc.Message().Contract().Version()) {
	case "everycolor30":
		return everycolor30.New(logger)
	case "myballot10":
		return myballot10.New(logger)
	case "mycoin10":
		return mycoin10.New(logger)
	case "mydice2win10":
		return mydice2win10.New(logger)
	case "mystorage10":
		return mystorage10.New(logger)
	default:
		logger.Fatal(fmt.Sprintf("NewStub error, contract=%s,version=%s", smc.Message().Contract().Name(), smc.Message().Contract().Version()))
	}

	return nil
}
