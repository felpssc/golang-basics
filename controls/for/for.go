package main

import "fmt"

// Go não tem while

func main() {
	c := 1

	for c <= 10 {
		fmt.Println(c)
		c++
	}

	for b := 10; b > 0; b-- {
		fmt.Println(b)
	}

	// Laço infinito

	// for {
	// 	fmt.Println("for ever...")
	// }
}
