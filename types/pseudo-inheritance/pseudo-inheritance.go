package main

import "fmt"

type Car struct {
	Nome        string
	ActualSpeed int
}

type Ferrari struct {
	Car
	IsTurboOn bool
}

func main() {
	f := Ferrari{}
	f.Nome = "F40"
	f.ActualSpeed = 0
	f.IsTurboOn = false

	fmt.Println(f)
}
