package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type UserEntity struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required, email"`
	Password    string `validate:"required"`
	Sex         string
	Address     string
	PhoneNumber string
	Balance     float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserRegister struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Sex         string `json:"sex" form:"sex"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

type UserResponse struct {
	ID          uint
	Name        string `json:"name"`
	Email       string `json:"email"`
	Sex         string `json:"sex"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

//go:generate mockery --name UserService_ --output ../../mocks
type UserServiceInterface_ interface {
	Login(email string, password string) (UserEntity, string, error)
	GetData(userID uint) (UserEntity, error)
	Create(input UserEntity) error
	ModifyData(userID uint, input UserEntity) error
	ModifyPassword(userID uint, input UserEntity) error
	Remove(userID uint) error
}

//go:generate mockery --name UserData_ --output ../../mocks
type UserDataInterface_ interface {
	Login(email string, password string) (UserEntity, string, error)
	Insert(input UserEntity) error
	SelectData(userID uint) (UserEntity, error)
	UpdateData(input UserEntity) error
	Delete(userID uint) error
}

//go:generate mockery --name UserDelivery_ --output ../../mocks
type UserDeliveryInterface_ interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetUserData(c echo.Context) error
	UpdateAccount(c echo.Context) error
	UpdatePassword(c echo.Context) error
	RemoveAccount(c echo.Context) error
}
