package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `json:"itemCode"`
	Descritpion string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"orderID"`
}
