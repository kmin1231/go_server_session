// lab/bst_main.go

package main
import (
	"fmt"
	"github.com/kmin1231/go_server_session/week02/tree"
)

/* to use 'bst' module,
	go get github.com/kmin1231/go_server_session/week02/tree
	* package name 'bst'
	* command execution directory: /week02/lab
*/

func main() {
	root := bst.MakeNode(50)
	root.InsertNode(54)
	root.InsertNode(76)
	root.InsertNode(45)
	root.InsertNode(24)
	root.InsertNode(47)
	root.InsertNode(94)

	fmt.Println("InOrder (Sorted Order):")
	bst.InOrder(root)
	fmt.Println()
}