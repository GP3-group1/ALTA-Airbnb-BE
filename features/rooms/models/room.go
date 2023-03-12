package room

import (
	_imageModel "alta-airbnb-be/features/images/models"
	_ratingModel "alta-airbnb-be/features/ratings/models"
	_reservationModel "alta-airbnb-be/features/reservations/models"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name         string                          `gorm:"not null;type:varchar(50)"`
	Description  string                          `gorm:"not null;type:text"`
	Location     string                          `gorm:"not null;type:text"`
	Price        float64                         `gorm:"not null"`
	Reservations []_reservationModel.Reservation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Ratings      []_ratingModel.Rating           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Images       []_imageModel.Image             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID       uint
}
