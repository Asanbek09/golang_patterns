package main

import (
	"factory/products"
	"fmt"
	"time"
)

func main() {
	factory := products.Product{}

	product := factory.New()

	fmt.Println("My product was created at", product.CreatedAt.UTC())

	// Второй пример
	product2 := products.Product {
		ProductName: "ball",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println("My product 2 was created at", product2.CreatedAt.UTC())
}