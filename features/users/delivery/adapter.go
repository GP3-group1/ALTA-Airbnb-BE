package delivery

import "alta-airbnb-be/features/users"

func registerToEntity(userRegister users.UserRegister) users.UserEntity {
	return users.UserEntity{
		Name:     userRegister.Name,
		Email:    userRegister.Email,
		Password: userRegister.Password,
	}
}

func requestToEntity(userRequest users.UserRequest) users.UserEntity {
	return users.UserEntity{
		Name:        userRequest.Name,
		Email:       userRequest.Email,
		Password:    userRequest.Password,
		Sex:         userRequest.Sex,
		Address:     userRequest.Address,
		PhoneNumber: userRequest.PhoneNumber,
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
	}
}
