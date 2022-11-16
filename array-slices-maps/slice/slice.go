package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [3]int{3, 2, 1}
	slice := []int{1, 2, 3}

	fmt.Println(array, slice)
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))

	// Slice não é um array! Slice define um pedaço de um array.

	slice2 := array[0:2]
	fmt.Println(slice2)
}
