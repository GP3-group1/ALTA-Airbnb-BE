package data

import (
	"alta-airbnb-be/app/storage"
	_imageModel "alta-airbnb-be/features/images/models"
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

	// Local
	imageUrl, err := storage.UploadFile(roomEntity.Image, roomEntity.ImageName)

	//GCS
	// imageUrl, err := storage.GetStorageClient().UploadFile(imageEntity.Image, imageEntity.ImageName)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx = txTransaction.Create(&_imageModel.Image{RoomID: roomGorm.ID, Url: imageUrl})
	if tx.Error != nil {
		txTransaction.Rollback()
		return errors.New(consts.SERVER_InternalServerError)
	}

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

	queryNormal := ""
	for key, val := range queryParams {
		if queryNormal != "" && key != "rating" {
			queryNormal += " AND "
		}
		if key == "price" {
			priceRange := strings.Split(val.(string), " - ")
			queryNormal += fmt.Sprintf("%s BETWEEN %s AND %s ", key, priceRange[0], priceRange[1])
		} else {
			queryNormal += fmt.Sprintf("%s LIKE %s%s%s ", key, "'%", val, "%'")
		}
	}

	tx := roomData.db.Where(queryNormal).Preload("Images").Find(&roomsGormOutput)
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

	tx := roomData.db.Preload("Images").First(&roomGorm)
	if tx.Error != nil {
		if tx.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, tx.Error
		}
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	roomEntity = convertToEntity(&roomGorm)
	row := roomData.db.Raw("SELECT u.name AS username, COALESCE(AVG(rs.rating), 0) AS avg_ratings FROM rooms LEFT JOIN reviews rs on rs.room_id  = rooms.id LEFT JOIN users u on u.id  = rooms.user_id WHERE rooms.id = ? AND rooms.deleted_AT IS NULL GROUP BY rooms.id, rs.id", roomEntity.ID).Row()
	row.Scan(&roomEntity.Username, &roomEntity.AVG_Ratings)
	if row.Err() != nil {
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	return roomEntity, nil
}

func (roomData *RoomData) SelectRoomsByUserId(roomEntity *rooms.RoomEntity) ([]*rooms.RoomEntity, error) {
	roomGorm := convertToGorm(roomEntity)
	fmt.Println(roomGorm)
	roomsGormOutput := []*_roomModel.Room{}

	tx := roomData.db.Where(&roomGorm).Preload("Images").Find(&roomsGormOutput)
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
	for _, val := range reviewEntities {
		tx = roomData.db.Raw("SELECT us.name FROM reviews LEFT JOIN users us on us.id = reviews.user_id WHERE room_id = ? AND reviews.deleted_AT IS NULL", val.RoomID).Scan(&val.Username)
		if tx.Error != nil {
			return nil, errors.New(consts.SERVER_InternalServerError)
		}
	}
	return reviewEntities, nil
}
