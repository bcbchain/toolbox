package genstub

import (
	"blockchain/smccheck/parsecode"
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

const iStubTemplate = `package {{$.PackageName}}stub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"
	"blockchain/smcsdk/sdk/types"
	{{if (hasStruct .Functions)}}. "contract/{{$.OrgID}}/code/{{$.DirectionName}}/v{{$.Version}}/{{$.DirectionName}}"{{end}}
	{{- if (hasInterface .Functions)}}
	"contract/{{$.OrgID}}/code/{{$.DirectionName}}/v{{$.Version}}/{{$.DirectionName}}"
	{{- end}}
	{{- if (hasResult .Functions)}}
	"blockchain/smcsdk/sdk/jsoniter"
	{{- end}}
)

{{$stubName := (printf "Interface%sStub" $.ContractStruct)}}

//{{$stubName}} interface stub
type {{$stubName}} struct {
	smc sdk.ISmartContract
}

var _ stubType.IContractIntfcStub = (*{{$stubName}})(nil)

//NewInterStub new interface stub
func NewInterStub(smc sdk.ISmartContract) stubType.IContractIntfcStub {
	return &{{$stubName}}{smc: smc}
}

//GetSdk get sdk
func (inter *{{$stubName}}) GetSdk() sdk.ISmartContract {
	return inter.smc
}

//SetSdk set sdk
func (inter *{{$stubName}}) SetSdk(smc sdk.ISmartContract) {
	inter.smc = smc
}

//Invoke invoke function
func (inter *{{$stubName}}) Invoke(methodID string, p interface{}) (response bcType.Response) {
	defer FuncRecover(&response)

	if len(inter.smc.Message().Origins()) > 9 {
		response.Code = types.ErrStubDefined
		response.Log = "invoke contract's interface steps beyond 8 step"
		return
	}

	// 生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(inter.smc, false)
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(inter.smc.Message(), response.Tags, "", fee, gasUsed, inter.smc.Tx().GasLimit(), err)
		return
	}

	var data string
	err = types.Error{ErrorCode:types.CodeOK}
	switch methodID {
	{{- range $i,$f := $.Functions}}
	{{- if $f.IGas}}
	case "{{$f.Method | createProto | calcMethodID | printf "%x"}}":	// prototype: {{createProto $f.Method}}
		{{- if eq (len $f.Results) 1}}
		data = inter._{{lowerFirst $f.Name}}(inter.smc)
		{{- else}}
		inter._{{lowerFirst $f.Name}}(p)
		{{- end}}
	{{- end}}
	{{- end}}
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(inter.smc.Message(), response.Tags, data, fee, gasUsed, inter.smc.Tx().GasLimit(), err)
	return
}

{{range $i0,$f := $.Functions}}
{{- if $f.IGas}}
func (inter *{{$stubName}}) _{{lowerFirst $f.Name}}(p interface{}) {{if (len $f.Results)}}string{{end}} {
	contractObj := new({{$.PackageName}}.{{$.ContractStruct}})
	contractObj.SetSdk(inter.smc)
	param := p.({{$.PackageName}}.{{$f.Name}}Param)
	{{$l := dec (len $f.Results)}}{{if (len $f.Results)}}{{range $i0,$sPara := $f.Results}}rst{{$i0}}{{if lt $i0 $l}},{{end}}{{end}} := {{end}}contractObj.{{$f.Name}}({{$l := (dec (len $f.SingleParams))}}{{range $i0,$sPara := $f.SingleParams}}param.{{$sPara|expandNames|upperFirst}} {{if lt $i0 $l}},{{end}}{{end}})
	{{- if (len $f.Results)}}
	resultList := make([]interface{}, 0)
	{{range $i0,$sPara := $f.Results}}resultList = append(resultList, rst{{$i0}})
	{{end}}
	resBytes, _ := jsoniter.Marshal(resultList)
	return string(resBytes)
	{{- end}}
}
{{- end}}
{{end}}
`

// GenInterfaceStub - generate the interface stub go source
func GenInterfaceStub(res *parsecode.Result, outDir string) error {
	newOutDir := filepath.Join(outDir, "v"+res.Version, res.ContractName)
	if err := os.MkdirAll(newOutDir, os.FileMode(0750)); err != nil {
		return err
	}
	filename := filepath.Join(newOutDir, res.PackageName+"stub_interface.go")

	funcMap := template.FuncMap{
		"upperFirst":   parsecode.UpperFirst,
		"lowerFirst":   parsecode.LowerFirst,
		"expandNames":  parsecode.ExpandNames,
		"expandType":   parsecode.ExpandType,
		"expandStruct": parsecode.ExpandStruct,
		"createProto":  parsecode.CreatePrototype,
		"calcMethodID": parsecode.CalcMethodID,
		"hasStruct":    hasStruct,
		"dec": func(i int) int {
			return i - 1
		},
		"hasInterface": func(functions []FatFunction) bool {
			for _, function := range functions {
				if function.IGas != 0 {
					return true
				}
			}

			return false
		},
		"hasResult": func(functions []FatFunction) bool {
			for _, function := range functions {
				if function.IGas != 0 && len(function.Results) > 0 {
					return true
				}
			}

			return false
		},
	}
	tmpl, err := template.New("interfaceStub").Funcs(funcMap).Parse(iStubTemplate)
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
