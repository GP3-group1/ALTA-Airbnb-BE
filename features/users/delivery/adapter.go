package delivery

import "alta-airbnb-be/features/users"

func registerToEntity(userRegister users.UserRegister) users.UserEntity {
	return users.UserEntity{
		Name:     userRegister.Name,
		Email:    userRegister.Email,
		Password: userRegister.Password,
	}
}

func requestUpdateToEntity(userUpdate users.UserUpdate) users.UserEntity {
	return users.UserEntity{
		Name:        userUpdate.Name,
		Email:       userUpdate.Email,
		Sex:         userUpdate.Sex,
		Address:     userUpdate.Address,
		PhoneNumber: userUpdate.PhoneNumber,
	}
}

func requestUpdatePasswordToEntity(UpdatePassword users.UserUpdatePassword) users.UserEntity {
	return users.UserEntity{
		Password:    UpdatePassword.Password,
		NewPassword: UpdatePassword.NewPassword,
	}
}

func entityToResponse(userEntity users.UserEntity) users.UserResponse {
	return users.UserResponse{
		ID:          userEntity.ID,
		Name:        userEntity.Name,
		Email:       userEntity.Email,
		Sex:         userEntity.Sex,
		Address:     userEntity.Address,
		PhoneNumber: userEntity.PhoneNumber,
		Balance:     userEntity.Balance,
	}
}
