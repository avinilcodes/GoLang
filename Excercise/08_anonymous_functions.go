package main

import (
	"fmt"
)

func c(a int) int{
	return a
}

func b(c func(int) int,val int) int{
	return c(val)
}


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

	d:= b(c,5)
	fmt.Println("calling a function from another function ,i.e. passing function as a parameter to another function ", d)
}
