package main

import "fmt"

func imprimirResultado(nota float64, hasCompletedHours bool) {
	if nota >= 7 && hasCompletedHours {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Reprovado com nota:", nota)
	}
}

func main() {
	imprimirResultado(6, true)
	imprimirResultado(7, true)
}
