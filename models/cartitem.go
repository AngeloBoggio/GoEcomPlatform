package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	ShoppingCartID uint
	WishlistID	uint
	ProductID	uint
	Quantity	int
	Product	Product `gorm:"foreignKey:ProductID"`
}