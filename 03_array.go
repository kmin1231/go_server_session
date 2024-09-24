package main
import "fmt"

func sumArray(array) {
	var sum int = 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func fundMaxMin(array) {
	var max int = -1
	var min int = -1
}


func main() {
	defer fmt.Println("프로그램이 종료되었습니다.")

	arr := []int{3, 5, 1, 2, 0}

	for i := range arr {
		fmt.Println(i)
	}

	var length int = len(arr)
	switch length {
		case length < 3: fmt.Println("배열의 길이가 짧습니다.")
		case length == 3: fmt.Println("배열의 길이가 적당합니다.")
		case length > 3: fmt.Println("배열의 길이가 깁니다.")	
	}

}