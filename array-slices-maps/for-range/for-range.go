package main

import "fmt"

func main() {
	numbers := [...]int{1, 3, 5, 7, 9}

	for index, number := range numbers {
		fmt.Println("Número", number, "na posição:", index)
	}
}
