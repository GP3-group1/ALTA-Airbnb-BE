package delivery

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/middlewares"
	"alta-airbnb-be/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_data_user = users.UserEntity{
		ID:          1,
		Name:        "Muhammad Ali",
		Email:       "ali@mail.com",
		Sex:         "Male",
		Address:     "Lewo",
		PhoneNumber: "08123456789",
		Balance:     5000,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
)

type ResponseGlobal struct {
	Message string
}

func TestGetAll(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.ReservationService)
	returnData := []reservations.ReservationEntity{
		{
			ID:     0,
			RoomID: 0,
			Room: reservations.RoomEntity{
				Name:  "",
				Price: 0,
			},
			CheckInDate:  time.Time{},
			CheckOutDate: time.Time{},
			TotalNight:   0,
			TotalPrice:   0,
		},
	}
	userData := mock_data_user

	t.Run("Success", func(t *testing.T) {
		usecase.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users/reservation", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/reservation")

		responseData := reservations.ReservationEntity{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0], responseData)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when func error", func(t *testing.T) {
		usecase.On("GetAll", mock.Anything, mock.Anything, mock.Anything).Return([]reservations.ReservationEntity{}, errors.New("error select")).Once()
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users/reservation", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/reservation")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error select", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}
