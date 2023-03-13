package delivery

import (
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/utils/consts"
	"alta-airbnb-be/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface_
}

// GetUserData implements users.UserDeliveryInterface_
func (userHandler *UserHandler) GetUserData(c echo.Context) error {
	panic("unimplemented")
}

// Login implements users.UserDeliveryInterface_
func (userHandler *UserHandler) Login(c echo.Context) error {
	panic("unimplemented")
}

// Register implements users.UserDeliveryInterface_
func (userHandler *UserHandler) Register(c echo.Context) error {
	userInput := users.UserRegister{}
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response(consts.USER_ErrorBindUserData))
	}
	userEntity := registerToEntity(userInput)
	errInsert := userHandler.userService.Create(userEntity)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response("error: "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.Response(consts.USER_RegisterSuccess))
}

// RemoveAccount implements users.UserDeliveryInterface_
func (userHandler *UserHandler) RemoveAccount(c echo.Context) error {
	panic("unimplemented")
}

// UpdateAccount implements users.UserDeliveryInterface_
func (userHandler *UserHandler) UpdateAccount(c echo.Context) error {
	panic("unimplemented")
}

// UpdatePassword implements users.UserDeliveryInterface_
func (userHandler *UserHandler) UpdatePassword(c echo.Context) error {
	panic("unimplemented")
}

func New(userService users.UserServiceInterface_) users.UserDeliveryInterface_ {
	return &UserHandler{
		userService: userService,
	}
}
