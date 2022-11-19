package main

import "fmt"

type Course struct {
	Name string
}

func main() {
	var thing interface{}

	fmt.Println(thing)

	thing = 3

	fmt.Println(thing)
	type dynamic interface{}

	var thing2 dynamic = "Hello"

	thing2 = Course{"Go"}

	fmt.Println(thing2)
}
