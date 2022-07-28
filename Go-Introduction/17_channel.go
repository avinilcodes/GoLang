package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int, 0)

	go func() {
		// ch <- 1
		for i := 0; i < 1000; i++ {
			fmt.Println("value i:", i)
		}
	}()
	time.Sleep(1 * time.Millisecond)
	// <-ch
}
