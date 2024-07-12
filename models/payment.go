package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	UserID	uint
	NameOnCard string `json:"nameoncard" binding:"required"`
	CardNumber string	`json:"cardnumber" binding:"required,len=16"`
	SecurityCode string `json:"securitycode" binding:"required,len=3"`
	ExpirationDate time.Time `json:"expirationdate" binding:"required"`
}