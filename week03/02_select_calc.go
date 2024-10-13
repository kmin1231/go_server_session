// 02_select_calc.go

package main
import "fmt"

// ref. Tanmay Teaches Go (Korean ver.) - p149

func main() {
	inputSumChan := make(chan int, 2)
	inputMulChan := make(chan int, 2)
	finishChan := make(chan int)
	outputSumChan := make(chan int)
	outputMulChan := make(chan int)
	
	go func(inputChan chan int, finishChan chan int) {
		for {
			select {
			case num1 := <- inputChan:
				num2 := <- inputChan
				outputSumChan <- num1 + num2
			case _ = <- finishChan:
				return
			}
		}
	}(inputSumChan, finishChan)

	go func(inputChan chan int, finishChan chan int) {
		for {
			select {
			case num1 := <- inputChan:
				num2 := <- inputChan
				outputMulChan <- num1 * num2
			case _ = <- finishChan:
				return
			}
		}
	}(inputMulChan, finishChan)

	var num1, num2 int

	fmt.Print("Please enter your  FIRST integer: ")
	fmt.Scan(&num1)

	fmt.Print("Please enter your SECOND integer: ")
	fmt.Scan(&num2)

	inputSumChan <- num1
	inputSumChan <- num2
	inputMulChan <- num1
	inputMulChan <- num2

	fmt.Println()

	for i := 0; i < 2; i++ {
		select {
		case sum := <- outputSumChan:
			fmt.Printf(">>   Sum   Result: %d + %d = %d\n", num1, num2, sum)
		case product := <- outputMulChan:
			fmt.Printf(">> Product Result: %d * %d = %d\n", num1, num2, product)
		}
	}
   
	finishChan <- 1
	// close(finishChan)
}