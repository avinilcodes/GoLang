package main

import (
	"fmt"
)

func main() {
	s := "This a string"
	fmt.Println("string :", s)
	b := []byte(s) // we can convert string to byte stream to give as response to the request
	fmt.Println("byte slice :", b)
	var r rune = 'a' //it is int32, notice rune is initialized using single quotes & they are tricky so we need to read documentation for using it
	fmt.Println("Rune r :", r)
}
