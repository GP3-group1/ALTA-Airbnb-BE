package service

import (
	"alta-airbnb-be/features/images"
	"alta-airbnb-be/features/rooms"
	"alta-airbnb-be/utils/consts"
	"errors"

	"github.com/go-playground/validator/v10"
)

type ImageService struct {
	imageData images.ImageData_
	roomData  rooms.RoomData_
	validate  *validator.Validate
}

func New(imageData images.ImageData_, roomData rooms.RoomData_) images.ImageService_ {
	return &ImageService{
		imageData: imageData,
		roomData:  roomData,
		validate:  validator.New(),
	}
}

func (imageService *ImageService) checkPermission(userId uint, imageEntity *images.ImageEntity) error {
	roomEntity := rooms.RoomEntity{UserID: userId}
	roomEntities, err := imageService.roomData.SelectRoomsByUserId(&roomEntity)
	if err != nil {
		return err
	}
	if len(roomEntities) != 0 {
		for _, val := range roomEntities {
			if val.ID == imageEntity.RoomID {
				return nil
			}
		}
	}
	return errors.New(consts.SERVER_ForbiddenRequest)
}

func (imageService *ImageService) CreateImage(userId uint, imageEntity *images.ImageEntity) (*images.ImageEntity, error) {
	err := imageService.checkPermission(userId, imageEntity)
	if err != nil {
		return nil, err
	}

	imageEntity, err = imageService.imageData.InsertImage(imageEntity)
	if err != nil {
		return nil, err
	}
	return imageEntity, nil
}

func (imageService *ImageService) ChangeImage(userId uint, imageEntity *images.ImageEntity) (*images.ImageEntity, error) {
	err := imageService.checkPermission(userId, imageEntity)
	if err != nil {
		return nil, err
	}

	imageEntity, err = imageService.imageData.UpdateImage(imageEntity)
	if err != nil {
		return nil, err
	}
	return imageEntity, nil
}

func (imageService *ImageService) RemoveImage(userId uint, imageEntity *images.ImageEntity) error {
	err := imageService.checkPermission(userId, imageEntity)
	if err != nil {
		return err
	}

	err = imageService.imageData.DeleteImage(imageEntity)
	if err != nil {
		return err
	}
	return nil
}
