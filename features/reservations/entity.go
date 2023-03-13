package reservations

import (
	"database/sql"
	"time"
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
