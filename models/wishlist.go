package models

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	UserID	uint
	CartItems	[]CartItem `gorm:"foreignkey:WishlistID"`
}