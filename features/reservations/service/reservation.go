package service

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/utils/consts"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationData reservations.ReservationData_
	validate        *validator.Validate
}

// CheckReservation implements reservations.ReservationService_
func (*reservationService) CheckReservation(CheckInDate time.Time, CheckOutDate time.Time, roomID uint) error {
	panic("unimplemented")
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
func (reservationService *reservationService) Create(userID, idParam uint, inputReservation reservations.ReservationEntity) error {
	errValidate := reservationService.validate.Struct(inputReservation)
	if errValidate != nil {
		return errValidate
	}

	inputReservation.UserID = userID
	inputReservation.RoomID = idParam

	diff := inputReservation.CheckOutDate.Sub(inputReservation.CheckInDate)
	inputReservation.TotalNight = int(diff.Hours() / 24)

	reservationEntity, errSelectRoom := reservationService.reservationData.SelectRoomPrice(inputReservation.RoomID)
	if errSelectRoom != nil {
		return errSelectRoom
	}

	inputReservation.TotalPrice = reservationEntity.Price * inputReservation.TotalNight

	reservationEntity, errSelectUser := reservationService.reservationData.SelectUserBalance(inputReservation.UserID)
	if errSelectUser != nil {
		return errSelectUser
	}

	if int(reservationEntity.Balance) < inputReservation.TotalPrice {
		return errors.New(consts.RESERVATION_InsertFailed)
	}

	inputUser := users.UserEntity{}
	inputUser.Balance = reservationEntity.Balance - inputReservation.TotalPrice

	errInsert := reservationService.reservationData.Insert(inputReservation, inputUser, userID)
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
