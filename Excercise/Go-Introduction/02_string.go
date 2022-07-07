package main

import (
	"fmt"
)

func main() {
	var s string = "This a string"
	fmt.Println("s is a string as follows ", s, "length of string is ", len(s))
	fmt.Println((true && false) || (true && false) || !(false && false))
}
