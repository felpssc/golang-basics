package main

import "fmt"

func main() {
	a := 5

	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &a
	*p++
	a = 2

	fmt.Println(*p)
}
