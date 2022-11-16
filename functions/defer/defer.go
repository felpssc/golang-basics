package main

import "fmt"

func obtainApprovedValue(n int) int {
	defer fmt.Println("This will be executed at the end.")

	if n > 5 {
		fmt.Println("Approved")
		return n
	} else {
		fmt.Println("Reproved")
		return n
	}
}

func main() {
	fmt.Println(obtainApprovedValue(6))
}
