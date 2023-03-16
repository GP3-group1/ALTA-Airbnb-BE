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
	mock_data_user = users.UserEntity{
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
		repo.On("Delete", mock.Anything).Return(errors.New("error delete data user")).Once()

		srv := New(repo)
		err := srv.Remove(id)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data user", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.Create(mock_data_user)
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
}
