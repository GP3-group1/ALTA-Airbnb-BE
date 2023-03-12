package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	RoomID uint
	Url    string `gorm:"not null;type:text"`
}
