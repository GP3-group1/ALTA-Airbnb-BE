package service

import (
	"alta-airbnb-be/features/reservations"

	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationData reservations.ReservationDataInterface_
	validate        *validator.Validate
}

// Create implements reservations.ReservationServiceInterface_
func (*reservationService) Create(input reservations.ReservationEntity) error {
	panic("unimplemented")
}

func New(reservationData reservations.ReservationDataInterface_) reservations.ReservationServiceInterface_ {
	return &reservationService{
		reservationData: reservationData,
		validate:        validator.New(),
	}
}
