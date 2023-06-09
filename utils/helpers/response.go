package helpers

import (
	"alta-airbnb-be/utils/consts"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Response(message string) map[string]any {
	return map[string]any{
		"message": message,
	}
}

func ResponseWithData(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}

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

func ValidateImageFailedResponse(c echo.Context, err error) (codeStatus int, failedMessage string) {
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

func ErrorResponse(err error) (int, interface{}) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError
	msg := err.Error()

	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "Atoi"):
		resp["message"] = "id must be number, cannot be string"
		code = http.StatusNotFound
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "you have insufficient balance"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "invalid input amount"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "not found"):
		resp["message"] = "data not found"
		code = http.StatusNotFound
	case strings.Contains(msg, "access"):
		resp["message"] = "restricted access"
		code = http.StatusInternalServerError
	case strings.Contains(msg, "deleted admin"):
		resp["message"] = "can't delete admin account"
		code = http.StatusInternalServerError
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "Duplicate"):
		if strings.Contains(msg, "username") {
			resp["message"] = "username is already in use"
			code = http.StatusConflict
		} else if strings.Contains(msg, "email") {
			resp["message"] = "email is already in use"
			code = http.StatusConflict
		} else {
			resp["message"] = "Internal server error"
			code = http.StatusInternalServerError
		}
	case strings.Contains(msg, "truncated"):
		if strings.Contains(msg, "team") {
			resp["message"] = "team input does not match category"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "status") {
			resp["message"] = "status input does not match category"
			code = http.StatusBadRequest
		}
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "hashedPassword"):
		resp["message"] = "password do not match"
		code = http.StatusForbidden
	case strings.Contains(msg, "validation"):
		resp["message"] = ValidationError(err)
		code = http.StatusBadRequest
	case strings.Contains(msg, "unmarshal"):
		if strings.Contains(msg, "fullname") {
			resp["message"] = "invalid fullname of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "username") {
			resp["message"] = "invalid username of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "gender") {
			resp["message"] = "invalid gender of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "email") {
			resp["message"] = "invalid email of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "password") {
			resp["message"] = "invalid password of type string"
			code = http.StatusBadRequest
		}
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	}
	return code, resp
}
