package genstub

import (
	"blockchain/smccheck/parsecode"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const templateText = `package stub
import (
	"blockchain/smcsdk/sdk"
	"contract/stubcommon/common"
	"contract/stubcommon/types"
	"fmt"
	"github.com/tendermint/tmlibs/log"

	{{- range $i,$directionName := $.DirectionNames}}
	{{getName $i $.PackageNames}}{{replace (version $i $.Versions)}} "contract/{{$.OrgID}}/stub/{{$directionName}}/v{{version $i $.Versions}}/{{$directionName}}"
	{{- end}}
)

func NewStub(smc sdk.ISmartContract, logger log.Logger) types.IContractStub {

	logger.Debug(fmt.Sprintf("NewStub error, contract=%s,version=%s", smc.Message().Contract().Name(), smc.Message().Contract().Version()))
	switch common.CalcKey(smc.Message().Contract().Name(), smc.Message().Contract().Version()) {
	{{- range $j,$contractName := $.ContractNames}}
	case "{{$contractName}}{{replace (version $j $.Versions)}}":
		return {{getName $j $.PackageNames}}{{replace (version $j $.Versions)}}.New(logger)
	{{- end}}
	default:
		logger.Fatal(fmt.Sprintf("NewStub error, contract=%s,version=%s", smc.Message().Contract().Name(), smc.Message().Contract().Version()))
	}

	return nil
}
`

type OrgContracts struct {
	OrgID          string
	DirectionNames []string
	ContractNames  []string
	PackageNames   []string
	Versions       []string
}

func res2factory(reses []*parsecode.Result) OrgContracts {

	factory := OrgContracts{
		OrgID:          reses[0].OrgID,
		DirectionNames: make([]string, 0),
		ContractNames:  make([]string, 0),
		PackageNames:   make([]string, 0),
		Versions:       make([]string, 0),
	}

	for _, res := range reses {
		factory.DirectionNames = append(factory.DirectionNames, res.DirectionName)
		factory.ContractNames = append(factory.ContractNames, res.ContractName)
		factory.PackageNames = append(factory.PackageNames, res.PackageName)
		factory.Versions = append(factory.Versions, res.Version)
	}

	return factory
}

// GenConStFactory - generate the contract stub factory go source
func GenConStFactory(reses []*parsecode.Result, outDir string) error {
	if err := os.MkdirAll(outDir, os.FileMode(0750)); err != nil {
		return err
	}
	filename := filepath.Join(outDir, "contractstubfactory.go")

	funcMap := template.FuncMap{
		"version": func(index int, versions []string) string {
			return versions[index]
		},
		"replace": func(version string) string {
			return strings.Replace(version, ".", "", -1)
		},
		"getName": func(index int, packageNames []string) string {
			return packageNames[index]
		},
	}

	tmpl, err := template.New("contractStubFactory").Funcs(funcMap).Parse(templateText)
	if err != nil {
		return err
	}

	factory := res2factory(reses)

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, factory); err != nil {
		return err
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}
	return nil
}
