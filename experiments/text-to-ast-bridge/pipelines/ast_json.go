package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Simplified AST Node for JSON serialization to demonstrate token inflation
type JSONNode struct {
	Type     string      `json:"type"`
	Name     string      `json:"name,omitempty"`
	Value    string      `json:"value,omitempty"`
	Children []JSONNode  `json:"children,omitempty"`
}

func astToJSON(node ast.Node) JSONNode {
	var j JSONNode
	switch n := node.(type) {
	case *ast.File:
		j.Type = "File"
		for _, decl := range n.Decls {
			j.Children = append(j.Children, astToJSON(decl))
		}
	case *ast.FuncDecl:
		j.Type = "FuncDecl"
		j.Name = n.Name.Name
		j.Children = append(j.Children, astToJSON(n.Type))
		if n.Body != nil {
			j.Children = append(j.Children, astToJSON(n.Body))
		}
	case *ast.FuncType:
		j.Type = "FuncType"
		if n.Params != nil {
			for _, p := range n.Params.List {
				j.Children = append(j.Children, astToJSON(p))
			}
		}
	case *ast.Field:
		j.Type = "Field"
		for _, name := range n.Names {
			j.Children = append(j.Children, JSONNode{Type: "Ident", Name: name.Name})
		}
	case *ast.BlockStmt:
		j.Type = "BlockStmt"
		for _, stmt := range n.List {
			j.Children = append(j.Children, astToJSON(stmt))
		}
	// Add more cases as needed for the experiment
	default:
		j.Type = fmt.Sprintf("%T", node)
	}
	return j
}

func main() {
	src := `package main
func Greeter(name string) string { return "Hello" }
`
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", src, 0)
	
	jsonAST := astToJSON(f)
	b, _ := json.MarshalIndent(jsonAST, "", "  ")
	fmt.Println(string(b))
}
