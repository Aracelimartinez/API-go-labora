package models

import ()

type Item struct {
	ID            string `json:"id"`
	CustomerName 	string `json:"customer_name"`
	OrderDate 		int 	 `json:"order_date"`
	Product       string `json:"product"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	Details       string `json:"details"`
}
