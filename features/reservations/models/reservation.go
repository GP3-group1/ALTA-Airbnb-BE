package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	CheckIn    sql.NullTime `gorm:"not null;type:date"`
	CheckOut   sql.NullTime `gorm:"not null;type:date"`
	TotalNight int          `gorm:"not null"`
	TotalPrice int          `gorm:"not null"`
	Rating     float64      `gorm:"not null"`
	RoomID     uint
	UserID     uint
}
