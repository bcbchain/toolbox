package gen

import (
	"blockchain/smccheck/parsecode"
	"bytes"
	"path/filepath"
	"text/template"
)

var sdkTemplate = `package {{.PackageName}}

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk - set sdk
func ({{$.ReceiverName}} *{{$.ContractName}})  SetSdk(sdk sdk.ISmartContract) {
	{{$.ReceiverName}}.sdk = sdk
}

//GetSdk - get sdk
func ({{$.ReceiverName}} *{{$.ContractName}})  GetSdk() sdk.ISmartContract {
	return {{$.ReceiverName}}.sdk
}`

type baseExport struct {
	PackageName  string
	ReceiverName string
	ContractName string
}

func res2Base(res *parsecode.Result) baseExport {
	base := baseExport{}
	base.PackageName = res.PackageName
	base.ReceiverName = res.InitChain.Receiver.Names[0]
	base.ContractName = res.ContractStructure
	return base
}

func GenSDK(inPath string, res *parsecode.Result) error {
	filename := filepath.Join(inPath, res.PackageName+"_autogen_sdk.go")

	tmpl, err := template.New("base").Parse(sdkTemplate)
	if err != nil {
		return err
	}

	base := res2Base(res)

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, base); err != nil {
		return err
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}
	return nil
}
