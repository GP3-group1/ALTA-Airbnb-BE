package models

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	RoomID uint
	Rating float64
}
