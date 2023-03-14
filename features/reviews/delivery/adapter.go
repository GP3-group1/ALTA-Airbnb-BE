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

func ConvertToResponses(reviewEntities []*reviews.ReviewEntity) []reviews.ReviewResponse {
	reviewResponses := []reviews.ReviewResponse{}
	for _, val := range reviewEntities {
		reviewResponses = append(reviewResponses, reviews.ReviewResponse{
			UserID:  val.UserID,
			RoomID:  val.RoomID,
			Comment: val.Comment,
			Rating:  val.Rating,
		})
	}
	return reviewResponses
}
