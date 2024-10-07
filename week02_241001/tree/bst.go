package bst
import "fmt"

type Node struct {
	value int
	Left *Node
	Right *Node
}

func MakeNode(num int) *Node {
	return &Node{ value: num, Left: nil, Right: nil }
}

func (tree *Node) InsertNode(num int) {
	if (num < tree.value) {
		if (tree.Left == nil) {
			tree.Left = MakeNode(num)
		} else {
			tree.Left.InsertNode(num)
		}
	} else if (num > tree.value) {
		if (tree.Right == nil) {
			tree.Right = MakeNode(num)
		} else {
			tree.Right.InsertNode(num)
		}
	}
}

func InOrder(n *Node) {
	if (n != nil) {
		InOrder(n.Left)
		fmt.Print(n.value, " ")
		InOrder(n.Right)
	}
}
