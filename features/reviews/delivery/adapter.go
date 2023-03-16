package delivery

import (
	"alta-airbnb-be/features/reviews"
)

func ConvertToEntity(reviewRequest *reviews.ReviewRequest) reviews.ReviewEntity {
	reviewEntity := reviews.ReviewEntity{
		UserID:  reviewRequest.UserID,
		RoomID:  reviewRequest.RoomID,
		Comment: reviewRequest.Comment,
		Rating:  reviewRequest.Rating,
	}
	return reviewEntity
}

func ConvertToResponse(reviewEntitiy *reviews.ReviewEntity) reviews.ReviewResponse {
	reviewResponse := reviews.ReviewResponse{
		UserID:    reviewEntitiy.UserID,
		Username:  reviewEntitiy.Username,
		RoomID:    reviewEntitiy.RoomID,
		Comment:   reviewEntitiy.Comment,
		Rating:    reviewEntitiy.Rating,
		CreatedAt: reviewEntitiy.CreatedAt.Format("2006-01-02"),
	}
	return reviewResponse
}

func ConvertToResponses(reviewEntities []*reviews.ReviewEntity) []reviews.ReviewResponse {
	reviewResponses := []reviews.ReviewResponse{}
	for _, val := range reviewEntities {
		reviewResponses = append(reviewResponses, ConvertToResponse(val))
	}
	return reviewResponses
}
