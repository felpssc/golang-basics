package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	switch {
	case now.Hour() < 12:
		fmt.Println("Bom dia")
	case now.Hour() > 12 && now.Hour() < 19:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
