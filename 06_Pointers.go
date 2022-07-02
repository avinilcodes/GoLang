package main

import (
	"fmt"
)

func main() {
	a := 5
	var ptr *int = &a
	ptr1 := &a
	fmt.Println("val of a: ", a)
	fmt.Println("val of a using Pointer: ", *ptr) //using Pointer
	fmt.Println("address of a: ", &a)
	fmt.Println("address of a using Pointer: ", ptr1) //using Pointer

}
