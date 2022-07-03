package main

import (
	"fmt"
)

func main() {
	a := 23
	defer fmt.Println("Value of a:", a) // print 23 as defer is evaluated here itself
	a = 33
}
