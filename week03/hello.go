package main

// #include <stdio.h>
// void printHelloWorld() {
//     printf("Hello, world!\n");
// }
import "C"

func main() {
	C.printHelloWorld()
}