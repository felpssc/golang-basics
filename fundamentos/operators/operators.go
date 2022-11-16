package main

import "fmt"

func main() {
	a := 500
	b := 275

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Módulo:", a%b)

	// Bitwise
	fmt.Println("E:", a&b)   // 0101 & 0111 = 0101
	fmt.Println("Ou:", a|b)  // 0101 | 0111 = 0111
	fmt.Println("Xor:", a^b) // 0101 ^ 0111 = 0010
}
