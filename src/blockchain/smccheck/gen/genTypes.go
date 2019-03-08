package gen

import (
	"blockchain/smccheck/genrpc"
	"blockchain/smccheck/parsecode"
	"bytes"
	"path/filepath"
	"text/template"
)

const typesTemplate = `package {{.PackageName}}
{{- if (hasInterface .Functions)}}
{{if .Imports}}import ({{end}}
  {{range $v,$vv := .Imports}}
{{$v.Name}} {{$v.Path}}{{end}}
{{if .Imports}}){{end}}
{{- end}}

{{range $i,$f := .Functions}}
{{- if $f.IGas}}
// {{$f.Name}}Param structure of parameters of {{$f.Name}}() of v2.0
type {{$f.Name}}Param struct { {{range $ii,$p := $f.SingleParams}}
	{{$p|expNames|upperFirst}} {{$p|expType}}{{end}}}
{{- end}}
{{end}}
`

func GenTypes(inPath string, res *parsecode.Result) error {
	filename := filepath.Join(inPath, res.PackageName+"_autogen_types.go")

	funcMap := template.FuncMap{
		"upperFirst": parsecode.UpperFirst,
		"expNames":   parsecode.ExpandNames,
		"expType":    parsecode.ExpandType,
		"hasInterface": func(functions []genrpc.FatFunction) bool {
			for _, function := range functions {
				if function.IGas != 0 {
					return true
				}
			}

			return false
		},
	}
	tmpl, err := template.New("types").Funcs(funcMap).Parse(typesTemplate)
	if err != nil {
		return err
	}

	types := genrpc.Res2rpc(res)

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, types); err != nil {
		return err
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		return err
	}
	return nil

}
