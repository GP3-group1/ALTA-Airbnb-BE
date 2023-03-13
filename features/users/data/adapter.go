package data

import (
	"alta-airbnb-be/features/users"
	_userModel "alta-airbnb-be/features/users/models"
)

func EntityToGorm(userEntity users.UserEntity) _userModel.User {
	userGorm := _userModel.User{
		Name:        userEntity.Name,
		Email:       userEntity.Email,
		Password:    userEntity.Password,
		Sex:         userEntity.Sex,
		Address:     userEntity.Address,
		PhoneNumber: userEntity.PhoneNumber,
		Balance:     userEntity.Balance,
	}
	return userGorm
}

func GormToEntity(userGorm _userModel.User) users.UserEntity {
	return users.UserEntity{
		ID:          userGorm.ID,
		Name:        userGorm.Name,
		Email:       userGorm.Email,
		Password:    userGorm.Password,
		Sex:         userGorm.Sex,
		Address:     userGorm.Address,
		PhoneNumber: userGorm.PhoneNumber,
		Balance:     userGorm.Balance,
		CreatedAt:   userGorm.CreatedAt,
		UpdatedAt:   userGorm.UpdatedAt,
	}
}
