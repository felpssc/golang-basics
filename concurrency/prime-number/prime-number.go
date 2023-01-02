package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primes(num int, ch chan int) {
	init := 2

	for i := init; i < num; i++ {
		for prime := init; ; prime++ {
			if isPrime(prime) {
				ch <- prime
				init = prime + 1
				time.Sleep(100 * time.Millisecond)
				break
			}
		}
	}

	close(ch)
}

func main() {
	c := make(chan int, 30)

	go primes(cap(c), c)

	for prime := range c {
		fmt.Println(prime)
	}
}
