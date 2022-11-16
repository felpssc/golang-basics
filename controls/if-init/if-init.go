package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	n := rand.NewSource(time.Now().UnixNano())
	number := rand.New(n)

	return number.Intn(10)
}

func main() {
	if n := randomNumber(); n > 5 {
		fmt.Println("aoooi")
	}
}
