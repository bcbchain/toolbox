package all

import (
	"fmt"
	"strconv"
	"testing"
	"unitest/bcctest/common"
)

//---------------------------------------------------ChainID--------------------------------------------------------
func TestVersionForChainID(t *testing.T) {

	//failedMsg := "invalid chainId"

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
		{chainid[r.Intn(13)], "", "正常流程", "1"},
		{common.RandomTool("list5", r.Intn(100)+1), "", "正常流程", "20"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.VersionF()
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Version:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}
