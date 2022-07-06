package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() { // here we are using defer to recover from panic as panic happens after defer stack becomes empty
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic("Something bad happend")
}
