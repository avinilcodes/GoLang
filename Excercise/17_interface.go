package main

import (
	"fmt"
)

//as long as struct implements all the methods of interface,
//that struct/type can access functions accepting that interface as parameter
type Socket string
type PowerDrawer interface {
	Draw()
}

func (s Socket) plug(pd PowerDrawer) {
	pd.Draw()
}

type Mobile struct {
	name string
}
type LEDTV struct {
	name string
}
type WashingMachine struct {
	name string
}

func (m Mobile) Draw() {
	fmt.Println("Mobile m draws less power")
}

func (wm WashingMachine) Draw() {
	fmt.Println("WashingMachine wm draws high power")
}

func (tv LEDTV) Draw() {
	fmt.Println("LEDTV tv draws less power")
}

func main() {
	mobile := Mobile{name: "Lenovo"}
	washingmachine := WashingMachine{name: "Samsung"}
	tv := LEDTV{name: "Sony"}
	fmt.Println(mobile)

	var s Socket
	s.plug(tv)
	s.plug(washingmachine)
	s.plug(mobile)
}
