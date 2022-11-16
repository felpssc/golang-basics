package main

import "fmt"

func media(numbers ...float64) float64 {
	len := len(numbers)

	var sum float64

	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len)
}

func main() {
	fmt.Println(media(3, 5, 1, 7, 9, 5, 7, 3, 8, 7))
}
