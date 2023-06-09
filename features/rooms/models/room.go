package models

import (
	_imageModel "alta-airbnb-be/features/images/models"
	_reservationModel "alta-airbnb-be/features/reservations/models"
	_reviewModel "alta-airbnb-be/features/reviews/models"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name         string                          `gorm:"not null;unique;type:varchar(50)"`
	Overview     string                          `gorm:"not null;type:text"`
	Description  string                          `gorm:"not null;type:text"`
	Location     string                          `gorm:"not null;type:text"`
	Price        int                             `gorm:"not null"`
	Facilities   string                          `gorm:"not null;type:varchar(255)"`
	Reservations []_reservationModel.Reservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Reviews      []_reviewModel.Review           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Images       []_imageModel.Image             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID       uint
}
