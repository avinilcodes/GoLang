package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) Perimeter() float64 {
	return math.Pi * 2 * c.r
}
func init() {}
func main() {
	var c Circle = Circle{r: 4.0}
	fmt.Println("Radius of Circle is ", c.r)
	fmt.Println("Area of Circle is", c.Area())
	fmt.Println("Perimeter of Circle is", c.Perimeter())
}
