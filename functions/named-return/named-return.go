package main

import "fmt"

func switchParams(p1, p2 int) (second, first int) {
	second = p2
	first = p1
	return
}

func main() {
	p1, p2 := switchParams(4, 5)

	fmt.Println(p1, p2)
}
