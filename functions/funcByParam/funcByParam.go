package main

import "fmt"

func increase(n1, n2 int) int {
	return n1 + n2
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func main() {
	result := exec(increase, 2, 1)

	fmt.Println(result)
}
