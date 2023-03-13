package service

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/utils/helpers"

	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationData reservations.ReservationDataInterface_
	validate        *validator.Validate
}

// GetAll implements reservations.ReservationServiceInterface_
func (reservationService *reservationService) GetAll(page int, limit int, userID uint) ([]reservations.ReservationEntity, error) {
	limit, offset := helpers.LimitOffsetConvert(page, limit)
	reservationEntity, errSelect := reservationService.reservationData.SelectAll(offset, limit, userID)
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

	roomModel, errSelectRoom := reservationService.reservationData.SelectData(input.RoomID)
	if errSelectRoom != nil {
		return errSelectRoom
	}

	input.TotalPrice = roomModel.Price * input.TotalNight

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
