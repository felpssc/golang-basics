package main

import "fmt"

func main() {
	var notas [3]float64

	notas[0], notas[1], notas[2] = 4.5, 7.2, 6.7

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total += notas[i]
		fmt.Printf("Nota: %.2f\n", notas[i])
	}

	length := float64(len(notas))

	fmt.Println("Quantidade de notas:", length)

	media := total / length

	fmt.Printf("Média: %2.f", media)
}
