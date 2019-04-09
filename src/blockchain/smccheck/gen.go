package smccheck

import (
	"blockchain/smccheck/gen"
	"blockchain/smccheck/gencmd"
	"blockchain/smccheck/genstub"
	"blockchain/smccheck/parsecode"
	"blockchain/smcsdk/sdk/std"
	"blockchain/smcsdk/sdk/types"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const orgGenesis = "orgJgaGConUyK81zibntUBjQ33PKctpk1K1G"

// Gen - walk contract path and generate code
func Gen(contractDir string, contractInfoList []gen.ContractInfo) (results []std.GenResult, err types.Error) {

	err.ErrorCode = types.CodeOK

	subDirs, er := ioutil.ReadDir(contractDir)
	if er != nil {
		err.ErrorCode = types.ErrInvalidParameter
		err.ErrorDesc = er.Error()
		return
	} else if len(subDirs) > 2 || (len(subDirs) == 2 && !(subDirs[0].Name() == orgGenesis || subDirs[1].Name() == orgGenesis)) {
		err.ErrorCode = types.ErrInvalidParameter
		err.ErrorDesc = "invalid directory"
		return
	}

	results = make([]std.GenResult, 0)
	for _, dir := range subDirs {
		totalResList := make([]*parsecode.Result, 0)

		firstContractPath := filepath.Join(contractDir, dir.Name())
		codePath := filepath.Join(firstContractPath, "code")
		contractDirs, er := ioutil.ReadDir(codePath)
		if er != nil {
			err.ErrorCode = types.ErrInvalidParameter
			err.ErrorDesc = er.Error()
			return
		}

		for _, contractPath := range contractDirs {
			resList, er := genWithContractPath(filepath.Join(codePath, contractPath.Name()))
			if er != nil {
				err.ErrorCode = types.ErrInvalidParameter
				err.ErrorDesc = er.Error()
				return
			}

			// get methods and interfaces
			results = append(results, getGenResult(resList))

			totalResList = append(totalResList, resList...)
		}

		// generate contract stub factory
		stubPath := filepath.Join(contractDir, dir.Name()+"/stub")
		er = genstub.GenConStFactory(totalResList, stubPath)
		if er != nil {
			err.ErrorCode = types.ErrInvalidParameter
			err.ErrorDesc = er.Error()
			return
		}

		for _, res := range totalResList {
			if res.ImportContract != "" {
				inPath := filepath.Join(filepath.Join(filepath.Join(codePath, res.DirectionName), "v"+res.Version), res.DirectionName)
				er = gen.GenImport(inPath, res, totalResList, contractInfoList)
				if er != nil {
					err.ErrorCode = types.ErrInvalidParameter
					err.ErrorDesc = er.Error()
					return
				}
			}
		}
	}

	// generate stub common
	er = genstub.GenStubCommon(filepath.Join(contractDir, "stubcommon"))
	if er != nil {
		err.ErrorCode = types.ErrInvalidParameter
		err.ErrorDesc = er.Error()
		return
	}

	// generate cmd
	stubName := subDirs[0].Name()
	if len(subDirs) == 2 {
		if subDirs[0].Name() == orgGenesis {
			stubName = subDirs[1].Name()
		}
	}
	er = gencmd.GenCmd(filepath.Dir(contractDir), stubName)
	if er != nil {
		err.ErrorCode = types.ErrInvalidParameter
		err.ErrorDesc = er.Error()
	}

	return
}

// genWithContractPath - generate auto gen code and stub code
func genWithContractPath(contractPath string) (resList []*parsecode.Result, err error) {

	versionDirs, err := ioutil.ReadDir(contractPath)
	if err != nil {
		return
	}

	resList = make([]*parsecode.Result, 0)
	for _, versionDir := range versionDirs {
		versionPath := filepath.Join(contractPath, versionDir.Name())
		var secDir []os.FileInfo
		secDir, err = ioutil.ReadDir(versionPath)
		if err != nil {
			return
		} else if len(secDir) != 1 {
			return nil, errors.New("invalid path " + versionPath)
		}

		secPath := filepath.Join(versionPath, secDir[0].Name())
		res, Err := parsecode.Check(secPath)
		if Err.ErrorCode != types.CodeOK {
			return nil, errors.New(Err.ErrorDesc)
		}
		res.DirectionName = secDir[0].Name()
		resList = append(resList, res)

		err = genAutoGenCode(secPath, res)
		if err != nil {
			return
		}

		stubPath := filepath.Join(filepath.Dir(filepath.Dir(contractPath)), "stub")
		Err = parsecode.Versions(contractPath, res)
		if Err.ErrorCode != types.CodeOK {
			return nil, errors.New(Err.ErrorDesc)
		}

		err = genStubCode(stubPath, res)
		if err != nil {
			return
		}
	}

	return
}

// genAutoGenCode - generate contract assist code
func genAutoGenCode(secPath string, res *parsecode.Result) (err error) {

	err = gen.GenReceipt(secPath, res)
	if err != nil {
		return
	}

	err = gen.GenSDK(secPath, res)
	if err != nil {
		return
	}

	err = gen.GenStore(secPath, res)
	if err != nil {
		return
	}

	err = gen.GenTypes(secPath, res)
	if err != nil {
		return
	}

	return
}

// genStubCode - generate stub code
func genStubCode(stubPath string, res *parsecode.Result) (err error) {
	stubConPath := filepath.Join(stubPath, res.DirectionName)

	err = genstub.GenMethodStub(res, stubConPath)
	if err != nil {
		return
	}

	err = genstub.GenInterfaceStub(res, stubConPath)
	if err != nil {
		return
	}

	err = genstub.GenStFactory(res, stubConPath)
	if err != nil {
		return
	}

	return
}

// getGenResult - get gen result
func getGenResult(resList []*parsecode.Result) (genResult std.GenResult) {
	for _, res := range resList {

		genResult.ContractName = res.ContractName
		genResult.Version = res.Version
		genResult.OrgID = res.OrgID
		genResult.Methods = make([]std.Method, 0)
		genResult.Interfaces = make([]std.Method, 0)

		for _, function := range res.Functions {
			proto := parsecode.CreatePrototype(function.Method)
			method := std.Method{
				ProtoType: proto,
				MethodID:  fmt.Sprintf("%x", parsecode.CalcMethodID(proto))}

			if function.MGas != 0 {
				method.Gas = function.MGas
				genResult.Methods = append(genResult.Methods, method)
			}

			if function.IGas != 0 {
				method.Gas = function.IGas
				genResult.Interfaces = append(genResult.Interfaces, method)
			}
		}
	}

	return
}
