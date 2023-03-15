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
	userData users.UserData_
	validate *validator.Validate
}

// UpdateBalance implements users.UserService_
func (userService *userService) UpdateBalance(userID uint, input users.UserEntity) error {
	if input.Balance < 1 {
		return errors.New(consts.USER_InvalidInput)
	}

	// get user balance
	userEntity, errSelect := userService.userData.SelectData(userID)
	if errSelect != nil {
		return errSelect
	}

	input.Balance += userEntity.Balance
	errUpdate := userService.userData.UpdateData(userID, input)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
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
	input.Balance = 0

	errInsert := userService.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

// GetData implements users.UserServiceInterface_
func (userService *userService) GetData(userID uint) (users.UserEntity, error) {
	userEntity, errSelect := userService.userData.SelectData(userID)
	if errSelect != nil {
		return users.UserEntity{}, errSelect
	}
	return userEntity, nil
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
	errValidate := userService.validate.StructExcept(input, "Password")
	if errValidate != nil {
		return errValidate
	}

	errUpdate := userService.userData.UpdateData(userID, input)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

// ModifyPassword implements users.UserServiceInterface_
func (userService *userService) ModifyPassword(userID uint, input users.UserEntity) error {
	if input.Password == "" || input.NewPassword == "" {
		return errors.New(consts.USER_EmptyUpdatePasswordError)
	}

	userEntity, errSelect := userService.userData.SelectData(userID)
	if errSelect != nil {
		return errSelect
	}

	if !helpers.CompareHashPassword(input.Password, userEntity.Password) {
		return errors.New(consts.USER_WrongPassword)
	}

	hashedPassword, errHash := helpers.HashPassword(input.NewPassword)
	if errHash != nil {
		return errHash
	}

	input.Password = hashedPassword

	errUpdatePassword := userService.userData.UpdateData(userID, input)
	if errUpdatePassword != nil {
		return errUpdatePassword
	}
	return nil
}

// Remove implements users.UserServiceInterface_
func (userService *userService) Remove(userID uint) error {
	errDelete := userService.userData.Delete(userID)

	if errDelete != nil {
		return errDelete
	}

	return nil
}

func New(userData users.UserData_) users.UserService_ {
	return &userService{
		userData: userData,
		validate: validator.New(),
	}
}
