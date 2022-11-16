package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"F": {
			"Felipe":   7000,
			"Fernanda": 8000,
		},
		"D": {
			"Danilo": 90000,
			"Junior": 10000,
		},
		"P": {
			"Paulo": 2000,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		for nome, salario := range funcs {
			fmt.Printf("\n[%s] %s %2.f", letra, nome, salario)
		}
	}
}
