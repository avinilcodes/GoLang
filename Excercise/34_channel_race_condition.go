//We use channel for avoiding race condition
package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x += 1
	wg.Done()
	<-ch
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, ch)
	}
	wg.Wait()
	fmt.Println("Done with incrementing, value of x is :", x)
}
