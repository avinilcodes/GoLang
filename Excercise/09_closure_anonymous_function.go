package main

import (
	"fmt"
)

func f() func() {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func main() {
	closure := f()
	closure()
	closure()
	closure()
}
