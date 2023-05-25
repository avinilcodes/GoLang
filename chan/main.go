package main

import "fmt"

func sayHello(ch chan<- string) {
	ch <- "Hello"
	fmt.Println("say Hello func done")
}

func main() {
	ch := make(chan string, 1)
	go sayHello(ch)

	data := <-ch
	fmt.Println("data received from channel")
	fmt.Printf("%v", data)
}
