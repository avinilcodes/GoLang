package main

import (
	"fmt"
)

func main() {
	v := 1
	p := &v
	s := "THis is a string"
	f := 45.23
	fmt.Printf("%#v %#v %#v", v, p, s)
	fmt.Printf("%p %s", p, s)
	fmt.Printf("%d", v)
	fmt.Printf("%g", f)
}
