package service

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/utils/consts"
	"errors"
	"fmt"

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

	reservationEntity, errSelectRoom := reservationService.reservationData.SelectRoomPrice(input.RoomID)
	if errSelectRoom != nil {
		return errSelectRoom
	}

	input.TotalPrice = int(reservationEntity.Price) * input.TotalNight

	reservationEntity, errSelectUser := reservationService.reservationData.SelectUserBalance(input.UserID)
	if errSelectUser != nil {
		return errSelectUser
	}

	fmt.Println(reservationEntity.Balance)
	fmt.Println(input.TotalPrice)

	if int(reservationEntity.Balance) < input.TotalPrice {
		return errors.New(consts.RESERVATION_InsertFailed)
	}

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
