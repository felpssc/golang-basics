package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 7.2
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println(strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	// string para boolean
	b, _ := strconv.ParseBool("1")

	if b {
		fmt.Println("Verdadeiro")
	}
}
