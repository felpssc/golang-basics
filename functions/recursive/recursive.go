package main

import "fmt"

func main() {
	r, e := fatorial(5)

	fmt.Println(r, e)
}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}
