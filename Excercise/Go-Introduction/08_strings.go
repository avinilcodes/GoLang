package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.Index("test", "t"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("t", 5))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d", "-"))
	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))
	//sometimes we need to work with strings as binary data.
	//To convert a string to a slice of bytes(and vice versa)
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println("to binary ", arr)
	fmt.Println("string from bytes ", str)
}