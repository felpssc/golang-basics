package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 9 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3
	fmt.Println(i)

	x, y := 10, 30
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
