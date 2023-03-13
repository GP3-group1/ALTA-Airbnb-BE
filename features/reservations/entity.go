package reservations

import (
	_modelRoom "alta-airbnb-be/features/rooms/models"
	"time"

	"github.com/labstack/echo/v4"
)

type ReservationEntity struct {
	ID           uint
	CheckInDate  time.Time `validate:"required"`
	CheckOutDate time.Time `validate:"required"`
	TotalNight   int
	TotalPrice   int
	RoomID       uint
	UserID       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ReservationInsert struct {
	CheckInDate  string `json:"check_in" form:"check_in"`
	CheckOutDate string `json:"check_out" form:"check_out"`
}

//go:generate mockery --name ReservationService_ --output ../../mocks
type ReservationServiceInterface_ interface {
	Create(userID, idParam uint, input ReservationEntity) error
}

//go:generate mockery --name ReservationData_ --output ../../mocks
type ReservationDataInterface_ interface {
	Insert(input ReservationEntity) error
	SelectData(roomID uint) (_modelRoom.Room, error)
}

//go:generate mockery --name ReservationDelivery_ --output ../../mocks
type ReservationDeliveryInterface_ interface {
	AddReservation(c echo.Context) error
	CheckReservation(c echo.Context) error
	GetAllReservation(c echo.Context) error
}
