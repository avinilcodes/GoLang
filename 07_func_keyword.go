package main

import (
	"fmt"
)

func calculatePrice(price, items int) int {
	return price * items
}
func main() {
	var price, items int = 90, 6
	fmt.Println("Total price of :", items, " is: ", calculatePrice(price, items))
}
