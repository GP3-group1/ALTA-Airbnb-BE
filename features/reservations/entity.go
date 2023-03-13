package reservations

import (
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
	CheckInDate  string
	CheckOutDate string
}

//go:generate mockery --name ReservationService_ --output ../../mocks
type ReservationServiceInterface_ interface {
	Create(input ReservationEntity) error
}

//go:generate mockery --name ReservationData_ --output ../../mocks
type ReservationDataInterface_ interface {
	Insert(input ReservationEntity) error
}

//go:generate mockery --name ReservationDelivery_ --output ../../mocks
type ReservationDeliveryInterface_ interface {
	AddReservation(c echo.Context) error
	CheckReservation(c echo.Context) error
	GetAllReservation(c echo.Context) error
}
