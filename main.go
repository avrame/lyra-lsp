package main

import (
	"fmt"

	"avrameisner.com/lyra-lsp/parser"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func main() {
	tree, err := parser.Parse("1 + 2 * 3")
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}
	rootNode := tree.RootNode()
	printNode(rootNode, "")

}

func printNode(node *sitter.Node, indent string) {
	if node.IsNamed() {
		fmt.Println(indent + node.Kind())
	}
	for i := uint(0); i < node.ChildCount(); i++ {
		printNode(node.Child(i), indent+"  ")
	}
}
