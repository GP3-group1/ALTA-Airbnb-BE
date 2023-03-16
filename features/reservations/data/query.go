package data

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/reservations/models"
	"alta-airbnb-be/features/users"
	_mapUser "alta-airbnb-be/features/users/data"
	"alta-airbnb-be/utils/consts"
	"errors"

	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

// CheckReservation implements reservations.ReservationData_
func (reservationQuery *reservationQuery) CheckReservation(input reservations.ReservationEntity, roomID uint) (int64, error) {
	// parsing date
	CheckInDate := input.CheckInDate.Format("2006-01-02")
	CheckOutDate := input.CheckOutDate.Format("2006-01-02")

	var row int64
	txSelect := reservationQuery.db.Model(&models.Reservation{}).Where("room_id = ?", roomID).Where("(check_in_date <= ? AND check_out_date >= ?) OR (check_in_date >= ? AND check_out_date <= ?) OR (check_in_date <= ? AND check_out_date >= ?)", CheckOutDate, CheckInDate, CheckInDate, CheckOutDate, CheckInDate, CheckOutDate).Count(&row)
	if txSelect.Error != nil {
		return 0, txSelect.Error
	}
	return row, nil
}

// SelectUserBalance implements reservations.ReservationData_
func (reservationQuery *reservationQuery) SelectUser(userID uint) (reservations.ReservationEntity, error) {
	reservationGorm := models.Reservation{}
	txSelect := reservationQuery.db.Where("id = ?", userID).First(&reservationGorm.User)
	if txSelect.Error != nil {
		return reservations.ReservationEntity{}, txSelect.Error
	}
	if txSelect.RowsAffected == 0 {
		return reservations.ReservationEntity{}, errors.New(consts.SERVER_ZeroRowsAffected)
	}
	return GormToEntity(reservationGorm), nil
}

// SelectRoomPrice implements reservations.ReservationData_
func (reservationQuery *reservationQuery) SelectRoom(roomID uint) (reservations.ReservationEntity, error) {
	reservationGorm := models.Reservation{}
	txSelect := reservationQuery.db.Where("id = ?", roomID).First(&reservationGorm.Room)
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
	txSelect := reservationQuery.db.Preload("Room").Limit(limit).Offset(offset).Order("created_at DESC").Where("user_id = ?", userID).Find(&reservationGorm)
	if txSelect.Error != nil {
		return nil, txSelect.Error
	}
	return ListGormToEntity(reservationGorm), nil
}

// Insert implements reservations.ReservationDataInterface_
func (reservationQuery *reservationQuery) Insert(inputReservation reservations.ReservationEntity, inputUser users.UserEntity, userID uint) error {
	reservationGorm := EntityToGorm(inputReservation)
	userGorm := _mapUser.EntityToGorm(inputUser)

	txTransaction := reservationQuery.db.Begin()
	if txTransaction.Error != nil {
		txTransaction.Rollback()
		return txTransaction.Error
	}

	tx := txTransaction.Model(&userGorm).Where("id = ?", userID).Update("balance", userGorm.Balance)
	if tx.Error != nil || tx.RowsAffected == 0 {
		txTransaction.Rollback()
		return txTransaction.Error
	}

	tx = txTransaction.Create(&reservationGorm)
	if tx.Error != nil || tx.RowsAffected == 0 {
		txTransaction.Rollback()
		return txTransaction.Error
	}

	tx = txTransaction.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return txTransaction.Error
	}

	return nil
}

func New(db *gorm.DB) reservations.ReservationData_ {
	return &reservationQuery{
		db: db,
	}
}
