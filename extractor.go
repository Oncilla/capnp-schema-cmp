package schemacmp

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

// Extractor extracts all schemas from a go file
type Extractor struct {
	File *ast.File
}

// Extract returns a mapping from all schema constant names to their values.
func (ext Extractor) Extract() map[string]string {
	v := make(visitor)
	ast.Walk(v, ext.File)
	return v
}

type visitor map[string]string

func (v visitor) Visit(n ast.Node) ast.Visitor {
	switch d := n.(type) {
	case *ast.GenDecl:
		if d.Tok != token.CONST {
			return v
		}
		for _, spec := range d.Specs {
			if value, ok := spec.(*ast.ValueSpec); ok {
				if !strings.HasPrefix(value.Names[0].Name, "schema") {
					continue
				}
				v[value.Names[0].Name] = concat(value.Values[0])
			}
		}
	}
	return v
}

func concat(value ast.Expr) string {
	switch e := value.(type) {
	case *ast.BasicLit:
		return strings.Trim(e.Value, `"`)
	case *ast.BinaryExpr:
		return concat(e.X) + concat(e.Y)
	default:
		panic(fmt.Sprintf("Invalid value type: %T", e))
	}
}
