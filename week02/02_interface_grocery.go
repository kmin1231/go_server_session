// 02_interface_grocery.go

package main
import "fmt"

type Snack struct {
	Name string
	Price int
}

type Drink struct {
	Name string
	Price int
}

type Shopping interface {
	SaleAndGetPrice() int
}

var snackDiscountRate = 0.1
var drinkDiscountRate = 0.2

func SaleAndGetPrice(item Shopping) int {
	return item.SaleAndGetPrice()
}

func (snack Snack) SaleAndGetPrice() int {
	discounted := int(float64(snack.Price) * (1 - snackDiscountRate))
	return discounted
}

func (drink Drink) SaleAndGetPrice() int {
	discounted := int(float64(drink.Price) * (1 - drinkDiscountRate))
	return discounted
}

func main() {
	chips := Snack{"Pringles", 4000}
	cracker := Snack{"Ace", 2500}

	soda := Drink{"Sprite", 1800}
	coffee := Drink{"TOP", 2700}

	var total int = 0
	total += SaleAndGetPrice(&chips)
	total += SaleAndGetPrice(&cracker)
	total += SaleAndGetPrice(&soda)
	total += SaleAndGetPrice(&coffee)

	fmt.Println(total)
}