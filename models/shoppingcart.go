package models

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	UserID	uint
	CartItems	[]CartItem `gorm:"foreignKey:ShoppingCartID"`
}
