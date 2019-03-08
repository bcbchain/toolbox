package parsecode

/*
 * pay attention to the receivers have star or not.. interesting ?
 */
import (
	"go/ast"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

const (
	methodGasPrefix    = "@:public:method:gas["
	interfaceGasPrefix = "@:public:interface:gas["
)

// Import - ..
type Import struct {
	Name string
	Path string
}

// Field - describe the Field of go ast
type Field struct {
	Names         []string // names have 0 to n member(s)
	FieldType     ast.Expr
	RelatedImport map[Import]struct{} // the field type imported package
}

// Method - the interface's method member has no receiver
type Method struct {
	Name    string
	Params  []Field
	Results []Field
}

// Function - describe the function in go ast
type Function struct {
	Method
	Comments string
	Receiver Field // go Function's receiver is an array, but I doubt how it works.
	MGas     int64
	IGas     int64
	pos      token.Pos

	GetTransferToMe bool
}

// Result is the parse result
type Result struct {
	DirectionName string
	PackageName   string
	Imports       map[Import]struct{} // current file parsed,
	AllImports    map[Import]struct{} // imports of all files, see function importsCollector

	ContractName      string
	OrgID             string
	Version           string
	Versions          []string
	Author            string
	ContractStructure string

	Stores      []Field
	StoreCaches []Field

	InitChain Function
	Functions []Function

	Receipts []Method

	UserStruct map[string]ast.GenDecl

	ErrFlag   bool
	ErrorDesc []string
	ErrorPos  []token.Pos
}

type visitor struct {
	res     *Result
	depth   int
	inClass bool // walk in the contract structure, it's a flag (parse store)
}

func newVisitor() visitor {
	errDesc := make([]string, 0)
	errPos := make([]token.Pos, 0)
	funcs := make([]Function, 0)
	receipts := make([]Method, 0)
	stores := make([]Field, 0)
	caches := make([]Field, 0)
	imp := make(map[Import]struct{})
	aImp := make(map[Import]struct{})
	us := make(map[string]ast.GenDecl)
	res := Result{
		Imports:     imp,
		AllImports:  aImp,
		Functions:   funcs,
		Receipts:    receipts,
		Stores:      stores,
		StoreCaches: caches,
		UserStruct:  us,
		ErrorDesc:   errDesc,
		ErrorPos:    errPos,
	}
	depth := 0
	return visitor{
		res:   &res,
		depth: depth,
	}
}

// Visit is a visitor method walk through the AST node in depth-first order, same to inspect
func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.Ident: // parse package name
		v.parsePackageName(d)
	case *ast.FuncDecl: // parse function annotation (gas,constructor)
		v.parseAllFunc(d)
	case *ast.Field: // parse function annotation (store)
		v.parseStoreField(d)
	case *ast.GenDecl:
		v.parseGenDeclare(d)
	}

	v.depth++
	return v
}

func (v *visitor) parseField(d *ast.Field) Field {
	f := Field{}
	names := make([]string, 0)
	for _, id := range d.Names {
		names = append(names, id.Name)
	}
	f.Names = names

	f.FieldType = d.Type

	imports := make(map[Import]struct{})
	getFieldImports(d.Type, v.res.Imports, imports)

	f.RelatedImport = imports

	return f
}

func getFieldImports(t ast.Expr, imports map[Import]struct{}, filtered map[Import]struct{}) {
	getImport := func(pkg string) Import {
		for imp := range imports {
			if pkg == imp.Name || strings.HasSuffix(imp.Path, "/"+pkg+"\"") {
				return imp
			}
		}
		return Import{}
	}
	doSel := func(sel *ast.SelectorExpr) {
		if id, okID := sel.X.(*ast.Ident); okID {
			pkg := id.Name
			imp := getImport(pkg)
			if imp.Path != "" {
				filtered[imp] = struct{}{}
			}
		}
	}

	if sel, okSel := t.(*ast.SelectorExpr); okSel {
		doSel(sel)
	} else if star, okStar := t.(*ast.StarExpr); okStar {
		if sel, okSel := star.X.(*ast.SelectorExpr); okSel {
			doSel(sel)
		}
	} else if arr, okArr := t.(*ast.ArrayType); okArr {
		if sel, okSel := arr.Elt.(*ast.SelectorExpr); okSel {
			doSel(sel)
		}
	} else if mt, okM := t.(*ast.MapType); okM {
		if sel, okSel := mt.Key.(*ast.SelectorExpr); okSel {
			doSel(sel)
		}
		if vt, okV := mt.Value.(*ast.SelectorExpr); okV {
			doSel(vt)
		} else if vs, okVS := mt.Value.(*ast.StarExpr); okVS {
			if sel, okSel := vs.X.(*ast.SelectorExpr); okSel {
				doSel(sel)
			}
		} else if ar, okA := mt.Value.(*ast.ArrayType); okA {
			if sel, okS := ar.Elt.(*ast.SelectorExpr); okS {
				doSel(sel)
			}
		}
	}
}

func (v *visitor) parseStoreField(d *ast.Field) {
	if v.inClass && d.Doc != nil {
		list := strings.Split(d.Doc.Text(), "\n")
		for _, doc := range list {
			doc = strings.TrimSpace(doc)
			if strings.HasPrefix(doc, "@:public:store:cache") {
				// fmt.Println("FIELD::: name(", d.Names, ") =>[[", d.Doc.Text(), "]]")
				cacheField := v.parseField(d)
				v.res.StoreCaches = append(v.res.StoreCaches, cacheField)
			} else if strings.HasPrefix(doc, "@:public:store") {
				// fmt.Println("FIELD::: name(", d.Names, ") =>[[", d.Doc.Text(), "]]")
				storeField := v.parseField(d)
				v.res.Stores = append(v.res.Stores, storeField)
			}
		}
	}
}

func (v *visitor) parseAllFunc(d *ast.FuncDecl) {
	if d.Doc != nil {
		// fmt.Println("FUNCTION::: name(", d.Name.Name, ") =>[[", d.Doc.Text(), "]]")
		if v.hasConstructorInComments(d) && d.Name.Name == "InitChain" {
			v.res.InitChain = v.parseInitChain(d)
		}
		mGas := v.getGasFromComments(d, methodGasPrefix)
		iGas := v.getGasFromComments(d, interfaceGasPrefix)
		if mGas+iGas > 0 {
			f := v.parseFunction(d)
			f.MGas = mGas
			f.IGas = iGas
			// fmt.Println(f)
			v.res.Functions = append(v.res.Functions, f)
		}
	}
}

func (v *visitor) parseFunction(d *ast.FuncDecl) Function {
	f := Function{}
	if d.Recv != nil {
		f.Receiver = v.parseField(d.Recv.List[0])
	}
	f.Name = d.Name.Name
	f.pos = d.Pos()
	f.Params = make([]Field, 0)
	for _, param := range d.Type.Params.List {
		f.Params = append(f.Params, v.parseField(param))
	}
	f.Results = make([]Field, 0)
	if d.Type.Results != nil {
		for _, res := range d.Type.Results.List {
			f.Results = append(f.Results, v.parseField(res))
		}
	}
	f.Comments = d.Doc.Text()

	return f
}

func (v *visitor) parseInitChain(d *ast.FuncDecl) Function {
	f := Function{}

	if len(d.Type.Params.List) > 0 {
		v.reportErr("InitChain must have no params", d.Type.Params.Pos())
	}
	if len(d.Recv.List) != 1 {
		v.reportErr("InitChain has wrong receiver", d.Recv.Pos())
	}
	f.Receiver = v.parseField(d.Recv.List[0])
	f.pos = d.Pos()

	return f
}

func (v *visitor) hasConstructorInComments(d *ast.FuncDecl) bool {
	l := strings.Split(d.Doc.Text(), "\n")
	for _, c := range l {
		c = strings.TrimSpace(c)
		if strings.HasPrefix(c, "@:constructor") {
			return true
		}
	}
	return false
}

func (v *visitor) getGasFromComments(d *ast.FuncDecl, prefix string) int64 {
	l := strings.Split(d.Doc.Text(), "\n")
	for _, c := range l {
		c = strings.TrimSpace(c)
		if strings.HasPrefix(c, prefix) {
			gas := c[len(prefix) : len(c)-1]
			i, e := strconv.ParseInt(gas, 10, 0)
			if e != nil {
				v.reportErr("method gas not a number", d.Pos()) // TODO report error?
			}
			return i
		}
	}
	return 0
}

// Id @ level 1 is package name
func (v *visitor) parsePackageName(d *ast.Ident) {
	if v.depth == 1 {
		v.res.PackageName = d.Name
	}
}

func (v *visitor) parseGenDeclare(d *ast.GenDecl) {
	if d.Tok == token.IMPORT {
		v.parseImport(d)
	} else if d.Tok == token.VAR {
		v.parseVar(d)
	} else {
		v.parseStructsNInterface(d)
	}
}

func (v *visitor) parseImport(d *ast.GenDecl) {
	for _, spec := range d.Specs {
		if imp, ok := spec.(*ast.ImportSpec); ok {
			path := imp.Path.Value
			if _, okPKG := WhiteListPKG[path]; !okPKG && !isWhitePrefix(path) {
				v.res.ErrFlag = true
				v.res.ErrorDesc = append(v.res.ErrorDesc, "INVALID IMPORT: "+path)
				v.res.ErrorPos = append(v.res.ErrorPos, d.Pos())
			} else {
				im := Import{Path: path}
				if imp.Name != nil {
					im.Name = imp.Name.Name
				}
				v.res.Imports[im] = struct{}{}
			}
		}
	}
}

func isWhitePrefix(path string) bool {
	for _, pre := range WhiteListPkgPrefix {
		if strings.HasPrefix(path, pre) {
			return true
		}
	}
	return false
}

func (v *visitor) parseVar(d *ast.GenDecl) {
	for _, spec := range d.Specs {
		if value, ok := spec.(*ast.ValueSpec); ok {
			for _, name := range value.Names {
				if name.Name == "_" {
					continue
				}
				if v.depth <= 1 {
					v.reportErr("GLOBAL VAR:"+name.Name, d.Pos())
				}
			}
		}
	}
}

func (v *visitor) parseStructsNInterface(d *ast.GenDecl) {
	for _, spec := range d.Specs {
		if typ, ok := spec.(*ast.TypeSpec); ok {
			if v.depth == 1 && d.Doc != nil {
				if _, ok := typ.Type.(*ast.InterfaceType); ok {
					v.parseInterface(d, typ)
				}
				if _, ok := typ.Type.(*ast.StructType); ok {
					v.parseStructs(d, typ)
				}
			}
		}
	}
}

// parse interface annotation (receipt)
func (v *visitor) parseInterface(d *ast.GenDecl, typ *ast.TypeSpec) {
	// fmt.Println("INTERFACE::: name(", typ.Name, ") =>[[", d.Doc.Text(), "]]")
	if v.isReceipt(d) {
		it, _ := typ.Type.(*ast.InterfaceType)
		for _, am := range it.Methods.List {
			if am.Names[0].Name[:4] != "emit" {
				v.reportErr("Method os receipt interface must start with 'emit'", typ.Pos())
			}
			if m, ok := am.Type.(*ast.FuncType); ok {
				params := make([]Field, 0)
				for _, p := range m.Params.List {
					params = append(params, v.parseField(p))
				}
				results := make([]Field, 0)
				if m.Results != nil {
					for _, r := range m.Results.List {
						results = append(results, v.parseField(r))
					}
				}
				v.res.Receipts = append(v.res.Receipts, Method{
					Name:    am.Names[0].Name,
					Params:  params,
					Results: results,
				})
			}
		}
	}
}

func (v *visitor) isReceipt(d *ast.GenDecl) bool {
	for _, l := range d.Doc.List {
		doc := strings.TrimSpace(l.Text)
		if strings.ToLower(doc) == "//@:public:receipt" {
			return true
		}
	}
	return false
}

// parse struct annotation (contract,version,organization,author)
// nolint cyclomatic ... 這復雜度不高啊，拆開反倒不利於閱讀了，我設置了 --cyclo-over=20，其實我可以不用這行注釋了，娃哈哈
func (v *visitor) parseStructs(d *ast.GenDecl, typ *ast.TypeSpec) {
	if v.depth == 1 {
		v.res.UserStruct[typ.Name.Name] = *d
	}
	docs := strings.Split(d.Doc.Text(), "\n")
	for _, comment := range docs {
		comment = strings.TrimSpace(comment)
		if len(comment) > 11 && comment[:11] == "@:contract:" {
			if v.res.ContractName != "" {
				v.reportErr("ContractName is already set", d.Pos())
			}
			v.res.ContractName = strings.TrimSpace(comment[11:])
			if !v.checkRegex(v.res.ContractName, contractNameExpr) {
				v.reportErr("ContractName -> invalid format:"+v.res.ContractName, d.Pos())
			}
			v.inClass = true
			//if v.res.ContractName != v.res.PackageName {
			//	v.reportErr("PackageName("+v.res.PackageName+")!=ContractName("+v.res.ContractName+")", d.Pos())
			//}
		}
		if len(comment) > 10 && comment[:10] == "@:version:" {
			v.res.Version = strings.TrimSpace(comment[10:])
			if !v.checkRegex(v.res.Version, versionExpr) {
				v.reportErr("contract Version -> invalid format:"+v.res.Version, d.Pos())
			}
		}
		if len(comment) > 15 && comment[:15] == "@:organization:" {
			v.res.OrgID = strings.TrimSpace(comment[15:])
			if !v.checkRegex(v.res.OrgID, organizationExpr) {
				v.reportErr("Organization -> invalid format:"+v.res.OrgID, d.Pos())
			}
		}
		if len(comment) > 9 && comment[:9] == "@:author:" {
			v.res.Author = strings.TrimSpace(comment[9:])
			if !v.checkRegex(v.res.Author, authorExpr) {
				v.reportErr("Author -> invalid format:"+v.res.Author, d.Pos())
			}
		}
	}
	if v.inClass && v.depth == 1 {
		if v.res.ContractStructure != "" {
			v.reportErr("You have more ContractStructure:"+v.res.ContractStructure+","+typ.Name.Name, d.Pos())
		}
		v.res.ContractStructure = typ.Name.Name
		if !v.checkRegex(v.res.ContractStructure, contractClassExpr) {
			v.reportErr("ContractStructure -> invalid format:"+v.res.ContractStructure, d.Pos())
		}
		if !v.checkSDKDeclare(typ) {
			v.reportErr("Contract's first field Must be 'sdk sdk.ISmartContract'", d.Pos())
		}
	}
}

func (v *visitor) reportErr(desc string, pos token.Pos) {
	v.res.ErrFlag = true
	v.res.ErrorDesc = append(v.res.ErrorDesc, desc)
	v.res.ErrorPos = append(v.res.ErrorPos, pos)
}

func (v *visitor) printContractInfo() {
	if v.res.ContractName != "" {
		//fmt.Println("PackageName:", v.res.PackageName)
		//fmt.Println("ContractName:", v.res.ContractName)
		//fmt.Println("ContractStructure:", v.res.ContractStructure)
		if len(v.res.InitChain.Receiver.Names) != 1 {
			v.reportErr("InitChain has no receiver", v.res.InitChain.pos)
		} else {
			//fmt.Println("InitChain's Receiver:", v.res.InitChain.Receiver.Names[0])
		}
	}
	if v.res.Version != "" {
		//fmt.Println("Version:", v.res.Version)
	}
	if v.res.OrgID != "" {
		//fmt.Println("Organization:", v.res.OrgID)
	}
	if v.res.Author != "" {
		//fmt.Println("Author:", v.res.Author)
	}
}

// all declare must obey our naming specification
func (v *visitor) checkRegex(obj string, regex string) bool {
	r, e := regexp.Compile(regex)
	if e != nil {
		return false
	}
	return r.MatchString(obj)
}

// contract structure's first field must be "sdk sdk.ISmartContract"
func (v *visitor) checkSDKDeclare(typ *ast.TypeSpec) bool {
	st, _ := typ.Type.(*ast.StructType)
	l := st.Fields.List
	if len(l) == 0 || len(l[0].Names) == 0 {
		return false
	}
	if l[0].Names[0].Name != "sdk" {
		return false
	}
	if id, ok := l[0].Type.(*ast.SelectorExpr); !ok {
		return false
	} else if id.Sel.Name != "ISmartContract" {
		return false
	} else if x, ok := id.X.(*ast.Ident); !ok {
		return false
	} else if x.Name != "sdk" {
		return false
	}

	return true
}

func importsCollector(res *Result) {
	for imp := range res.Imports {
		res.AllImports[imp] = struct{}{}
		delete(res.Imports, imp)
	}
}
