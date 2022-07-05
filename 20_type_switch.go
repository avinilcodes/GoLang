package main

import (
	"fmt"
)

func main() {
	var x interface{} = "foo"

	switch v := x.(type) {
	case nil:
		fmt.Println("v is nil")
	case int:
		fmt.Println("v is integer", v)
	case bool, string:
		fmt.Println("v is bool or string", v)
	default:
		fmt.Println("v is unknown", v)
	}
}
