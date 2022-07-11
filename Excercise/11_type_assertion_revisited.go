package main

import (
	"fmt"
)

func main() {
	var a interface{} = "An interface"

	v := a.(string)
	fmt.Println("Type assertion done :", v)
}
