package models

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	UserID	uint
	Items []WishlistItem `gorm:"foreignKey:WishlistID"`
}