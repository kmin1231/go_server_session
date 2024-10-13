// 01_goroutine_calc.go

package main
import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Please enter your  FIRST integer: ")
	fmt.Scan(&num1)

	fmt.Print("Please enter your SECOND integer: ")
	fmt.Scan(&num2)

	fmt.Println()

	// to specify 'buffer size' -> 'make(chan, int 2)' form
	inputSumChan := make(chan int, 2)
	outputSumChan := make(chan int, 2)
	inputMulChan := make(chan int, 2)
	outputMulChan := make(chan int, 2)

	// creates a 'goroutine' using an anonymous function
	go func() {
		inputSumChan <- num1
		inputSumChan <- num2

		a := <- inputSumChan
		b := <- inputSumChan
		sum := a + b

		outputSumChan <- sum
	}()	// execute the function

	go func() {
		inputMulChan <- num1
		inputMulChan <- num2

		a := <- inputMulChan
		b := <- inputMulChan
		product := a * b

		outputMulChan <- product
	}()	// execute the function

	// receives results from each goroutine in the 'main' goroutine
	sum := <- outputSumChan
	product := <- outputMulChan

    fmt.Printf(">>   Sum   Result: %d + %d = %d\n", num1, num2, sum)
    fmt.Printf(">> Product Result: %d * %d = %d\n", num1, num2, product)
}