package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example of defer function and the output is", foo())
}

func foo() (result string) {
	defer func() {
		result = "Change World"
	}()
	return "Hello World"
}
