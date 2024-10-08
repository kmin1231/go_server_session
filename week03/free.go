package main

// #include <stdio.h>
// #include <stdlib.h>
// void printString(char *str) {
//     printf("%s\n", str);
// }
import "C"
import "unsafe" // to avoid memory leak

func main() {
    a := C.CString("This is from Golang")
    C.printString(a)
    C.free(unsafe.Pointer(a))  // 메모리 할당 해제 (generic pointer: type 가리지 않음)
}