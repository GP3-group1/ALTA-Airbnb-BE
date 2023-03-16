// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	reservations "alta-airbnb-be/features/reservations"

	mock "github.com/stretchr/testify/mock"

	users "alta-airbnb-be/features/users"
)

// ReservationData is an autogenerated mock type for the ReservationData_ type
type ReservationData struct {
	mock.Mock
}

// CheckReservation provides a mock function with given fields: input, roomID
func (_m *ReservationData) CheckReservation(input reservations.ReservationEntity, roomID uint) (int64, error) {
	ret := _m.Called(input, roomID)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(reservations.ReservationEntity, uint) (int64, error)); ok {
		return rf(input, roomID)
	}
	if rf, ok := ret.Get(0).(func(reservations.ReservationEntity, uint) int64); ok {
		r0 = rf(input, roomID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(reservations.ReservationEntity, uint) error); ok {
		r1 = rf(input, roomID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: inputReservation, inputUser, userID
func (_m *ReservationData) Insert(inputReservation reservations.ReservationEntity, inputUser users.UserEntity, userID uint) error {
	ret := _m.Called(inputReservation, inputUser, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(reservations.ReservationEntity, users.UserEntity, uint) error); ok {
		r0 = rf(inputReservation, inputUser, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields: limit, offset, userID
func (_m *ReservationData) SelectAll(limit int, offset int, userID uint) ([]reservations.ReservationEntity, error) {
	ret := _m.Called(limit, offset, userID)

	var r0 []reservations.ReservationEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, uint) ([]reservations.ReservationEntity, error)); ok {
		return rf(limit, offset, userID)
	}
	if rf, ok := ret.Get(0).(func(int, int, uint) []reservations.ReservationEntity); ok {
		r0 = rf(limit, offset, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reservations.ReservationEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, uint) error); ok {
		r1 = rf(limit, offset, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectRoom provides a mock function with given fields: roomID
func (_m *ReservationData) SelectRoom(roomID uint) (reservations.ReservationEntity, error) {
	ret := _m.Called(roomID)

	var r0 reservations.ReservationEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (reservations.ReservationEntity, error)); ok {
		return rf(roomID)
	}
	if rf, ok := ret.Get(0).(func(uint) reservations.ReservationEntity); ok {
		r0 = rf(roomID)
	} else {
		r0 = ret.Get(0).(reservations.ReservationEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(roomID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUser provides a mock function with given fields: userID
func (_m *ReservationData) SelectUser(userID uint) (reservations.ReservationEntity, error) {
	ret := _m.Called(userID)

	var r0 reservations.ReservationEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (reservations.ReservationEntity, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) reservations.ReservationEntity); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(reservations.ReservationEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReservationData interface {
	mock.TestingT
	Cleanup(func())
}

// NewReservationData creates a new instance of ReservationData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReservationData(t mockConstructorTestingTNewReservationData) *ReservationData {
	mock := &ReservationData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}