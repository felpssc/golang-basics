package main

import "fmt"

func clojure() func() {
	x := 10
	var function = func() {
		fmt.Println(x)
	}

	return function
}

func main() {
	x := 5
	fmt.Println(x)

	imprimeX := clojure()
	imprimeX()
}
