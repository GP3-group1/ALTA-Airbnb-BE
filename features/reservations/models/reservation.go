package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	CheckInDate  sql.NullTime `gorm:"not null;type:date"`
	CheckOutDate sql.NullTime `gorm:"not null;type:date"`
	TotalNight   int          `gorm:"not null"`
	TotalPrice   int          `gorm:"not null"`
	RoomID       uint
	UserID       uint
}
