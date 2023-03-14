package data

import (
	"alta-airbnb-be/features/reviews"
	_reviewModel "alta-airbnb-be/features/reviews/models"
)

func ConvertToEntity(reviewModel _reviewModel.Review) *reviews.ReviewEntity {
	reviewEntity := reviews.ReviewEntity{
		UserID:  reviewModel.UserID,
		RoomID:  reviewModel.RoomID,
		Comment: reviewModel.Comment,
		Rating:  reviewModel.Rating,
	}
	return &reviewEntity
}

func ConvertToEntities(reviewModels []_reviewModel.Review) []*reviews.ReviewEntity {
	reviewEntities := []*reviews.ReviewEntity{}
	for _, val := range reviewModels {
		reviewEntities = append(reviewEntities, &reviews.ReviewEntity{
			UserID:  val.UserID,
			RoomID:  val.RoomID,
			Comment: val.Comment,
			Rating:  val.Rating,
		})
	}
	return reviewEntities
}

func ConvertToGorm(reviewEntity *reviews.ReviewEntity) _reviewModel.Review {
	reviewModel := _reviewModel.Review{
		UserID:  reviewEntity.UserID,
		RoomID:  reviewEntity.RoomID,
		Comment: reviewEntity.Comment,
		Rating:  reviewEntity.Rating,
	}
	return reviewModel
}
