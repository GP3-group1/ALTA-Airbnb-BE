package delivery

import (
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

	mock_insert_user = users.UserEntity{
		Name:     "Muhammad Ali",
		Email:    "ali@mail.com",
		Password: "thegreatest",
	}
)

type ResponseGlobal struct {
	Message string
}

func TestGetUser(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UserService)
	returnData := mock_data_user

	t.Run("Success", func(t *testing.T) {
		usecase.On("GetData", mock.Anything).Return(returnData, nil).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := mock_data_user

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetUserData))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData.Name, responseData.Name)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when select data", func(t *testing.T) {
		usecase.On("GetData", mock.Anything).Return(users.UserEntity{}, errors.New("error select data")).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		var responseData ResponseGlobal

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetUserData))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error select data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	reqBody, err := json.Marshal(mock_insert_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	e := echo.New()
	usecase := new(mocks.UserService)

	t.Run("Success", func(t *testing.T) {
		usecase.On("Create", mock.Anything).Return(nil).Once()

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Register(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "succesfully insert user data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Add User when bind error", func(t *testing.T) {

		var dataFail = map[string]int{
			"name": 134,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Register(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "error bind user data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}
