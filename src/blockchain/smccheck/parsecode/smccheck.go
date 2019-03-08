package parsecode

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"blockchain/smcsdk/sdk/types"
)

// Check 分析该目录下的合约代码，进行各种规范检查，提取关键信息
func Check(inPath string) (res *Result, err types.Error) {
	err.ErrorCode = 200

	fSet := token.NewFileSet()

	// parseDir 實際並不能遞歸檢查多層級目錄，需要自己去遞歸檢查。實際解析到的也只有一個pkg
	pkgMap, err0 := parser.ParseDir(fSet, inPath, isContractFile, parser.ParseComments)
	if err0 != nil {
		ErrorTransfer(err0, &err)
	}

	if len(pkgMap) != 1 {
		err.ErrorCode = 500
		err.ErrorDesc = "parse failed, no pkg or more than 1 pkg"
	}

	v := newVisitor()
	for _, pkg := range pkgMap {
		//fmt.Println(pkgName)
		for _, node := range pkg.Files {
			//fmt.Println(fName)
			ast.Walk(v, node)
			importsCollector(v.res)
		}
	}
	for _, pkg := range pkgMap {
		for _, tc := range transferCallee {
			v.parseCall(tc, pkg.Files)
		}
	}

	v.printContractInfo()
	checkImportConflict(v.res)
	if v.res.ErrFlag {
		for idx, pos := range v.res.ErrorPos {
			fmt.Println(v.res.ErrorDesc[idx])
			fmt.Println(fSet.Position(pos).Filename, fSet.Position(pos).Line)
		}
		err.ErrorCode = 500
		err.ErrorDesc = "check failed"
		return
	}
	res = v.res

	//if err0 = genSDK(inPath, v.res); err0 != nil {
	//	ErrorTransfer(err0, &err)
	//}
	//if err0 = genReceipt(inPath, v.res); err0 != nil {
	//	ErrorTransfer(err0, &err)
	//}
	//if err0 = genStore(inPath, v.res); err0 != nil {
	//	ErrorTransfer(err0, &err)
	//}

	return
}

func isContractFile(d os.FileInfo) bool {
	return !d.IsDir() &&
		strings.HasSuffix(d.Name(), ".go") &&
		!strings.Contains(d.Name(), "autogen") &&
		!strings.HasSuffix(d.Name(), "_test.go")
}

func checkImportConflict(res *Result) {
	for imp := range res.Imports {
		for im := range res.Imports {
			if imp.Name == "." {
				res.ErrFlag = true
				res.ErrorDesc = append(res.ErrorDesc, "dot Import not allowed")
				res.ErrorPos = append(res.ErrorPos, 0)
			}
			if imp != im && imp.Name != "" && imp.Name == im.Name && imp.Path != im.Path {
				res.ErrFlag = true
				res.ErrorDesc = append(res.ErrorDesc, "Import conflict:"+imp.Name+" has more than one path:"+imp.Path+","+im.Path)
				res.ErrorPos = append(res.ErrorPos, 0)
			}
			// TODO alias name is in other path
		}
	}
}

// FmtAndWrite - go fmt content and write to filename
func FmtAndWrite(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if ee := f.Close(); ee != nil {
			fmt.Println(ee)
		}
	}()

	// Create a FileSet for node. Since the node does not come
	// from a real source file, fSet will be empty.
	fSet := token.NewFileSet()

	// parser.ParseExpr parses the argument and returns the
	// corresponding ast.Node.
	node, err := parser.ParseFile(fSet, "", content, parser.ParseComments)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = format.Node(&buf, fSet, node)
	if err != nil {
		return err
	}

	n, err := f.WriteString(buf.String())
	if err != nil {
		return err
	}
	fmt.Println(n, "byte write to file")
	return nil
}

// Versions - 阿朱你也不寫注釋，顯然沒裝 meta linter , 哈哈哈
func Versions(firstContractPath string, res *Result) types.Error {
	retErr := types.Error{ErrorCode: types.CodeOK}

	fInfoS, err := ioutil.ReadDir(firstContractPath)
	if err != nil {
		retErr.ErrorCode = types.ErrInvalidParameter
		retErr.ErrorDesc = err.Error()
		return retErr
	}

	res.Versions = make([]string, 0)
	for _, fInfo := range fInfoS {
		if fInfo.IsDir() && fInfo.Name() != "." && fInfo.Name() != ".." {
			res.Versions = append(res.Versions, fInfo.Name())
		}
	}

	return retErr
}
