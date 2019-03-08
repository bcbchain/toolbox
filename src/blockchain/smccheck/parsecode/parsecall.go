package parsecode

import (
	"go/ast"
	"strings"
)

// Callee 被調用的方法
type Callee struct {
	Select     []string
	LastImport Import
}

var transferCallee = []Callee{
	{
		Select: []string{
			"GetTransferToMe",
			"Message",
			"ISmartContract"},
		LastImport: Import{
			Name: "",
			Path: "blockchain/smcsdk/sdk",
		},
	},
	{
		Select: []string{
			"GetTransferToMe",
			"IMessage"},
		LastImport: Import{
			Name: "",
			Path: "blockchain/smcsdk/sdk",
		},
	},
}

func (v *visitor) parseCall(callee Callee, fileMap map[string]*ast.File) {
	for _, f := range fileMap {
		ast.Inspect(f, func(node ast.Node) bool {
			switch node.(type) {
			case *ast.FuncDecl:
				fun := node.(*ast.FuncDecl)
				if v.isCalling(callee, fun) {
					newCallee := v.parseFunction(fun)
					for idx, contractFunction := range v.res.Functions {
						if newCallee.Name == contractFunction.Name &&
							newCallee.Receiver.FieldType == contractFunction.Receiver.FieldType {
							v.res.Functions[idx].GetTransferToMe = true
							return false
						}
					}
					caller := v.parseCaller(fun)
					v.parseCall(caller, fileMap)
				}
			}
			return true
		})
	}
}

func (v *visitor) parseCaller(d *ast.FuncDecl) Callee {
	names := make([]string, 0)
	names = append(names, d.Name.Name)
	if d.Recv != nil && len(d.Recv.List) > 0 {
		receiverType := d.Recv.List[0].Type
		typ := ExpandTypeNoStar(Field{FieldType: receiverType})
		names = append(names, typ)
	}
	return Callee{Select: names}
}

// nolint cyclomatic
func (v *visitor) isCalling(callee Callee, d *ast.FuncDecl) bool {
	exists := false
	ast.Inspect(d, func(n ast.Node) bool {
		if _, ok0 := n.(*ast.CallExpr); !ok0 {
			return true
		}

		target := n
		for _, sel := range callee.Select {
			if call, ok1 := target.(*ast.CallExpr); ok1 {
				if id, okID := call.Fun.(*ast.Ident); okID {
					if id.Name == sel {
						exists = true
						return false
					}
				} else if fun, ok2 := call.Fun.(*ast.SelectorExpr); !ok2 || fun.Sel.Name != sel {
					return true
				} else {
					target = fun.X
				}

			} else if id, ok3 := target.(*ast.Ident); ok3 {
				if assign, ok4 := id.Obj.Decl.(*ast.AssignStmt); ok4 {
					for i := 0; i < len(assign.Lhs); i++ {
						aid, ok5 := assign.Lhs[i].(*ast.Ident)
						if !ok5 {
							return true
						}
						if id.Name == aid.Name {
							if assignF, okf := assign.Rhs[i].(*ast.CallExpr); okf {
								if assignSel, okSel := assignF.Fun.(*ast.SelectorExpr); okSel {
									target = assignSel.X
								}
							}
						}
					}
				} else if field, ok6 := id.Obj.Decl.(*ast.Field); ok6 {
					if sel2, ok7 := field.Type.(*ast.SelectorExpr); ok7 {
						if sel != sel2.Sel.Name {
							return true
						}
						target = sel2.X
					} else if sel3, okSel3 := field.Type.(*ast.StarExpr); okSel3 {
						if id, okID := sel3.X.(*ast.Ident); okID && id.Name == sel {
							exists = true
							return false
						}
					} else if id, okID := field.Type.(*ast.Ident); okID && id.Name == sel {
						exists = true
						return false
					}
				} else if val, okVal := id.Obj.Decl.(*ast.ValueSpec); okVal {
					if idt, okt := val.Type.(*ast.Ident); okt && idt.Name == sel {
						exists = true
						return false
					}
				}
			} else if sel2, okT := target.(*ast.SelectorExpr); okT {
				if id, okID := sel2.X.(*ast.Ident); okID {
					if d.Recv != nil && id.Name != d.Recv.List[0].Names[0].Name {
						return true
					}
					target = sel2.Sel
				}
			}
		}

		pkg, ok8 := target.(*ast.Ident)
		if !ok8 {
			return true
		}
		paths := strings.Split(callee.LastImport.Path, "/")
		pkgName := paths[len(paths)-1]
		if pkg.Name == callee.LastImport.Name || pkg.Name == pkgName {
			exists = true
			return false
		}
		return true
	})

	return exists
}
