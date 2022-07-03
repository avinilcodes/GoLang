package main

import (
	"fmt"
)

func main() {
	arr := [10]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //basic syntax of slice
	slice := arr[:]                           // 0 to len-1
	slice1 := arr[0:]                         // 0 to len-1
	slice2 := arr[:6]                         //0 to high-1
	slice3 := arr[3:8]                        //3 to 8-1
	fmt.Printf("slice length :%d, slice capacity :%d\n", len(slice2), cap(slice2))
	slice2 = arr[:8]                                                               //resizing slice
	fmt.Printf("slice length :%d, slice capacity :%d\n", len(slice2), cap(slice2)) //after resizing
	fmt.Println(len(arr))
	fmt.Println("slice :", slice)
	fmt.Println("s :", s)
	fmt.Println("slice1 :", slice1)
	fmt.Println("slice2 :", slice2)
	fmt.Println("slice3 :", slice3)
	slice[0] = 9
	fmt.Println("slice after reassignment as it is reference type :", slice)

	//making a slice using make function

	slice4 := make([]int, 5, 10)
	fmt.Println("slice4 :", slice4) //default value of slice is 0

	//making slice a nil slice
	var slice5 []int
	fmt.Println("slice5 :", slice5)
}
