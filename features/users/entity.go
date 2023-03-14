package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type UserEntity struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	NewPassword string
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

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserUpdate struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Sex         string `json:"sex" form:"sex"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

type UserUpdatePassword struct {
	Password    string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
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
type UserService_ interface {
	Login(email string, password string) (UserEntity, string, error)
	GetData(userID uint) (UserEntity, error)
	Create(input UserEntity) error
	ModifyData(userID uint, input UserEntity) error
	ModifyPassword(userID uint, input UserEntity) error
	Remove(userID uint) error
}

//go:generate mockery --name UserData_ --output ../../mocks
type UserData_ interface {
	Login(email string) (UserEntity, error)
	Insert(input UserEntity) error
	SelectData(userID uint) (UserEntity, error)
	UpdateData(userID uint, input UserEntity) error
	Delete(userID uint) error
}

//go:generate mockery --name UserDelivery_ --output ../../mocks
type UserDelivery_ interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetUserData(c echo.Context) error
	UpdateAccount(c echo.Context) error
	UpdatePassword(c echo.Context) error
	RemoveAccount(c echo.Context) error
}
