package main

import "fmt"

var sum = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(sum(9, 1))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(10, 1))
}
