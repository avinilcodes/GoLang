package main

import "fmt"

func g() (int, error) {
	return 0, fmt.Errorf("x")
}
func f() (err error) {
	if n, err := g(); err == nil {
		return err // here only return statement was there so it does not compile
	} else {
		return fmt.Errorf("%w:%d", err, n)
	}
}

func main() {
	fmt.Println(f() != nil)
}
