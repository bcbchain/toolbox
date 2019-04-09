package all

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unitest/bcctest/common"
)

//---------------------------------------------------Account--------------------------------------------------------
func TestNonceForAccount(t *testing.T) {
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
		//{"", password, failedMsg1, "用户名为空", "1"},
		{common.NameofTest("list1", 1), password, "", "正常流程", "20"},
		{common.NameofTest("list1", 40), password, "", "正常流程", "20"},
		{common.NameofTest("list1", r.Intn(38)+2), password, "", "正常流程", "20"},
		{common.NameofTest("list2", r.Intn(40)+1), password, failedMsg1, "异常流程，非法字符", "20"},
		{common.NameofTest("list1", r.Intn(300)+40), password, failedMsg1, "账户名称长度错误", "20"},

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
	common.PrintTitle("TestNonceForAccount")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Noncce("", test.Name, test.Password, chainID, keystorePath)
			common.PrintResult(err)
			jsIndent := []byte{}
			if result != nil {
				jsIndent, _ = json.MarshalIndent(&result.Nonce, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Nonce:"+string(jsIndent), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------ChainID--------------------------------------------------------
func TestNonceForChainID(t *testing.T) {

	failedMsg := "invalid chainId"

	chainid := make([]string, 0)
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
		{chainid[r.Intn(13)], "", "无效的chainID", "1"},
		{common.RandomTool("list5", r.Intn(100)+1), failedMsg, "无效的chainID", "20"},
	}

	success := common.Success
	common.PrintTitle("TestNonceForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Noncce("", name, password, test.ChainID, keystorePath)
			common.PrintResult(err)
			jsIndent := []byte{}
			if result != nil {
				jsIndent, _ = json.MarshalIndent(&result.Nonce, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Nonce:"+string(jsIndent), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------KeyStorePath--------------------------------------------------------
func TestNonceForKeyStorePath(t *testing.T) {

	failedMsg := "KeyStorePath does not exist"

	path := make([]string, 0)
	for m := 0; m < r.Intn(5)+1; m++ {
		path = append(path, common.RandomTool("list5", r.Intn(10)+1))
	}
	Path1 := strings.Join(path, "/")

	var tests = []struct {
		Name         string
		Pwd          string
		Addr         string
		KeyStorePath string
		Msg          string
		Desc         string
		Count        string //循环次数
	}{
		{"", "", Address, "", "", "正常流程", "1"},
		{"", "", Address, Path1, "", "正常流程(地址正确)", "20"},
		{"", "", Address, "/home/lzw/svn/code/v2.0/bctoolbox/src/unitest/bcctest/all_test/", "", "正常流程(地址正确)", "1"},
		{"", "", Address, ".keystore/", "", "正常流程(地址正确)", "1"},

		{"", "", "", "", "", "正常流程", "1"},
		{"", "", "", Path1, failedMsg, "钱包地址不存在", "20"},
		{name, password, "", "/home/lzw/svn/code/v2.0/bctoolbox/src/unitest/bcctest/all_test/", "", "正常流程", "1"},
		{name, password, "", ".keystore/", "", "正常流程", "1"},
	}

	success := common.Success
	common.PrintTitle("TestNonceForKeyStorePath")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Nonce(test.Addr, test.Name, test.Pwd, chainID, test.KeyStorePath)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Nonce:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------Address--------------------------------------------------------
func TestNonceForKeyAddress(t *testing.T) {

	failedMsg := "invalid Address"

	var tests = []struct {
		Address string
		Name    string
		Pwd     string
		Msg     string
		Desc    string
		Count   string //循环次数
	}{
		{"", name, password, "", "正常流程", "1"},
		{Address, name, password, "", "正常流程", "1"},
		{"qqqqL9BzYNYns5dfgdfgf@@HEBJLzS1bhpHjx7j", name, password, failedMsg, "非法地址", "1"},

		{"", "", "", failedMsg, "地址和用户名为空", "1"},
		{Address, "", "", "", "正常流程", "1"},
		{"qqqqL9BzYNYns5dfgdfgf@@HEBJLzS1bhpHjx7j", "", "", failedMsg, "地址错误，用户名为空", "1"},
	}

	success := common.Success
	common.PrintTitle("TestNonceForKeyAddress")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Nonce(test.Address, test.Name, test.Pwd, chainID, keystorePath)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, "Nonce:"+result, err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}
