package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	area := math.Pi * math.Pow(r, 2)
	fmt.Println("Area of circle with radius :", r, " is ", area)
}
