package main

import "fmt"

func main() {
	ch := make(chan int, 3) // Channel with type int and buffer size 3

	ch <- 5 // Send 5 to channel
	<-ch    // Reading channel

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 20
	fmt.Println(<-ch) // 30
}
