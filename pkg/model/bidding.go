package model

import "gorm.io/gorm"

type Bid struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
}
