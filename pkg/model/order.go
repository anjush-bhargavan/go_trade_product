package model

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	ProductID uint    `gorm:"not null"`
	UserID    uint    `gorm:"not null"`
	SellerID  uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	PaymentID string  `gorm:"not null"`
}
