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
		CreatedAt:    reservationGorm.CreatedAt,
		UpdatedAt:    reservationGorm.UpdatedAt,
		Room: reservations.RoomEntity{
			Name:  reservationGorm.Room.Name,
			Price: reservationGorm.Room.Price,
		},
		User: reservations.UserEntity{
			Name:        reservationGorm.User.Name,
			Email:       reservationGorm.User.Email,
			Address:     reservationGorm.User.Address,
			PhoneNumber: reservationGorm.User.PhoneNumber,
			Balance:     reservationGorm.User.Balance,
		},
	}
}

func ListGormToEntity(reservationGorm []_reservationModel.Reservation) []reservations.ReservationEntity {
	var reservationEntities []reservations.ReservationEntity
	for _, v := range reservationGorm {
		reservationEntities = append(reservationEntities, GormToEntity(v))
	}
	return reservationEntities
}
