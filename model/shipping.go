package model

import (
	"gorm.io/gorm"
	"time"
)

type Shipping struct {
	gorm.Model
	ID        uint
	UserID    uint
	User      User
	Status    string
	CreatedAt time.Time
}
