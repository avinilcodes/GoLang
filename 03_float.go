package main

import (
	"fmt"
)

func main() {
	a := 5.56
	b := 6.23
	fmt.Printf("type of a :%T b :%T\n", a, b)
	sum := a + b
	fmt.Println("Sum a,b :", sum)
	diff := a - b
	fmt.Println("Diff a,b :", diff)

}
