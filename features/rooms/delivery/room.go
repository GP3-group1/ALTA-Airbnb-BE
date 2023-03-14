package delivery

import (
	"alta-airbnb-be/features/reviews"
	_reviewDelivery "alta-airbnb-be/features/reviews/delivery"
	"alta-airbnb-be/features/rooms"
	"alta-airbnb-be/middlewares"
	"alta-airbnb-be/utils/consts"
	"alta-airbnb-be/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoomDelivery struct {
	roomService rooms.RoomService_
}

func New(roomService rooms.RoomService_) rooms.RoomDelivery_ {
	return &RoomDelivery{
		roomService: roomService,
	}
}

func (roomDelivery *RoomDelivery) AddRoom(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	roomRequest := rooms.RoomRequest{}
	err := c.Bind(&roomRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.ROOM_ErrorBindRoomData))
	}
	roomRequest.UserID = userId

	file, fileName, err := helpers.ExtractImage(c, "image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(err.Error()))
	}
	roomRequest.Image = file
	roomRequest.ImageName = fileName

	roomEntity := convertToEntity(&roomRequest)
	err = roomDelivery.roomService.CreateRoom(&roomEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateRoomFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.ROOM_SuccessInsertRoomData))
}

func (roomDelivery *RoomDelivery) ModifyRoom(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	roomId, err := helpers.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}

	roomRequest := rooms.RoomRequest{}
	err = c.Bind(&roomRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.ROOM_ErrorBindRoomData))
	}
	roomRequest.ID = roomId
	roomRequest.UserID = userId

	file, _, err := helpers.ExtractImage(c, "image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(err.Error()))
	}
	roomRequest.Image = file

	roomEntity := convertToEntity(&roomRequest)
	err = roomDelivery.roomService.ChangeRoom(&roomEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateRoomFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.ROOM_SuccessUpdateRoomData))
}

func (roomDelivery *RoomDelivery) GetRooms(c echo.Context) error {
	page, limit, err := helpers.ExtractPageLimit(c)
	if err != nil {
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
		}
	}

	queryParams := c.QueryParams()
	limit, offset := helpers.LimitOffsetConvert(page, limit)

	roomEntities, err := roomDelivery.roomService.GetRooms(limit, offset, queryParams)
	if err != nil {
		codeStatus, message := helpers.ValidateRoomFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	roomResponses := convertsToResponses(roomEntities)
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.ROOM_SuccesReadRoomData, roomResponses))
}

func (roomDelivery *RoomDelivery) RemoveRoom(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	roomId, err := helpers.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}

	roomRequest := rooms.RoomRequest{}
	err = c.Bind(&roomRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.ROOM_ErrorBindRoomData))
	}
	roomRequest.UserID = userId
	roomRequest.ID = roomId

	roomEntity := convertToEntity(&roomRequest)
	err = roomDelivery.roomService.RemoveRoom(&roomEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateRoomFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.ROOM_SuccesDeleteRoomData))
}

func (roomDelivery *RoomDelivery) GetRoomByRoomId(c echo.Context) error {
	roomId, err := helpers.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}

	roomRequest := rooms.RoomRequest{}
	err = c.Bind(&roomRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.ROOM_ErrorBindRoomData))
	}
	roomRequest.ID = roomId

	roomEntity := convertToEntity(&roomRequest)
	roomEntityResponse, err := roomDelivery.roomService.GetRoomByRoomId(&roomEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateRoomFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	roomResponses := convertToResponse(roomEntityResponse)
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.ROOM_SuccesReadRoomData, roomResponses))
}

func (roomDelivery *RoomDelivery) GetRoomsByUserId(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	roomRequest := rooms.RoomRequest{}
	err := c.Bind(&roomRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.ROOM_ErrorBindRoomData))
	}
	roomRequest.UserID = userId

	roomEntity := convertToEntity(&roomRequest)
	roomEntities, err := roomDelivery.roomService.GetRoomsByUserId(&roomEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateRoomFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	roomResponses := convertsToResponses(roomEntities)
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.ROOM_SuccesReadRoomData, roomResponses))
}

func (roomDelivery *RoomDelivery) AddReview(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	roomId, err := helpers.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}

	reviewRequest := reviews.ReviewRequest{}
	err = c.Bind(&reviewRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.REVIEW_ErrorBindReviewData))
	}
	reviewRequest.UserID = userId
	reviewRequest.RoomID = roomId

	reviewEntity := _reviewDelivery.ConvertToEntity(&reviewRequest)
	err = roomDelivery.roomService.CreateReview(&reviewEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateReviewFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.REVIEW_SuccessInsertReviewData))
}

func (roomDelivery *RoomDelivery) GetReviewsByRoomId(c echo.Context) error {
	roomId, err := helpers.ExtractIDParam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(err.Error()))
	}

	reviewRequest := reviews.ReviewRequest{}
	reviewRequest.RoomID = roomId

	reviewEntity := _reviewDelivery.ConvertToEntity(&reviewRequest)
	reviewEntitiesResponse, err := roomDelivery.roomService.GetReviewsByRoomId(&reviewEntity)
	if err != nil {
		codeStatus, message := helpers.ValidateRoomFailedResponse(c, err)
		return c.JSON(codeStatus, helpers.Response(message))
	}

	reviewResponses := _reviewDelivery.ConvertToResponses(reviewEntitiesResponse)
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.REVIEW_SuccesReadReviewData, reviewResponses))
}
