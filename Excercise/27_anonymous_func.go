package main

import (
	"fmt"
)

func outer() func(){
	return func(){
		fmt.Println("Anonymous function")
	}
}

func main()  {
	f:=outer()
	f()
}