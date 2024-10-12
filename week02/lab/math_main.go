// lab/bst_main.go

package main
import (
	"fmt"
	"github.com/kmin1231/go_server_session/week02/calculator"
)


func main() {
	a, b := 5, 2
    fmt.Printf("%d + %d = %d\n", a, b, calculator.Add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, calculator.Subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, calculator.Multiply(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, calculator.Divide(a, b))

	// TRY: Zero Division
	c := 0
	fmt.Printf("%d / %d = %s\n", b, c, calculator.Divide(b, c))
}