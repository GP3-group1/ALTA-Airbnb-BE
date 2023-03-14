package service

import (
	"alta-airbnb-be/features/reviews"
	"alta-airbnb-be/features/rooms"

	"github.com/go-playground/validator/v10"
)

type RoomService struct {
	roomData rooms.RoomData_
	validate *validator.Validate
}

// func New(roomData rooms.RoomData_) rooms.RoomService_ {
// 	return &RoomService{
// 		roomData: roomData,
// 		validate: validator.New(),
// 	}
// }

// func (roomService *RoomService) CreateRoom(roomEntity *rooms.RoomEntity) error {
// 	err := roomService.validate.Struct(roomEntity)
// 	if err != nil {
// 		return errors.New(consts.ROOM_InvalidInput)
// 	}

// 	err = roomService.roomData.InsertRoom(roomEntity)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (roomService *RoomService) RemoveRoom(roomEntity *rooms.RoomEntity) error {
// 	err := roomService.roomData.DeleteRoom(roomEntity)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (roomService *RoomService) GetRooms(limit, offset int, queryParams url.Values) ([]*rooms.RoomEntity, error) {
// 	extractedQueryParams := helpers.ExtractQueryParams(queryParams)
// 	roomEntities, err := roomService.roomData.SelectRooms(limit, offset, extractedQueryParams)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return roomEntities, nil
// }

// func (roomService *RoomService) GetRoomByRoomId(roomEntity *rooms.RoomEntity) (*rooms.RoomEntity, error) {
// 	roomEntity, err := roomService.roomData.SelectRoomByRoomId(roomEntity)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return roomEntity, nil
// }

func (roomService *RoomService) GetRoomsByUserId(roomEntity *rooms.RoomEntity) ([]*rooms.RoomEntity, error) {
	roomEntities, err := roomService.roomData.SelectRoomsByUserId(roomEntity)
	if err != nil {
		return nil, err
	}
	return roomEntities, nil
}

func (roomService *RoomService) CreateReview(reviewEntity *reviews.ReviewEntity) error {
	err := Validate(roomService, reviewEntity)
	if err != nil {
		return err
	}

	err = roomService.roomData.InsertReview(reviewEntity)
	if err != nil {
		return err
	}
	return nil
}

func (roomService *RoomService) GetReviewsByRoomId(reviewEntity *reviews.ReviewEntity) ([]*reviews.ReviewEntity, error) {
	reviewEntities, err := roomService.roomData.SelectReviewsByRoomId(reviewEntity)
	if err != nil {
		return nil, err
	}
	return reviewEntities, nil
}
