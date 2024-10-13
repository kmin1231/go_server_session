// calculator/math_module.go

package calculator
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

func Divide(a int, b int) string {
	if (b == 0) {
		fmt.Print("Error: division by ZERO --- ")
		return "?"
	}
	return fmt.Sprintf("%d", a/b)
}