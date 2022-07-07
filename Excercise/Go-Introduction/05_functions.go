package main

import (
	"fmt"
)

func evenOdd(a int) (int, bool) {
	if a%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

func main() {
	a, b := evenOdd(2)
	fmt.Println(a, b)
	var t int = 25
	var p *int = &t
	fmt.Println(*p)
}
