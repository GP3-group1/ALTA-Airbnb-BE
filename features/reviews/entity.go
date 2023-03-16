package reviews

import "time"

type ReviewEntity struct {
	UserID    uint
	Username  string
	RoomID    uint
	Comment   string  `validate:"required"`
	Rating    float64 `validate:"required"`
	CreatedAt time.Time
}

type ReviewRequest struct {
	UserID  uint    `json:"user_id" form:"user_id"`
	RoomID  uint    `json:"room_id" form:"room_id"`
	Comment string  `json:"comment" form:"comment"`
	Rating  float64 `json:"rating" form:"rating"`
}

type ReviewResponse struct {
	UserID    uint      `json:"user_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	RoomID    uint      `json:"room_id,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	Rating    float64   `json:"rating,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
