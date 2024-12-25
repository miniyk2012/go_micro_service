package main

import (
	"demo/proto/book"
	"fmt"
)

func main() {
	price := &book.Price{
		MarketPrice: 100,
		SalePrice:   80,
	}
	fmt.Println(price)
}
