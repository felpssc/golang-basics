package main

import (
	"fmt"
	"time"
)

func talk(person, txt string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, txt, i+1)
	}
}

func main() {
	// talk("Maria", "Hi!", 2)
	// talk("João", "Hello!", 1)

	// go talk("Maria", "Hi!", 2)
	// go talk("João", "Hello!", 2)

	go talk("Maria", "Hi!", 10)
	talk("João", "Hello!", 5)
}
