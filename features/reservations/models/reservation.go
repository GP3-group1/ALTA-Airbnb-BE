package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	CheckInDate  time.Time `gorm:"not null;type:date"`
	CheckOutDate time.Time `gorm:"not null;type:date"`
	TotalNight   int       `gorm:"not null"`
	TotalPrice   int       `gorm:"not null"`
	RoomID       uint
	UserID       uint
	RoomName     string
	Price        int
	Balance      int
}
