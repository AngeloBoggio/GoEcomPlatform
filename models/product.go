package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	StoreID  uint
	Name	string
	Description	string
	Price	float64
	StockStatus	string
}