package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
// START OMIT
fset := token.NewFileSet()
astFile, _ := parser.ParseFile(fset, "ast/short_funcs.go", nil, 0)

ast.Inspect(astFile, func(n ast.Node) bool {
	funcDecl, _ := n.(*ast.FuncDecl)
	if funcDecl == nil {
		return true
	}

	if len(funcDecl.Name.Name) < 3 {
		fmt.Printf("function %q name too short (%s)\n",
			funcDecl.Name,
			fset.Position(funcDecl.Pos()),
		)
	}

	return true
})
// END OMIT
}
