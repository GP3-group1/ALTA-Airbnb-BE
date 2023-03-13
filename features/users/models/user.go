package models

import (
	"gorm.io/gorm"

	_reservationModel "alta-airbnb-be/features/reservations/models"
	_roomModel "alta-airbnb-be/features/rooms/models"
)

type User struct {
	gorm.Model
	Name         string            `gorm:"not null;type:varchar(50)"`
	Email        string            `gorm:"not null;unique;type:varchar(50)"`
	Password     string            `gorm:"not null;type:text"`
	Sex          string            `gorm:"not null;type:enum('Male','Female');default:'Male'"`
	Address      string            `gorm:"type:varchar(100)"`
	PhoneNumber  string            `gorm:"type:varchar(12)"`
	Balance      int               `gorm:"not null;default:1000"`
	Room         []_roomModel.Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Reservations []_reservationModel.Reservation
}
