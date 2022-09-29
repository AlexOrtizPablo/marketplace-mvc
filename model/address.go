package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID     uint
	Street string
	Number string
	CP     string
	State  string
}
