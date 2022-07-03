package main

import (
	"fmt"
)

func main() {
	arr := [3]int32{1, 2, 3}
	fmt.Println("array arr: ", arr)
	fmt.Printf("type of arr: %T\n", arr) //Type of an array is size and data type
	arr1 := arr
	fmt.Println("array arr1:", arr1)
	fmt.Println("If arr1 equals to arr: ", arr == arr1) //for array to be equal their type and elements should be same
	arr[2] = 24
	fmt.Println("If arr1 equals to arr: ", arr == arr1) //for array to be equal their type and elements should be same

	arr2 := [3][4]int{{1, 2, 3, 4}, {3, 2, 1}}
	fmt.Println("two D array arr2:", arr2)

}
