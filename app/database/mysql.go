package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"alta-airbnb-be/app/config"
	_imageModel "alta-airbnb-be/features/images/models"
	_ratingModel "alta-airbnb-be/features/ratings/models"
	_reservationModel "alta-airbnb-be/features/reservations/models"
	_roomModel "alta-airbnb-be/features/rooms/models"
	_userModel "alta-airbnb-be/features/users/models"
)

func InitDB(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to DB", err.Error())
		return nil
	}

	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(_userModel.User{}, _roomModel.Room{}, _reservationModel.Reservation{}, _imageModel.Image{}, _ratingModel.Rating{})
}
