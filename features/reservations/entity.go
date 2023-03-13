package reservations

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
)

type ReservationEntity struct {
	ID           uint
	CheckInDate  sql.NullTime `validate:"required"`
	CheckOutDate sql.NullTime `validate:"required"`
	TotalNight   int
	TotalPrice   int
	RoomID       uint
	UserID       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ReservationInsert struct {
	CheckInDate  sql.NullTime
	CheckOutDate sql.NullTime
}

//go:generate mockery --name ReservationService_ --output ../../mocks
type UserServiceInterface_ interface {
	Create(input ReservationEntity) error
}

//go:generate mockery --name ReservationData_ --output ../../mocks
type UserDataInterface_ interface {
	Insert(input ReservationEntity) error
}

//go:generate mockery --name ReservationDelivery_ --output ../../mocks
type UserDeliveryInterface_ interface {
	AddReservation(c echo.Context) error
	CheckReservation(c echo.Context) error
	GetAllReservation(c echo.Context) error
}
