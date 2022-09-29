package model

import (
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	gorm.Model
	ID             uint
	UserID         uint
	User           User
	ProductRatedID uint
	ProductRated   Product
	Stars          uint8
	Title          string
	Comment        string
	CreatedAt      time.Time
}
