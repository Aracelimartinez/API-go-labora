package models

import (
	"math"
)

type Item struct {
	ID            int    	`json:"id"`
	CustomerName 	string 	`json:"customer_name"`
	OrderDate 		string	`json:"order_date"`
	Product       string	`json:"product"`
	Quantity      int64  	`json:"quantity"`
	Price         float64	`json:"price"`
	Details       string	`json:"details"`
	TotalPrice    float64 `json:"total_price"`
	ViewCounter   int64 	`json:"view_counter"`
}

//Calcula el precio total de un item
func (item *Item) CalculateTotalPrice() (float64){
	totalPrice := item.Price * float64(item.Quantity)

	item.TotalPrice = math.Round(totalPrice*100) / 100
	return item.TotalPrice
}

//Aumenta la vizualización de un item
// func (item *Item) IncrementViewCounter() int64 {
// 	item.ViewCounter ++
// 	return item.ViewCounter
// }
