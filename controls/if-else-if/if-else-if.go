package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota < 9 && nota >= 7 {
		return "B"
	} else {
		return "C"
	}
}

func main() {
	c := notaParaConceito(9.5)

	fmt.Println(c)
}
