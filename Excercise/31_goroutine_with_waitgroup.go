package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //Wait Group of sync package

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Add(1)
	}
	//time.Sleep(2 * time.Second) // This just to wait for two seconds
	wg.Wait()
	fmt.Println("main is going to end")
}
