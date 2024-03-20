package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name             string `gorm:"not null"`
	SellerID         uint   `gorm:"not null"`
	Category         uint   `gorm:"not null"`
	BasePrice        uint   `gorm:"not null"`
	Details          string
	Description      string
	Image            string
	BiddingType      bool      `gorm:"default:false"`
	BiddingStartTime time.Time `gorm:"not null"`
	BiddingEndTime   time.Time `gorm:"not null"`
	ListedOn         time.Time `gorm:"not null"`
}
