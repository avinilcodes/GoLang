package main

import (
	"fmt"
)

type Receiver struct {
	id   int
	name string
}

func (p Receiver) someFunc() { //accepts normal receiver
	p.id = 10
	fmt.Println(p)
}

func (p *Receiver) someFunc1() { //accepts receiver pointer
	p.id = 10
	fmt.Println(p)
}

func main() {
	p1 := Receiver{1, "Sachin"}
	p2 := &Receiver{1, "Sachin"}

	fmt.Println("before function execution p1: ", p1)
	p1.someFunc()
	fmt.Println("before function execution p1: ", p1)
	fmt.Println("before function execution p2: ", p2)
	p2.someFunc1()
	fmt.Println("before function execution p2: ", p2)
}
