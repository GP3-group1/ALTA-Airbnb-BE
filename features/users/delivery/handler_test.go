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

	t.Run("Failed when create", func(t *testing.T) {
		usecase.On("Create", mock.Anything).Return(errors.New("error create")).Once()

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
			assert.Equal(t, "error create", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestGetBalance(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodGet, "/users/balances", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/balances")

		responseData := mock_data_user

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetUserBalance))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData.Balance, responseData.Balance)
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

		req := httptest.NewRequest(http.MethodGet, "/users/balances", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/balances")

		var responseData ResponseGlobal

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetUserBalance))(echoContext)
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

func TestDelete(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UserService)
	returnData := mock_data_user

	t.Run("Success", func(t *testing.T) {
		usecase.On("Remove", mock.Anything).Return(nil).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodDelete, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.RemoveAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "succesfully delete user", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when remove data", func(t *testing.T) {
		usecase.On("Remove", mock.Anything).Return(errors.New("error select data")).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodDelete, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		var responseData ResponseGlobal

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.RemoveAccount))(echoContext)
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

func TestUpdateAccount(t *testing.T) {
	reqBody, err := json.Marshal(mock_insert_user)
	if err != nil {
		t.Error(t, err, "error")
	}
	returnData := mock_data_user

	e := echo.New()
	usecase := new(mocks.UserService)

	t.Run("Success", func(t *testing.T) {
		usecase.On("ModifyData", mock.Anything, mock.Anything).Return(nil).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdateAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "succesfully update user data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed update User when bind error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		var dataFail = map[string]int{
			"name": 134,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdateAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error bind user data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed when modify data", func(t *testing.T) {
		usecase.On("ModifyData", mock.Anything, mock.Anything).Return(errors.New("error update")).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdateAccount))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error update", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestUpdatePassword(t *testing.T) {
	input := users.UserUpdatePassword{
		Password:    "thegreatest",
		NewPassword: "goat",
	}
	reqBody, err := json.Marshal(input)
	if err != nil {
		t.Error(t, err, "error")
	}
	returnData := mock_data_user

	e := echo.New()
	usecase := new(mocks.UserService)

	t.Run("Success", func(t *testing.T) {
		usecase.On("ModifyPassword", mock.Anything, mock.Anything).Return(nil).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users/password", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/password")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdatePassword))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "succesfully update user data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Update password when func error", func(t *testing.T) {
		usecase.On("ModifyPassword", mock.Anything, mock.Anything).Return(errors.New("error update")).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users/password", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/password")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdatePassword))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error update", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed update password when bind error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		var dataFail = map[string]int{
			"old_password": 1234,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users/password", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/password")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdatePassword))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error bind user data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestUpdateBalance(t *testing.T) {
	input := users.UserUpdate{
		Balance: 5000,
	}
	reqBody, err := json.Marshal(input)
	if err != nil {
		t.Error(t, err, "error")
	}
	returnData := mock_data_user

	e := echo.New()
	usecase := new(mocks.UserService)

	t.Run("Success", func(t *testing.T) {
		usecase.On("UpdateBalance", mock.Anything, mock.Anything).Return(nil).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users/balances", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/balances")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdateBalance))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "succesfully update user balance data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Update password when func error", func(t *testing.T) {
		usecase.On("UpdateBalance", mock.Anything, mock.Anything).Return(errors.New("error update")).Once()
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users/balances", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/balances")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdateBalance))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error update", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed update balance when bind error", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(returnData.ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}

		var dataFail = map[string]string{
			"amount": "asdfasd",
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPut, "/users/balances", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users/balances")

		responseData := ResponseGlobal{}

		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.UpdateBalance))(echoContext)
		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "error bind user data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	input := users.UserLogin{
		Email:    "ali@mail.com",
		Password: "thegreatest",
	}

	reqBody, err := json.Marshal(input)
	if err != nil {
		t.Error(t, err, "error")
	}

	returnData := mock_data_user
	token := "asasdasd"

	e := echo.New()
	usecase := new(mocks.UserService)

	t.Run("Success", func(t *testing.T) {
		usecase.On("Login", mock.Anything, mock.Anything).Return(returnData, token, nil).Once()

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := mock_data_user

		if assert.NoError(t, srv.Login(echoContext)) {
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

	t.Run("Failed login when bind error", func(t *testing.T) {

		var dataFail = map[string]int{
			"email": 134,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Login(echoContext)) {
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

	t.Run("Failed login when func error", func(t *testing.T) {
		usecase.On("Login", mock.Anything, mock.Anything).Return(users.UserEntity{}, "", errors.New("errror login")).Once()

		srv := New(usecase)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/login")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.Login(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "errror login", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}
