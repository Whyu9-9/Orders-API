package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID      uint      `json:"orderId" gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID"`
}
