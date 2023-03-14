package data

import (
	"alta-airbnb-be/features/reservations"
	_reservationModel "alta-airbnb-be/features/reservations/models"
)

func EntityToGorm(reservationEntity reservations.ReservationEntity) _reservationModel.Reservation {
	reservationGorm := _reservationModel.Reservation{
		CheckInDate:  reservationEntity.CheckInDate,
		CheckOutDate: reservationEntity.CheckOutDate,
		TotalNight:   reservationEntity.TotalNight,
		TotalPrice:   reservationEntity.TotalPrice,
		RoomID:       reservationEntity.RoomID,
		UserID:       reservationEntity.UserID,
		RoomName:     reservationEntity.RoomName,
		Price:        reservationEntity.Price,
		Balance:      reservationEntity.Balance,
	}
	return reservationGorm
}

func GormToEntity(reservationGorm _reservationModel.Reservation) reservations.ReservationEntity {
	return reservations.ReservationEntity{
		ID:           reservationGorm.ID,
		CheckInDate:  reservationGorm.CheckInDate,
		CheckOutDate: reservationGorm.CheckOutDate,
		TotalNight:   reservationGorm.TotalNight,
		TotalPrice:   reservationGorm.TotalPrice,
		RoomID:       reservationGorm.RoomID,
		UserID:       reservationGorm.UserID,
		RoomName:     reservationGorm.RoomName,
		Price:        reservationGorm.Price,
		Balance:      reservationGorm.Balance,
		CreatedAt:    reservationGorm.CreatedAt,
		UpdatedAt:    reservationGorm.UpdatedAt,
	}
}

func ListGormToEntity(reservationGorm []_reservationModel.Reservation) []reservations.ReservationEntity {
	var reservationEntities []reservations.ReservationEntity
	for _, v := range reservationGorm {
		reservationEntities = append(reservationEntities, GormToEntity(v))
	}
	return reservationEntities
}
