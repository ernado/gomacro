package typeutil

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
)

func CreateInfoAndPkg(path, name string) (*types.Package, *types.Info) {
	pkg := types.NewPackage(path, name)
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Scopes:     make(map[ast.Node]*types.Scope),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}

	return pkg, info
}

func ParseTypes(src string) (Types, error) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return Types{}, err
	}

	t := newTypes(".", "expr")
	conf := &types.Config{}
	err = types.NewChecker(conf, fset, t.pkg, t.info).Files([]*ast.File{file})
	if err != nil {
		return Types{}, err
	}

	return t, nil
}
