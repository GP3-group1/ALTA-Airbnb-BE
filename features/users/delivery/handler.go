package delivery

import (
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/middlewares"
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
	userID := middlewares.ExtractTokenUserId(c)
	userEntity, errSelect := userHandler.userService.GetData(userID)
	if errSelect != nil {
		return c.JSON(helpers.ErrorResponse(errSelect))
	}
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.USER_SuccessReadUserData, entityToResponse(userEntity)))
}

// Login implements users.UserDeliveryInterface_
func (userHandler *UserHandler) Login(c echo.Context) error {
	loginInput := users.UserLogin{}
	err := c.Bind(&loginInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.USER_ErrorBindUserData))
	}
	userEntity, token, errLogin := userHandler.userService.Login(loginInput.Email, loginInput.Password)
	if errLogin != nil {
		return c.JSON(helpers.ErrorResponse(errLogin))
	}
	dataResponse := map[string]any{
		"id":    userEntity.ID,
		"name":  userEntity.Name,
		"tokne": token,
	}
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.USER_LoginSuccess, dataResponse))
}

// Register implements users.UserDeliveryInterface_
func (userHandler *UserHandler) Register(c echo.Context) error {
	userInput := users.UserRegister{}
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.USER_ErrorBindUserData))
	}
	errInsert := userHandler.userService.Create(registerToEntity(userInput))
	if errInsert != nil {
		return c.JSON(helpers.ErrorResponse(errInsert))
	}

	return c.JSON(http.StatusCreated, helpers.Response(consts.USER_RegisterSuccess))
}

// RemoveAccount implements users.UserDeliveryInterface_
func (userHandler *UserHandler) RemoveAccount(c echo.Context) error {
	panic("unimplemented")
}

// UpdateAccount implements users.UserDeliveryInterface_
func (userHandler *UserHandler) UpdateAccount(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	userInput := users.UserUpdate{}
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.USER_ErrorBindUserData))
	}
	errUpdate := userHandler.userService.ModifyData(userID, requestToEntity(userInput))
	if errUpdate != nil {
		return c.JSON(helpers.ErrorResponse(errUpdate))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.USER_SuccessUpdateUserData))
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
