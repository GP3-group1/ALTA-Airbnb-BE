package data

import (
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/features/users/models"
	"alta-airbnb-be/utils/consts"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Delete implements users.UserDataInterface_
func (userQuery *userQuery) Delete(userID uint) error {
	panic("unimplemented")
}

// Insert implements users.UserDataInterface_
func (userQuery *userQuery) Insert(input users.UserEntity) error {
	userGorm := EntityToGorm(input)
	txInsert := userQuery.db.Create(&userGorm)
	if txInsert.Error != nil {
		return txInsert.Error
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.SERVER_ZeroRowsAffected)
	}
	return nil
}

// Login implements users.UserDataInterface_
func (userQuery *userQuery) Login(email string) (users.UserEntity, error) {
	userLogin := models.User{}
	txSelect := userQuery.db.Where("email = ?", email).First(&userLogin)
	if txSelect.Error != nil {
		return users.UserEntity{}, txSelect.Error
	}
	return GormToEntity(userLogin), nil
}

// SelectData implements users.UserDataInterface_
func (userQuery *userQuery) SelectData(userID uint) (users.UserEntity, error) {
	userGorm := models.User{}
	txSelect := userQuery.db.Where("id = ?", userID).First(&userGorm)
	if txSelect.Error != nil {
		return users.UserEntity{}, txSelect.Error
	}
	return GormToEntity(userGorm), nil
}

// UpdateData implements users.UserDataInterface_
func (userQuery *userQuery) UpdateData(input users.UserEntity) error {
	userGorm := EntityToGorm(input)
	txUpdate := userQuery.db.Updates(&userGorm)
	if txUpdate.Error != nil {
		return txUpdate.Error
	}
	if txUpdate.RowsAffected == 0 {
		return errors.New(consts.SERVER_ZeroRowsAffected)
	}
	return nil
}

func New(db *gorm.DB) users.UserDataInterface_ {
	return &userQuery{
		db: db,
	}
}
