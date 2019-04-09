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
func TestTranferForAccount(t *testing.T) {
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

	name1, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password1)
	name2, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password2)
	name3, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password3)
	name4, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password4)
	name5, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password5)
	name6, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password6)
	name7, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password7)
	name8, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password8)
	name9, _ := common.PassWordOfTest(common.RandomTool("list1", 10), password9)
	var tests = []struct {
		Name     string
		Password string
		Msg      string
		Desc     string
		Count    string //循环次数
	}{
		{"", password, failedMsg1, "用户名为空", "1"},
		{common.NameofTest("list1", 1), password, "", "正常流程", "20"},
		{common.NameofTest("list1", 40), password, "", "正常流程", "20"},
		{common.NameofTest("list1", r.Intn(38)+2), password, "", "正常流程", "20"},
		{common.NameofTest("list2", r.Intn(40)+1), password, failedMsg1, "异常流程，非法字符", "20"},
		{common.NameofTest("list1", r.Intn(300)+40), password, failedMsg1, "账户名称长度错误", "20"},

		{name, "", failedMsg2, "密码不能为空", "1"},
		{name, common.RandomTool("list3", 1), failedMsg2, "密码长度错误", "20"},
		{name, common.RandomTool("list3", 7), failedMsg2, "密码长度错误", "20"},
		{name1, password1, "", "正常流程", "20"},
		{name2, password2, "", "正常流程", "20"},
		{name3, password3, "", "正常流程", "20"},
		{name4, password4, failedMsg2, "不满足密码要求，缺数字", "20"},
		{name5, password5, failedMsg2, "不满足密码要求，缺小写字母", "20"},
		{name6, password6, failedMsg2, "不满足密码要求，缺小写字母", "20"},
		{name7, password7, failedMsg2, "不满足密码要求，缺特殊字符", "20"},
		{name8, password8, failedMsg2, "含有异常（不可打印）字符", "20"},
		{name9, password9, failedMsg2, "密码长度错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForAccount")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(test.Name, test.Password, token, gasLimit, note, to, value, keystorePath, chainID)
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------gasLimit--------------------------------------------------------
func TestTranferForGasLimit(t *testing.T) {

	//failedMsg1 := "GasLimit has to be a positive integer!"
	failedMsg1 := "gasLimit cannot be emtpy"
	//failedMsg2 := "success or failed"
	failedMsg3 := "Gas Limit is not enough"
	failedMsg4 := "GasLimit has to be a positive integer!"

	var tests = []struct {
		GasLimit string
		Msg      string
		//Code uint32
		Desc  string
		Count string //循环次数
	}{
		{"", failedMsg1 /*, bcerrors.ErrCodeInterContractsInvalidGasLimit*/, "gasLimit为空", "1"},
		//{"100000", ""/*, bcerrors.ErrCodeInterContractsInvalidGasLimit*/,"gasLimit为空", "2"},
		{"0", failedMsg3, "gasLimit 不足", "1"},
		{strconv.Itoa(r.Intn(1000000000) + 1), "", "成功或者失败", "20"},
		{"-" + strconv.Itoa(r.Intn(1000000000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(1000000000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(1000000000)+1) + common.RandomTool("list9", 1), failedMsg4, "非正整数，数据类型错误", "20"},
		{"-" + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(1000000000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(1000000000)+1) + common.RandomTool("list9", 1), failedMsg4, "非正整数，数据类型错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForGasLimit")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(name, password, token, test.GasLimit, note, to, value, keystorePath, chainID)
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------Note--------------------------------------------------------
func TestTranferForNote(t *testing.T) {

	failedMsg := "Invalid transaction note"

	var tests = []struct {
		Note  string
		Msg   string
		Desc  string
		Count string //循环次数
	}{
		{"", "", "正常流程", "1"},
		{"a", "", "正常流程", "1"},
		{common.RandomTool("list1", r.Intn(253+2)), "", "正常流程", "20"},
		{common.RandomTool("list1", 256), "", "正常流程", "20"},
		{common.RandomTool("list1", 257), failedMsg, "note太长，最大256byte", "20"},
		{common.RandomTool("list1", r.Intn(1000)+257), failedMsg, "note太长，最大256byte", "20"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForNote")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(name, password, token, gasLimit, test.Note, to, value, keystorePath, chainID)
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------ChainID--------------------------------------------------------
func TestTranferForChainID(t *testing.T) {

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
	common.PrintTitle("TestTranferForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(name, password, token, gasLimit, note, to, value, keystorePath, test.ChainID)
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------KeyStorePath--------------------------------------------------------
func TestTranferForKeyStorePath(t *testing.T) {

	failedMsg := "KeyStorePath does not exist"

	path := make([]string, 0)
	for m := 0; m < r.Intn(5)+1; m++ {
		path = append(path, common.RandomTool("list5", r.Intn(10)+1))
	}
	Path1 := strings.Join(path, "/")

	var tests = []struct {
		KeyStorePath string
		Msg          string
		Desc         string
		Count        string //循环次数
	}{
		{"", "", "正常流程", "1"},
		{Path1, failedMsg, "钱包路径不存在", "20"},
		{"/home/lzw/svn/code/v2.0/bctoolbox/src/unitest/bcctest/all_test/", "", "正常流程", "1"},
		{".keystore/", "", "正常流程", "1"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForKeyStorePath")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(name, password, token, gasLimit, note, to, value, test.KeyStorePath, chainID)
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------Token--------------------------------------------------------
func TestTranferForKeyToken(t *testing.T) {

	failedMsg := "return value is empty, please check key(Token)"

	var tests = []struct {
		Token string
		Msg   string
		Desc  string
		Count string //循环次数
	}{
		{"", failedMsg, "token为空", "1"},
		{"LOC", "", "正常流程", "1"},
		{"xyz", failedMsg, "非法代币", "1"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForKeyStorePath")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(name, password, test.Token, gasLimit, note, to, value, keystorePath, chainID)
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------To--------------------------------------------------------
func TestTranferForTo(t *testing.T) {

	//failedMsg := "Address cannot be emtpy"
	failedMsg2 := "Invalid Address Format"

	var tests = []struct {
		To    string
		Msg   string
		Desc  string
		Count string //循环次数
	}{
		//{"", failedMsg, "地址为空", "1"},
		{to, "", "正常流程", "1"},
		{"qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j", failedMsg2, "非法地址-chainID", "1"},
		{"qqqqL9BzYNYns5dfgdfgfHEBJLzS1bhpHjx7j", failedMsg2, "非法地址-base58", "1"},
		{"qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHssss", failedMsg2, "非法地址-4位验证码", "1"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForTo")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(name, password, token, gasLimit, note, test.To, value, keystorePath, chainID)
			if err != nil && test.Msg == "" {
				panic(err)
			}
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------Value--------------------------------------------------------
func TestTranferForValue(t *testing.T) {

	failedMsg := "Value cannot be empty!"
	failedMsg2 := "Value has to be a positive integer!"

	var tests = []struct {
		Value string
		Msg   string
		Desc  string
		Count string //循环次数
	}{
		{"", failedMsg, "金额为空", "1"},
		{"0", failedMsg2, "金额为0", "1"},
		{strconv.Itoa(r.Intn(10000000) + 1), "", "正常流程", "20"},
		{"-" + strconv.Itoa(r.Intn(10000000)+1), failedMsg2, "金额为负数", "20"},
		{common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000000)+1), failedMsg2, "非法金额", "20"},
		{strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg2, "非法金额", "20"},
		{strconv.Itoa(r.Intn(10000000)+1) + common.RandomTool("list9", 1), failedMsg2, "非法金额", "20"},
		{"-" + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000000)+1), failedMsg2, "非法金额", "20"},
		{"-" + strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg2, "非法金额", "20"},
		{"-" + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1) + strconv.Itoa(r.Intn(10000)+1), failedMsg2, "非法金额", "20"},
		{"-" + strconv.Itoa(r.Intn(10000000)+1) + common.RandomTool("list9", 1), failedMsg2, "非法金额", "20"},
	}

	success := common.Success
	common.PrintTitle("TestTranferForValue")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.Transfer(name, password, token, gasLimit, note, to, test.Value, keystorePath, chainID)
			if err != nil && test.Msg == "" {
				panic(err)
			}
			common.PrintResult(err)
			jsIndent1 := []byte{}
			jsIndent2 := []byte{}
			if result != nil {
				jsIndent1, _ = json.MarshalIndent(&result.Code, "", "\t")
				jsIndent2, _ = json.MarshalIndent(&result.TxHash, "", "\t")
			}
			common.PrintCase(index, test.Desc, test.Msg, "Code:"+string(jsIndent1)+"\t TxHash:"+string(jsIndent2), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}
