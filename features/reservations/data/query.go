package data

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/reservations/models"
	"alta-airbnb-be/utils/consts"
	"errors"

	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

// SelectRoomPrice implements reservations.ReservationData_
func (reservationQuery *reservationQuery) SelectRoomPrice(roomID uint) (reservations.ReservationEntity, error) {
	reservationGorm := models.Reservation{}
	txSelect := reservationQuery.db.Where("rooms.id = ?", roomID).Select("rooms.price").First(&reservationGorm, roomID)
	if txSelect.Error != nil {
		return reservations.ReservationEntity{}, txSelect.Error
	}
	if txSelect.RowsAffected == 0 {
		return reservations.ReservationEntity{}, errors.New(consts.SERVER_ZeroRowsAffected)
	}
	return GormToEntity(reservationGorm), nil
}

// SelectAll implements reservations.ReservationDataInterface_
func (reservationQuery *reservationQuery) SelectAll(limit, offset int, userID uint) ([]reservations.ReservationEntity, error) {
	reservationGorm := []models.Reservation{}
	txSelect := reservationQuery.db.Limit(limit).Offset(offset).Where("reservations.user_id = ?", userID).Select("reservations.id, rooms.name AS room_name, reservations.check_in_date, reservations.check_out_date, rooms.price, reservations.total_night, reservations.total_price").Joins("JOIN rooms ON reservations.room_id = rooms.id").Find(&reservationGorm)
	if txSelect.Error != nil {
		return nil, txSelect.Error
	}
	return ListGormToEntity(reservationGorm), nil
}

// Insert implements reservations.ReservationDataInterface_
func (reservationQuery *reservationQuery) Insert(input reservations.ReservationEntity) error {
	reservationGorm := EntityToGorm(input)
	txInsert := reservationQuery.db.Create(&reservationGorm)
	if txInsert.Error != nil {
		return txInsert.Error
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.SERVER_ZeroRowsAffected)
	}
	return nil
}

func New(db *gorm.DB) reservations.ReservationData_ {
	return &reservationQuery{
		db: db,
	}
}
