package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	mes := <-pings
	pongs <- mes
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "This is a message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
