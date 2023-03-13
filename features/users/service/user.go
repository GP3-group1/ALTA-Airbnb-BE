package service

import (
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/middlewares"
	"alta-airbnb-be/utils/consts"
	"alta-airbnb-be/utils/helpers"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData users.UserDataInterface_
	validate *validator.Validate
}

// Create implements users.UserServiceInterface_
func (userService *userService) Create(input users.UserEntity) error {
	errValidate := userService.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	hashedPassword, errHash := helpers.HashPassword(input.Password)
	if errHash != nil {
		return errHash
	}
	input.Password = hashedPassword

	errInsert := userService.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

// GetData implements users.UserServiceInterface_
func (userService *userService) GetData(userID uint) (users.UserEntity, error) {
	panic("unimplemented")
}

// Login implements users.UserServiceInterface_
func (userService *userService) Login(email string, password string) (users.UserEntity, string, error) {
	if email == "" || password == "" {
		return users.UserEntity{}, "", errors.New(consts.USER_EmptyCredentialError)
	}

	userEntity, errLogin := userService.userData.Login(email)
	if errLogin != nil {
		return users.UserEntity{}, "", errLogin
	}

	if !helpers.CompareHashPassword(password, userEntity.Password) {
		return users.UserEntity{}, "", errors.New(consts.USER_WrongPassword)
	}

	token, errToken := middlewares.CreateToken(userEntity.ID)
	if errToken != nil {
		return users.UserEntity{}, "", errToken
	}

	return userEntity, token, nil
}

// ModifyData implements users.UserServiceInterface_
func (userService *userService) ModifyData(userID uint, input users.UserEntity) error {
	panic("unimplemented")
}

// ModifyPassword implements users.UserServiceInterface_
func (userService *userService) ModifyPassword(userID uint, input users.UserEntity) error {
	panic("unimplemented")
}

// Remove implements users.UserServiceInterface_
func (userService *userService) Remove(userID uint) error {
	panic("unimplemented")
}

func New(userData users.UserDataInterface_) users.UserServiceInterface_ {
	return &userService{
		userData: userData,
		validate: validator.New(),
	}
}
