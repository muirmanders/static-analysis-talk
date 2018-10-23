package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
// SETUP_START OMIT
fset := token.NewFileSet()
astFile, _ := parser.ParseFile(fset, "iterator/iterator.go", nil, 0)

defaultImporter := importer.Default()

typesConfig := &types.Config{
	Importer: defaultImporter,
}

typesInfo := &types.Info{
	Types: make(map[ast.Expr]types.TypeAndValue),
	Uses:  make(map[*ast.Ident]types.Object),
}

iterPkg, _ := typesConfig.Check("iterator", fset, []*ast.File{astFile}, typesInfo)

var (
	iterType          = iterPkg.Scope().Lookup("Iter").Type()
	nextMethod, _, _  = types.LookupFieldOrMethod(iterType, true, iterPkg, "Next")
	closeMethod, _, _ = types.LookupFieldOrMethod(iterType, true, iterPkg, "Close")
	nextCalls         = make(map[types.Object]bool)
)
// SETUP_END OMIT

// CHECK_START OMIT
ast.Inspect(astFile, func(n ast.Node) bool {
	call, _ := n.(*ast.CallExpr) // Looking for .Next() or .Close()
	if call == nil {
		return true
	}

	sel, _ := call.Fun.(*ast.SelectorExpr) // foo.bar
	if sel == nil {
		return true
	}

	if typesInfo.ObjectOf(sel.Sel) == nextMethod {
		nextCalls[typesInfo.ObjectOf(sel.X.(*ast.Ident))] = true
	}

	if typesInfo.ObjectOf(sel.Sel) == closeMethod {
		delete(nextCalls, typesInfo.ObjectOf(sel.X.(*ast.Ident)))
	}
	return true
})

for missingClose := range nextCalls {
	fmt.Printf("Iter missing Close() at %s\n", fset.Position(missingClose.Pos()))
}
// CHECK_END OMIT
}
