package main

// #cgo LDFLAGS: -L. -lchello
// #include <stdlib.h>
// int hello();
import "C"

func main() {
	C.hello()
}