package main

import "fmt"

func main() {
	fmt.Print("a")
	fmt.Print("b")

	fmt.Println("c")
	fmt.Println("d")

	x := 1.23434

	fmt.Println("o valor de x é", x)

	fmt.Printf("o valor de x é %.3f", x)

	a := 1
	fmt.Printf("\no valor de a é %d", a)
	b := 1.9999
	fmt.Printf("\no valor de b é %.1f", b)
	c := false
	fmt.Printf("\no valor de c é %t", c)
	d := "bye"
	fmt.Printf("\no valor de d é %s", d)
}
