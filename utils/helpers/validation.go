package helpers

import (
	"alta-airbnb-be/utils/consts"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ValidateRoomFailedResponse(c echo.Context, err error) (codeStatus int, failedMessage string) {
	if err.Error() == consts.ECHO_InvaildIdParam {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == consts.ROOM_InvalidInput {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == consts.ROOM_UserNotExisted {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == consts.ROOM_RoomNameAlreadyExisted {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == gorm.ErrRecordNotFound.Error() {
		return http.StatusBadRequest, err.Error()
	}
	return http.StatusInternalServerError, err.Error()
}

func ValidateReviewFailedResponse(c echo.Context, err error) (codeStatus int, failedMessage string) {
	if err.Error() == consts.ECHO_InvaildIdParam {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == consts.REVIEW_InvalidInput {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == consts.REVIEW_InvalidRatingInputRange {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == consts.REVIEW_UserNotExisted {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == consts.REVIEW_RoomNotExisted {
		return http.StatusBadRequest, err.Error()
	} else if err.Error() == gorm.ErrRecordNotFound.Error() {
		return http.StatusBadRequest, err.Error()
	}
	return http.StatusInternalServerError, err.Error()
}

func ValidationError(err error) string {
	reports := []string{}

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range castedObject {
			switch v.Tag() {
			case "required":
				reports = append(reports, fmt.Sprintf("%s is required", v.Field()))
			case "min":
				reports = append(reports, fmt.Sprintf("%s value must be greater than %s character", v.Field(), v.Param()))
			case "max":
				reports = append(reports, fmt.Sprintf("%s value must be lower than %s character", v.Field(), v.Param()))
			case "email":
				reports = append(reports, fmt.Sprintf("%s is not valid", v.Field()))
			case "lte":
				reports = append(reports, fmt.Sprintf("%s value must be below %s", v.Field(), v.Param()))
			case "gte":
				reports = append(reports, fmt.Sprintf("%s value must be above %s", v.Field(), v.Param()))
			case "numeric":
				reports = append(reports, fmt.Sprintf("%s value must be numeric", v.Field()))
			case "url":
				reports = append(reports, fmt.Sprintf("%s value must be url", v.Field()))
			}
		}
	}
	report := strings.Join(reports, ", ")
	return report
}
