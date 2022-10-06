package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemID      string `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Descritpion string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order       Order  `json:"order"`
}
