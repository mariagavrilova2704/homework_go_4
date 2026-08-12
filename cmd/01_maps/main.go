package main

import (
	"fmt"

	"github.com/rinatkh/homework_go_4/internal/maps"
)

type Product struct {
	SKU      string
	Quantity int
	Price    int
}

func main() {
	fmt.Println(maps.Example())
	//	product1 := Product{
	//		SKU:      "a",
	//		Quantity: 0,
	//	}
	//
	//	product2 := Product{
	//		SKU:      "A",
	//		Quantity: 0,
	//	}
	//
	//	inventory := make(map[string]Product)
	//	inventory["a"] = product1
	//	inventory["A"] = product2
	//	fmt.Println(LowStockSKUs(inventory, 0))
	//}
	//
	//func LowStockSKUs(inventory map[string]Product, limit int) []string {
	//	// TODO: вернуть SKU товаров с остатком не больше limit.
	//	// Результат должен иметь стабильный алфавитный порядок.
	//	sku := make([]string, 0, len(inventory))
	//	for _, v := range inventory {
	//		if v.Quantity <= limit {
	//			sku = append(sku, v.SKU)
	//		}
	//	}
	//	slices.Sort(sku)
	//	return sku
	//}
}
