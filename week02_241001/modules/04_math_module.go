package module
import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a *  b
}

func Divide(a int, b int) int {
	if (b == 0) {
		fmt.Println("Error: division by ZERO!")
	}
	return a / b
}