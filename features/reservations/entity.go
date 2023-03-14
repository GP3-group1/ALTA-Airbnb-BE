package reservations

import (
	_modelRoom "alta-airbnb-be/features/rooms/models"
	"time"

	"github.com/labstack/echo/v4"
)

type ReservationEntity struct {
	ID           uint
	CheckInDate  time.Time
	CheckOutDate time.Time
	TotalNight   int
	TotalPrice   int
	RoomID       uint
	UserID       uint
	RoomName     string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ReservationInsert struct {
	CheckInDate  string `json:"check_in" form:"check_in"`
	CheckOutDate string `json:"check_out" form:"check_out"`
}

type ReservationRequest struct {
	CheckInDate  time.Time `json:"check_in_date" form:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date" form:"check_out_date"`
	TotalNight   int       `json:"total_night" form:"total_night"`
	TotalPrice   int       `json:"total_price" form:"total_price"`
	RoomID       uint      `json:"room_id" form:"room_id"`
	UserID       uint      `json:"user_id" form:"user_id"`
}

type ReservationResponse struct {
	ID           uint    `json:"id"`
	RoomName     string  `json:"room_name"`
	CheckInDate  string  `json:"check_in"`
	CheckOutDate string  `json:"check_out"`
	Price        float64 `json:"price"`
	TotalNight   int     `json:"total_night"`
	TotalPrice   int     `json:"total_price"`
}

//go:generate mockery --name ReservationService_ --output ../../mocks
type ReservationServiceInterface_ interface {
	Create(userID, idParam uint, input ReservationEntity) error
	GetAll(page, limit int, userID uint) ([]ReservationEntity, error)
}

//go:generate mockery --name ReservationData_ --output ../../mocks
type ReservationDataInterface_ interface {
	SelectData(roomID uint) (_modelRoom.Room, error)
	Insert(input ReservationEntity) error
	SelectAll(limit, offset int, userID uint) ([]ReservationEntity, error)
}

//go:generate mockery --name ReservationDelivery_ --output ../../mocks
type ReservationDeliveryInterface_ interface {
	AddReservation(c echo.Context) error
	CheckReservation(c echo.Context) error
	GetAllReservation(c echo.Context) error
}
