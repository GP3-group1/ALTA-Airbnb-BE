package delivery

import (
	"alta-airbnb-be/features/reservations"
	"errors"
	"time"
)

func insertToEntity(reservationInsert *reservations.ReservationInsert) (reservations.ReservationEntity, error) {
	CheckInDate, err := time.Parse("2006-01-02", reservationInsert.CheckInDate)
	if err != nil {
		return reservations.ReservationEntity{}, errors.New("invalid check in date format must YYYY-MM-DD")
	}
	CheckOutDate, err := time.Parse("2006-01-02", reservationInsert.CheckOutDate)
	if err != nil {
		return reservations.ReservationEntity{}, errors.New("invalid check out date format must YYYY-MM-DD")
	}
	return reservations.ReservationEntity{
		CheckInDate:  CheckInDate,
		CheckOutDate: CheckOutDate,
	}, nil
}
