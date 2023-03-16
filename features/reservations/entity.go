package reservations

import (
	"alta-airbnb-be/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type ReservationEntity struct {
	ID           uint
	CheckInDate  time.Time `validate:"required"`
	CheckOutDate time.Time `validate:"required"`
	TotalNight   int
	TotalPrice   float64
	RoomID       uint
	UserID       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Room         RoomEntity
	User         UserEntity
}

type UserEntity struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Address     string
	PhoneNumber string
	Balance     float64
}

type RoomEntity struct {
	ID    uint
	Name  string `validate:"required"`
	Price int    `validate:"required"`
}

type ReservationInsert struct {
	CheckInDate  string `json:"check_in" form:"check_in"`
	CheckOutDate string `json:"check_out" form:"check_out"`
}

type ReservationRequest struct {
	CheckInDate  time.Time `json:"check_in_date" form:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date" form:"check_out_date"`
	TotalNight   int       `json:"total_night" form:"total_night"`
	TotalPrice   float64   `json:"total_price" form:"total_price"`
	RoomID       uint      `json:"room_id" form:"room_id"`
	UserID       uint      `json:"user_id" form:"user_id"`
}

type ReservationResponse struct {
	ID           uint    `json:"id"`
	RoomID       uint    `json:"room_id"`
	RoomName     string  `json:"room_name"`
	CheckInDate  string  `json:"check_in"`
	CheckOutDate string  `json:"check_out"`
	Price        float64 `json:"price"`
	TotalNight   int     `json:"total_night"`
	TotalPrice   float64 `json:"total_price"`
}

type MidtransResponse struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}

//go:generate mockery --name ReservationService_ --output ../../mocks
type ReservationService_ interface {
	Create(userID, idParam uint, inputReservation ReservationEntity) (MidtransResponse, error)
	GetAll(page, limit int, userID uint) ([]ReservationEntity, error)
	CheckReservation(input ReservationEntity, roomID uint) (int, error)
}

//go:generate mockery --name ReservationData_ --output ../../mocks
type ReservationData_ interface {
	SelectRoom(roomID uint) (ReservationEntity, error)
	SelectUser(userID uint) (ReservationEntity, error)
	Insert(inputReservation ReservationEntity, inputUser users.UserEntity, userID uint) error
	SelectAll(limit, offset int, userID uint) ([]ReservationEntity, error)
	CheckReservation(input ReservationEntity, roomID uint) (int, error)
}

//go:generate mockery --name ReservationDelivery_ --output ../../mocks
type ReservationDelivery_ interface {
	AddReservation(c echo.Context) error
	CheckReservation(c echo.Context) error
	GetAllReservation(c echo.Context) error
}
