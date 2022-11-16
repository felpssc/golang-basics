package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Discount float64
}

// Methods...

func (p Product) priceWithDiscount() float64 {
	return p.Price - p.Discount
}

func main() {
	brinquedo := Product{
		Name:     "Lego Star Wars",
		Price:    156.90,
		Discount: 10.90,
	}

	fmt.Println(brinquedo.priceWithDiscount())
}
