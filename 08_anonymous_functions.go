package main

import (
	"fmt"
)

func main() {
	func() { // this function is anonymous function
		fmt.Println("Inside a anonymous function")
	}()
}
