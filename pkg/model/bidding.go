package model

import "gorm.io/gorm"

type Bid struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
}


type Bids struct {
	UserID    uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
}

type Bidding struct {
	gorm.Model
	ProductID uint    `gorm:"not null"`
	Bids []byte
}



