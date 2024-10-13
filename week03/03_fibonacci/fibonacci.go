package main

// #cgo LDFLAGS: -L. -lcfibonacci
// int fibonacci(int n);
import ("C")
import("fmt")

func main() {
	var n int

	fmt.Print("Please enter your integer n: ")
	fmt.Scanf("%d", &n)

	fmt.Println(C.fibonacci(C.int(n)))
}