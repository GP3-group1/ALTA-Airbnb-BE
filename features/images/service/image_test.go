package service

import (
	"alta-airbnb-be/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateImage(t *testing.T) {
	table := CreateImageTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			imageDataMock := new(mocks.ImageData_)
			imageDataMock.On("InsertImage", mock.Anything).Return(v.Output.imageEntity, v.Output.errResult)

			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("SelectRoomsByUserId", mock.Anything).Return(v.Output.roomEntity, v.Output.errResult)
			imageService := New(imageDataMock, roomDataMock)
			_, err := imageService.CreateImage(v.Input.userId, &v.Input.imageEntity)
			fmt.Println(v.Output.IsError)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestChangeImage(t *testing.T) {
	table := CreateImageTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			imageDataMock := new(mocks.ImageData_)
			imageDataMock.On("UpdateImage", mock.Anything).Return(v.Output.imageEntity, v.Output.errResult)

			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("SelectRoomsByUserId", mock.Anything).Return(v.Output.roomEntity, v.Output.errResult)
			imageService := New(imageDataMock, roomDataMock)
			_, err := imageService.ChangeImage(v.Input.userId, &v.Input.imageEntity)
			fmt.Println(v.Output.IsError)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestRemoveImage(t *testing.T) {
	table := RemoveImageTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			imageDataMock := new(mocks.ImageData_)
			imageDataMock.On("DeleteImage", mock.Anything).Return(v.Output.errResult)

			roomDataMock := new(mocks.RoomData_)
			roomDataMock.On("SelectRoomsByUserId", mock.Anything).Return(v.Output.roomEntity, v.Output.errResult)
			imageService := New(imageDataMock, roomDataMock)
			err := imageService.RemoveImage(v.Input.userId, &v.Input.imageEntity)
			fmt.Println(v.Output.IsError)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
