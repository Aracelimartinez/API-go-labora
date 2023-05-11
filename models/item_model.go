package models

import ()

type Item struct {
	ID            int64  `json:"id"`
	CustomerName 	string `json:"customer_name"`
	OrderDate 		string `json:"order_date"`
	Product       string `json:"product"`
	Quantity      int64  `json:"quantity"`
	Price         int64  `json:"price"`
	Details       string `json:"details"`
}
