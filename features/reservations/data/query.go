package data

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/reservations/models"
	_modelRoom "alta-airbnb-be/features/rooms/models"
	"alta-airbnb-be/utils/consts"
	"alta-airbnb-be/utils/helpers"
	"errors"

	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

// SelectAll implements reservations.ReservationDataInterface_
func (reservationQuery *reservationQuery) SelectAll(page int, limit int, userID uint) ([]reservations.ReservationEntity, error) {
	reservationGorm := []models.Reservation{}
	limit, offset := helpers.LimitOffsetConvert(page, limit)
	txSelect := reservationQuery.db.Offset(offset).Limit(limit).Select("reservations.id, room.name AS room_name, reservations.check_in_date, reservations.check_out_date, rooms.price, reservations.total_night, reservations.total_price").Joins("JOIN rooms ON reservations.room_id = rooms.id").Find(&reservationGorm, "user_id = ?", userID)
	if txSelect.Error != nil {
		return []reservations.ReservationEntity{}, txSelect.Error
	}
	if txSelect.RowsAffected == 0 {
		return []reservations.ReservationEntity{}, errors.New(consts.SERVER_ZeroRowsAffected)
	}
	return ListGormToEntity(reservationGorm), nil
}

// SelectRoom implements reservations.ReservationDataInterface_
func (reservationQuery *reservationQuery) SelectData(roomID uint) (_modelRoom.Room, error) {
	roomGorm := _modelRoom.Room{}
	txSelect := reservationQuery.db.First(&roomGorm, roomID)
	if txSelect.Error != nil {
		return _modelRoom.Room{}, txSelect.Error
	}
	if txSelect.RowsAffected == 0 {
		return _modelRoom.Room{}, errors.New(consts.SERVER_ZeroRowsAffected)
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
