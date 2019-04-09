package all

import (
	"fmt"
	"strconv"
	"testing"
	"unitest/bcctest/common"
)

//---------------------------------------------------ChainID--------------------------------------------------------
func TestBlockForChainID(t *testing.T) {

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
	common.PrintTitle("TestBlockForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Block(test.ChainID, height)
			if result != "" && test.Msg != "" {
				panic(err)
			}
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "BLKHeight:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------Height--------------------------------------------------------
func TestBlockForHeight(t *testing.T) {

	failedMsg := "invalid Height"

	var tests = []struct {
		Height string
		Msg    string
		Desc   string
		Count  string //循环次数
	}{
		{"", "", "正常流程", "1"},
		{"-1", failedMsg, "Height为-1", "1"},
		{"0", failedMsg, "Height为0", "1"},
		{string(common.GetBlockHeight()), "", "正常流程", "1"},

		{strconv.Itoa(r.Intn(999) + int(common.GetBlockHeight())), failedMsg, "大于区块高度", "20"},
		{strconv.Itoa(r.Intn(int(common.GetBlockHeight())) + 1), "", "正常流程", "20"},

		{common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(1000000000)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(1000000000)+1) + common.RandomTool("list9", 1), failedMsg, "非正整数，数据类型错误", "20"},

		{"-" + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(1000000000)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(1000000000)+1) + common.RandomTool("list9", 1), failedMsg, "非正整数，数据类型错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestBlockForHeight")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Block(chainID, test.Height)
			if result != "" && test.Msg != "" {
				panic(err)
			}
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "BLKHeight:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}
