package main

import (
	"fmt"
	"week02_241001/modules"
)

func main() {
	a := 2
	b := 3
	// sum := math_module.Add(2, 3)
	sum := modules.Add(a, b)
	fmt.Println("Sum:", sum)
}

// ex4 폴더 내에 main.go
// golangStudy 내의 ex3, ex4 폴더
// github 시간 차

// bst
/*
type Node struct
Value int, Left *Node, Right *Node
*/
(tree *Node) Insert(num int)