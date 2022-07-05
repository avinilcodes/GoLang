package main

import (
	"fmt"
)

func main() {
	mp := make(map[string]int)
	mp["avinil"] = 4
	mp["aniket"] = 2

	fmt.Println("my first map in Go: ", mp)

	fmt.Println("Size of map mp :", len(mp))

	fmt.Println("Getting an element from a map :", mp["avinil"])

	delete(mp, "avinil")
	fmt.Println("map :", mp)

	_, prs := mp["avinil"]
	fmt.Println("prs :", prs)

	mp1 := map[string]int{"avinil": 1, "aniket": 2}
	fmt.Println("map mp1 :", mp1)
}
