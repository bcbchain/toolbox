package all

import (
	"fmt"
	"strconv"
	"unitest/bcctest/common"

	"testing"
)

//---------------------------------------------------orgName------------------------------------------------------------
func TestContractInfoForOrgName(t *testing.T) {
	failedMsg3 := "Insufficient input parameters"

	var tests = []struct {
		OrgName      string
		contractName string
		orgID        string
		contractAddr string

		Msg   string
		Desc  string
		Count string //循环次数
	}{
		{"", "", "", "", failedMsg3, "无参数", "1"},
		{"中文组织名", "", "", "", failedMsg3, "参数不足", "1"},
		{common.RandomTool("list10", 1), "", "", "", failedMsg3, "参数不足", "1"},
		{common.RandomTool("list10", 256), "", "", "", failedMsg3, "参数不足", "1"},
		{common.RandomTool("list10", 257), "", "", "", failedMsg3, "参数不足", "1"},
		{common.RandomTool("list10", r.Intn(254)+2), "", "", "", failedMsg3, "参数不足", "20"},

		{"", "", "", contractAddr, "", "正常流程1", "1"},
		{"中文组织名", "", "", contractAddr, "", "正常流程2", "1"},
		{common.RandomTool("list10", 1), "", "", contractAddr, "", "正常流程3", "1"},
		{common.RandomTool("list10", 256), "", "", contractAddr, "", "正常流程4", "1"},
		{common.RandomTool("list10", 257), "", "", contractAddr, "", "正常流程5", "1"},
		{common.RandomTool("list10", r.Intn(254)+2), "", "", contractAddr, "", "正常流程6", "20"},

		{"", contractName, orgID, "", "", "正常流程1", "1"},
		{"中文组织名", contractName, orgID, "", "", "正常流程2", "1"},
		{common.RandomTool("list10", 1), contractName, orgID, "", "", "正常流程3", "1"},
		{common.RandomTool("list10", 256), contractName, orgID, "", "", "正常流程4", "1"},
		{common.RandomTool("list10", 257), contractName, orgID, "", "", "正常流程5", "1"},
		{common.RandomTool("list10", r.Intn(254)+2), contractName, orgID, "", "", "正常流程6", "20"},
	}

	success := common.Success
	common.PrintTitle("TestContractInfoForOrgName")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			err := common.ContractInfo(test.OrgName, test.contractName, test.orgID, test.contractAddr)
			common.PrintResult(err)

			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("complete call"), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------ContractName-------------------------------------------------------
func TestContractInfoForContractName(t *testing.T) {

	failedMsg2 := "Invalid ContractName"
	failedMsg3 := "Insufficient input parameters"

	var tests = []struct {
		OrgName      string
		contractName string
		orgID        string
		contractAddr string

		Msg   string
		Desc  string
		Count string //循环次数
	}{
		{"", "", "", "", failedMsg3, "无参数", "1"},
		{"", "中文合约名", "", "", failedMsg3, "参数不足", "1"},
		{"", common.RandomTool("list10", 1), "", "", failedMsg3, "参数不足", "20"},
		{"", common.RandomTool("list10", 256), "", "", failedMsg3, "参数不足", "1"},
		{"", common.RandomTool("list10", 257), "", "", failedMsg3, "参数不足", "1"},
		{"", common.RandomTool("list10", r.Intn(254)+2), "", "", failedMsg3, "参数不足", "20"},

		{orgName, "", "", "", "", "参数不足", "1"},
		{orgName, "中文合约名", "", "", "", "正常流程", "1"},
		{orgName, common.RandomTool("list10", 1), "", "", "", "正常流程", "20"},
		{orgName, common.RandomTool("list10", 256), "", "", "", "正常流程", "1"},
		{orgName, common.RandomTool("list10", 257), "", "", failedMsg2, "合约名称长度错误", "1"},
		{orgName, common.RandomTool("list10", r.Intn(254)+2), "", "", "", "正常流程", "20"},

		{"", "", "", contractAddr, "", "正常流程", "1"},
		{"", "中文合约名", "", contractAddr, "", "正常流程", "1"},
		{"", common.RandomTool("list10", 1), "", contractAddr, "", "正常流程", "20"},
		{"", common.RandomTool("list10", 256), "", contractAddr, "", "正常流程", "1"},
		{"", common.RandomTool("list10", 257), "", contractAddr, "", "正常流程", "1"},
		{"", common.RandomTool("list10", r.Intn(254)+2), "", contractAddr, "", "正常流程", "20"},

		{"", "", orgID, "", "", "参数不足", "1"},
		{"", "中文合约名", orgID, "", "", "正常流程", "1"},
		{"", common.RandomTool("list10", 1), orgID, "", "", "正常流程", "20"},
		{"", common.RandomTool("list10", 256), orgID, "", "", "正常流程", "1"},
		{"", common.RandomTool("list10", 257), orgID, "", failedMsg2, "合约名称长度错误", "1"},
		{"", common.RandomTool("list10", r.Intn(254)+2), orgID, "", "", "正常流程", "20"},
	}

	success := common.Success
	common.PrintTitle("TestContractInfoForContractName")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			err := common.ContractInfo(test.OrgName, test.contractName, test.orgID, test.contractAddr)
			common.PrintResult(err)

			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("complete call"), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------OrgID--------------------------------------------------------------
func TestContractInfoForOrgID(t *testing.T) {
	failedMsg1 := "Insufficient input parameters"
	failedMsg2 := "Address chainid is error!"
	failedMsg3 := "Base58Addr parse error!"
	failedMsg4 := "Address checksum is error!"

	var tests = []struct {
		OrgName      string
		contractName string
		orgID        string
		contractAddr string

		Msg   string
		Desc  string
		Count string //循环次数
	}{
		{"", "", "", "", "", "正常流程", "1"},
		{"", "", orgID, "", failedMsg1, "参数不足", "1"},
		{"", "", "oppBtjfCSPCAJ84uQWcpNr74NLMWYm5SXzer", "", failedMsg1, "参数不足", "1"},
		{"", "", "orgBtjfCSPCAJ84uWYm5SXzer", "", failedMsg1, "参数不足", "1"},
		{"", "", "orgBtjfCSPCAJ84uQWcpNr74NLMWYm5Sqqqq", "", failedMsg1, "参数不足", "1"},

		{orgName, "", "", "", failedMsg1, "参数不足", "1"},
		{orgName, "", orgID, "", failedMsg1, "参数不足", "1"},
		{orgName, "", "oppBtjfCSPCAJ84uQWcpNr74NLMWYm5SXzer", "", failedMsg1, "参数不足", "1"},
		{orgName, "", "orgBtjfCSPCAJ84uWYm5SXzer", "", failedMsg1, "参数不足", "1"},
		{orgName, "", "orgBtjfCSPCAJ84uQWcpNr74NLMWYm5Sqqqq", "", failedMsg1, "参数不足", "1"},

		{"", contractName, "", "", failedMsg1, "参数不足", "1"},
		{"", contractName, orgID, "", "", "正常流程", "1"},
		{"", contractName, "oppBtjfCSPCAJ84uQWcpNr74NLMWYm5SXzer", "", failedMsg2, "非法OrgID-chainID", "1"},
		{"", contractName, "orgBtjfCSPCAJ84uWYm5SXzer", "", failedMsg3, "非法OrgID-base8", "1"},
		{"", contractName, "orgBtjfCSPCAJ84uQWcpNr74NLMWYm5Sqqqq", "", failedMsg4, "非法OrgID-4位验证码", "1"},

		{"", "", "", contractAddr, "", "正常流程", "1"},
		{"", "", orgID, contractAddr, "", "正常流程", "1"},
		{"", "", "oppBtjfCSPCAJ84uQWcpNr74NLMWYm5SXzer", contractAddr, failedMsg2, "非法OrgID-chainID", "1"},
		{"", "", "orgBtjfCSPCAJ84uWYm5SXzer", contractAddr, failedMsg3, "非法OrgID-base8", "1"},
		{"", "", "orgBtjfCSPCAJ84uQWcpNr74NLMWYm5Sqqqq", contractAddr, failedMsg4, "非法OrgID-4位验证码", "1"},
	}

	success := common.Success
	common.PrintTitle("TestContractInfoForOrgID")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			err := common.ContractInfo(test.OrgName, test.contractName, test.orgID, test.contractAddr)
			common.PrintResult(err)

			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("complete call"), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))
		}
	}
}

//---------------------------------------------------ContractAddr--------------------------------------------------------
func TestContractInfoForContractAddr(t *testing.T) {
	failedMsg1 := "Insufficient input parameters"
	failedMsg2 := "Invalid Address Format"

	var tests = []struct {
		OrgName      string
		contractName string
		orgID        string
		contractAddr string

		Msg   string
		Desc  string
		Count string //循环次数
	}{
		{"", "", "", "", "", "正常流程", "1"},
		{"", "", "", contractAddr, "", "正常流程", "1"},
		{"", "", "", "qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j", failedMsg2, "非法地址-chainID", "1"},
		{"", "", "", "local2H5hxpzfuDqvgwe7GyDEGP", failedMsg2, "非法地址-base58", "1"},
		{"", "", "", "local2H5hxpzfuDqvd2uc8DqhF5sgwe7G11111", failedMsg2, "非法地址-4位验证码", "1"},

		{orgName, "", "", "", failedMsg1, "参数不足", "1"},
		{orgName, "", "", contractAddr, "", "正常流程", "1"},
		{orgName, "", "", "qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j", failedMsg2, "非法地址-chainID", "1"},
		{orgName, "", "", "local2H5hxpzfuDqvgwe7GyDEGP", failedMsg2, "非法地址-base58", "1"},
		{orgName, "", "", "local2H5hxpzfuDqvd2uc8DqhF5sgwe7G11111", failedMsg2, "非法地址-4位验证码", "1"},

		{"", contractName, orgID, "", "", "正常流程", "1"},
		{"", contractName, orgID, contractAddr, "", "正常流程", "1"},
		{"", contractName, orgID, "qqqqL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j", failedMsg2, "非法地址-chainID", "1"},
		{"", contractName, orgID, "local2H5hxpzfuDqvgwe7GyDEGP", failedMsg2, "非法地址-base58", "1"},
		{"", contractName, orgID, "local2H5hxpzfuDqvd2uc8DqhF5sgwe7G11111", failedMsg2, "非法地址-4位验证码", "1"},
	}

	success := common.Success
	common.PrintTitle("TestContractInfoForContractAddr")

	for index, test := range tests {
		num, _ := strconv.Atoi(test.Count)
		for i := 0; i < num; i++ {
			err := common.ContractInfo(test.OrgName, test.contractName, test.orgID, test.contractAddr)
			common.PrintResult(err)

			common.PrintCase(index, test.Desc, test.Msg, fmt.Sprintf("complete call"), err)

			fmt.Println(test.Desc)
			common.AssertSuccess(common.Success-success, len(tests))

		}
	}
}
