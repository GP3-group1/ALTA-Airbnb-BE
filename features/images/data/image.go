package data

import (
	"alta-airbnb-be/app/storage"
	"alta-airbnb-be/features/images"
	"alta-airbnb-be/utils/consts"
	"errors"

	"gorm.io/gorm"
)

type ImageData struct {
	db *gorm.DB
}

func New(db *gorm.DB) images.ImageData_ {
	return &ImageData{
		db: db,
	}
}

func (imageData *ImageData) InsertImage(imageEntity *images.ImageEntity) (*images.ImageEntity, error) {
	imageGorm := convertToGorm(imageEntity)

	imageUrl, err := storage.GetStorageClient().UploadFile(imageEntity.Image, imageEntity.ImageName)
	if err != nil {
		return nil, err
	}
	imageGorm.Url = imageUrl

	tx := imageData.db.Create(&imageGorm)
	if tx.Error != nil {
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	imageEntityOutput := ConvertToEntity(imageGorm)
	return &imageEntityOutput, nil
}

func (imageData *ImageData) UpdateImage(imageEntity *images.ImageEntity) (*images.ImageEntity, error) {
	imageGorm := convertToGorm(imageEntity)

	imageUrl, err := storage.GetStorageClient().UploadFile(imageEntity.Image, imageEntity.ImageName)
	if err != nil {
		return nil, err
	}
	imageGorm.Url = imageUrl

	tx := imageData.db.Model(&imageGorm).Where("id = ?", imageGorm.ID).Updates((map[string]interface{}{"url": imageGorm.Url}))
	if tx.Error != nil {
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	imageEntityOutput := ConvertToEntity(imageGorm)
	return &imageEntityOutput, nil
}

func (imageData *ImageData) DeleteImage(imageEntity *images.ImageEntity) error {
	imageGorm := convertToGorm(imageEntity)

	tx := imageData.db.Where("id = ?", imageEntity.ID).Delete(&imageGorm)
	if tx.Error != nil {
		return errors.New(consts.SERVER_InternalServerError)
	}

	return nil
}
