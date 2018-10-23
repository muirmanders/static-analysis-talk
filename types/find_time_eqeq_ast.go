package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func astThenWhat() {
fset := token.NewFileSet()
astFile, _ := parser.ParseFile(fset, "types/time_eqeq.go", nil, 0)

// START OMIT
ast.Inspect(astFile, func(n ast.Node) bool {
	binary, _ := n.(*ast.BinaryExpr)
	if binary == nil {
		return true
	}

	// we have a binary expression... now what?

	return true
})
// END OMIT
}
