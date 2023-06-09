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
	userService users.UserService_
}

// UpdateBalance implements users.UserDelivery_
func (userHandler *UserHandler) UpdateBalance(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	userInput := users.UserUpdate{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.USER_ErrorBindUserData))
	}
	errUpdate := userHandler.userService.UpdateBalance(userID, requestUpdateToEntity(userInput))
	if errUpdate != nil {
		return c.JSON(helpers.ErrorResponse(errUpdate))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.USER_SuccessUpdateBalance))
}

// GetUserBalance implements users.UserDelivery_
func (userHandler *UserHandler) GetUserBalance(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	userEntity, errSelect := userHandler.userService.GetData(userID)
	if errSelect != nil {
		return c.JSON(helpers.ErrorResponse(errSelect))
	}
	response := entityToResponse(userEntity)
	dataResponse := map[string]any{
		"id":      response.ID,
		"balance": response.Balance,
	}
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.USER_SuccessReadBalance, dataResponse))
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
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.USER_ErrorBindUserData))
	}
	userEntity, token, errLogin := userHandler.userService.Login(loginInput.Email, loginInput.Password)
	if errLogin != nil {
		return c.JSON(helpers.ErrorResponse(errLogin))
	}
	dataResponse := map[string]any{
		"id":    userEntity.ID,
		"name":  userEntity.Name,
		"token": token,
	}
	return c.JSON(http.StatusOK, helpers.ResponseWithData(consts.USER_LoginSuccess, dataResponse))
}

// Register implements users.UserDeliveryInterface_
func (userHandler *UserHandler) Register(c echo.Context) error {
	userInput := users.UserRegister{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
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
	userID := middlewares.ExtractTokenUserId(c)
	errDelete := userHandler.userService.Remove(userID)
	if errDelete != nil {
		return c.JSON(helpers.ErrorResponse(errDelete))
	}
	return c.JSON(http.StatusOK, helpers.Response(consts.USER_SuccessDelete))
}

// UpdateAccount implements users.UserDeliveryInterface_
func (userHandler *UserHandler) UpdateAccount(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	userInput := users.UserUpdate{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.USER_ErrorBindUserData))
	}
	errUpdate := userHandler.userService.ModifyData(userID, requestUpdateToEntity(userInput))
	if errUpdate != nil {
		return c.JSON(helpers.ErrorResponse(errUpdate))
	}

	return c.JSON(http.StatusOK, helpers.Response(consts.USER_SuccessUpdateUserData))
}

// UpdatePassword implements users.UserDeliveryInterface_
func (userHandler *UserHandler) UpdatePassword(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	input := users.UserUpdatePassword{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(consts.USER_ErrorBindUserData))
	}
	errUpdatePassword := userHandler.userService.ModifyPassword(userID, requestUpdatePasswordToEntity(input))
	if errUpdatePassword != nil {
		return c.JSON(helpers.ErrorResponse(errUpdatePassword))
	}
	return c.JSON(http.StatusOK, helpers.Response(consts.USER_SuccessUpdateUserData))
}

func New(userService users.UserService_) users.UserDelivery_ {
	return &UserHandler{
		userService: userService,
	}
}
