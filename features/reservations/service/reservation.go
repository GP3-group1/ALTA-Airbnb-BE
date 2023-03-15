package service

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/utils/consts"
	"errors"

	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationData reservations.ReservationData_
	validate        *validator.Validate
}

// CheckReservation implements reservations.ReservationService_
func (reservationService *reservationService) CheckReservation(input reservations.ReservationEntity, roomID uint) ([]reservations.ReservationEntity, error) {
	errValidate := reservationService.validate.StructExcept(input, "Room", "User")
	if errValidate != nil {
		return nil, errValidate
	}

	diff := input.CheckOutDate.Sub(input.CheckInDate)
	totalNight := int(diff.Hours() / 24)
	if totalNight < 1 {
		return nil, errors.New(consts.RESERVATION_InvalidInput)
	}

	reservationEntity, errSelect := reservationService.reservationData.CheckReservation(input, roomID)
	if errSelect != nil {
		return nil, errSelect
	}
	return reservationEntity, nil
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
	errValidate := reservationService.validate.StructExcept(inputReservation, "Room", "User")
	if errValidate != nil {
		return errValidate
	}

	inputReservation.UserID = userID
	inputReservation.RoomID = idParam

	diff := inputReservation.CheckOutDate.Sub(inputReservation.CheckInDate)
	inputReservation.TotalNight = int(diff.Hours() / 24)

	if inputReservation.TotalNight < 1 {
		return errors.New(consts.RESERVATION_InvalidInput)
	}

	selectRoom, errSelectRoom := reservationService.reservationData.SelectRoomPrice(idParam)
	if errSelectRoom != nil {
		return errSelectRoom
	}

	inputReservation.TotalPrice = float64(selectRoom.Room.Price) * float64(inputReservation.TotalNight)

	selectUser, errSelectUser := reservationService.reservationData.SelectUserBalance(inputReservation.UserID)
	if errSelectUser != nil {
		return errSelectUser
	}

	if selectUser.User.Balance < inputReservation.TotalPrice {
		return errors.New(consts.RESERVATION_InsertFailed)
	}

	inputUser := users.UserEntity{}
	inputUser.Balance = selectUser.User.Balance - inputReservation.TotalPrice

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
