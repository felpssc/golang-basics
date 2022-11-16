package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	var slice []int

	slice = append(slice, 4, 5, 6)
	fmt.Println(array, slice)

	slice2 := make([]int, 2)
	copy(slice2, slice)
	fmt.Println(slice2)
}
