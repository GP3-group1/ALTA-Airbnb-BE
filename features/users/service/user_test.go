package service

import (
	"alta-airbnb-be/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
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
		repo.On("Delete", id).Return(errors.New("error delete data user")).Once()

		srv := New(repo)
		err := srv.Remove(id)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data user", err.Error())
		repo.AssertExpectations(t)
	})
}
