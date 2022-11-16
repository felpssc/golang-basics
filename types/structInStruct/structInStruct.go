package main

import "fmt"

type Item struct {
	ProductID int
	Quantity  int
	Price     float64
}

type Order struct {
	OrderID int
	UserID  int
	Items   []Item
}

func (o Order) totalPrice() float64 {
	var sum float64

	for _, item := range o.Items {
		sum += item.Price * float64(item.Quantity)
	}

	return sum
}

func main() {
	order := Order{
		OrderID: 22321,
		UserID:  22391,
		Items: []Item{
			{
				ProductID: 343,
				Quantity:  2,
				Price:     12.50,
			},
			{
				ProductID: 323,
				Quantity:  5,
				Price:     93.25,
			},
		},
	}

	fmt.Println(order.totalPrice())
}
