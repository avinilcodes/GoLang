//Function that accepts variable number of arguments

package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	n := 2
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}
