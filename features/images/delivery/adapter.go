package delivery

import "alta-airbnb-be/features/images"

func ConvertToEntity(imageRequest *images.ImageRequest) images.ImageEntity {
	imageEntity := images.ImageEntity{
		ID:        imageRequest.ID,
		RoomID:    imageRequest.RoomID,
		Image:     imageRequest.Image,
		ImageName: imageRequest.ImageName,
	}
	return imageEntity
}

func ConvertToResponse(imageEntity *images.ImageEntity) images.ImageResponse {
	imageRequest := images.ImageResponse{
		ID:     imageEntity.ID,
		RoomID: imageEntity.RoomID,
		Url:    imageEntity.Url,
	}
	return imageRequest
}

func ConvertToResponses(imageEntity []images.ImageEntity) []images.ImageResponse {
	imageResponses := []images.ImageResponse{}
	for _, val := range imageEntity {
		imageResponses = append(imageResponses, ConvertToResponse(&val))
	}
	return imageResponses
}
