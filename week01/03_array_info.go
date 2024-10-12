// 03_array_info.go

package main
import "fmt"

func sumArray(array []int) int {
	// input parameter: array
	// return type: int
	var sum int = 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func findMaxMin(array []int) (int, int) {
	var max int = array[0]	// initialize
	var min int = array[0]	// initialize

	for _, value := range array {
		// not using index of the array (underscore)
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return max, min
}

func main() {
	// print termination message, using 'defer'
	/* DEFER
	delays the execution of a function until the surrounding function returns;
	ensures that the deferred function is always executed, no matter how the function ends
	*/
	defer fmt.Println(">> The program has finished!")
	// termination message to be printed just before the 'main' function exits
	// even if an error occurs or the function returns early

	// declare and initialize array 'arr'
	arr := []int{3, 5, 1, 2, 0}

	fmt.Println(">> Array Info")

	// print all elements of the array
	fmt.Print("Elements: ")	// cf. Print, Println, Printf
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Println()	// cf. [Python] print()

	// print the sum of the array
	sum := sumArray(arr)
	fmt.Println("Sum:", sum)
	
	// print the max & min values of the array
	max, min := findMaxMin(arr)
	fmt.Println("Max value:", max)
	fmt.Println("Min value:", min)

	// check the 'length' of the array
	var length int = len(arr)
	fmt.Print("Length:", length)

	// (Atn.) switch without expression to use comparison operators
	switch {
		case length < 3: fmt.Println(" -- The array is too SHORT.")
		case length == 3: fmt.Println(" -- The array length is ADEQUATE.")
		case length > 3: fmt.Println(" -- The array is too LONG.")	
	}
}