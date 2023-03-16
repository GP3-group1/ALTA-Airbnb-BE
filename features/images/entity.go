package images

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type ImageEntity struct {
	ID        uint
	RoomID    uint `validate:"required"`
	Url       string
	Image     multipart.File
	ImageName string
}

type ImageRequest struct {
	ID        uint           `json:"id" form:"id"`
	RoomID    uint           `json:"room_id" form:"room_id"`
	Image     multipart.File `json:"image" form:"image"`
	ImageName string         `json:"image_name" form:"image_name"`
}

type ImageResponse struct {
	ID     uint   `json:"id,omitempty"`
	RoomID uint   `json:"room_id,omitempty"`
	Url    string `json:"url_image,omitempty"`
}

type ImageData_ interface {
	InsertImage(imageEntity *ImageEntity) (*ImageEntity, error)
	UpdateImage(imageEntity *ImageEntity) (*ImageEntity, error)
	DeleteImage(imageEntity *ImageEntity) error
}

//go:generate mockery --name MenteeService_ --output ../../mocks
type ImageService_ interface {
	CreateImage(userId uint, imageEntity *ImageEntity) (*ImageEntity, error)
	ChangeImage(userId uint, imageEntity *ImageEntity) (*ImageEntity, error)
	RemoveImage(userId uint, imageEntity *ImageEntity) error
}

//go:generate mockery --name MenteeDelivery_ --output ../../mocks
type ImageDelivery_ interface {
	AddImage(c echo.Context) error
	ModifyImage(c echo.Context) error
	RemoveImage(c echo.Context) error
}
