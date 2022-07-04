package main

import (
	"fmt"
)

func main() {
	var x interface{} = "foo"
	s := x.(string)
	fmt.Println("Type assertion on x interface, s: ", s)

	t, ok := x.(string) // handling panic situation
	fmt.Println("Type assertion on x interface, t: ", t, ok)

	u, ok := x.(int)
	fmt.Println("Type assertion on x interface, u: ", u, ok)

	//commenting as it causes panic
	// v := x.(int)
	// fmt.Println("Type assertion on x interface v: ", v)

}
