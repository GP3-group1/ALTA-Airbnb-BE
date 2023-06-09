package delivery

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/middlewares"
	"alta-airbnb-be/utils/consts"
	"alta-airbnb-be/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	reservationService reservations.ReservationService_
}

// AddReservation implements reservations.ReservationDeliveryInterface_
func (reservationHandler *ReservationHandler) AddReservation(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	idParam, errParam := helpers.ExtractIDParam(c)
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(errParam.Error()))
	}

	input := reservations.ReservationInsert{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.RESERVATION_ErrorBindReservationData))
	}

	reservationEntity, errMapping := insertToEntity(input)
	if errMapping != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(errMapping.Error()))
	}

	midTransResponse, errInsert := reservationHandler.reservationService.Create(userID, uint(idParam), reservationEntity)
	if errInsert != nil {
		return c.JSON(helpers.ErrorResponse(errInsert))
	}

	dataResponse := map[string]any{
		"token":        midTransResponse.Token,
		"redirect_url": midTransResponse.RedirectUrl,
	}
	return c.JSON(http.StatusCreated, helpers.ResponseWithData(consts.RESERVATION_InsertSuccess, dataResponse))
}

// CheckReservation implements reservations.ReservationDeliveryInterface_
func (reservationHandler *ReservationHandler) CheckReservation(c echo.Context) error {
	idParam, errParam := helpers.ExtractIDParam(c)
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(errParam.Error()))
	}

	input := reservations.ReservationInsert{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.RESERVATION_ErrorBindReservationData))
	}

	inputReservationEntity, errMapping := insertToEntity(input)
	if errMapping != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(errMapping.Error()))
	}

	row, errSelect := reservationHandler.reservationService.CheckReservation(inputReservationEntity, idParam)
	if errSelect != nil {
		return c.JSON(helpers.ErrorResponse(errSelect))
	}

	if row == 0 {
		return c.JSON(http.StatusOK, helpers.Response(consts.RESERVATION_RoomAvailable))
	}

	return c.JSON(http.StatusBadRequest, helpers.Response(consts.RESERVATION_RoomNotAvailable))
}

// GetAllReservation implements reservations.ReservationDeliveryInterface_
func (reservationHandler *ReservationHandler) GetAllReservation(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)

	var page int = 1
	pageParam := c.QueryParam("page")
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(consts.ECHO_InvaildPageParam))
		} else {
			page = pageConv
		}
	}

	var limit int = 8
	limitParam := c.QueryParam("limit")
	if limitParam != "" {
		limitConv, errConv := strconv.Atoi(limitParam)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(consts.ECHO_InvaildLimitParam))
		} else {
			limit = limitConv
		}
	}

	reservationEntity, errSelect := reservationHandler.reservationService.GetAll(page, limit, userID)
	if errSelect != nil {
		return c.JSON(helpers.ErrorResponse(errSelect))
	}
	return c.JSON(http.StatusOK, helpers.ResponseWithData("Success", entityToResponseList(reservationEntity)))
}

func New(reservationService reservations.ReservationService_) reservations.ReservationDelivery_ {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}
