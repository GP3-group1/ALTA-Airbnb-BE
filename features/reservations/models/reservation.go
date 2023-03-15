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
	TotalPrice   float64   `gorm:"not null"`
	RoomID       uint
	UserID       uint
	Room         Room
	User         User
}

type Room struct {
	gorm.Model
	Name  string `gorm:"not null;unique;type:varchar(50)"`
	Price int    `gorm:"not null"`
}

type User struct {
	gorm.Model
	Name        string  `gorm:"not null;type:varchar(50)"`
	Email       string  `gorm:"not null;unique;type:varchar(50)"`
	Address     string  `gorm:"type:varchar(100)"`
	PhoneNumber string  `gorm:"type:varchar(12)"`
	Balance     float64 `gorm:"type:float not null"`
}
