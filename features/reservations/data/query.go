package data

import (
	"alta-airbnb-be/features/reservations"
	_modelRoom "alta-airbnb-be/features/rooms/models"
	"alta-airbnb-be/utils/consts"
	"errors"

	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

// SelectRoom implements reservations.ReservationDataInterface_
func (reservationQuery *reservationQuery) SelectData(roomID uint) (_modelRoom.Room, error) {
	roomGorm := _modelRoom.Room{}
	txSelect := reservationQuery.db.First(&roomGorm, roomID)
	if txSelect.Error != nil {
		return _modelRoom.Room{}, txSelect.Error
	}
	return roomGorm, nil
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

func New(db *gorm.DB) reservations.ReservationDataInterface_ {
	return &reservationQuery{
		db: db,
	}
}
