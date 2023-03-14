package facilities

type FacilityEntity struct {
	RoomID uint
	Name   string
}

type FacilityRequest struct {
	RoomID uint   `json:"room_id" form:"room_id"`
	Name   string `json:"name" form:"name"`
}

type FacilityResponse struct {
	RoomID uint   `json:"room_id"`
	Name   string `json:"name"`
}
