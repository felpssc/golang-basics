package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"José":   7500.43,
		"Felipe": 18000,
	}

	funcionarios["Henrique"] = 5600.41

	delete(funcionarios, "INEXISTENTE")

	for nome, salario := range funcionarios {
		fmt.Println(nome, salario)
	}
}
