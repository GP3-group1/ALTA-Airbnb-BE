package images

import "mime/multipart"

type ImageEntity struct {
	RoomID uint
	Url    string
	Image  multipart.File
}

type ImageRequest struct {
	RoomID uint           `json:"room_id" form:"room_id"`
	Image  multipart.File `json:"image" form:"image"`
}

type ImageResponse struct {
	RoomID uint           `json:"room_id"`
	Image  multipart.File `json:"image"`
}
