package model

import (
	"gorm.io/gorm"
	"time"
)

type ShoppingCart struct {
	gorm.Model
	ID         uint
	BuyerID    uint
	Buyer      User
	Products   []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalPrice int64
	UpdatedAt  time.Time
	Quantities []uint `gorm:"-:migration"`
}
