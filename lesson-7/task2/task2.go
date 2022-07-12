package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	amount, err := getAmount("file.go", "somePrint1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("amount: ", amount)
}

func getAmount(srcFileName, funcName string) (int, error) {
	var i int
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, srcFileName, nil, 0)
	if err != nil {
		return i, err
	}

	var requiredFun *ast.FuncDecl
	for _, decl := range astFile.Decls {
		if fn, isFn := decl.(*ast.FuncDecl); isFn {
			if fn.Name.Name == funcName {
				requiredFun = fn
			}
		}
	}
	if requiredFun == nil {
		return 0, errors.New("function " + funcName + " not found in file " + srcFileName)
	}

	ast.Inspect(requiredFun, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.GoStmt:
			i++
			fmt.Println(n)
		}
		return true
	})

	return i, nil

}
