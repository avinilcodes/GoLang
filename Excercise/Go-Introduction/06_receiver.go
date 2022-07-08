package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func init() {

}

func main() {
	var c Circle = Circle{r: 4.0}
	fmt.Println("Radius of Circle is ", c.r)
	fmt.Println("Area of Circle is", c.area())
}
