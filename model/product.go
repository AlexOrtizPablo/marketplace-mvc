package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID             uint
	Name           string
	Price          string
	Description    string
	SellerID       uint
	Seller         User
	Stock          int
	CreatedAt      time.Time
	ShoppingCartID uint
}
