package delivery

import (
	"alta-airbnb-be/features/reservations"
	"errors"
	"time"
)

func insertToEntity(reservationInsert reservations.ReservationInsert) (reservations.ReservationEntity, error) {
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

func entityToResponse(reservationEntity reservations.ReservationEntity) reservations.ReservationResponse {
	CheckInDate := reservationEntity.CheckInDate.Format("2006-01-02")
	CheckOutDate := reservationEntity.CheckOutDate.Format("2006-01-02")
	return reservations.ReservationResponse{
		ID:           reservationEntity.ID,
		RoomName:     reservationEntity.Room.Name,
		CheckInDate:  CheckInDate,
		CheckOutDate: CheckOutDate,
		Price:        float64(reservationEntity.Room.Price),
		TotalNight:   reservationEntity.TotalNight,
		TotalPrice:   reservationEntity.TotalPrice,
	}
}

func entityToResponseList(reservationEntity []reservations.ReservationEntity) []reservations.ReservationResponse {
	var dataResponse []reservations.ReservationResponse
	for _, v := range reservationEntity {
		dataResponse = append(dataResponse, entityToResponse(v))
	}
	return dataResponse
}
