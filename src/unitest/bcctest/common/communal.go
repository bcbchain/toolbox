package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	// 合法参数定义
	name         = "qingzhong"
	password     = "qingzhong361001G@"
	gasLimit     = "1000000"
	note         = ""
	chainID      = "local"
	keystorePath = ""
	PubKey       = "0x01bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e"
	TokenName    = "xxt"

	// 工具定义
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// List组
	Lists  = GetLists()
	List1  = Lists[0]
	List2  = Lists[1]
	List3  = Lists[2]
	List4  = Lists[3]
	List5  = Lists[4]
	List6  = Lists[5]
	List7  = Lists[6]
	List8  = Lists[7]
	List9  = Lists[8] // len = 86
	List10 = Lists[9]
	List11 = Lists[10]
)

type tAccount struct {
	Name     string
	Password string
	Msg      string
	Desc     string
}

type tgasLimit struct {
	GasLimit string
	Msg      string
	Desc     string
}

type tnote struct {
	Note string
	Msg  string
	Desc string
}

type tChainId struct {
	ChainID string
	Msg     string
	Desc    string
}

type tKeyStorePath struct {
	KeyStorePath string
	Msg          string
	Desc         string
}

func TestOfAccount() []tAccount {
	tests := make([]tAccount, 0)
	failedMsg1 := `Name contains by [letters, numbers, "_", "@", "." and "-"] and length must be [1-40]`
	failedMsg2 := "Password contains by [letters, numbers, ASCII 32 through 127] and length must be [8-20]"
	//----------------------------------------------------name----------------------------------------------------------

	// name = ""
	test0 := tAccount{"", password, failedMsg1, "用户名为空"}
	tests = append(tests, test0)

	// name = list中随机1位（循环20次）
	for i := 0; i < 20; i++ {
		num := r.Intn(66)
		test1 := tAccount{string(List1[num]), password, "", "正常流程"}
		tests = append(tests, test1)
	}

	// name = list中随机40位（循环20次）
	for i := 0; i < 20; i++ {
		var names []string
		for j := 0; j < 40; j++ {
			num := r.Intn(65)
			name := string(List1[num])
			names = append(names, name)
		}
		name := strings.Join(names, "")
		test2 := tAccount{name, password, "", "正常流程"}
		tests = append(tests, test2)
	}

	// name = list中随机2-39位（循环20次）
	for i := 0; i < 20; i++ {
		var names []string
		for j := 0; j < r.Intn(38)+2; j++ {
			num := r.Intn(65)
			name := string(List1[num])
			names = append(names, name)
		}
		name := strings.Join(names, "")
		test3 := tAccount{name, password, "", "正常流程"}
		tests = append(tests, test3)
	}

	// name = 随机出现一个或多个在list2中1-40位（循环20次）
	for i := 0; i < 20; i++ {
		var names []string
		for j := 0; j < r.Intn(40)+1; j++ {
			num := r.Intn(190)
			name := string(List2[num])
			names = append(names, name)
		}
		name := strings.Join(names, "")
		test4 := tAccount{name, password, failedMsg1, "异常流程，非法字符"}
		tests = append(tests, test4)
	}

	// name = List中随机大于40位（循环20次）
	for i := 0; i < 20; i++ {
		var names []string
		for j := 0; j < 300; j++ {
			num := r.Intn(65)
			name := string(List1[num])
			names = append(names, name)
		}
		name := strings.Join(names, "")
		test5 := tAccount{name, password, failedMsg1, "账户名称长度错误"}
		tests = append(tests, test5)
	}

	//--------------------------------------------------password--------------------------------------------------------

	// password = ""
	test00 := tAccount{name, "", failedMsg2, "密码不能为空"}
	tests = append(tests, test00)

	// password =list3中随机1位（循环20次）
	for i := 0; i < 20; i++ {
		num := r.Intn(96)
		test1 := tAccount{name, string(List3[num]), failedMsg2, "密码长度错误"}
		tests = append(tests, test1)
	}

	// password =list3中随机7位（循环20次）
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < 7; j++ {
			num := r.Intn(96)
			pwd := string(List3[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, failedMsg2, "密码长度错误"}
		tests = append(tests, test1)
	}

	// password =满足密码要求随机8位（循环20次）
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < 2; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 2; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 2; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 2; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, "", "正常流程"}
		tests = append(tests, test1)
	}

	// password =满足密码要求随机20位（循环20次）
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < 5; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 5; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 5; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 5; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, "", "正常流程"}
		tests = append(tests, test1)
	}

	// password =满足密码要求随机9-19位（循环20次）
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < r.Intn(3)+2; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+2; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+2; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, "", "正常流程"}
		tests = append(tests, test1)
	}

	// password =不满足密码要求随机8-20位（循环20次）-缺数字
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < r.Intn(3)+2; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, failedMsg2, "不满足密码要求，缺数字"}
		tests = append(tests, test1)
	}

	// password =不满足密码要求随机8-20位（循环20次）-缺小写字母
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+2; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, failedMsg2, "不满足密码要求，缺小写字母"}
		tests = append(tests, test1)
	}

	// password =不满足密码要求随机8-20位（循环20次）-缺大写字母
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+2; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, failedMsg2, "不满足密码要求，缺小写字母"}
		tests = append(tests, test1)
	}

	// password =不满足密码要求随机8-20位（循环20次）-缺特殊字符
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+2; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(3)+3; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, failedMsg2, "不满足密码要求，缺特殊字符"}
		tests = append(tests, test1)
	}

	// password =随机出现一个或多个不在list3中8-20位（循环20次）
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < 2; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 2; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 2; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < 2; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(12)+1; j++ {
			num := r.Intn(160)
			pwd := string(List4[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, failedMsg2, "含有异常（不可打印）字符"}
		tests = append(tests, test1)
	}

	// password = list3中随机大于20位（循环20次）-满足密码要求
	for i := 0; i < 20; i++ {
		var passwords []string
		for j := 0; j < r.Intn(100)+5; j++ {
			num := r.Intn(25)
			pwd := string(List5[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(100)+5; j++ {
			num := r.Intn(25)
			pwd := string(List6[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(100)+5; j++ {
			num := r.Intn(9)
			pwd := string(List7[num])
			passwords = append(passwords, pwd)
		}
		for j := 0; j < r.Intn(100)+6; j++ {
			num := r.Intn(33)
			pwd := string(List8[num])
			passwords = append(passwords, pwd)
		}
		passwordss := strings.Join(passwords, "")
		test1 := tAccount{name, passwordss, failedMsg2, "密码长度错误"}
		tests = append(tests, test1)
	}
	return tests
}

func TestOfgasLimit() []tgasLimit {

	tests := make([]tgasLimit, 0)
	failedMsg11 := "gasLimit cannot be emtpy"
	failedMsg12 := "success or failed"
	failedMsg13 := "Gas Limit is not enough"
	failedMsg14 := "GasLimit has to be a positive integer!"

	// gasLimit=""
	test1 := tgasLimit{"", failedMsg11, "密码长度错误"}
	tests = append(tests, test1)

	// gasLimit="0"
	test2 := tgasLimit{"0", failedMsg13, "gasLimit 不足"}
	tests = append(tests, test2)

	// gasLimit=随机正整数（循环20次）
	for i := 0; i < 20; i++ {
		test1 := tgasLimit{string(r.Intn(1000000000) + 1), failedMsg12, "成功或者失败"}
		tests = append(tests, test1)
	}

	// gasLimit=随机负整数（循环20次）
	for i := 0; i < 20; i++ {
		test1 := tgasLimit{"-" + string(r.Intn(1000000000)+1), failedMsg14, "非正整数，数据类型错误"}
		tests = append(tests, test1)
	}

	// gasLimit = 在正整数前出现一个非数字字符（循环20次）

	for i := 0; i < 20; i++ {
		test1 := tgasLimit{string(List9[r.Intn(86)]) + string(r.Intn(1000000000)+1), failedMsg14, "非正整数，数据类型错误"}
		tests = append(tests, test1)
	}

	// gasLimit = 在正整数中出现一个非数字字符（循环20次）
	for i := 0; i < 20; i++ {
		test1 := tgasLimit{string(r.Intn(10000)+1) + string(List9[r.Intn(86)]) + string(r.Intn(10000)+1), failedMsg14, "非正整数，数据类型错误"}
		tests = append(tests, test1)
	}

	// gasLimit = 在正整数后出现一个非数字字符（循环20次）
	for i := 0; i < 20; i++ {
		test1 := tgasLimit{string(r.Intn(1000000000)+1) + string(List9[r.Intn(86)]), failedMsg14, "非正整数，数据类型错误"}
		tests = append(tests, test1)
	}

	// gasLimit = 在负整数前出现一个非数字字符（循环20次）
	for i := 0; i < 20; i++ {
		test1 := tgasLimit{"-" + string(List9[r.Intn(86)]) + string(r.Intn(1000000000)+1), failedMsg14, "非正整数，数据类型错误"}
		tests = append(tests, test1)
	}

	// gasLimit = 在负整数中出现一个非数字字符（循环20次）
	for i := 0; i < 20; i++ {
		test1 := tgasLimit{"-" + string(r.Intn(10000)+1) + string(List9[r.Intn(86)]) + string(r.Intn(10000)+1), failedMsg14, "非正整数，数据类型错误"}
		tests = append(tests, test1)
	}

	// gasLimit = 在负整数后出现一个非数字字符（循环20次）
	for i := 0; i < 20; i++ {
		test1 := tgasLimit{"-" + string(r.Intn(1000000000)+1) + string(List9[r.Intn(86)]), failedMsg14, "非正整数，数据类型错误"}
		tests = append(tests, test1)
	}

	return tests

}

func TestOfNote() []tnote {

	tests := make([]tnote, 0)

	failedMsg := "Invalid transaction note"

	// note= 0 byte
	testNote1 := tnote{"", "", "正常流程"}
	tests = append(tests, testNote1)

	// note= 1 byte
	testNote2 := tnote{"a", "", "正常流程"}
	tests = append(tests, testNote2)

	// note = 随机2-255 byte （循环20次）
	for i := 0; i < 20; i++ {
		var str string
		for j := 0; j < r.Intn(253)+2; j++ {
			str = str + "a"
		}
		testNote := tnote{str, "", "正常流程"}
		tests = append(tests, testNote)
	}

	// note = 256 byte
	for i := 0; i < 1; i++ {
		var str string
		for j := 0; j < 256; j++ {
			str = str + "a"
		}
		testNote := tnote{str, "", "正常流程"}
		tests = append(tests, testNote)
	}

	// note = 257 byte
	for i := 0; i < 1; i++ {
		var str string
		for j := 0; j < 257; j++ {
			str = str + "a"
		}
		testNote := tnote{str, failedMsg, "note太长，最大256byte"}
		tests = append(tests, testNote)
	}

	// note>257字节的随机串  (循环20次)
	for i := 0; i < 20; i++ {
		var str string
		for j := 0; j < r.Intn(1000)+257; j++ {
			str = str + "a"
		}
		testNote := tnote{str, failedMsg, "note太长，最大256byte"}
		tests = append(tests, testNote)
	}

	return tests

}

func TestOfChainId() []tChainId {

	tests := make([]tChainId, 0)
	failedMsg := "invalid chainId"

	// chainId=“”
	testChainID1 := tChainId{"", "", "正常流程"}
	tests = append(tests, testChainID1)

	//chainId=“配置文件中任意chainId”
	chainid := make([]string, 0)
	chainid = append(chainid, NodeAddrSliceToTest("")...)
	chainid = append(chainid, NodeAddrSliceToTest("bcb")...)
	chainid = append(chainid, NodeAddrSliceToTest("bcbtest")...)
	chainid = append(chainid, NodeAddrSliceToTest("devtest")...)
	chainid = append(chainid, NodeAddrSliceToTest("local")...)
	testChainID2 := tChainId{chainid[r.Intn(13)], "", "正常流程"}
	tests = append(tests, testChainID2)

	// chainId=“随机一个字符串”（循环20次）-非配置文件中chainId
	for i := 0; i < 20; i++ {
		str := make([]string, 0)
		for j := 0; j < r.Intn(100)+1; j++ {
			str = append(str, string(List5[r.Intn(26)]))
		}
		Str := strings.Join(str, "")
		testChainID2 := tChainId{Str, failedMsg, "无效的chainID"}
		tests = append(tests, testChainID2)
	}

	return tests
}

func TestOfKeyStorePath() []tKeyStorePath {

	tests := make([]tKeyStorePath, 0)

	failedMsg := "KeyStorePath does not exist"

	// keystorepath=“”
	pathTest := tKeyStorePath{"", "", "正常流程"}
	tests = append(tests, pathTest)

	// keystorepath= 不存在指定钱包的路径（循环20次）
	for i := 0; i < 20; i++ {
		path := make([]string, 0)
		for m := 0; m < r.Intn(5)+1; m++ {
			for j := 0; j < r.Intn(10)+1; j++ {
				path = append(path, string(List5[r.Intn(26)]))
			}
		}
		Path := strings.Join(path, "/")
		pathTest2 := tKeyStorePath{Path, failedMsg, "钱包地址不存在"}
		tests = append(tests, pathTest2)
	}

	// keystorepath= 存在指定钱包的路径 -非.keystore
	Path, err := GetCurrentPath()
	if err != nil {
		fmt.Println("获取当前文件路径失败-keyStorePath")
		return nil
	}
	pathTest3 := tKeyStorePath{Path + name + ".wal", "", "正常流程"}
	tests = append(tests, pathTest3)

	// keystorepath=.keystore - 有指定钱包
	pathTest4 := tKeyStorePath{".keystore/", "", "正常流程"}
	tests = append(tests, pathTest4)

	return tests
}

func RandomTool(list string, Num int) string {

	switch list {
	case "list1":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List1[r.Intn(len(List1))])
		}
		return Str
	case "list2":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List2[r.Intn(len(List2))])
		}
		return Str
	case "list3":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List3[r.Intn(len(List3))])
		}
		return Str
	case "list4":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List4[r.Intn(len(List4))])
		}
		return Str
	case "list5":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List5[r.Intn(len(List5))])
		}
		return Str
	case "list6":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List6[r.Intn(len(List6))])
		}
		return Str
	case "list7":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List7[r.Intn(len(List7))])
		}
		return Str
	case "list8":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List8[r.Intn(len(List8))])
		}
		return Str
	case "list9":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List9[r.Intn(len(List9))])
		}
		return Str
	case "list10":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List10[r.Intn(len(List10))])
		}
		return Str
	case "list11":
		var Str string
		for i := 0; i < Num; i++ {
			Str = Str + string(List11[r.Intn(len(List11))])
		}
		return Str
	default:
		return "输入的查询List有误，请检查"
	}
	return ""
}
