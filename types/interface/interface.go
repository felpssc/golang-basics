package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Price float64
}

func (p Person) ToString() string {
	return p.Name + " - " + strconv.Itoa(p.Age) + " years old"
}

func (p Product) ToString() string {
	return p.Name + " - R$ " + strconv.FormatFloat(p.Price, 'f', 2, 64)
}

func PrintX(x Printable) {
	fmt.Println(x.ToString())
}

func main() {
	p1 := Person{"John", 20}
	p2 := Person{"Mary", 25}

	prod1 := Product{"Notebook", 1899.90}
	prod2 := Product{"Pen", 8.59}

	PrintX(p1)
	PrintX(p2)
	PrintX(prod1)
	PrintX(prod2)
}
