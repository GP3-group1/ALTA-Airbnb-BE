package delivery

import (
	"alta-airbnb-be/features/rooms"
	"strings"
)

func convertToEntity(roomRequest *rooms.RoomRequest) rooms.RoomEntity {
	// imageEntity := _imageDelivery.RequestToEntity(&roomRequest.ImageRequest)
	roomEntity := rooms.RoomEntity{
		ID:          roomRequest.ID,
		UserID:      roomRequest.UserID,
		Name:        roomRequest.Name,
		Overview:    roomRequest.Overview,
		Description: roomRequest.Description,
		Location:    roomRequest.Location,
		Price:       roomRequest.Price,
		Facilities:  roomRequest.Facilities,
	}
	return roomEntity
}

func convertToResponse(roomEntity *rooms.RoomEntity) *rooms.RoomResponse {
	facilities := []string{}
	for _, val := range strings.Split(roomEntity.Facilities, ", ") {
		facilities = append(facilities, val)
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
