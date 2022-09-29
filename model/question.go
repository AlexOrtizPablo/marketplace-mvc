package model

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	gorm.Model
	ID          uint
	ProductID   uint
	Product     Product
	UserID      uint
	User        User
	Description string
	Answer      string
	CreatedAt   time.Time
	AnsweredAt  time.Time
}
