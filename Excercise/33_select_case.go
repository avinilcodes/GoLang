package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(ch chan string) {
		time.Sleep(1 * time.Second)
		ch <- "One"
	}(ch1)
	go func(ch chan string) {
		time.Sleep(2 * time.Second)
		ch <- "One"
	}(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Data read from channel 1", msg1, i)
		case msg2 := <-ch2:
			fmt.Println("Data read from channel 2", msg2, i)
		}
	}
}
