package router

import (
	_userData "alta-airbnb-be/features/users/data"
	_userDelivery "alta-airbnb-be/features/users/delivery"
	_userService "alta-airbnb-be/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userDelivery.New(userService)

	e.POST("/login", userHandler.Login)
	e.POST("/users", userHandler.Register)
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
}
