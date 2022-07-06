package main

import (
	"fmt"
	"time"
)

func main() {
	var msg string = "Hello"
	go func() {
		fmt.Println("Inside Anonymous function - \nMessage is printed from outer function because of closure msg: ", msg)
	}()
	time.Sleep(1000 * time.Millisecond)
}
