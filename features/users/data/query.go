package data

import (
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/features/users/models"
	"alta-airbnb-be/utils/consts"
	"errors"
	"strings"

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
		if strings.Contains(txInsert.Error.Error(), "Error 1062 (23000)") {
			return errors.New(consts.USER_EmailAlreadyUsed)
		}
		return errors.New(consts.SERVER_InternalServerError)
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.SERVER_ZeroRowsAffected)
	}
	return nil
}

// Login implements users.UserDataInterface_
func (userQuery *userQuery) Login(email string, password string) (users.UserEntity, error) {
	userLogin := models.User{}
	txSelect := userQuery.db.Where("email = ?", email).First(&userLogin)
	if txSelect.Error != nil {
		if txSelect.Error == gorm.ErrRecordNotFound {
			return users.UserEntity{}, errors.New(gorm.ErrRecordNotFound.Error())
		}
		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
	}
	userEntity := GormToEntity(userLogin)
	return userEntity, nil
}

// SelectData implements users.UserDataInterface_
func (userQuery *userQuery) SelectData(userID uint) (users.UserEntity, error) {
	panic("unimplemented")
}

// UpdateData implements users.UserDataInterface_
func (userQuery *userQuery) UpdateData(input users.UserEntity) error {
	panic("unimplemented")
}

func New(db *gorm.DB) users.UserDataInterface_ {
	return &userQuery{
		db: db,
	}
}
