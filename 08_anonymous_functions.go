package main

import (
	"fmt"
)

func main() {
	func() { // this function is anonymous function
		fmt.Println("Inside anonymous function")
	}()

	s := func(a, b int) int { // this function is anonymous function
		fmt.Println("Inside another anonymous function")
		return a + b
	}
	fmt.Println("Sum of 3,2 is :", s(2, 3))

	f := func(a, b int) int { // this function is anonymous function
		return a - b
	}(5, 2)
	fmt.Println("diff between 5,2 is: ", f)
}
