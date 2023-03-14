package delivery

import (
	_facilityDelivery "alta-airbnb-be/features/facilities/delivery"
	"alta-airbnb-be/features/images"
	_imageDelivery "alta-airbnb-be/features/images/delivery"
	"alta-airbnb-be/features/rooms"
)

func convertToEntity(roomRequest *rooms.RoomRequest) rooms.RoomEntity {
	imageEntity := _imageDelivery.RequestToEntity(&roomRequest.ImageRequest)
	facilityEntities := _facilityDelivery.ConvertToEntities(roomRequest)
	roomEntity := rooms.RoomEntity{
		ID:          roomRequest.ID,
		UserID:      roomRequest.UserID,
		Name:        roomRequest.Name,
		Overview:    roomRequest.Overview,
		Description: roomRequest.Description,
		Location:    roomRequest.Location,
		Price:       roomRequest.Price,
		Images:      []images.ImageEntity{imageEntity},
		Facilities:  facilityEntities,
	}
	return roomEntity
}

func convertToResponse(roomEntity *rooms.RoomEntity) *rooms.RoomResponse {
	facilities := []string{}
	for _, val := range roomEntity.Facilities {
		facilities = append(facilities, val.Name)
	}
	roomResponse := rooms.RoomResponse{
		ID:          roomEntity.ID,
		UserID:      roomEntity.UserID,
		Username:    roomEntity.Username,
		Name:        roomEntity.Name,
		Overview:    roomEntity.Overview,
		Description: roomEntity.Description,
		Location:    roomEntity.Location,
		Price:       roomEntity.Price,
		AVG_Ratings: roomEntity.AVG_Ratings,
		Facilities:  facilities,
	}
	return &roomResponse
}

func convertsToResponses(roomEntities []*rooms.RoomEntity) []*rooms.RoomResponse {
	roomResponses := []*rooms.RoomResponse{}
	for _, val := range roomEntities {
		roomResponses = append(roomResponses, convertToResponse(val))
	}
	return roomResponses
}
