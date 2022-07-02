package main

import (
	"fmt"
)

func main() {
	var a bool = true
	b := false
	c := a && b
	d := a || b
	fmt.Println("value of a :", a, " value of b: ", b)
	fmt.Println("c: ", c, " d: ", d)
}
