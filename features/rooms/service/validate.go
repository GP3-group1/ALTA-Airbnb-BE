package service

import (
	"alta-airbnb-be/features/reviews"
	"alta-airbnb-be/utils/consts"
	"errors"
)

func Validate(roomService *RoomService, reviewEntity *reviews.ReviewEntity) error {
	err := roomService.validate.Struct(reviewEntity)
	if err != nil {
		return errors.New(consts.REVIEW_InvalidInput)
	}

	if reviewEntity.Rating > 5.0 {
		return errors.New(consts.REVIEW_InvalidRatingInputRange)
	}
	return nil
}