package main

import (
	"fmt"
)

func main() {
	v := 1
	p := &v
	s := "THis is a string"
	fmt.Printf("%#v %#v %#v", v, p, s)
}
