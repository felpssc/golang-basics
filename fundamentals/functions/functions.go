package main

import "fmt"

func somar(a int, b float64) float64 {
	return float64(a) + b
}

func main() {
	fmt.Println(somar(3, 4))
}
