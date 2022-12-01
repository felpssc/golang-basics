package main

import "fmt"

// Não tem operador ternário

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	result := obterResultado(6)

	fmt.Println(result)
}
