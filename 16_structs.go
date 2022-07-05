package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Avinil", 28})

	fmt.Println(person{name: "Aniket", age: 22})

	fmt.Println(person{name: "Fred"}) //omitted field are zero valued

	fmt.Println(person{name: "Fred", age: 77}) //& operator gives a pointer to struct

	fmt.Println(newPerson("John"))

	s := person{name: "Jayesh", age: 49}

	fmt.Println(s.age)

	sp := &s
	sp.age = 50

	fmt.Println(sp.age)
	fmt.Println(s.age)
}
