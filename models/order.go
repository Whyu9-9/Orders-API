package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `json:"customerName" gorm:"not null" binding:"required"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"not null" binding:"required"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID" binding:"required"`
}
