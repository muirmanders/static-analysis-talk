package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// START OMIT
func main() {
	fset := token.NewFileSet()
	astFile, _ := parser.ParseFile(fset, "ast/square.go", nil, 0)

	var depth int
	ast.Inspect(astFile, func(n ast.Node) bool {
		if n == nil {
			depth--
		} else {
			depth++
			fmt.Printf("%s%T\n", strings.Repeat("  ", depth), n)
		}

		return true
	})
}

// END OMIT
