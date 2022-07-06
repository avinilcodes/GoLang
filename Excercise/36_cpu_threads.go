//Wite a program to find number of threads present in the system
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of threads available on this system are :", runtime.GOMAXPROCS(-1))
}
