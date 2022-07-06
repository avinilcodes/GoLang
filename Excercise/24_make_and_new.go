package main

import (
	"fmt"
)

type Avinil struct{
	name string
	age uint
}

func main()  {
	arr:=[5]int{1,2,3,4,5}
	s:= make([]int, 4,5)
	fmt.Println(arr,s)
	//below two init are identical
	a := &Avinil{}
	fmt.Println(a)
	b:= new(Avinil)
	fmt.Println(b)
}
