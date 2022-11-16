package main

import "fmt"

func basicFunction() {
	fmt.Println("This is a basic function!")
}

func funcWithParams(p1, p2 string) {
	fmt.Println(p1, p2)
}

func funcWithReturn() string {
	return "This is a function with return!"
}

func funcWithParamsAndReturn(p string) string {
	return p
}

func funcWithMultipleReturns(n1, n2 int) (int, int, int) {
	sum := n1 + n2
	sub := n1 - n2
	mult := n1 * n2

	return sum, sub, mult
}

func main() {
	basicFunction()
	funcWithParams("o", "i")
	fmt.Println(funcWithReturn())
	fmt.Println(funcWithParamsAndReturn("abc"))
	fmt.Println(funcWithMultipleReturns(5, 2))
}
