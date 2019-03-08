package genstub

import (
	"blockchain/smccheck/parsecode"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const stubTemplate = `package {{$.PackageName}}stub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"
	"blockchain/smcsdk/sdk/types"

	"github.com/tendermint/tmlibs/log"
	{{- if (hasMethod .Functions)}}
	"blockchain/smcsdk/sdk/rlp"
	"contract/{{$.OrgID}}/code/{{$.DirectionName}}/v{{$.Version}}/{{$.DirectionName}}"
	{{- end}}
	{{- if (hasResult .Functions)}}
	"blockchain/smcsdk/sdk/jsoniter"
	{{- end}}

  	{{- range $v,$vv := .Imports}}
	{{- if (filterImport $v.Path)}}
  	{{$v.Name}} {{$v.Path}}
	{{- end}}
	{{- end}}
)

{{$stubName := (printf "%sStub" $.ContractStruct)}}

//{{$stubName}} an object
type {{$stubName}} struct {
	logger log.Logger
}

var _ stubType.IContractStub = (*{{$stubName}})(nil)

//New generate a stub
func New(logger log.Logger) stubType.IContractStub {
	return &{{$stubName}}{logger: logger}
}

//FuncRecover recover panic by Assert
func FuncRecover(response *bcType.Response) {
	if err := recover(); err != nil {
		if _, ok := err.(types.Error); ok {
			error := err.(types.Error)
			response.Code = error.ErrorCode
			response.Log = error.Error()
		} else {
			panic(err)
		}
	}
}

//Invoke invoke function
func (pbs *{{$stubName}}) Invoke(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	// 生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(smc, true)
	response.Fee = fee
	response.GasUsed = gasUsed
 	response.Tags = append(response.Tags, tmcommon.KVPair{Key:feeReceipt.Key, Value:feeReceipt.Value})
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(smc.Message(), response.Tags, "", fee, gasUsed, smc.Tx().GasLimit(), err)
		return
	}


	var data string
	err = types.Error{ErrorCode:types.CodeOK}
	switch smc.Message().MethodID() {
	{{- range $i,$f := $.Functions}}
	case "{{$f.Method | createProto | calcMethodID | printf "%x"}}":	// prototype: {{createProto $f.Method}}
		{{if eq (len $f.Results) 1}}data = {{end}}{{lowerFirst $f.Name}}(smc)
	{{- end}}
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(smc.Message(), response.Tags, data, fee, gasUsed, smc.Tx().GasLimit(), err)
	return
}

{{range $i0,$f := $.Functions}}
func {{lowerFirst $f.Name}}(smc sdk.ISmartContract) {{if (len $f.Results)}}string{{end}} {
	items := smc.Message().Items()
	sdk.Require(len(items) == {{paramsLen $f.Method}}, types.ErrStubDefined, "Invalid message data")

	{{- if len $f.SingleParams}}
	var err error
	{{- end}}
	{{range $i1,$param := $f.SingleParams}}
	var v{{$i1}} {{$param | expandType}}
	err = rlp.DecodeBytes(items[{{$i1}}], &v{{$i1}})
	sdk.RequireNotError(err, types.ErrInvalidParameter)
	{{end}}

	contractObj := new({{$.PackageName}}.{{$.ContractStruct}})
	contractObj.SetSdk(smc)
	{{$l := dec (len $f.Results)}}{{if (len $f.Results)}}{{range $i0,$sPara := $f.Results}}rst{{$i0}}{{if lt $i0 $l}},{{end}}{{end}} := {{end}}contractObj.{{$f.Name}}{{$l2 := dec (len $f.SingleParams)}}({{range $i2,$sPara := $f.SingleParams}}v{{$i2}}{{if lt $i2 $l2}},{{end}}{{end}})
	{{- if (len $f.Results)}}
	resultList := make([]interface{}, 0)
	{{range $i0,$sPara := $f.Results}}resultList = append(resultList, rst{{$i0}})
	{{end}}
	resBytes, _ := jsoniter.Marshal(resultList)
	return string(resBytes)
	{{- end}}
}
{{end}}
`

// FatFunction - flat params
type FatFunction struct {
	parsecode.Function
	SingleParams []parsecode.Field
}

// RPCExport - the functions for rpc & autogen types
type StubExport struct {
	DirectionName  string
	PackageName    string
	ReceiverName   string
	ContractName   string
	ContractStruct string
	Version        string
	Versions       []string
	OrgID          string
	Owner          string
	Imports        map[parsecode.Import]struct{}
	Functions      []FatFunction
	Port           int

	PlainUserStruct []string
}

// Res2rpc - transform the parsed result to RPC Export struct
func Res2stub(res *parsecode.Result) StubExport {
	exp := StubExport{}
	exp.DirectionName = res.DirectionName
	exp.PackageName = res.PackageName
	exp.ReceiverName = res.InitChain.Receiver.Names[0]
	exp.ContractName = res.ContractName
	exp.ContractStruct = res.ContractStructure
	exp.OrgID = res.OrgID
	exp.Version = res.Version
	exp.Versions = res.Versions
	imports := make(map[parsecode.Import]struct{})

	fatFunctions := make([]FatFunction, 0)
	pus := make([]string, 0)
	for _, f := range res.Functions {
		fat := FatFunction{
			Function: f,
		}
		singleParams := make([]parsecode.Field, 0)
		for _, para := range f.Params {
			for imp := range para.RelatedImport {
				imports[imp] = struct{}{}
			}
			singleParams = append(singleParams, parsecode.FieldsExpand(para)...)
			t := parsecode.ExpandTypeNoStar(para)
			t = strings.TrimSpace(t)
			if u, ok := res.UserStruct[t]; ok {
				pus = append(pus, parsecode.ExpandStruct(u))
			}
		}
		fat.SingleParams = singleParams
		fatFunctions = append(fatFunctions, fat)
	}
	exp.Functions = fatFunctions
	exp.Imports = imports
	exp.PlainUserStruct = pus
	return exp
}

// GenMethodStub - generate the method stub go source
func GenMethodStub(res *parsecode.Result, outDir string) error {
	newOutDir := filepath.Join(outDir, "v"+res.Version, res.DirectionName)
	if err := os.MkdirAll(newOutDir, os.FileMode(0750)); err != nil {
		return err
	}
	filename := filepath.Join(newOutDir, res.PackageName+"stub_method.go")

	funcMap := template.FuncMap{
		"upperFirst":   parsecode.UpperFirst,
		"lowerFirst":   parsecode.LowerFirst,
		"expandNames":  parsecode.ExpandNames,
		"expandType":   parsecode.ExpandType,
		"createProto":  parsecode.CreatePrototype,
		"paramsLen":    parsecode.ParamsLen,
		"calcMethodID": parsecode.CalcMethodID,
		"filterImport": parsecode.FilterImports,
		"dec": func(i int) int {
			return i - 1
		},
		"hasMethod": func(functions []FatFunction) bool {
			for _, function := range functions {
				if function.MGas != 0 {
					return true
				}
			}

			return false
		},
		"hasResult": func(functions []FatFunction) bool {
			for _, function := range functions {
				if function.MGas != 0 && len(function.Results) > 0 {
					return true
				}
			}

			return false
		},
	}
	tmpl, err := template.New("methodStub").Funcs(funcMap).Parse(stubTemplate)
	if err != nil {
		return err
	}

	obj := Res2stub(res)

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, obj); err != nil {
		return err
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}
	return nil
}
