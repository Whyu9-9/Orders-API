package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `json:"itemCode" binding:"required, min=3, max=10, alphanum, unique"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	OrderID     uint   `json:"orderID" binding:"required"`
}
