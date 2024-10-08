package main

// #include <stdio.h>
// int fibonacci(int n) {
// 	if (n <= 0) return 0;
// 	if (n == 1) return 1;
// 	return fibonacci(n-1) + fibonacci(n-2);
// }
import ("C")
import("fmt")

func main() {
	var n int

	fmt.Println("Please enter your integer n:")
	fmt.Scanf("%d", &n)

	fmt.Println(C.fibonacci(C.int(n)))
}