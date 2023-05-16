package models

import (
	"fmt"
	"testing"
)

type totalPriceTest struct {
	price    float64
	quantity int64
	expected float64
}

var totalPriceTests []totalPriceTest = []totalPriceTest{
	{140.0, 5, 700},
	{375.0, 3, 1125},
	{105.75, 2, 211.5},
	{160.0, 4, 640},
	{0.0, 10, 0.0},
	{2.99, 1, 2.99},
}

func TestCalculateTotalPrice(t *testing.T) {
	for _, test := range totalPriceTests {
		item := Item {
			Price:    test.price,
			Quantity: test.quantity,
		}

		item.CalculateTotalPrice()

		if item.TotalPrice != test.expected {
			t.Errorf("Para el precio %f y cantidad %d, se esperaba un total de %f pero se obtuvo %f",
				test.price, test.quantity, test.expected, item.TotalPrice)
		}

	}
}

func BenchmarkCalculateTotalPrice(b *testing.B) {
	item := Item{
		Price:    107.55,
		Quantity: 3,
	}

	for n := 0; n < b.N; n++ {
		item.CalculateTotalPrice()
	}
}

func BenchmarkCalculateTotalPriceLargeQuantity(b *testing.B) {
	item := Item{
		Price:    10.5,
		Quantity: 1000,
	}

	for n := 0; n < b.N; n++ {
		item.CalculateTotalPrice()
	}
}

func ExampleCalculateTotalPrice() {
	item := Item{
		Price:    107.55,
		Quantity: 3,
	}

	fmt.Println(item.CalculateTotalPrice())
	//Output: 322.65
}
