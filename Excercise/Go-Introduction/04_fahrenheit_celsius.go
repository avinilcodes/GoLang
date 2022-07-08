package main

import (
	"fmt"
)

func main() {
	var a float32
	fmt.Scan(&a)
	fmt.Println("fahrenheit ", a, " equals to - ", (a-32.0)*(5.0/9.0), " celsius")
}
