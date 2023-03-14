package rooms

import (
	"alta-airbnb-be/features/facilities"
	"alta-airbnb-be/features/images"
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/reviews"

	"github.com/labstack/echo/v4"
)

type RoomEntity struct {
	ID           uint
	UserID       uint
	Username     string
	Name         string `validate:"required"`
	Overview     string `validate:"required"`
	Description  string `validate:"required"`
	Location     string `validate:"required"`
	Price        int    `validate:"required"`
	Reservations []reservations.ReservationEntity
	Images       []images.ImageEntity
	Reviews      []reviews.ReviewEntity
	AVG_Ratings  float64
	Facilities   []facilities.FacilityEntity
}

type RoomRequest struct {
	ID     uint `json:"id" form:"id"`
	UserID uint `json:"user_id" form:"user_id"`
	images.ImageRequest
	Name        string `json:"name" form:"name"`
	Overview    string `json:"overview" form:"overview"`
	Description string `json:"description" form:"description"`
	Location    string `json:"location" form:"location"`
	Price       int    `json:"price" form:"price"`
	Facilities  string `json:"facilities" form:"facilities"`
}

type RoomResponse struct {
	ID           uint                             `json:"id,omitempty" form:"id"`
	Username     string                           `json:"username,omitempty" form:"username"`
	UserID       uint                             `json:"user_id,omitempty" form:"user_id"`
	Name         string                           `json:"name,omitempty" form:"name"`
	Overview     string                           `json:"overview,omitempty" form:"overview"`
	Description  string                           `json:"description,omitempty" form:"description"`
	Location     string                           `json:"location,omitempty" form:"location"`
	Price        int                              `json:"price,omitempty" form:"price"`
	Reservations []reservations.ReservationEntity `json:"reservations,omitempty"`
	Images       []images.ImageEntity             `json:"ratings,omitempty"`
	Reviews      []reviews.ReviewEntity           `json:"reviews,omitempty"`
	AVG_Ratings  float64                          `json:"avg_ratings"`
	Facilities   []string                         `json:"facilities,omitempty"`
}

//go:generate mockery --name MenteeData_ --output ../../mocks
type RoomData_ interface {
	// InsertRoom(roomEntity *RoomEntity) error
	// UpdateRoom(roomEntity *RoomEntity) error
	// DeleteRoom(RoomEntity *RoomEntity) error
	// SelectRooms(limit, offset int, extractedQueryParams map[string]any) ([]*RoomEntity, error)
	// SelectRoomByRoomId(roomEntity *RoomEntity) (*RoomEntity, error)
	SelectRoomsByUserId(roomEntity *RoomEntity) ([]*RoomEntity, error)
	InsertReview(reviewEntity *reviews.ReviewEntity) error
	SelectReviewsByRoomId(reviewEntity *reviews.ReviewEntity) ([]*reviews.ReviewEntity, error)
}

//go:generate mockery --name MenteeService_ --output ../../mocks
type RoomService_ interface {
	// CreateRoom(roomEntity *RoomEntity) error
	// ChangeRoom(roomEntity *RoomEntity) error
	// RemoveRoom(RoomEntity *RoomEntity) error
	// GetRooms(limit, offset int, queryParams url.Values) ([]*RoomEntity, error)
	// GetRoomByRoomId(RoomEntity *RoomEntity) (*RoomEntity, error)
	GetRoomsByUserId(RoomEntity *RoomEntity) ([]*RoomEntity, error)
	CreateReview(reviewEntity *reviews.ReviewEntity) error
	GetReviewsByRoomId(reviewEntity *reviews.ReviewEntity) ([]*reviews.ReviewEntity, error)
}

//go:generate mockery --name MenteeDelivery_ --output ../../mocks
type RoomDelivery_ interface {
	// AddRoom(c echo.Context) error
	// ModifyRoom(c echo.Context) error
	// RemoveRoom(c echo.Context) error
	// GetRooms(c echo.Context) error
	// GetRoomByRoomId(c echo.Context) error
	GetRoomsByUserId(c echo.Context) error
	AddReview(c echo.Context) error
	GetReviewsByRoomId(c echo.Context) error
}
