package service

import (
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_insert_user = users.UserEntity{
		ID:          1,
		Name:        "Muhammad Ali",
		Email:       "ali@mail.com",
		Password:    "thegreatest",
		NewPassword: "",
		Sex:         "",
		Address:     "",
		PhoneNumber: "",
		Balance:     0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock_data_user = users.UserEntity{
		ID:          1,
		Name:        "Muhammad Ali",
		Email:       "ali@mail.com",
		Password:    "$2a$14$J4AF0twBNp2Pxx/LY5McIu/0v0FEvP.T7TOU/ozo.afMSDA03aBZ6",
		NewPassword: "",
		Sex:         "",
		Address:     "",
		PhoneNumber: "",
		Balance:     0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
)

func TestDelete(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Delete", func(t *testing.T) {
		id := uint(1)
		repo.On("Delete", id).Return(nil).Once()

		srv := New(repo)
		err := srv.Remove(id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Delete return error", func(t *testing.T) {
		id := uint(1)
		repo.On("Delete", mock.Anything).Return(errors.New("error delete data")).Once()

		srv := New(repo)
		err := srv.Remove(id)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.Create(mock_insert_user)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := users.UserEntity{
			Name:  "Muhammad Ali",
			Email: "ali@mail.com",
		}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Insert return error", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.Create(mock_insert_user)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := new(mocks.UserData)
	returnCore := mock_data_user
	email := mock_data_user.Email
	password := "thegreatest"

	t.Run("Success", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(returnCore, nil).Once()

		srv := New(repo)
		core, token, err := srv.Login(email, password)
		assert.Nil(t, err)
		assert.Equal(t, returnCore.Name, core.Name)
		assert.Equal(t, token, token)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when select user by email return error", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(users.UserEntity{}, errors.New("error login")).Once()

		srv := New(repo)
		core, token, err := srv.Login("dummy@mail.com", password)
		assert.NotNil(t, err)
		assert.Equal(t, users.UserEntity{}.Name, core.Name)
		assert.Equal(t, "", token)
		repo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(returnCore, nil).Once()

		srv := New(repo)
		core, token, err := srv.Login(email, "asdasd")
		assert.NotNil(t, err)
		assert.Equal(t, "", core.Name)
		assert.Equal(t, "", token)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		input := users.UserEntity{
			Email:    "",
			Password: "",
		}
		srv := New(repo)
		core, token, err := srv.Login(input.Email, input.Password)
		assert.NotNil(t, err)
		assert.NotNil(t, err)
		assert.Equal(t, users.UserEntity{}.Name, core.Name)
		assert.Equal(t, "", token)
		repo.AssertExpectations(t)
	})
}

func TestSelectData(t *testing.T) {
	repo := new(mocks.UserData)
	returnCore := mock_data_user
	id := mock_data_user.ID

	t.Run("Success", func(t *testing.T) {
		repo.On("SelectData", mock.Anything).Return(returnCore, nil).Once()

		srv := New(repo)
		core, err := srv.GetData(id)
		assert.Nil(t, err)
		assert.Equal(t, returnCore.Name, core.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("SelectData", mock.Anything).Return(users.UserEntity{}, errors.New("error select")).Once()

		srv := New(repo)
		core, err := srv.GetData(id)
		assert.NotNil(t, err)
		assert.Equal(t, users.UserEntity{}.Name, core.Name)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	repo := new(mocks.UserData)
	core := mock_data_user
	id := mock_data_user.ID

	t.Run("Success", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.ModifyData(id, core)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := users.UserEntity{
			Name: "Muhammad Ali",
		}
		srv := New(repo)
		err := srv.ModifyData(id, inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Update return error", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(errors.New("error update data")).Once()

		srv := New(repo)
		err := srv.ModifyData(id, core)
		assert.NotNil(t, err)
		assert.Equal(t, "error update data", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpdatePassword(t *testing.T) {
	repo := new(mocks.UserData)
	core := mock_data_user
	id := mock_data_user.ID

	t.Run("Success", func(t *testing.T) {
		input := users.UserEntity{
			Password:    "thegreatest",
			NewPassword: "goat",
		}
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("SelectData", mock.Anything).Return(core, nil).Once()

		srv := New(repo)
		err := srv.ModifyPassword(id, input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when select data", func(t *testing.T) {
		input := users.UserEntity{
			Password:    "thegreatest",
			NewPassword: "goat",
		}
		repo.On("SelectData", mock.Anything).Return(users.UserEntity{}, errors.New("error select")).Once()

		srv := New(repo)
		err := srv.ModifyPassword(id, input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		input := users.UserEntity{
			Password:    "asdasd",
			NewPassword: "goat",
		}
		repo.On("SelectData", mock.Anything).Return(core, nil).Once()

		srv := New(repo)
		err := srv.ModifyPassword(id, input)
		assert.NotNil(t, err)
		assert.Equal(t, "wrong password", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		input := users.UserEntity{
			Password:    "",
			NewPassword: "",
		}
		srv := New(repo)
		err := srv.ModifyPassword(id, input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdateBalance(t *testing.T) {
	repo := new(mocks.UserData)
	core := mock_data_user
	id := mock_data_user.ID

	t.Run("Success", func(t *testing.T) {
		input := users.UserEntity{
			Balance: 5000,
		}
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("SelectData", mock.Anything).Return(core, nil).Once()

		srv := New(repo)
		err := srv.UpdateBalance(id, input)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Error when select data", func(t *testing.T) {
		input := users.UserEntity{
			Balance: 5000,
		}
		repo.On("SelectData", mock.Anything).Return(users.UserEntity{}, errors.New("error select")).Once()

		srv := New(repo)
		err := srv.UpdateBalance(id, input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Success", func(t *testing.T) {
		input := users.UserEntity{
			Balance: 5000,
		}
		repo.On("SelectData", mock.Anything).Return(core, nil).Once()
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(errors.New("error update")).Once()

		srv := New(repo)
		err := srv.UpdateBalance(id, input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		input := users.UserEntity{
			Balance: -1,
		}
		srv := New(repo)
		err := srv.UpdateBalance(id, input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
