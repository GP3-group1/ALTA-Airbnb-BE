package images

import "mime/multipart"

type ImageEntity struct {
	ID     uint
	RoomID uint
	Url    string
	Image  multipart.File
}

type ImageRequest struct {
	ID     uint           `json:"id" form:"id"`
	RoomID uint           `json:"room_id" form:"room_id"`
	Image  multipart.File `json:"image" form:"image"`
}

type ImageResponse struct {
	ID     uint   `json:"id,omitempty"`
	RoomID uint   `json:"room_id,omitempty"`
	Url    string `json:"url_image,omitempty"`
}
