package data

import (
	"alta-airbnb-be/features/images"
	_imageModel "alta-airbnb-be/features/images/models"
)

func ConvertToEntity(imageModel _imageModel.Image) images.ImageEntity {
	imageEntity := images.ImageEntity{
		ID:     imageModel.ID,
		RoomID: imageModel.ID,
		Url:    imageModel.Url,
	}
	return imageEntity
}

func ConvertToEntities(imageModel []_imageModel.Image) []images.ImageEntity {
	imageEntities := []images.ImageEntity{}
	for _, val := range imageModel {
		imageEntities = append(imageEntities,ConvertToEntity(val))
	}
	return imageEntities
}
