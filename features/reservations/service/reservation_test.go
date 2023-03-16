package service

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_room = reservations.RoomEntity{
		ID:    1,
		Name:  "Villa Indah Biru",
		Price: 100,
	}
)

func TestCheckReservation(t *testing.T) {
	repo := new(mocks.ReservationData)
	id := mock_room.ID
	returnRow := int64(1)

	t.Run("Success", func(t *testing.T) {
		input := reservations.ReservationEntity{
			CheckInDate:  time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
		}
		repo.On("CheckReservation", mock.Anything, mock.Anything).Return(returnRow, nil).Once()

		srv := New(repo)
		row, err := srv.CheckReservation(input, id)
		assert.Nil(t, err)
		assert.Equal(t, returnRow, row)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		input := reservations.ReservationEntity{
			CheckInDate:  time.Time{},
			CheckOutDate: time.Time{},
		}
		srv := New(repo)
		row, err := srv.CheckReservation(input, id)
		assert.NotNil(t, err)
		assert.Equal(t, int64(0), row)
		repo.AssertExpectations(t)
	})

}
