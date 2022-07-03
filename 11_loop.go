//Go has only for loop
package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	n := 2
	//while like for
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
	//infinite loop
	// for {
	// 	fmt.Println("Hello Go")
	// }
}
