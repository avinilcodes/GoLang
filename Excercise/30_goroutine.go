package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) { // Goroutine - returns immediately it does not wait for function to execute
			fmt.Println(i) //- more like background execution for this function
		}(i)
	}
	time.Sleep(2 * time.Second) // This just to wait for two seconds
	fmt.Println("main is going to end")
}
