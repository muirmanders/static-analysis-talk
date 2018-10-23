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
astFile, _ := parser.ParseFile(fset, "types/time_eqeq.go", nil, 0)

defaultImporter := importer.Default()

typesConfig := &types.Config{
	Importer: defaultImporter,
}

typesInfo := &types.Info{
	Types: make(map[ast.Expr]types.TypeAndValue),
}

typesConfig.Check("time_eqeq", fset, []*ast.File{astFile}, typesInfo)
// SETUP_END OMIT

// CHECK_START OMIT
timePkg, _ := defaultImporter.Import("time")
timeDotTime := timePkg.Scope().Lookup("Time").Type()

ast.Inspect(astFile, func(n ast.Node) bool {
	binary, _ := n.(*ast.BinaryExpr)
	if binary == nil {
		return true
	}

	if typesInfo.Types[binary.X].Type == timeDotTime {
		fmt.Printf("use Equal() to compare time.Time at %s\n", fset.Position(n.Pos()))
	}

	return true
})
// CHECK_END OMIT
}
