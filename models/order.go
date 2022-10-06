package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	OrderID      string `json:"order_id"`
}
