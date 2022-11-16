package main

import "fmt"

func inc(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc(n)

	fmt.Println(n)

	inc2(&n)

	fmt.Println(n)
}
