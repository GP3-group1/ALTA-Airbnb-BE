package models

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	rating float64
	RoomID uint
}
