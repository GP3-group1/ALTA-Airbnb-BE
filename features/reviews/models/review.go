package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  uint
	RoomID  uint
	Comment string  `gorm:"not null;type:text"`
	Rating  float64 `gorm:"not null"`
}
