package main

import (
	"fmt"
)

type User struct {
	name       string
	occupation string
}

func (u User) String() string {
	return fmt.Sprintf("%s is a %s", u.name, u.occupation)
}

func main() {
	u := User{"John Doe", "Pilot"}
	fmt.Println(u)
}
