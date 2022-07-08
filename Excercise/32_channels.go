package main

import (
	"fmt"
	"time"
)

func writeToChannel(ch chan int, s string) {
	ch <- 8
	fmt.Println("Done writing --> ", s)
}
func readFromChannel(ch chan int, s string) {
	v := <-ch
	fmt.Println("Read ", v)
	fmt.Println("Done reading --> ", s)
}
func main() {
	ch := make(chan int) //unbuffered channel
	go readFromChannel(ch, "World")
	go writeToChannel(ch, "Hello, ")
	time.Sleep(1 * time.Second)
}
