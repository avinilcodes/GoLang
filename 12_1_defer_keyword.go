package main

import (
	"fmt"
)

//defer stack example
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
