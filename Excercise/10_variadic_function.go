//Function that accepts variable number of arguments

package main

import (
	"fmt"
)

func main() {
	fmt.Println(variadicFunction(1, 2, 3))
	fmt.Println(variadicFunction(1, 2, 5, 3, 7))
	fmt.Println(variadicFunction(11, 2, 43, 77, 120))
}

func variadicFunction(nums ...int) int {
	var total int = 0
	for _, n := range nums {
		total += n
	}
	return total
}
