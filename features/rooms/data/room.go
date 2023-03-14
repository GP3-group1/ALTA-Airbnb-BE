package data

import (
	"alta-airbnb-be/features/reviews"
	_reviewData "alta-airbnb-be/features/reviews/data"
	"alta-airbnb-be/features/reviews/models"
	"alta-airbnb-be/features/rooms"
	_roomModel "alta-airbnb-be/features/rooms/models"
	"alta-airbnb-be/utils/consts"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type RoomData struct {
	db *gorm.DB
}

func New(db *gorm.DB) rooms.RoomData_ {
	return &RoomData{
		db: db,
	}
}

func (roomData *RoomData) InsertRoom(roomEntity *rooms.RoomEntity) error {
	roomGorm := convertToGorm(roomEntity)

	txTransaction := roomData.db.Begin()
	if txTransaction.Error != nil {
		txTransaction.Rollback()
		return errors.New(consts.SERVER_InternalServerError)
	}

	tx := txTransaction.Create(&roomGorm)
	if tx.Error != nil {
		txTransaction.Rollback()
		if strings.Contains(tx.Error.Error(), "Error 1452 (23000)") {
			return errors.New(consts.ROOM_UserNotExisted)
		}
		if strings.Contains(tx.Error.Error(), "Error 1062 (23000)") {
			return errors.New(consts.ROOM_RoomNameAlreadyExisted)
		}
		return errors.New(consts.SERVER_InternalServerError)
	}

	// err := storage.GetStorageClient().UploadFile(roomEntity.Image, roomEntity.ImageName)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	tx = txTransaction.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New(consts.SERVER_InternalServerError)
	}

	return nil
}

func (roomData *RoomData) UpdateRoom(roomEntity *rooms.RoomEntity) error {
	roomGorm := convertToGorm(roomEntity)

	txTransaction := roomData.db.Begin()
	if txTransaction.Error != nil {
		txTransaction.Rollback()
		return errors.New(consts.SERVER_InternalServerError)
	}

	tx := txTransaction.Model(&roomGorm).Where("id = ? AND user_id = ?", roomGorm.ID, roomGorm.UserID).Updates(&roomGorm)
	if tx.Error != nil {
		txTransaction.Rollback()
		if strings.Contains(tx.Error.Error(), "Error 1452 (23000)") {
			return errors.New(consts.ROOM_UserNotExisted)
		}
		if strings.Contains(tx.Error.Error(), "Error 1062 (23000)") {
			return errors.New(consts.ROOM_RoomNameAlreadyExisted)
		}
		return errors.New(consts.SERVER_InternalServerError)
	}

	tx = txTransaction.Commit()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New(consts.SERVER_InternalServerError)
	}

	return nil
}

func (roomData *RoomData) DeleteRoom(roomEntity *rooms.RoomEntity) error {
	tx := roomData.db.Where("id = ?", roomEntity.ID).Delete(&_roomModel.Room{})
	if tx.Error != nil {
		return errors.New(consts.SERVER_InternalServerError)
	}
	return nil
}

func (roomData *RoomData) SelectRooms(limit, offset int, queryParams map[string]any) ([]*rooms.RoomEntity, error) {
	roomsGormOutput := []*_roomModel.Room{}

	query := ""
	for key, val := range queryParams {
		if query != "" {
			query += " AND "
		}
		if key != "price" {
			query += fmt.Sprintf("%s LIKE %s%s%s ", key, "'%", val, "%'")
		} else {
			priceRange := strings.Split(val.(string), " - ")
			query += fmt.Sprintf("%s BETWEEN %s AND %s ", key, priceRange[0], priceRange[1])
		}
	}

	tx := roomData.db.Where(query).Find(&roomsGormOutput)
	if tx.Error != nil {
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	roomEntities := convertToEntities(roomsGormOutput)
	for _, val := range roomEntities {
		tx = roomData.db.Raw("SELECT COALESCE(AVG(rs.rating), 0) AS avg_ratings FROM rooms LEFT JOIN reviews rs on rs.room_id = rooms.id WHERE rooms.id = ? AND rooms.deleted_AT IS NULL", val.ID).Scan(&val.AVG_Ratings)
		if tx.Error != nil {
			return nil, errors.New(consts.SERVER_InternalServerError)
		}
	}

	return roomEntities, nil
}

func (roomData *RoomData) SelectRoomByRoomId(roomEntity *rooms.RoomEntity) (*rooms.RoomEntity, error) {
	roomGorm := convertToGorm(roomEntity)

	tx := roomData.db.First(&roomGorm)
	if tx.Error != nil {
		if tx.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, tx.Error
		}
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	roomEntity = convertToEntity(&roomGorm)
	row := roomData.db.Raw("SELECT u.name AS username, COALESCE(AVG(rs.rating), 0) AS avg_ratings FROM rooms LEFT JOIN reviews rs on rs.room_id  = rooms.id LEFT JOIN users u on u.id  = rooms.user_id WHERE rooms.id = ? AND rooms.deleted_AT IS NULL", roomEntity.ID).Row()
	row.Scan(&roomEntity.Username, &roomEntity.AVG_Ratings)
	if tx.Error != nil {
		if tx.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, tx.Error
		}
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	return roomEntity, nil
}

func (roomData *RoomData) SelectRoomsByUserId(roomEntity *rooms.RoomEntity) ([]*rooms.RoomEntity, error) {
	roomGorm := convertToGorm(roomEntity)
	roomsGormOutput := []*_roomModel.Room{}

	tx := roomData.db.Where(&roomGorm).Find(&roomsGormOutput)
	if tx.Error != nil {
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	roomEntities := convertToEntities(roomsGormOutput)
	for _, val := range roomEntities {
		tx = roomData.db.Raw("SELECT COALESCE(AVG(rs.rating), 0) AS avg_ratings FROM rooms LEFT JOIN reviews rs on rs.room_id = rooms.id WHERE rooms.id = ? AND rooms.deleted_AT IS NULL", val.ID).Scan(&val.AVG_Ratings)
		if tx.Error != nil {
			return nil, errors.New(consts.SERVER_InternalServerError)
		}
	}

	return roomEntities, nil
}

func (roomData *RoomData) InsertReview(reviewEntity *reviews.ReviewEntity) error {
	reviewGorm := _reviewData.ConvertToGorm(reviewEntity)

	tx := roomData.db.Create(&reviewGorm)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "users") {
			return errors.New(consts.REVIEW_UserNotExisted)
		}
		if strings.Contains(tx.Error.Error(), "rooms") {
			return errors.New(consts.REVIEW_RoomNotExisted)
		}
		return tx.Error
	}

	return nil
}

func (roomData *RoomData) SelectReviewsByRoomId(reviewEntity *reviews.ReviewEntity) ([]*reviews.ReviewEntity, error) {
	reviewGorm := _reviewData.ConvertToGorm(reviewEntity)
	reviewsGormOutput := []models.Review{}

	tx := roomData.db.Where(&reviewGorm).Find(&reviewsGormOutput)
	if tx.Error != nil {
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	reviewEntities := _reviewData.ConvertToEntities(reviewsGormOutput)
	return reviewEntities, nil
}
