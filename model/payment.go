package model

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	ID             uint
	PayerID        uint
	Payer          User
	Amount         int64
	Status         string
	ShoppingCartID uint
	ShoppingCart   *ShoppingCart
	CreatedAt      time.Time
}
