package main

import "fmt"

func routine(ch chan int) {
	ch <- 1
	fmt.Println("routine: sent 1")
	ch <- 2
	fmt.Println("routine: sent 2")
	ch <- 3
	fmt.Println("routine: sent 3")
	ch <- 4
	fmt.Println("routine: sent 4")
	ch <- 5
	fmt.Println("routine: sent 5")
	ch <- 6
	fmt.Println("routine: sent 6")
}

func main() {
	c := make(chan int, 3)

	go routine(c)

	fmt.Println("main: received", <-c)
}
