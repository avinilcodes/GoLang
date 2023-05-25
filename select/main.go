package main

import (
	"fmt"
	"time"
)

var greetings []string = []string{"Hello", "Ola", "Namaskara", "Jai Jinendra"}
var goodbyes []string = []string{"Goodbye", "Adios", "Saiyonara"}

func sayHello(greetings []string, ch chan<- string) {
	for _, greeting := range greetings {
		ch <- greeting
	}
	close(ch)
	fmt.Println("say hello exited")
}

func main() {
	ch := make(chan string, 1)
	ch1 := make(chan string, 1)

	go sayHello(greetings, ch)
	go sayHello(goodbyes, ch1)
	time.Sleep(time.Second * 1)
	for {
		select {
		case greet, ok := <-ch:
			if !ok {
				ch = nil
				break
			}
			fmt.Println(greet)
		case greet, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Println(greet)
		default:
			fmt.Println("Nothing received")
			return
		}
	}
}
