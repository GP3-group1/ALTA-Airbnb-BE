package service

import (
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/mocks"
	"errors"
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

	t.Run("Failed when check reservation error", func(t *testing.T) {
		input := reservations.ReservationEntity{
			CheckInDate:  time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
		}
		repo.On("CheckReservation", mock.Anything, mock.Anything).Return(int64(0), errors.New("error check reservation")).Once()

		srv := New(repo)
		row, err := srv.CheckReservation(input, id)
		assert.NotNil(t, err)
		assert.Equal(t, int64(0), row)
		assert.Equal(t, "error check reservation", err.Error())
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

	t.Run("Failed validate date", func(t *testing.T) {
		input := reservations.ReservationEntity{
			CheckInDate:  time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
		}
		srv := New(repo)
		row, err := srv.CheckReservation(input, id)
		assert.NotNil(t, err)
		assert.Equal(t, int64(0), row)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.ReservationData)
	returnData := []reservations.ReservationEntity{
		{
			ID:     1,
			RoomID: 1,
			Room: reservations.RoomEntity{
				Name:  "Villa Biru Laut",
				Price: 200,
			},
			CheckInDate:  time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
			TotalNight:   1,
			TotalPrice:   200,
		},
	}
	page := 1
	limit := 10
	userID := 1

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything, mock.Anything, mock.Anything).Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetAll(page, limit, uint(userID))
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].RoomID, response[0].RoomID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything, mock.Anything, mock.Anything).Return([]reservations.ReservationEntity{}, errors.New("error select")).Once()

		srv := New(repo)
		response, err := srv.GetAll(page, limit, uint(userID))
		assert.NotNil(t, err)
		assert.Equal(t, []reservations.ReservationEntity{}, response)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.ReservationData)
	returnDataMidtrans := reservations.MidtransResponse{
		Token:       "",
		RedirectUrl: "",
	}

	selectUser := reservations.ReservationEntity{
		User: reservations.UserEntity{
			ID:          1,
			Name:        "Muhammad Ali",
			Email:       "ali@mail.com",
			Address:     "Lewo",
			PhoneNumber: "08123456789",
			Balance:     1000000,
		},
	}

	selectRoom := reservations.ReservationEntity{
		Room: reservations.RoomEntity{
			ID:    1,
			Name:  "Villa Indah Hijau",
			Price: 200,
		},
	}

	userID := uint(1)
	roomID := uint(1)

	t.Run("Success", func(t *testing.T) {
		input := reservations.ReservationEntity{
			UserID:       1,
			RoomID:       1,
			CheckInDate:  time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
			TotalNight:   1,
			TotalPrice:   200,
		}

		repo.On("SelectRoom", mock.Anything).Return(selectRoom, nil).Once()

		repo.On("SelectUser", mock.Anything).Return(selectUser, nil).Once()

		repo.On("Insert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repo)

		response, err := srv.Create(userID, roomID, input)
		assert.NotNil(t, err)
		assert.Equal(t, returnDataMidtrans, response)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		input := reservations.ReservationEntity{
			CheckInDate:  time.Time{},
			CheckOutDate: time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC)}
		srv := New(repo)
		response, err := srv.Create(userID, roomID, input)
		assert.NotNil(t, err)
		assert.Equal(t, returnDataMidtrans, response)
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate date", func(t *testing.T) {
		input := reservations.ReservationEntity{
			CheckInDate:  time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
		}
		srv := New(repo)
		response, err := srv.Create(userID, roomID, input)
		assert.NotNil(t, err)
		assert.Equal(t, returnDataMidtrans, response)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when select room", func(t *testing.T) {
		input := reservations.ReservationEntity{
			UserID:       1,
			RoomID:       1,
			CheckInDate:  time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
			TotalNight:   1,
			TotalPrice:   200,
		}

		repo.On("SelectRoom", mock.Anything).Return(reservations.ReservationEntity{}, errors.New("error select")).Once()

		srv := New(repo)

		response, err := srv.Create(userID, roomID, input)
		assert.NotNil(t, err)
		assert.Equal(t, returnDataMidtrans, response)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when select user", func(t *testing.T) {
		input := reservations.ReservationEntity{
			UserID:       1,
			RoomID:       1,
			CheckInDate:  time.Date(2023, time.March, 16, 0, 0, 0, 0, time.UTC),
			CheckOutDate: time.Date(2023, time.March, 17, 0, 0, 0, 0, time.UTC),
			TotalNight:   1,
			TotalPrice:   200,
		}

		repo.On("SelectRoom", mock.Anything).Return(selectRoom, nil).Once()
		repo.On("SelectUser", mock.Anything).Return(reservations.ReservationEntity{}, errors.New("error select")).Once()

		srv := New(repo)

		response, err := srv.Create(userID, roomID, input)
		assert.NotNil(t, err)
		assert.Equal(t, returnDataMidtrans, response)
		repo.AssertExpectations(t)
	})
}
