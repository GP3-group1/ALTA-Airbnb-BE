package models

import (
	"gorm.io/gorm"

	_reservationModel "alta-airbnb-be/features/reservations/models"
	_reviewModel "alta-airbnb-be/features/reviews/models"
	_roomModel "alta-airbnb-be/features/rooms/models"
)

type User struct {
	gorm.Model
	Name         string                          `gorm:"not null;type:varchar(50)"`
	Email        string                          `gorm:"not null;unique;type:varchar(50)"`
	Password     string                          `gorm:"not null;type:text"`
	Sex          string                          `gorm:"type:varchar(10)"`
	Address      string                          `gorm:"type:varchar(100)"`
	PhoneNumber  string                          `gorm:"type:varchar(12)"`
	Balance      float64                         `gorm:"type:float not null"`
	Room         []_roomModel.Room               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Reservations []_reservationModel.Reservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Reviews      []_reviewModel.Review           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
