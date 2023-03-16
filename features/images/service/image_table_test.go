package service

import (
	"alta-airbnb-be/features/images"
	"alta-airbnb-be/features/rooms"
	"alta-airbnb-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		userId      uint
		imageEntity images.ImageEntity
	}
	Output struct {
		IsError     bool
		imageEntity interface{}
		roomEntity  interface{}
		errResult   error
	}
}

func CreateImageTestTable() []TestTable {
	tname := "test create image"
	return []TestTable{
		{
			Name: tname + "expect failed - empty room id",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 1,
				imageEntity: images.ImageEntity{},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(consts.IMAGE_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - error select rooms by user id",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 1,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect succes - non forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: false,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 5,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: nil,
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: errors.New(""),
			},
		},
	}
}

func ChangeImageTestTable() []TestTable {
	tname := "test change image"
	return []TestTable{
		{
			Name: tname + "expect failed - empty room id",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 1,
				imageEntity: images.ImageEntity{},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(consts.IMAGE_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - error select rooms by user id",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 1,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect succes - non forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: false,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 5,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: nil,
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: errors.New(""),
			},
		},
	}
}

func RemoveImageTestTable() []TestTable {
	tname := "test remove room"
	return []TestTable{
		{
			Name: tname + "expect failed - empty room id",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 1,
				imageEntity: images.ImageEntity{},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(consts.IMAGE_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - error select rooms by user id",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 1,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect succes - non forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: false,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 5,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: &images.ImageEntity{},
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: nil,
			},
		},
		{
			Name: tname + "expect failed - forbidden request",
			Input: struct {
				userId      uint
				imageEntity images.ImageEntity
			}{
				userId: 2,
				imageEntity: images.ImageEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError bool; 
				imageEntity interface{}; 
				roomEntity interface{}; 
				errResult error;
			}{
				IsError: true,
				imageEntity: nil,
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      2,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				&rooms.RoomEntity{
					ID:          2,
					UserID:      3,
					Name:        "Villa Kencana",
					Overview:    "Villa dekat pantai",
					Description: "Villa ini mempunyai suasana dan lingkungan yang fresh",
					Location:    "Jakarta",
					Price:       200,
				}},
				errResult: errors.New(""),
			},
		},
	}
}
