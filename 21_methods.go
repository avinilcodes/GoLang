package main

import (
	"fmt"
)

type Receiver int

func (r Receiver) thisFunctionAcceptsReceiver() { //this method accepts a receiver
	r = 10
	fmt.Println("R: ", r)
}

func main() {
	var r Receiver
	r.thisFunctionAcceptsReceiver()
}
