package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)

	ch <- "Buffered"
	ch <- "Channel"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
