package parsecode

import (
	"blockchain/smcsdk/sdk/types"
	"bytes"
	"encoding/binary"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

// ErrorTransfer - transform common err to types.Error
func ErrorTransfer(err0 error, err *types.Error) {
	err.ErrorCode = 500
	err.ErrorDesc = err0.Error()
}

// FieldsExpand - expand a multi name field to multi single name field
func FieldsExpand(f Field) []Field {
	list := make([]Field, 0)
	if len(f.Names) > 1 {
		for _, name := range f.Names {
			ff := Field{Names: []string{name}}
			ff.FieldType = f.FieldType
			list = append(list, ff)
		}
	} else {
		list = append(list, f)
	}
	return list
}

// UpperFirst - first letter to upper
func UpperFirst(word string) string {
	if len(word) == 0 {
		return ""
	}
	return strings.ToUpper(word[:1]) + word[1:]
}

// LowerFirst - first letter to lower
func LowerFirst(word string) string {
	if len(word) == 0 {
		return ""
	}
	return strings.ToLower(word[:1]) + word[1:]
}

// ExpandType - write field type to string
// nolint unhandled ..
func ExpandType(t Field) string {
	var buf bytes.Buffer

	fSet := token.NewFileSet()
	if err := format.Node(&buf, fSet, t.FieldType); err != nil {
		return ""
	}

	return buf.String()
}

// ExpandTypeNoStar - write field type to string without star, ...
func ExpandTypeNoStar(t Field) string {
	s := ExpandType(t)
	if strings.HasPrefix(s, "*") {
		return s[1:]
	}
	return s
}

// ExpandMapFieldKey - write map field key type to string
func ExpandMapFieldKey(f Field, index int) string {
	m, ok := f.FieldType.(*ast.MapType)
	if !ok {
		return ""
	}

	vF := Field{FieldType: m.Value}
	str := ExpandMapFieldKey(vF, index+1)

	var buf bytes.Buffer
	fSet := token.NewFileSet()
	if err := format.Node(&buf, fSet, m.Key); err != nil {
		return ""
	}

	if len(str) > 0 {
		return fmt.Sprintf("k%d %s, %s", index, buf.String(), str)
	} else {
		return fmt.Sprintf("k%d %s", index, buf.String())
	}
}

// ExpandMapFieldKeyToKey - write map field key make to string as access key
func ExpandMapFieldKeyToKey(f Field, index int) string {
	str := expandMapFieldKeyToKey(f, index)

	strSplit := strings.Split(str, ",")
	fmtStr := ""
	for range strSplit {
		fmtStr += "/%v"
	}
	fmtStr += "\", "

	return fmtStr + str
}

func expandMapFieldKeyToKey(f Field, index int) string {
	m, ok := f.FieldType.(*ast.MapType)
	if !ok {
		return ""
	}

	vF := Field{FieldType: m.Value}
	str := expandMapFieldKeyToKey(vF, index+1)

	if len(str) > 0 {
		return fmt.Sprintf("k%d,%s", index, str)
	} else {
		return fmt.Sprintf("k%d", index)
	}
}

// ExpandMapFieldVal - write map field val type to string
func ExpandMapFieldVal(f Field) string {
	m, ok := f.FieldType.(*ast.MapType)
	if !ok {
		return ""
	}

	if m1, ok := m.Value.(*ast.MapType); ok {
		m = m1
	}
	var buf bytes.Buffer
	fSet := token.NewFileSet()
	if err := format.Node(&buf, fSet, m.Value); err != nil {
		return ""
	}
	return buf.String()
}

// ExpandMapFieldValNoStar - write map field val type to string but not star, ...
func ExpandMapFieldValNoStar(f Field) string {
	v := ExpandMapFieldVal(f)
	if strings.HasPrefix(v, "*") {
		return v[1:]
	}
	return v
}

// ExpandNames - write field names to string
// nolint unhandled ..
func ExpandNames(t Field) string {
	var buf bytes.Buffer

	l := len(t.Names)
	for idx, name := range t.Names {
		buf.WriteString(name)
		if idx < l-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

// ExpandStruct - actually expand GenDecl
// nolint unhandled ..
func ExpandStruct(s ast.GenDecl) string {
	var buf0 bytes.Buffer
	for _, spec := range s.Specs {
		var buf bytes.Buffer

		fSet := token.NewFileSet()
		if err := format.Node(&buf, fSet, spec); err != nil {
			continue
		}
		buf0.WriteString(buf.String())
	}
	return buf0.String()
}

// nolint unhandled
func ExpandMethodPrototype(f Function) string {
	var buf bytes.Buffer

	buf.WriteString(f.Name)
	buf.WriteString("(")
	l := len(f.Params)
	for idx, p := range f.Params {
		l2 := len(p.Names)
		for idx2 := range p.Names {
			buf.WriteString(ExpandType(p))
			if idx2 < l2 {
				buf.WriteString(",")
			}
		}
		if idx < l {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")
	l = len(f.Results)
	if l > 1 || len(f.Results[0].Names) > 1 {
		buf.WriteString("(")
	}
	for idx, p := range f.Results {
		l2 := len(p.Names)
		for idx2 := range p.Names {
			buf.WriteString(ExpandType(p))
			if idx2 < l2 {
				buf.WriteString(",")
			}
		}
		if idx < l {
			buf.WriteString(",")
		}
	}

	if l > 1 || len(f.Results[0].Names) > 1 {
		buf.WriteString(")")
	}
	return buf.String()
}

// GetGas - get gas value
func GetGas(comment string) string {
	lines := strings.Split(comment, "\n")

	var gas string
	for _, line := range lines {
		if strings.HasPrefix(line, "@:public:method:gas") {
			startIndex := strings.Index(line, "[")
			endIndex := strings.Index(line, "]")

			gas = line[startIndex+1 : endIndex]
			break
		}
	}

	return gas
}

// CreatePrototype - create method prototype
func CreatePrototype(item Method) string {
	proto := item.Name + "("

	for index, param := range item.Params {
		var buf bytes.Buffer

		fSet := token.NewFileSet()
		if err := format.Node(&buf, fSet, param.FieldType); err != nil {
			return ""
		}

		for index := range param.Names {
			proto += buf.String()
			if index < len(param.Names)-1 {
				proto += ","
			}
		}

		if index < len(item.Params)-1 {
			proto += ","
		}
	}

	proto += ")"

	return proto
}

// CreatePrototype - create method prototype
func ParamsLen(item Method) int {
	var count = 0

	for _, param := range item.Params {
		count += len(param.Names)
	}

	return count
}

// CalcMethodID - calculate method id with method prototype
func CalcMethodID(protoType string) int64 {
	// 计算sha3-256, 取前4字节
	d := sha3.New256()
	if _, err := d.Write([]byte(protoType)); err != nil {
		return 0
	}
	b := d.Sum(nil)
	return int64(binary.BigEndian.Uint32(b[:4]))
}

// nolint unhandled
// CalcContractAddress calculate contract address from name、version and owner
func CalcContractAddress(name string, version string, owner types.Address) types.Address {
	chainID := "local"

	hasherSHA3256 := sha3.New256()
	hasherSHA3256.Write([]byte(chainID))
	hasherSHA3256.Write([]byte(name))
	hasherSHA3256.Write([]byte(version))
	hasherSHA3256.Write([]byte(owner))
	sha := hasherSHA3256.Sum(nil)

	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha) // does not error
	rpd := hasherRIPEMD160.Sum(nil)

	hasher := ripemd160.New()
	hasher.Write(rpd)
	md := hasher.Sum(nil)

	addr := make([]byte, 0, 0)
	addr = append(addr, rpd...)
	addr = append(addr, md[:4]...)

	return chainID + base58.Encode(addr)
}

// GetContractAddress get contract address from name、version and owner
func GetContractAddress(name string, version string) types.Address {

	return "去查：" + name + ":" + version
}

// FilterImports filter import
func FilterImports(importPath string) bool {
	if importPath == `"blockchain/smcsdk/sdk/types"` {
		return false
	}

	return true
}
