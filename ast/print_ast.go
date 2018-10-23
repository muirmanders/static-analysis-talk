package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// START OMIT
func main() {
	fset := token.NewFileSet()
	astFile, _ := parser.ParseFile(fset, "ast/single_var.go", nil, 0)

	ast.Print(fset, astFile)
}

// END OMIT
