package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	product := Product{
		ID:    1,
		Name:  "Product 1",
		Price: 10.5,
		Tags:  []string{"tag1", "tag2"},
	}

	fmt.Println(product)

	productJson, err := json.Marshal(product)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(productJson))

	var product2 Product

	jsonString := `{"id":2,"name":"Product 2","price":20.5,"tags":["tag3","tag4"]}`

	json.Unmarshal([]byte(jsonString), &product2)

	fmt.Println(product2)
}
