// 01_struct_sorting.go

package main
import "fmt"

type User struct {
	Name string
	Age int
}

func change(a *User, b *User) {
	temp := a.Name
	a.Name = b.Name
	b.Name = temp
}

// selection sort version
/*
func sorting(users []User) {
	n := len(users)
	for i := 0; i < n - 1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if users[j].Name < users[minIndex].Name {
				minIndex = j
			}
		}
		if minIndex != i {
			change(&users[i], &users[minIndex])
		}
	}
}
*/

func sorting(users []User) {
	n := len(users)
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if users[i].Name > users[j].Name {
				change(&users[i], &users[j])
			}
		}
	}
}

func main() {
	list := []User {
		{"Paul", 19},
		{"John", 21},
		{"Jane", 35},
		{"Abraham", 25},
	}

	sorting(list)
	for _, user := range list {
		fmt.Println(user.Name)
	}
}