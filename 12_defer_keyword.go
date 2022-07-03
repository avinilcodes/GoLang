//Function that accepts variable number of arguments

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello defer - Go")
	defer fmt.Println("First defer")
	defer fmt.Println("second defer")
	fmt.Println("Good bye")
}
