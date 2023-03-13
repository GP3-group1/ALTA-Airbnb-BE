package router

import (
	_reservationData "alta-airbnb-be/features/reservations/data"
	_reservationDelivery "alta-airbnb-be/features/reservations/delivery"
	_reservationService "alta-airbnb-be/features/reservations/service"
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
	e.PUT("/users/password", userHandler.UpdatePassword, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandler.RemoveAccount, middlewares.JWTMiddleware())
}

func initReservationRouter(db *gorm.DB, e *echo.Echo) {
	reservationData := _reservationData.New(db)
	reservationService := _reservationService.New(reservationData)
	reservationHandler := _reservationDelivery.New(reservationService)

	e.POST("/rooms/:id/reservations", reservationHandler.AddReservation, middlewares.JWTMiddleware())
	e.GET("/users/reservations", reservationHandler.GetAllReservation, middlewares.JWTMiddleware())
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
	initReservationRouter(db, e)
}
