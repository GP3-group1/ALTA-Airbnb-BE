package router

import (
	_reservationData "alta-airbnb-be/features/reservations/data"
	_reservationDelivery "alta-airbnb-be/features/reservations/delivery"
	_reservationService "alta-airbnb-be/features/reservations/service"
	_roomData "alta-airbnb-be/features/rooms/data"
	_roomDelivery "alta-airbnb-be/features/rooms/delivery"
	_roomService "alta-airbnb-be/features/rooms/service"
	_userData "alta-airbnb-be/features/users/data"
	_userDelivery "alta-airbnb-be/features/users/delivery"
	_userService "alta-airbnb-be/features/users/service"
	_imageData "alta-airbnb-be/features/images/data"
	_imageDelivery "alta-airbnb-be/features/images/delivery"
	_imageService "alta-airbnb-be/features/images/service"
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
	e.GET("/users/balances", userHandler.GetUserBalance, middlewares.JWTMiddleware())
	e.PUT("/users/balances", userHandler.UpdateBalance, middlewares.JWTMiddleware())
}

func initReservationRouter(db *gorm.DB, e *echo.Echo) {
	reservationData := _reservationData.New(db)
	reservationService := _reservationService.New(reservationData)
	reservationHandler := _reservationDelivery.New(reservationService)

	e.POST("/rooms/:id/reservations/check", reservationHandler.CheckReservation, middlewares.JWTMiddleware())
	e.POST("/rooms/:id/reservations", reservationHandler.AddReservation, middlewares.JWTMiddleware())
	e.GET("/users/reservations", reservationHandler.GetAllReservation, middlewares.JWTMiddleware())
}

func initRoomRouter(db *gorm.DB, e *echo.Echo) {
	roomData := _roomData.New(db)
	roomService := _roomService.New(roomData)
	roomHandler := _roomDelivery.New(roomService)

	e.POST("/rooms", roomHandler.AddRoom, middlewares.JWTMiddleware())
	e.PUT("/rooms/:id", roomHandler.ModifyRoom, middlewares.JWTMiddleware())
	e.DELETE("/rooms/:id", roomHandler.RemoveRoom, middlewares.JWTMiddleware())
	e.GET("/rooms", roomHandler.GetRooms)
	e.GET("/rooms/users", roomHandler.GetRoomsByUserId, middlewares.JWTMiddleware())
	e.GET("/rooms/:id", roomHandler.GetRoomByRoomId)
	e.POST("/rooms/:id/reviews", roomHandler.AddReview, middlewares.JWTMiddleware())
	e.GET("/rooms/:id/reviews", roomHandler.GetReviewsByRoomId)
}

func initImageRouter(db *gorm.DB, e *echo.Echo) {
	imageData := _imageData.New(db)
	roomData := _roomData.New(db)
	imageService := _imageService.New(imageData, roomData)
	imageHandler := _imageDelivery.New(imageService)

	e.POST("/images", imageHandler.AddImage, middlewares.JWTMiddleware())
	e.PUT("/images/:id", imageHandler.ModifyImage, middlewares.JWTMiddleware())
	e.DELETE("/images/:id", imageHandler.RemoveImage, middlewares.JWTMiddleware())
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	initUserRouter(db, e)
	initReservationRouter(db, e)
	initRoomRouter(db, e)
	initImageRouter(db, e)
}
