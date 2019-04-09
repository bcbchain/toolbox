package all

import (
	"fmt"
	"strconv"
	"strings"
	"unitest/bcctest/common"

	"testing"
)

//---------------------------------------------------Account--------------------------------------------------------
func TestBalanceForAccount(t *testing.T) {
	failedMsg1 := `Name contains by [letters, numbers, "_", "@", "." and "-"] and length must be [1-40]`
	failedMsg2 := "Password contains by [letters, numbers, ASCII 32 through 127] and length must be [8-20]"

	password1 := common.RandomTool("list5", 2) + common.RandomTool("list6", 2) + common.RandomTool("list7", 2) + common.RandomTool("list8", 2)
	password2 := common.RandomTool("list5", 5) + common.RandomTool("list6", 5) + common.RandomTool("list7", 5) + common.RandomTool("list8", 5)
	password3 := common.RandomTool("list5", r.Intn(3)+2) + common.RandomTool("list6", r.Intn(3)+2) + common.RandomTool("list7", r.Intn(3)+2) + common.RandomTool("list8", r.Intn(3)+3)
	password4 := common.RandomTool("list5", r.Intn(3)+2) + common.RandomTool("list6", r.Intn(3)+3) + common.RandomTool("list8", r.Intn(3)+3)
	password5 := common.RandomTool("list7", r.Intn(3)+2) + common.RandomTool("list6", r.Intn(3)+3) + common.RandomTool("list8", r.Intn(3)+3)
	password6 := common.RandomTool("list5", r.Intn(3)+2) + common.RandomTool("list7", r.Intn(3)+3) + common.RandomTool("list8", r.Intn(3)+3)
	password7 := common.RandomTool("list5", r.Intn(3)+2) + common.RandomTool("list6", r.Intn(3)+3) + common.RandomTool("list7", r.Intn(3)+3)
	password8 := common.RandomTool("list5", 2) + common.RandomTool("list6", 2) + common.RandomTool("list7", 2) + common.RandomTool("list8", 2) + common.RandomTool("list4", r.Intn(12)+1)
	password9 := common.RandomTool("list5", r.Intn(100)+5) + common.RandomTool("list6", r.Intn(100)+5) + common.RandomTool("list7", r.Intn(100)+5) + common.RandomTool("list8", r.Intn(100)+6)

	var tests = []struct {
		Name     string
		Password string
		Msg      string
		Desc     string
		Count    string //循环次数
	}{
		{"", password, failedMsg1, "用户名为空", "1"},
		{common.RandomTool("list1", 1), password, "", "正常流程", "20"},
		{common.RandomTool("list1", 40), password, "", "正常流程", "20"},
		{common.RandomTool("list1", r.Intn(38)+2), password, "", "正常流程", "20"},
		{common.RandomTool("list2", r.Intn(40)+1), password, failedMsg1, "异常流程，非法字符", "20"},
		{common.RandomTool("list1", r.Intn(300)+40), password, failedMsg1, "账户名称长度错误", "20"},

		{name, "", failedMsg2, "密码不能为空", "1"},
		{name, common.RandomTool("list3", 1), failedMsg2, "密码长度错误", "20"},
		{name, common.RandomTool("list3", 7), failedMsg2, "密码长度错误", "20"},
		{name, password1, "", "正常流程", "20"},
		{name, password2, "", "正常流程", "20"},
		{name, password3, "", "正常流程", "20"},
		{name, password4, failedMsg2, "不满足密码要求，缺数字", "20"},
		{name, password5, failedMsg2, "不满足密码要求，缺小写字母", "20"},
		{name, password6, failedMsg2, "不满足密码要求，缺小写字母", "20"},
		{name, password7, failedMsg2, "不满足密码要求，缺特殊字符", "20"},
		{name, password8, failedMsg2, "含有异常（不可打印）字符", "20"},
		{name, password9, failedMsg2, "密码长度错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestBalanceForAccount")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Balance(accAddr, test.Name, test.Password, token, all, chainID, keystorePath)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Balance:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
			//
		}
	}
}

//---------------------------------------------------ChainID--------------------------------------------------------
func TestBalanceForChainID(t *testing.T) {

	failedMsg := "invalid chainId"

	chainid := make([]string, 0)
	chainid = append(chainid, common.NodeAddrSliceToTest("")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("bcb")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("bcbtest")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("devtest")...)
	chainid = append(chainid, common.NodeAddrSliceToTest("local")...)

	var tests = []struct {
		Addr    string
		Name    string
		ChainID string
		Msg     string
		Desc    string
		Count   string //循环次数
	}{
		{"", "", "", failedMsg, "用户名和地址为空", "1"},
		{"", "", chainid[r.Intn(13)], failedMsg, "用户名和地址为空", "1"},
		{"", "", common.RandomTool("list5", r.Intn(100)+1), failedMsg, "用户名和地址为空", "20"},

		{accAddr, "", "", "", "正常流程", "1"},
		{accAddr, "", chainid[r.Intn(13)], failedMsg, "无效的chainID", "1"},
		{accAddr, "", common.RandomTool("list5", r.Intn(100)+1), failedMsg, "无效的chainID", "20"},

		{"", name, "", "", "正常流程", "1"},
		{"", name, chainid[r.Intn(13)], failedMsg, "无效的chainID", "1"},
		{"", name, common.RandomTool("list5", r.Intn(100)+1), failedMsg, "无效的chainID", "20"},
	}

	success := common.Success
	common.PrintTitle("TestBalanceForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Balance(test.Addr, test.Name, password, token, all, test.ChainID, keystorePath)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Balance:"+result, err)
			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------KeyStorePath--------------------------------------------------------
func TestBalanceForKeyStorePath(t *testing.T) {

	failedMsg := "KeyStorePath does not exist"

	path := make([]string, 0)
	for m := 0; m < r.Intn(5)+1; m++ {
		path = append(path, common.RandomTool("list5", r.Intn(10)+1))
	}
	Path1 := strings.Join(path, "/")

	var tests = []struct {
		Addr         string
		Name         string
		KeyStorePath string
		Msg          string
		Desc         string
		Count        string //循环次数
	}{
		{"", "", "", failedMsg, "用户名和地址为空", "1"},
		{"", "", Path1, failedMsg, "钱包地址不存在(用户名和地址为空)", "1"},
		{"", "", "/home/lzw/svn/code/v2.0/bctoolbox/src/unitest/bcctest/all_test/", failedMsg, "用户名和地址为空", "1"},
		{"", "", ".keystore/", failedMsg, "用户名和地址为空", "1"},

		{accAddr, "", "", "", "正常流程", "1"},
		{accAddr, "", Path1, "", "钱包地址不存在-(合法地址)", "1"},
		{accAddr, "", "/home/lzw/svn/code/v2.0/bctoolbox/src/unitest/bcctest/all_test/", "", "正常流程", "1"},
		{accAddr, "", ".keystore/", "", "正常流程", "1"},

		{"", name, "", "", "正常流程", "1"},
		{"", name, Path1, "", "钱包地址不存在-(合法用户名)", "1"},
		{"", name, "/home/lzw/svn/code/v2.0/bctoolbox/src/unitest/bcctest/all_test/", "", "正常流程", "1"},
		{"", name, ".keystore/", "", "正常流程", "1"},
	}

	success := common.Success
	common.PrintTitle("TestBalanceForKeyStorePath")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Balance(test.Addr, test.Name, password, token, all, chainID, test.KeyStorePath)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Balance:"+result, err)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------AccAddress--------------------------------------------------------
func TestBalanceForTo(t *testing.T) {

	//failedMsg := "Address cannot be emtpy"
	failedMsg2 := "Invalid Address Format"

	var tests = []struct {
		Name       string
		AccAddress string
		Msg        string
		Desc       string
		Count      string //循环次数
	}{
		//{"","", failedMsg, "用户名和地址为空", "1"},
		//{"",accAddr, "", "正常流程", "1"},
		{"", "qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j", failedMsg2, "非法地址-chainID-(用户名为空)", "1"},
		{"", "qqqqL9BzYNYns5dfgdfgfHEBJLzS1bhpHjx7j", failedMsg2, "非法地址-base58-(用户名为空)", "1"},
		{"", "qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHssss", failedMsg2, "非法地址-4位验证码-(用户名为空)", "1"},

		{name, "", "", "地址为空", "1"},
		{name, accAddr, "", "正常流程", "1"},
		{name, "qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j", "", "非法地址-chainID", "1"},
		{name, "qqqqL9BzYNYns5dfgdfgfHEBJLzS1bhpHjx7j", "", "非法地址-base58", "1"},
		{name, "qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHssss", "", "非法地址-4位验证码", "1"},
	}

	success := common.Success
	common.PrintTitle("TestBalanceForTo")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Balance(test.AccAddress, test.Name, password, token, all, chainID, keystorePath)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Balance:"+result, err)
			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------TokenName--------------------------------------------------------
func TestBalanceForTokenName(t *testing.T) {

	failedMsg2 := "tokenName cannot be emtpy"
	failedMsg3 := "Invalid tokenName."

	var tests = []struct {
		TokenName string
		Msg       string
		Desc      string
		Count     string //循环次数
	}{
		{"", failedMsg2, "代币名为空", "1"},
		{"xtt", "", "正常流程", "1"},
		{common.RandomTool("list3", 1), failedMsg3, "名称太短", "20"},
		{common.RandomTool("list3", 2), failedMsg3, "名称太短", "20"},
		{common.RandomTool("list3", 3), failedMsg3, "名称太短", "20"},
		{common.TestTokenName("list3", r.Intn(36)+4), "", "正常流程", "20"},
		{common.TestTokenName("list3", 40), "", "正常流程", "20"},
		{common.RandomTool("list3", r.Intn(200)+41), failedMsg3, "名称太长", "20"},
	}

	success := common.Success
	common.PrintTitle("TestBalanceForTokenName")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Balance(accAddr, name, password, test.TokenName, all, chainID, keystorePath)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Balance:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}
