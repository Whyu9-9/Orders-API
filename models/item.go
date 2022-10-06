package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemID      uint   `json:"itemId" gorm:"primaryKey"`
	ItemCode    string `json:"itemcCode"`
	Descritpion string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"orderID"`
}
