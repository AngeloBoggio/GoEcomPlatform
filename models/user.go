package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"` 
	Password string `json:"password" binding:"required"`
	ShoppingCart ShoppingCart `gorm:"foreignkey:UserID"`
	Wishlist	[]Wishlist	`gorm:"foreignkey:UserID"`
	PaymentMethod []PaymentMethod `gorm:"foreignkey:UserID"`
}
