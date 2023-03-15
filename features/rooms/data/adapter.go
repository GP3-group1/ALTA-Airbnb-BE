package data

import (
	_imageData "alta-airbnb-be/features/images/data"
	"alta-airbnb-be/features/rooms"
	_roomModel "alta-airbnb-be/features/rooms/models"
)

func convertToGorm(roomEntity *rooms.RoomEntity) _roomModel.Room {
	roomModel := _roomModel.Room{
		UserID:      roomEntity.UserID,
		Name:        roomEntity.Name,
		Overview:    roomEntity.Overview,
		Description: roomEntity.Description,
		Location:    roomEntity.Location,
		Price:       roomEntity.Price,
		Facilities:  roomEntity.Facilities,
	}
	if roomEntity.ID != 0 {
		roomModel.ID = roomEntity.ID
	}
	return roomModel
}

func convertToEntity(roomModels *_roomModel.Room) *rooms.RoomEntity {
	roomEntity := rooms.RoomEntity{
		ID:          roomModels.ID,
		UserID:      roomModels.UserID,
		Name:        roomModels.Name,
		Overview:    roomModels.Overview,
		Description: roomModels.Description,
		Location:    roomModels.Location,
		Price:       roomModels.Price,
		Facilities:  roomModels.Facilities,
		Images:      _imageData.ConvertToEntities(roomModels.Images),
	}
	return &roomEntity
}

func convertToEntities(roomModels []*_roomModel.Room) []*rooms.RoomEntity {
	roomEntities := []*rooms.RoomEntity{}
	for _, val := range roomModels {
		roomEntities = append(roomEntities, convertToEntity(val))
	}
	return roomEntities
}
