package main

import "fmt"

func main() {
	p1 := Point{X: 10, Y: 20}
	p2 := Point{X: 15, Y: 25}

	fmt.Println(catetos(p1, p2))

	fmt.Println(Distance(p1, p2))
}
