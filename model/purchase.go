package model

import (
	"gorm.io/gorm"
	"time"
)

type Purchase struct {
	gorm.Model
	ID        uint
	BuyerID   uint
	Buyer     User
	ProductID uint
	Product   Product
	CreatedAt time.Time
}
