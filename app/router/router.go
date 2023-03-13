package router

import (
	_userData "alta-airbnb-be/features/users/data"
	_userDelivery "alta-airbnb-be/features/users/delivery"
	_userService "alta-airbnb-be/features/users/service"
	"alta-airbnb-be/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userDelivery.New(userService)

	e.POST("/login", userHandler.Login)
	e.POST("/users", userHandler.Register)
	e.GET("/users", userHandler.GetUserData, middlewares.JWTMiddleware())
	e.PUT("/users", userHandler.UpdateAccount, middlewares.JWTMiddleware())
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
}
