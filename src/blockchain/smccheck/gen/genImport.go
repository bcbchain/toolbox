package gen

import (
	"blockchain/smccheck/parsecode"
	"blockchain/smcsdk/sdk/std"
	"bytes"
	"fmt"
	"go/ast"
	"path/filepath"
	"strings"
	"text/template"
)

var importTemplate = `package {{.PackageName}}

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	{{if (isEx $.Contracts $.ImportContract $.ImportInterfaces)}}"common/jsoniter"{{end}}
	"contract/{{$.OrgID}}/stub/{{$.ImportContract}}"
	types2 "contract/stubcommon/types"

	{{range $i, $contract := $.Contracts}}{{if eq $contract.Name $.ImportContract}}
	{{$contract.Name}}v{{vEx $contract.Version}} "contract/{{$.OrgID}}/code/{{$.ImportContract}}/v{{$contract.Version}}/{{$.ImportContract}}"
	{{- end}}{{- end}}
)

//InterfaceStub interface stub of playerbook
type InterfaceStub struct {
    stub types2.IContractIntfcStub
}

const importContractName = "{{$.ImportContract}}"

func (s *{{.ContractStructure}}) {{$.ImportContract}}Stub() *InterfaceStub {
    return &InterfaceStub{ {{$.ImportContract}}.NewInterfaceStub(s.GetSdk(), importContractName)}
}

{{range $j, $method := $.ImportInterfaces}}
// {{$method.Name}}
func (intfc *InterfaceStub) {{$method.Name}}({{range $i0, $param := $method.Params}}{{$param | expNames}} {{$param | expType}}{{if lt $i0 (dec (len $method.Params))}},{{end}}{{end}}) {{if (len $method.Results)}}string{{end}} {

    methodID := "{{$method | createProto | calcMethodID | printf "%x"}}" // prototype: {{createProto $method}}
    oldSmc := intfc.stub.GetSdk()
    defer intfc.stub.SetSdk(oldSmc)
    //合约调用时的输入收据，同时可作为跨合约调用的输入收据
    contract := oldSmc.Helper().ContractHelper().ContractOfName(importContractName)
    newSmc := sdkhelper.OriginNewMessage(oldSmc, contract, methodID, oldSmc.Message().(*object.Message).OutputReceipts())
    intfc.stub.SetSdk(newSmc)

    //TODO 编译时builder从数据库已获取合约版本和失效高度，直接使用
    height := newSmc.Block().Height()
    var rn interface{}
	{{createVar $.Contracts $.ImportContract $method}}

    response := intfc.stub.Invoke(methodID, rn)
    if response.Code != types.CodeOK {
        panic(response.Log)
    }
    oldmsg := oldSmc.Message()
    oldmsg.(*object.Message).AppendOutput(intfc.stub.GetSdk().Message().(*object.Message).OutputReceipts())
    {{if (len $method.Results)}}return response.Data{{end}}
}

{{- end}}
`

type ContractInfo struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	EffectHeight int64  `json:"effectHeight"`
	LoseHeight   int64  `json:"loseHeight"`
}

type OtherContract struct {
	OrgID         string
	DirectionName string
	Name          string
	PackageName   string
	Version       string
	LoseHeight    int64
	EffectHeight  int64
	Functions     []parsecode.Function
	UserStruct    map[string]ast.GenDecl
}

type ImportContract struct {
	Contracts []OtherContract

	OrgID              string
	PackageName        string
	ContractStructure  string
	ImportContract     string
	ImportInterfaces   []parsecode.Method
	ImportContractInfo std.ContractVersionList
}

func res2importContract(res *parsecode.Result, reses []*parsecode.Result, contractInfoList []ContractInfo) ImportContract {

	importContract := ImportContract{
		Contracts:        make([]OtherContract, 0),
		ImportInterfaces: make([]parsecode.Method, 0),
	}

	for _, item := range reses {
		contract := OtherContract{
			OrgID:         item.OrgID,
			DirectionName: item.DirectionName,
			Name:          item.ContractName,
			PackageName:   item.PackageName,
			Version:       item.Version,
			Functions:     item.Functions,
			UserStruct:    item.UserStruct,
		}
		contract.EffectHeight, contract.LoseHeight = contractInfoOfNameVersion(contract.Name, contract.Version, contractInfoList)

		importContract.Contracts = append(importContract.Contracts, contract)

	}
	importContract.OrgID = res.OrgID
	importContract.ImportContract = res.ImportContract
	importContract.ImportInterfaces = res.ImportInterfaces
	importContract.ContractStructure = res.ContractStructure
	importContract.PackageName = res.PackageName

	importContract.ImportContractInfo.Name = importContract.ImportContract
	importContract.ImportContractInfo.EffectHeights = []int64{1000, -1}

	return importContract
}

func contractInfoOfNameVersion(name, version string, contractInfoList []ContractInfo) (effectHeight, loseHeight int64) {

	for _, contractInfo := range contractInfoList {
		if contractInfo.Name == name && contractInfo.Version == version {
			return contractInfo.EffectHeight, contractInfo.LoseHeight
		}
	}

	return 0, 0
}

// GenImport - generate import code from source smart contract to destination smart contract
func GenImport(inPath string, res *parsecode.Result, reses []*parsecode.Result, contractInfoList []ContractInfo) error {
	filename := filepath.Join(inPath, res.PackageName+"_autogen_import_"+res.ImportContract+".go")

	funcMap := template.FuncMap{
		"upperFirst":   parsecode.UpperFirst,
		"lowerFirst":   parsecode.LowerFirst,
		"expNames":     parsecode.ExpandNames,
		"expType":      parsecode.ExpandType,
		"expNoS":       parsecode.ExpandTypeNoStar,
		"expK":         parsecode.ExpandMapFieldKey,
		"expV":         parsecode.ExpandMapFieldVal,
		"expVNoS":      parsecode.ExpandMapFieldValNoStar,
		"createProto":  parsecode.CreatePrototype,
		"calcMethodID": parsecode.CalcMethodID,
		"createVar":    createVar,
		"isEx":         isEx,
		"dec": func(i int) int {
			return i - 1
		},
		"vEx": vEx,
	}
	tmpl, err := template.New("import").Funcs(funcMap).Parse(importTemplate)
	if err != nil {
		return err
	}

	importContract := res2importContract(res, reses, contractInfoList)

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, importContract); err != nil {
		return err
	}

	if err = parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}
	return nil
}

func vEx(version string) string {
	return strings.Replace(version, ".", "", -1)
}

// createVar - create string of stride smart contract code about parameter for method
func createVar(allContracts []OtherContract, contractName string, method parsecode.Method) string {
	contracts := getContracts(allContracts, contractName)

	formatStr := ""

	for _, contract := range contracts {
		if isOK(contract.Functions, method) {
			formatStr += exchangeVar(contract, method)
			break
		}
	}

	for index, contract := range contracts {
		item := getArg(contract.Functions, method)
		if index == 0 {
			if contract.LoseHeight == 0 {
				formatStr += fmt.Sprintf("\tif height >= %d {\n", contract.EffectHeight)
			} else {
				formatStr += fmt.Sprintf("\tif height < %d {\n", contract.LoseHeight)
			}
		} else if index < len(contracts)-1 {
			formatStr += fmt.Sprintf(" else if height < %d {\n", contract.LoseHeight)
		} else {
			formatStr += fmt.Sprintf(" else {\n")
		}
		formatStr += "\t\t"
		if isOK(contract.Functions, method) {
			formatStr += fmt.Sprintf("rn = %sv%s.%sParam", contract.Name, vEx(contract.Version), method.Name)
		}
		formatStr += item
		formatStr += "\n\t}"
	}

	return formatStr
}

var baseTypes = []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64",
	"bool", "[]byte", "types.Address", "bn.Number", "types.HexBytes", "types.Hash", "types.PubKey", "string",
	"float32", "float64", "byte"}

func exchangeVar(contract OtherContract, method parsecode.Method) string {

	exchangeStr := ""
	isF := true
	for _, function := range contract.Functions {
		if function.Name == method.Name {
			for index, filed := range function.Params {
				isBase := false
				varType := strings.TrimLeft(parsecode.ExpandType(filed), "*")
				for _, typeStr := range baseTypes {
					if varType == typeStr {
						isBase = true
						break
					}
				}

				if isBase == false {
					var names []string
					for indexN, name := range method.Params[index].Names {
						if isF {
							exchangeStr += "var err error\n"
							exchangeStr += "\tvar resBytes []byte\n"
							isF = false
						}
						typeTemp := varType
						if isTypeIn(contract, varType) {
							typeTemp = fmt.Sprintf("%sv%s.%s", contract.Name, vEx(contract.Version), varType)
						}
						exchangeStr += fmt.Sprintf("\tvar p%d%d %s\n", index, indexN, typeTemp)
						exchangeStr += fmt.Sprintf("\tresBytes, err = jsoniter.Marshal(%s)\n", name)
						exchangeStr += "\tif err != nil {\n"
						exchangeStr += "\t\tpanic(err)\n"
						exchangeStr += "\t}\n"
						exchangeStr += fmt.Sprintf("\terr = jsoniter.Unmarshal(resBytes, &p%d%d)\n", index, indexN)
						exchangeStr += "\tif err != nil {\n"
						exchangeStr += "\t\tpanic(err)\n"
						exchangeStr += "\t}\n\n"
						names = append(names, fmt.Sprintf("p%d%d", index, indexN))
					}
					method.Params[index].Names = names
				}
			}
		}
	}

	return exchangeStr
}

func isTypeIn(contract OtherContract, typeStr string) bool {
	for key := range contract.UserStruct {
		if key == typeStr {
			return true
		}
	}

	return false
}

func getContracts(allContracts []OtherContract, contractName string) []OtherContract {
	contracts := make([]OtherContract, 0)

	for _, contract := range allContracts {
		if contract.Name == contractName {
			contracts = append(contracts, contract)
		}
	}

	return contracts
}

func isOK(functions []parsecode.Function, method parsecode.Method) bool {
	for _, function := range functions {
		if function.Name == method.Name && function.IGas != 0 {
			// step 1. check the parameter's count
			// step 2. check the parameter's name and type to same, different type will make different methodID
			mLenParams := 0
			for index, param := range method.Params {
				mLenParams += len(param.Names)

				if parsecode.ExpandType(param) != parsecode.ExpandType(function.Params[index]) {
					return false
				}
			}

			fLenParams := 0
			for _, param := range function.Params {
				fLenParams += len(param.Names)
			}

			if fLenParams == mLenParams {
				return true
			} else {
				return false
			}
		}
	}

	return false
}

func getArg(functions []parsecode.Function, method parsecode.Method) string {
	for _, function := range functions {
		if function.Name == method.Name && function.IGas != 0 {
			mLenParams := 0
			for _, param := range method.Params {
				mLenParams += len(param.Names)
			}

			paramStr := ""
			fLenParams := 0
			for index1, param := range function.Params {
				fLenParams += len(param.Names)
				for index2, name := range param.Names {
					paramStr += parsecode.UpperFirst(name) + ":" + method.Params[index1].Names[index2] + ","
				}
			}
			paramStr = paramStr[:len(paramStr)-1]
			paramStr = "{" + paramStr + "}"

			if fLenParams == mLenParams {
				return paramStr
			} else {
				return `panic("Invalid parameter")`
			}
		}
	}

	return `panic("Invalid parameter")`
}

func isEx(allContracts []OtherContract, contractName string, methods []parsecode.Method) bool {
	for _, method := range methods {
		contracts := getContracts(allContracts, contractName)

		for _, contract := range contracts {
			if isOK(contract.Functions, method) {
				temp := exchangeVar(contract, method)
				if len(temp) != 0 {
					return true
				}
			}
		}
	}

	return false
}
