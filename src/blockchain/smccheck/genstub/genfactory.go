package genstub

import (
	"blockchain/smccheck/parsecode"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const factoryTemplate = `package {{$.PackageName}}

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdkimpl/helper"
	"contract/stubcommon/types"

	{{- range $i,$v := $.Versions}}
	{{$v | replace}} "contract/{{$.OrgID}}/stub/{{$.DirectionName}}/{{$v}}/{{$.DirectionName}}"
	{{- end}}
)

//NewInterfaceStub new interface stub
func NewInterfaceStub(smc sdk.ISmartContract, contractName string) types.IContractIntfcStub {
	//Get contract with ContractName
	ch := helper.ContractHelper{}
	ch.SetSMC(smc)
	contract := ch.ContractOfName(contractName)

	switch contract.Version() {
	{{- range $i1,$v1 := $.Versions}}
	case "{{$v1}}":
		return {{replace $v1}}.NewInterStub(smc)
	{{- end}}
	}
	return nil
}
`

// GenStubFactory - generate the interface stub factory go source
func GenStFactory(res *parsecode.Result, outDir string) error {
	if err := os.MkdirAll(outDir, os.FileMode(0750)); err != nil {
		return err
	}
	filename := filepath.Join(outDir, "interfacestubfactory.go")

	funcMap := template.FuncMap{
		"replace": func(version string) string {
			return strings.Replace(version, ".", "_", -1)
		},
	}

	tmpl, err := template.New("interfaceStubFactory").Funcs(funcMap).Parse(factoryTemplate)
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
