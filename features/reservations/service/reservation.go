package service

import (
	"alta-airbnb-be/features/reservations"

	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationData reservations.ReservationData_
	validate        *validator.Validate
}

// GetAll implements reservations.ReservationServiceInterface_
func (reservationService *reservationService) GetAll(page, limit int, userID uint) ([]reservations.ReservationEntity, error) {
	offset := (page - 1) * limit
	reservationEntity, errSelect := reservationService.reservationData.SelectAll(limit, offset, userID)
	if errSelect != nil {
		return []reservations.ReservationEntity{}, errSelect
	}
	return reservationEntity, nil
}

// Create implements reservations.ReservationServiceInterface_
func (reservationService *reservationService) Create(userID, idParam uint, input reservations.ReservationEntity) error {
	errValidate := reservationService.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	input.UserID = userID
	input.RoomID = idParam

	diff := input.CheckOutDate.Sub(input.CheckInDate)
	input.TotalNight = int(diff.Hours() / 24)

	// reservationService.roomData.SelectRoomByRoomId()

	// input.TotalPrice = roomModel.Price * input.TotalNight

	errInsert := reservationService.reservationData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func New(reservationData reservations.ReservationData_) reservations.ReservationService_ {
	return &reservationService{
		reservationData: reservationData,
		validate:        validator.New(),
	}
}
