package main

import (
	"fmt"
	"go/parser"
	"go/token"

	schemacmp "github.com/Oncilla/capnp-schema-cmp"
)

func main() {

	fset := token.NewFileSet() // positions are relative to fset

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "input.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	ext := schemacmp.Extractor{File: f}
	schemas := ext.Extract()

	for name, value := range schemas {
		fmt.Printf("const %s = %s\n", name, value)
	}

}
