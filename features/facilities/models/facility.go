package models

import (
	"gorm.io/gorm"
)

type Facility struct {
	gorm.Model
	RoomID uint
	Name   string `gorm:"not null;type:varchar(30)"`
}
