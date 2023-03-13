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
func (reservationService *reservationService) Create(userID, idParam uint, input reservations.ReservationEntity) error {
	errValidate := reservationService.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	input.UserID = userID
	input.RoomID = idParam

	errInsert := reservationService.reservationData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func New(reservationData reservations.ReservationDataInterface_) reservations.ReservationServiceInterface_ {
	return &reservationService{
		reservationData: reservationData,
		validate:        validator.New(),
	}
}
