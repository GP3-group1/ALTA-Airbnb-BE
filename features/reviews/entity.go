package reviews

type ReviewEntity struct {
	UserID  uint
	RoomID  uint
	Comment string  `validate:"required"`
	Rating  float64 `validate:"required"`
}

type ReviewRequest struct {
	UserID  uint    `json:"user_id" form:"user_id"`
	RoomID  uint    `json:"room_id" form:"room_id"`
	Comment string  `json:"comment" form:"comment"`
	Rating  float64 `json:"rating" form:"rating"`
}

type ReviewResponse struct {
	UserID  uint    `json:"user_id"`
	RoomID  uint    `json:"room_id"`
	Comment string  `json:"comment"`
	Rating  float64 `json:"rating"`
}
