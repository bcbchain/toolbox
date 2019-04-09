package common

import (
	"blockchain/abciapp_v1.0/types"
	"cmd/bcc/cache"
	"cmd/bcc/core"
	"common/wal"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os/exec"
	"strconv"

	"cmd/bcc/common"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	geneAddress     = ""
	geneName        = ""
	CreatePassword  = "qingzhong361001G@"
	PasswordOfOwner = "Ab1@Cd3$"
	Success         = 0
	reportFile      = ""
	Total           = 0
	Failed          = 0
	hexStr          = "abcdefghijlkmnopqrstuvwxyz9876543210"
	privKey         = "0x832c3477b3a730e9601ed774a88c27ca43112ff7d0718686a33c21721dd3369701bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e"
	//random num
	lowerList       = "abcdefghijklmnopqrstuvwxyz"
	upperList       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numList         = "0123456789"
	SpecList        = "@._-"
	charList        = lowerList + upperList + numList + SpecList
	IllegalCharList = ` !"#$%&'()*+,/:;<=>?[]\^{|}~`
	BackQuote       = "`"

	ownerAddr = ""
)

// Get test Lists
func GetLists() [][]byte {
	// 各类List定义
	var Lists [][]byte
	var List1 = []byte{} //[a-z, A-Z, 0-9，"_", "@", "." and "-"]
	var List2 = []byte{} //not [a-z, A-Z, 0-9，"_", "@", "." and "-"]
	var List3 = []byte{} //可打印字符     len=96
	var List4 = []byte{} //不可打印字符   len=160

	var List5 = []byte{} //小写字母       len=26
	var List6 = []byte{} //大写字母		len=26
	var List7 = []byte{} //数字			len=10
	var List8 = []byte{} //可打印特殊字符  len=34

	var List9 = []byte{}  // 不包含数字的可打印字符
	var List10 = []byte{} //ASCII 字符
	var List11 = []byte{} //16进制字符

	for i := 65; i < 91; i++ {
		List1 = append(List1, byte(i))
	}
	for j := 97; j < 123; j++ {
		List1 = append(List1, byte(j))
	}
	List1 = append(List1, 45, 46, 95, 64)

	for z := 0; z < 256; z++ {
		for _, v := range List1 {
			if v == byte(z) {
				break
			}
		}
		List2 = append(List2, byte(z))
	}

	for m := 32; m < 128; m++ {
		List3 = append(List3, byte(m))
	}
	for m := 0; m < 32; m++ {
		List4 = append(List4, byte(m))
	}
	for m := 128; m < 256; m++ {
		List4 = append(List4, byte(m))
	}

	for m := 97; m < 123; m++ {
		List5 = append(List5, byte(m))
	}
	for m := 65; m < 91; m++ {
		List6 = append(List6, byte(m))
	}
	for m := 48; m < 58; m++ {
		List7 = append(List7, byte(m))
	}

	for m := 32; m < 48; m++ {
		List8 = append(List8, byte(m))
	}
	for m := 58; m < 64; m++ {
		List8 = append(List8, byte(m))
	}
	for m := 91; m < 96; m++ {
		List8 = append(List8, byte(m))
	}
	for m := 123; m < 127; m++ {
		List8 = append(List8, byte(m))
	}
	for m := 0; m < len(hexStr); m++ {
		List11 = append(List11, byte(hexStr[m]))
	}

	Lists = append(Lists, List1)
	Lists = append(Lists, List2)
	Lists = append(Lists, List3)
	Lists = append(Lists, List4)
	Lists = append(Lists, List5)
	Lists = append(Lists, List6)
	Lists = append(Lists, List7)
	Lists = append(Lists, List8)

	List9 = append(List9, List5...)
	List9 = append(List9, List6...)
	List9 = append(List9, List8...)
	Lists = append(Lists, List9)

	List10 = append(List3, List4...)
	Lists = append(Lists, List10)

	Lists = append(Lists, List11)

	return Lists
}

func GenesisToken() (string, string) {
	if geneAddress == "" {
		key := "/genesis/token"

		token := new(IssueToken)

		err := core.DoHttpQueryAndParse(nodeAddrSlice(chainID), key, token)
		if err != nil {
			fmt.Println(err.Error())
			return geneAddress, geneName
		}

		geneAddress = token.Address
		geneName = token.Name
	}

	return geneAddress, geneName
}

func NameofTest(list string, num int) string {
	testName := RandomTool(list, num)
	_, err := Accounts(testName, 1)
	if err != nil {
		return err.Error()
	}
	return testName
}

func PassWordOfTest(name, password string) (Name string, err error) {

	acct, err := wal.NewAccount(keyStorePath, name, password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = Transfer("owner", "Ab1@Cd3$", "LOC", "100000", "", acct.Address(chainID), "100000000", "", "")
	if err != nil {
		fmt.Println(err.Error())
	}

	return acct.Name, err
}

// Test name of SetOrgSigners
func TestNameofSetOrgSigners(list string, num int, orgName string) string {
	testName := RandomTool(list, num)
	Name, err := PassWordOfTest(testName, password)
	if err != nil {
		return err.Error()
	}
	_, err = RegisterOrg(Name, password, orgName, "1000000", "", "", "")
	if err != nil {
		fmt.Println(err.Error())
	}

	return Name
}

// Test name of DeployContract
func TestNameofDeployContract(list string, num int) string {

	acct, err := wal.NewAccount(keyStorePath, RandomTool(list, num), password)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	_, err = Transfer("owner", "Ab1@Cd3$", "LOC", "100000", "", acct.Address(chainID), "100000000", "", "")
	if err != nil {
		fmt.Println(err.Error())
	}

	orgNameTest := RandomTool("list3", 5)
	_, err = RegisterOrg(acct.Name, password, orgNameTest, "1000000", "", "", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	code, err := SetOrgDeployer("owner", "Ab1@Cd3$", orgNameTest, acct.Address(chainID), "1000000", "", "", "")
	if code != "200" {
		fmt.Println("授权合约发布者失败")
	}
	return acct.Name
}

//privateKey中保存accessKey
func Accounts(name string, count int) (accounts []wal.Account, err error) {
	defer FuncRecover(&err)
	accounts = make([]wal.Account, 0)
	acct, err := wal.NewAccount(keyStorePath, name, CreatePassword)
	if err != nil {
		return
	}
	_, err = Transfer("owner", "Ab1@Cd3$", "loc", "100000", "", acct.Address(chainID), "100000000", "", "")

	acct1 := wal.Account{Name: acct.Name, PrivateKey: acct.PrivateKey, Hash: acct.Hash, KeyStoreFile: acct.KeyStoreFile}
	accounts = append(accounts, acct1)
	return
}

// 获取chainID
func NodeAddrSliceToTest(chainID string) []string {
	if len(chainID) == 0 {
		chainID = common.GetBCCConfig().DefaultChainID
	}

	switch chainID {
	case "bcb":
		return common.GetBCCConfig().Bcb
	case "bcbtest":
		return common.GetBCCConfig().Bcbtest
	case "devtest":
		return common.GetBCCConfig().Devtest
	case "local":
		return common.GetBCCConfig().Local
	default:
		return []string{}
	}
}

// 获取当前文件路径
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	_, err = filepath.Abs(file)
	if err != nil {
	}
	return filepath.Abs(file)
}

// 打印测试标题
func PrintTitle(funcName string) {
	inter := strings.Replace(funcName, "Test", "", -1)
	inter = "BCC_" + fmt.Sprintf("%c", inter[0]+0x20) + inter[1:]

	starS := "\n******************************************************************"
	title := fmt.Sprintf("\nTEST CASE: 测试接口%s -- %s\n", inter, funcName)
	PrintAndWriteToFile(starS)
	PrintAndWriteToFile(title)
}

// 打印、写入日志
func PrintAndWriteToFile(caseOutput string) error {

	fmt.Printf("%s", caseOutput)
	if reportFile == "" {
		err := InitReportFile()
		if err != nil {
			return err
		}
	}
	fd, err := os.OpenFile(reportFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer fd.Close()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = fd.Write([]byte(caseOutput))
	common.GetLogger().Info(caseOutput)
	return nil
}

func InitReportFile() error {
	outDir := "./test"
	if err := os.MkdirAll(outDir, os.FileMode(0750)); err != nil {
		panic(err)
	}

	reportFile = filepath.Join(outDir, "report_"+Version1()+"_"+time.Now().Format("20060102150405"))

	fd, err := os.OpenFile(reportFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer fd.Close()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	_, err = fd.Write([]byte("Running Test"))

	return err
}

// 获取版本信息
func Version1() string {

	result, err := core.Version()
	if err != nil {
		fmt.Println("查询版本信息错误")
		return ""
	}

	return result.Version
}

func PrintResult(err error) {
	Total += 1
	var result string
	if err == nil {
		result = fmt.Sprintf("PASS")
		Success += 1
	} else {
		result = fmt.Sprintf("FAIL")
		Failed += 1
	}
	PrintAndWriteToFile(result)
}

func PrintCase(index int, desc, msg, successDesc string, err error) {
	caseOut := fmt.Sprintf(fmt.Sprintf(" -- Case  %2d : %s", index, desc))
	PrintAndWriteToFile(caseOut)
	descLen := utf8.RuneCountInString(desc) + (len(desc)-utf8.RuneCountInString(desc))/2
	printLen := 25 - descLen
	var spaceStr string
	for printLen > 0 {
		spaceStr += " "
		printLen--
	}
	PrintAndWriteToFile(spaceStr)

	var expectedResult string
	if len(msg) == 0 {
		expectedResult = fmt.Sprintf("\t预期：成功 ")
	} else {
		expectedResult = fmt.Sprintf("\t预期：失败 ")
	}
	PrintAndWriteToFile(expectedResult)

	var lineTail string
	if err != nil {
		lineTail = fmt.Sprintf("\t%s\n", err.Error())
	} else {
		lineTail = fmt.Sprintf("\t%s\n", successDesc)
	}
	PrintAndWriteToFile(lineTail)
}

func AssertSuccess(resultS, expectedS int) {
	msg := fmt.Sprintf("\ntotal pass: %d/%d\n\n", resultS, expectedS)
	PrintAndWriteToFile(msg)
}

func GetBlockHeight() uint64 {
	blkHeight, err := core.BlockHeight(chainID)
	if err != nil {
		fmt.Println("查询区块高度错误")
		return 0
	}
	return uint64(blkHeight.LastBlock)
}

func assert(expr bool, key, expectedValue, actualValue interface{}) {
	if expr == false {
		panic(fmt.Sprintf("expected%v: %v,\t actual%v: %v", key, expectedValue, key, actualValue))
	}
}

func assertIsNil(err error, key string) {
	if err != nil {
		panic(fmt.Sprintf("expected%s is Nil,\t actual%s: %s", key, key, err.Error()))
	}
}

// random number
func RandString(length int) string {
	pwCharList := lowerList + upperList + numList + IllegalCharList + SpecList + BackQuote

	password := ""
	if length > 3 {
		temp := ""
		rand.Seed(time.Now().UnixNano())
		index := rand.Int() % len(lowerList)
		temp += lowerList[index : index+1]

		index = rand.Int() % len(upperList)
		temp += upperList[index : index+1]

		index = rand.Int() % len(numList)
		temp += numList[index : index+1]

		printChars := IllegalCharList + SpecList + BackQuote
		index = rand.Int() % len(printChars)
		temp += printChars[index : index+1]

		length -= 4
		count := 4
		for count > 0 {
			if count == 1 {
				password += temp
			} else {
				rand.Seed(time.Now().UnixNano())
				index := rand.Int() % count
				password += temp[index : index+1]
				temp = temp[:index] + temp[index+1:]
			}

			count--
		}
	}

	for length > 0 {
		rand.Seed(time.Now().UnixNano())
		index := rand.Int() % len(pwCharList)
		password += pwCharList[index : index+1]

		length--
	}

	return password
}

func GenesisOwner() string {
	if ownerAddr == "" {
		key := "/genesis/token"

		addr := nodeAddrSlice(chainID)
		token := new(types.IssueToken)

		err := core.DoHttpQueryAndParse(addr, key, token)
		if err != nil {
			fmt.Println(err.Error())
			return ownerAddr
		}

		ownerAddr = token.Owner
	}

	return ownerAddr
}

func CheckErr(errStr, msg string) {
	if (errStr != "" && msg == "") || (errStr == "" && msg != "") {
		panic(errStr)
	}
}

func OrgNameOfTest(list string, num int) string {
	if num == 0 {
		_, err := RegisterOrg(name, password, list, gasLimit, note, keystorePath, chainID)
		if err != nil {
			fmt.Println("register Org of Name err")
			return err.Error()
		}
		return list
	} else {
		orgname := RandomTool(list, num)
		_, err := RegisterOrg(name, password, orgname, gasLimit, note, keystorePath, chainID)
		if err != nil {
			fmt.Println("register Org of Name err")
			return err.Error()
		}
		return orgname
	}
}

// Hexadecimal to decimal
func HexDec(h string) (n int64) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			log.Println("Hexadecimal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}

// get Nonce
func GetNonce(keyStorePath, chainID, name, password string, bNonceErr bool) (nonce uint64, err error) {

	nonce, err = cache.Nonce(name, keyStorePath)
	if err != nil || bNonceErr {
		var nonceResult *core.NonceResult

		nonceResult, err = core.Nonce("", name, password, chainID, keyStorePath)
		if err != nil {
			return
		}
		nonce = nonceResult.Nonce

		err = cache.SetNonce(name, nonce, keyStorePath)
		if err != nil {
			return
		}
	}

	return
}

// TestTokenName of Balance TokenName
func TestTokenName(list string, num int) string {
	TokenName := RandomTool(list, num)
	_, err := RegisterToken(name, password, TokenName, RandomTool("list3", 4), "10000000000000", "2500", "10000000", "", "", "", "true", "true")
	if err != nil {
		fmt.Println("注册测试代币错误")
		return ""
	}
	return TokenName
}
