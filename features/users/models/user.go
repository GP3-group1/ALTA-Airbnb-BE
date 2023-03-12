package models

import (
	"gorm.io/gorm"

	_reservationModel "alta-airbnb-be/features/reservations/models"
	_roomModel "alta-airbnb-be/features/rooms/models"
)

type User struct {
	gorm.Model
	FullName     string            `gorm:"not null;type:varchar(50)"`
	Email        string            `gorm:"not null;unique;type:varchar(50)"`
	Password     string            `gorm:"not null;type:text"`
	Balance      float64           `gorm:"not null;default:1000"`
	Room         []_roomModel.Room `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Reservations []_reservationModel.Reservation
}
