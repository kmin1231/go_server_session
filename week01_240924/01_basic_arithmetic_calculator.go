package main
import "fmt"

func main() {
	var num1 int
	var num2 int
	var opr string
	var result int

	fmt.Println("Please enter your FIRST integer:")
	fmt.Scanf("%d", &num1)

	fmt.Println("Please enter your SECOND integer:")
	fmt.Scanf("%d", &num2)

	fmt.Println("Please enter an OPERATOR (+, -, *, /):")
	fmt.Scanf("%s", &opr)

	switch opr {
		case "+": result = num1 + num2
		case "-": result = num1 - num2
		case "*": result = num1 * num2
		case "/":
			if (opr == "/") && (num2 == 0) {
				fmt.Println("Error: division by ZERO!")
				}
			result = num1 / num2
		default: fmt.Println("Error: wrong operator!")
	}

	fmt.Println("The result is:", result)
}