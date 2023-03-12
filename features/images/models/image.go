package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	url    string `gorm:"not null;type:text"`
	RoomID uint
}
