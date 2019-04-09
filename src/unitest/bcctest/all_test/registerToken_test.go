package all

import (
	"fmt"
	"strconv"
	"strings"
	"unitest/bcctest/common"

	"testing"
)

//---------------------------------------------------Account--------------------------------------------------------
func TestRegisterTokenForAccount(t *testing.T) {
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
	common.PrintTitle("TestRegisterTokenForAccount")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(test.Name, test.Password, tokenName, tokenSymbol, totalSupply, gasprice, gasLimit, note, keystorePath, chainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------gasLimit--------------------------------------------------------
func TestRegisterTokenForGasLimit(t *testing.T) {

	failedMsg1 := "gasLimit cannot be emtpy"
	failedMsg3 := "Gas Limit is not enough"
	failedMsg4 := "GasLimit has to be a positive integer!"

	var tests = []struct {
		GasLimit string
		Msg      string
		Desc     string
		Count    string //循环次数
	}{
		{"", failedMsg1, "gasLimit为空", "1"},
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
	common.PrintTitle("TestRegisterTokenForGasLimit")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, gasprice, test.GasLimit, note, keystorePath, chainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------Note--------------------------------------------------------
func TestRegisterTokenForNote(t *testing.T) {

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
	common.PrintTitle("TestRegisterTokenForNote")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, gasprice, gasLimit, test.Note, keystorePath, chainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------ChainID--------------------------------------------------------
func TestRegisterTokenForChainID(t *testing.T) {

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
		{chainid[r.Intn(13)], "", "正常流程", "1"},
		{common.RandomTool("list5", r.Intn(100)+1), failedMsg, "无效的chainID", "20"},
	}

	success := common.Success
	common.PrintTitle("TestRegisterTokenForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, gasprice, gasLimit, note, keystorePath, test.ChainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------KeyStorePath--------------------------------------------------------
func TestRegisterTokenForKeyStorePath(t *testing.T) {

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
		{Path1, failedMsg, "钱包地址不存在", "20"},
		{"/home/lzw/svn/code/v2.0/bctoolbox/src/unitest/bcctest/all_test/", "", "正常流程", "1"},
		{".keystore/", "", "正常流程", "1"},
	}

	success := common.Success
	common.PrintTitle("TestRegisterTokenForKeyStorePath")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, gasprice, gasLimit, note, test.KeyStorePath, chainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------TokenName--------------------------------------------------------
func TestRegisterTokenForTokenName(t *testing.T) {

	failedMsg2 := "tokenName cannot be emtpy"
	failedMsg3 := "Invalid tokenName."

	TokenName1 := common.RandomTool("list3", 1)
	TokenName2 := common.RandomTool("list3", 2)
	TokenName3 := common.RandomTool("list3", 3)
	TokenName4 := common.RandomTool("list3", r.Intn(36)+4)
	TokenName5 := common.RandomTool("list3", 40)
	TokenName6 := common.RandomTool("list3", r.Intn(200)+41)

	var tests = []struct {
		TokenName string
		Msg       string
		Desc      string
		Count     string //循环次数
	}{
		{"", failedMsg2, "代币名为空", "1"},
		{TokenName1, failedMsg3, "名称太短", "1"},
		{TokenName2, failedMsg3, "名称太短", "1"},
		{TokenName3, "", "正常流程", "1"},
		{TokenName4, "", "正常流程", "1"},
		{TokenName5, "", "正常流程", "1"},
		{TokenName6, failedMsg3, "名称太长", "1"},
	}

	success := common.Success
	common.PrintTitle("TestRegisterTokenForTokenName")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, test.TokenName, tokenSymbol, totalSupply, gasprice, gasLimit, note, keystorePath, chainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------TokenSymbol--------------------------------------------------------
func TestRegisterTokenForTokenSymbol(t *testing.T) {

	failedMsg2 := "TokenSymbol cannot be emtpy"
	failedMsg3 := "Invalid TokenSymbol."

	var tests = []struct {
		TokenSymbol string
		Msg         string
		Desc        string
		Count       string //循环次数
	}{
		{"", failedMsg2, "代币符号为空", "1"},
		{"xtt", "", "正常流程", "1"},
		{common.RandomTool("list3", 1), "", "正常流程", "20"},
		{common.RandomTool("list3", r.Intn(18)+2), "", "正常流程", "20"},
		{common.RandomTool("list3", 20), "", "正常流程", "20"},
		{common.RandomTool("list3", r.Intn(200)+21), failedMsg3, "代币符号太长", "20"},
	}

	success := common.Success
	common.PrintTitle("TestRegisterTokenForTokenSymbol")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, test.TokenSymbol, totalSupply, gasprice, gasLimit, note, keystorePath, chainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------GasPrice--------------------------------------------------------
func TestRegisterTokenForGasPrice(t *testing.T) {

	failedMsg := "Invalid gas price"

	var tests = []struct {
		GasPrice string
		Msg      string
		Desc     string
		Count    string //循环次数
	}{
		{"", failedMsg, "GasPrice为空", "1"},
		{"0", failedMsg, "GasPrice为0", "1"},
		{"999999999", "", "正常流程", "1"},
		{"1000000001", failedMsg, "GasPrice数值过大", "1"},

		{strconv.Itoa(r.Intn(999997500) + 1), "", "正常流程", "20"},
		{strconv.Itoa(r.Intn(2500) + 1), failedMsg, "GasPrice数值过小", "20"},
		{"-" + strconv.Itoa(r.Intn(1000000000)+1), failedMsg, "非正整数，数据类型错误", "20"},

		{common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(999997500)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(1000000000)+1) + common.RandomTool("list9", 1), failedMsg, "非正整数，数据类型错误", "20"},

		{common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(2500)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(2500)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(2500)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(2500)+1) + common.RandomTool("list9", 1), failedMsg, "非正整数，数据类型错误", "20"},

		{"-" + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(999997500)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(999997500)+1) + common.RandomTool("list9", 1), failedMsg, "非正整数，数据类型错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestRegisterTokenForGasPrice")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, test.GasPrice, gasLimit, note, keystorePath, chainID, addSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------AddSupplyEnabled --------------------------------------------------------
func TestRegisterTokenForAddSupplyEnabled(t *testing.T) {

	failedMsg := "invalid AddSupplyEnabled"

	var tests = []struct {
		AddSupplyEnabled string
		Msg              string
		Desc             string
		Count            string //循环次数
	}{
		{"", failedMsg, "AddSupplyEnabled为空", "1"},
		{"true", "", "正常流程", "1"},
		{"false", "", "正常流程", "1"},
		{common.RandomTool("list3", r.Intn(100)), failedMsg, "AddSupplyEnabled数据错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestRegisterTokenForAddSupplyEnabled")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, gasprice, gasLimit, note, keystorePath, chainID, test.AddSupplyEnabled, burnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------BurnEnabled --------------------------------------------------------
func TestRegisterTokenForBurnEnabled(t *testing.T) {

	failedMsg := "invalid BurnEnabled"

	var tests = []struct {
		BurnEnabled string
		Msg         string
		Desc        string
		Count       string //循环次数
	}{
		{"", failedMsg, "BurnEnabled为空", "1"},
		{"true", "", "正常流程", "1"},
		{"false", "", "正常流程", "1"},
		{common.RandomTool("list3", r.Intn(100)), failedMsg, "BurnEnabled数据错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestRegisterTokenForBurnEnabled")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			tokenName := common.RandomTool("list3", 3)
			tokenSymbol := common.RandomTool("list3", 3)
			result, err := common.RegisterToken(name, password, tokenName, tokenSymbol, totalSupply, gasprice, gasLimit, note, keystorePath, chainID, addSupplyEnabled, test.BurnEnabled)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("TokenAddress : %s,\t", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}
