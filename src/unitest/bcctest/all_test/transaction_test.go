package all

import (
	"fmt"
	"strconv"
	"unitest/bcctest/common"

	"testing"
)

//---------------------------------------------------ChainID--------------------------------------------------------
func TestTransactionForChainID(t *testing.T) {

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
	common.PrintTitle("TestTransactionForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transaction(test.ChainID, txHash)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "TxHash:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------TxHash--------------------------------------------------------
func TestTransactionForTxHash(t *testing.T) {

	failedMsg1 := "invalid TxHash"
	failedMsg2 := "TxHash cannot be emtpy"

	num := r.Intn(200) + 1
	if num == 64 {
		num = num + r.Intn(50)
	}

	var tests = []struct {
		TxHash string
		Msg    string
		Desc   string
		Count  string //循环次数
	}{
		{"", failedMsg2, "txHash为空", "1"},
		{common.RandomTool("list11", num), failedMsg1, "交易哈希长度错误", "20"},
		{common.RandomTool("list11", 64), "", "正常流程-交易不存在", "20"},
		{common.RandomTool("list3", 64), failedMsg1, "非法交易哈希数据", "20"},
	}

	success := common.Success
	common.PrintTitle("TestTransactionForTxHash")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transaction(chainID, test.TxHash)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "TxHash:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}
