package all

import (
	"fmt"
	"strconv"
	"strings"
	"unitest/bcctest/common"

	"testing"
)

var (
	Address   = "localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j"
	base58Str = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

//---------------------------------------------------Account--------------------------------------------------------
func TestSetOrgDeployerForAccount(t *testing.T) {
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
	common.PrintTitle("TestSetOrgDeployerForAccount")
	fmt.Println("\n非链Owner调用")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.SetOrgDeployer(test.Name, test.Password, orgName, deployer, gasLimit, note, keystorePath, chainID)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("Code : %s", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------gasLimit--------------------------------------------------------
func TestSetOrgDeployerForGasLimit(t *testing.T) {

	failedMsg1 := "gasLimit cannot be emtpy"
	failedMsg2 := "success or failed"
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
		{strconv.Itoa(r.Intn(1000000000) + 1), failedMsg2, "成功或者失败", "20"},
		{"-" + strconv.Itoa(r.Intn(1000000000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(1000000000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{strconv.Itoa(r.Intn(1000000000)+1) + common.RandomTool("list9", 1), failedMsg4, "非正整数，数据类型错误", "20"},
		{"-" + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(1000000000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(10000)+1) + common.RandomTool("list9", 1) + strconv.Itoa(r.Intn(10000)+1), failedMsg4, "非正整数，数据类型错误", "20"},
		{"-" + strconv.Itoa(r.Intn(1000000000)+1) + common.RandomTool("list9", 1), failedMsg4, "非正整数，数据类型错误", "20"},
	}

	success := common.Success
	common.PrintTitle("TestSetOrgDeployerForGasLimit")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.SetOrgDeployer(ownerName, ownerPassword, orgName, deployer, test.GasLimit, note, "", chainID)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("Code : %s", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------Note--------------------------------------------------------
func TestSetOrgDeployerForNote(t *testing.T) {

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
	common.PrintTitle("TestSetOrgDeployerForNote")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.SetOrgDeployer(ownerName, ownerPassword, orgName, deployer, gasLimit, test.Note, keystorePath, chainID)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("Code : %s", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------ChainID--------------------------------------------------------
func TestSetOrgDeployerForChainID(t *testing.T) {

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
	common.PrintTitle("TestSetOrgDeployerForChainID")
	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.SetOrgDeployer(ownerName, ownerPassword, orgName, deployer, gasLimit, note, keystorePath, test.ChainID)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("Code : %s", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------KeyStorePath--------------------------------------------------------
func TestSetOrgDeployerForKeyStorePath(t *testing.T) {

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
	common.PrintTitle("TestSetOrgDeployerForKeyStorePath")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.SetOrgDeployer(ownerName, ownerPassword, orgName, deployer, gasLimit, note, test.KeyStorePath, chainID)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("Code : %s", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}

//---------------------------------------------------orgName-------------------------------------------------------
func TestSetOrgDeployerForOrgName(t *testing.T) {

	failedMsg1 := "orgName cannot be emtpy"
	failedMsg2 := "Invalid orgName"

	var tests = []struct {
		OrgName string
		Msg     string
		Desc    string
		Count   string //循环次数
	}{
		{common.OrgNameOfTest("", 0), failedMsg1, "组织名为空", "1"},
		{common.OrgNameOfTest("中文组织名", 0), "", "正常流程", "1"},
		{common.OrgNameOfTest("list3", 1), "", "正常流程", "1"},
		{common.OrgNameOfTest("list3", 256), "", "正常流程", "1"},
		{common.OrgNameOfTest("list3", 257), failedMsg2, "长度错误", "1"},
		{common.OrgNameOfTest("list3", r.Intn(254)+2), "", "正常流程", "100"},
	}

	success := common.Success
	common.PrintTitle("TestSetOrgDeployerForOrgName")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.SetOrgDeployer(ownerName, ownerPassword, test.OrgName, deployer, gasLimit, note, keystorePath, chainID)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("Code : %s", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------Deployer-------------------------------------------------------
func TestSetOrgDeployerForDeployer(t *testing.T) {

	failedMsg1 := "Deployer cannot be emtpy"
	failedMsg2 := "Invalid orgName"

	var tests = []struct {
		Deployer string
		Msg      string
		Desc     string
		Count    string //循环次数
	}{
		{"", failedMsg1, "地址为空", "1"},
		{Address, "", "正常流程", "1"},
		{"sadfasdfawersd23356", failedMsg2, "非法地址", "1"},
	}

	success := common.Success
	common.PrintTitle("TestSetOrgDeployerForDeployer")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			result, err := common.SetOrgDeployer(ownerName, ownerPassword, orgName, test.Deployer, gasLimit, note, keystorePath, chainID)
			common.PrintResult(err)
			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("Code : %s", result), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}
