package main

import (
	"fmt"
)

type Entity string

type Person interface{
	Walk()
	Eat()
	Work()
	Sleep()
}

type Avinil struct{
	name string
	age uint
}

func (e Entity) usePersonFuncWalk(p Person)  {
	p.Walk()
}

func (e Entity) usePersonFuncEat(p Person)  {
	p.Eat()
}

func (e Entity) usePersonFuncWork(p Person)  {
	p.Work()
}

func (e Entity) usePersonFuncSleep(p Person)  {
	p.Sleep()
}


func (a Avinil) Walk(){
	fmt.Println(a.name," is walking")
}

func (a Avinil) Eat(){
	fmt.Println(a.name," is eating")
}

func (a Avinil) Work(){
	fmt.Println(a.name," is working")
}

func (a Avinil) Sleep(){
	fmt.Println(a.name," is sleeping")
}

func main()  {
	var e Entity
	a := Avinil{"Avinil",28}
	e.usePersonFuncWalk(a)
	e.usePersonFuncEat(a)
	e.usePersonFuncWork(a)
	e.usePersonFuncSleep(a)
}