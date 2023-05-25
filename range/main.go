package main

import (
	"fmt"
	"time"
)

var greetings []string = []string{"Hello", "Ola", "Namaskara", "Jai Jinendra"}

func sayHello(greetings []string, ch chan<- string) {
	for _, greeting := range greetings {
		ch <- greeting
	}
	close(ch)
	fmt.Println("say hello exited")
}

func main() {
	ch := make(chan string, 1)
	go sayHello(greetings, ch)
	time.Sleep(time.Millisecond * 500)
	for greet := range ch {
		if greet == "" {
			return
		}
		fmt.Println("Received greeting :")
		fmt.Printf("Greeting %v\n", greet)
	}
	fmt.Println("Main exited")
}
