package main

import "fmt"

func main() {
	// Maps devem ser inicializados
	// var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[23213123123] = "Maria"
	aprovados[12949023839] = "Felipe"
	aprovados[94839284294] = "Jurema"

	fmt.Println(aprovados)

	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(aprovados[23213123123])
	delete(aprovados, 23213123123)

	fmt.Print(aprovados[23213123123])
}
