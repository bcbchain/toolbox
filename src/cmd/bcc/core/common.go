package core

import (
	"cmd/bcc/common"
	"errors"
	"fmt"
	"strconv"
)

func nodeAddrSlice(chainID string) []string {
	if len(chainID) == 0 {
		chainID = common.GetBCCConfig().DefaultChainID
	}

	switch chainID {
	case "bcb":
		return common.GetBCCConfig().Bcb
	case "bcbtest":
		return common.GetBCCConfig().Bcbtest
	case "devtest":
		return common.GetBCCConfig().Devtest
	case "local":
		return common.GetBCCConfig().Local
	default:
		return []string{}
	}
}

func RequireNoEmpty(id, data string) (err error) {

	if len(data) == 0 {
		err = errors.New(fmt.Sprintf("%s cannot be emtpy", id))
	}

	return
}

func RequireUint64(valueStr string) (uint64, error) {
	value, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func CheckPay(pay string) (err error) {

	return
}
