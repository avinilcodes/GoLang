package main

import (
	"fmt"
)

func main() {
	var principal float64
	var rate float64
	var time float64
	fmt.Scan(&principal)
	fmt.Scan(&rate)
	fmt.Scan(&time)
	interest := (principal * time * rate) / 100.0
	fmt.Printf("Total interest on amount %f with rate %f for time %f is : %f\n", principal, rate, time, interest)
}
