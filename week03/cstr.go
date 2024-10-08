package main

// #include <stdio.h>
// char *getName(int idx) {
// 	if (idx == 1)
// 		return "Tanmay Bakshi";
// 	if (idx == 2)
// 		return "Baheer Kamal";
// 	return "Invalid index";
// }
import "C"
import "fmt"

func main() {
	cstr := C.getName(C.int(2))
	fmt.Println(C.GoString(cstr))
}