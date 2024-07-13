package models

import (
	"gorm.io/gorm"
)

type WishlistItem struct {
	gorm.Model
	WishlistID	uint
	ProductID uint
	Product Product `gorm:"foreignKey:ProductID"`
}

