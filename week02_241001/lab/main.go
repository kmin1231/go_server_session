package main

import (
	"fmt"
	"github.com/kmin1231/go_server_session/week02_241001/tree"
)

func main() {
	root := tree.MakeNode(50)
	root.InsertNode(54)
	root.InsertNode(76)
	root.InsertNode(45)
	root.InsertNode(24)
	root.InsertNode(47)
	root.InsertNode(94)

	fmt.Println("InOrder (Sorted Order):")
	tree.InOrder(root)
	fmt.Println()
}