// 02_odd_or_even.go

package main
import "fmt"

func main() {
	var num int
	fmt.Print("Please enter an integer: ")
	fmt.Scanf("%d", &num)

	if ((num%2) == 0) {
		fmt.Println(">> Even!")
	} else {fmt.Println(">> Odd!")}
}