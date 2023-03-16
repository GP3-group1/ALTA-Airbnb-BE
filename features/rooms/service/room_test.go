package service

import (
	"alta-airbnb-be/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRoom(t *testing.T) {
	table := CreateRoomTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("InsertRoom", mock.Anything).Return(v.Output.errResult)

			roomService := New(roomDataMock)
			err := roomService.CreateRoom(&v.Input.roomEntity)
			fmt.Println(v.Output.IsError)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestChangeRoom(t *testing.T) {
	table := ChangeRoomTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("UpdateRoom", mock.Anything).Return(v.Output.errResult)

			roomService := New(roomDataMock)
			err := roomService.ChangeRoom(&v.Input.roomEntity)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestRemoveRoom(t *testing.T) {
	table := RemoveRoomTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("DeleteRoom", mock.Anything).Return(v.Output.errResult)

			roomService := New(roomDataMock)
			err := roomService.RemoveRoom(&v.Input.roomEntity)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetRooms(t *testing.T) {
	table := GetRoomsTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("SelectRooms", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.Anything).Return(v.Output.roomEntity, v.Output.errResult)

			roomService := New(roomDataMock)
			_, err := roomService.GetRooms(v.Input.limit, v.Input.offset, v.Input.queryParams)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetRoomByRoomId(t *testing.T) {
	table := GetRoomByRoomIdTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("SelectRoomByRoomId", mock.Anything).Return(v.Output.roomEntity, v.Output.errResult)

			roomService := New(roomDataMock)
			_, err := roomService.GetRoomByRoomId(&v.Input.roomEntity)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetRoomsByUserId(t *testing.T) {
	table := GetRoomsByUserIdTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("SelectRoomsByUserId", mock.Anything).Return(v.Output.roomEntity, v.Output.errResult)

			roomService := New(roomDataMock)
			_, err := roomService.GetRoomsByUserId(&v.Input.roomEntity)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCreateReview(t *testing.T) {
	table := CreateReviewTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("InsertReview", mock.Anything).Return(v.Output.errResult)

			roomService := New(roomDataMock)
			err := roomService.CreateReview(&v.Input.reviewEntity)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetReviewsByRoomId(t *testing.T) {
	table := GetReviewsByRoomIdTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("SelectReviewsByRoomId", mock.Anything).Return(v.Output.reviewEntity, v.Output.errResult)

			roomService := New(roomDataMock)
			_, err := roomService.GetReviewsByRoomId(&v.Input.reviewEntity)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}