package all

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unitest/bcctest/common"

	"testing"
)

//---------------------------------------------------ChainID--------------------------------------------------------
func TestBlockHeightForChainID(t *testing.T) {

	failedMsg := "invalid chainId"

	chainid := make([]string, 0)
	chainid = append(chainid, common.NodeAddrSliceToTest("")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("bcb")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("bcbtest")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("devtest")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("local")...)

	var tests = []struct {
		ChainID string
		Msg     string
		Desc    string
		Count   string //循环次数
	}{
		{"", "", "正常流程", "1"},
		{chainid[r.Intn(13)], failedMsg, "无效的chainID", "1"},
		{common.RandomTool("list5", r.Intn(100)+1), failedMsg, "无效的chainID", "20"},
	}

	success := common.Success
	common.PrintTitle("TestBlockHeightForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.BlockHeight(test.ChainID)
			jsIndent, _ := json.MarshalIndent(&result.LastBlock, "", "\t")
			if string(jsIndent) != "0" && test.Msg != "" {
				panic(err)
			}
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "BLKHeight:"+string(jsIndent), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}
