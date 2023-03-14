package delivery

import "alta-airbnb-be/features/images"

func RequestToEntity(imageRequest *images.ImageRequest) (images.ImageEntity) {
	imageEntity := images.ImageEntity{
		RoomID: imageRequest.RoomID,
		Image:  imageRequest.Image,
	}
	return imageEntity
}
