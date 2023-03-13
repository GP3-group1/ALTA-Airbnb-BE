package data

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/utils/consts"
	"errors"

	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
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
