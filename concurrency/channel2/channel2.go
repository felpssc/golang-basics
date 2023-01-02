package main

import "time"

func someOperations(base int, c chan int) {
	time.Sleep(2 * time.Second)
	c <- base * 2

	time.Sleep(3 * time.Second)
	c <- base * 3

	time.Sleep(4 * time.Second)
	c <- base * 4
}

func main() {
	c := make(chan int)
	go someOperations(2, c)

	// This will block until the first value is received
	result := <-c
	println(result)
	println("first value received")

	// This will block until the second value is received
	result = <-c
	println(result)
	println("second value received")

	// This will block until the third value is received
	result = <-c
	println(result)
	println("third value received")
}
