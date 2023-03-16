package data

import (
	"alta-airbnb-be/features/images"
	_imageModel "alta-airbnb-be/features/images/models"
)

func convertToGorm(imageEntity *images.ImageEntity) _imageModel.Image {
	imageModel := _imageModel.Image{
		RoomID: imageEntity.RoomID,
		Url:    imageEntity.Url,
	}
	if imageEntity.ID != 0 {
		imageModel.ID = imageEntity.ID
	}
	return imageModel
}

func ConvertToEntity(imageModel _imageModel.Image) images.ImageEntity {
	imageEntity := images.ImageEntity{
		ID:     imageModel.ID,
		RoomID: imageModel.RoomID,
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
