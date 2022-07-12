package main

import (
	"fmt"
)

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base num %d", b.num)
}

func main() {
	co := container{
		base: base{
			num: 25,
		},
		str: "This a string",
	}
	fmt.Printf("container has num %d and a string str %s", co.num, co.str)
	fmt.Println(" also num :", co.base.num)
	fmt.Println("Calling method describe ", co.base.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("calling describe now :", d.describe())
}
