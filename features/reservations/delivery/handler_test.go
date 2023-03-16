package delivery

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/middlewares"
	"alta-airbnb-be/mocks"
	"bytes"
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

		req := httptest.NewRequest(http.MethodGet, "/users/reservation?page=2&limit=4", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/reservationpage=2&limit=4")

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

	t.Run("Failed when page query error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users/reservations?page=asdfas", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/reservations?page=asdfas")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "invalid page parameter", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when limit query error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users/reservations?limit=asdfas", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/reservations?limit=asdfas")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "invalid limit parameter", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestCheckReservation(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.ReservationService)
	userData := mock_data_user
	input := reservations.ReservationInsert{
		CheckInDate:  "2023-03-16",
		CheckOutDate: "2023-03-07",
	}
	reqBody, err := json.Marshal(input)
	if err != nil {
		t.Error(t, err, "error")
	}

	t.Run("Room Available", func(t *testing.T) {
		usecase.On("CheckReservation", mock.Anything, mock.Anything).Return(int64(0), nil).Once()
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations/check", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations/check")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.CheckReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "room available", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Room Not Available", func(t *testing.T) {
		usecase.On("CheckReservation", mock.Anything, mock.Anything).Return(int64(1), nil).Once()
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations/check", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations/check")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.CheckReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "room not available on inputed date", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed check reservation when bind error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		var dataFail = map[string]int{
			"check_in": 134,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations/check", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations/check")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.CheckReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error bind reservation data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when param error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations/check", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations/check")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("asdasd")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.CheckReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "invaild id parameter", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when mapping to entity", func(t *testing.T) {
		input := reservations.ReservationInsert{
			CheckInDate:  "2023-3-6",
			CheckOutDate: "2023-03-07",
		}
		reqBody, err := json.Marshal(input)
		if err != nil {
			t.Error(t, err, "error")
		}

		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations/check", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations/check")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.CheckReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "invalid check in date format must YYYY-MM-DD", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when func error", func(t *testing.T) {
		usecase.On("CheckReservation", mock.Anything, mock.Anything).Return(int64(1), errors.New("error check reservation")).Once()
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations/check", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations/check")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.CheckReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error check reservation", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestAddReservation(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.ReservationService)
	userData := mock_data_user
	input := reservations.ReservationInsert{
		CheckInDate:  "2023-03-16",
		CheckOutDate: "2023-03-07",
	}
	reqBody, err := json.Marshal(input)
	if err != nil {
		t.Error(t, err, "error")
	}
	returnData := reservations.MidtransResponse{
		Token:       "",
		RedirectUrl: "",
	}

	t.Run("Success", func(t *testing.T) {
		usecase.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := reservations.MidtransResponse{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.AddReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, returnData.RedirectUrl, responseData.RedirectUrl)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when param error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("assdas")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.AddReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "invaild id parameter", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed check reservation when bind error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		var dataFail = map[string]int{
			"check_in": 134,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.AddReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error bind reservation data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when check in date mapping to entity", func(t *testing.T) {
		input := reservations.ReservationInsert{
			CheckInDate:  "2023-3-6",
			CheckOutDate: "2023-03-07",
		}
		reqBody, err := json.Marshal(input)
		if err != nil {
			t.Error(t, err, "error")
		}

		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.AddReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "invalid check in date format must YYYY-MM-DD", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when check out date mapping to entity", func(t *testing.T) {
		input := reservations.ReservationInsert{
			CheckInDate:  "2023-03-06",
			CheckOutDate: "2023-3-7",
		}
		reqBody, err := json.Marshal(input)
		if err != nil {
			t.Error(t, err, "error")
		}

		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.AddReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "invalid check out date format must YYYY-MM-DD", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when func error", func(t *testing.T) {
		usecase.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(reservations.MidtransResponse{}, errors.New("error insert")).Once()
		token, errToken := middlewares.CreateToken(userData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/rooms/:id/reservations", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/rooms/:id/reservations")
		echoContext.SetParamNames("id")
		echoContext.SetParamValues("1")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.AddReservation))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error insert", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}
