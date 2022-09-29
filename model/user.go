package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint
	Name      string
	LastName  string
	Password  string
	Mail      string
	AddressID uint
	Address   Address
	CreatedAt time.Time
}
