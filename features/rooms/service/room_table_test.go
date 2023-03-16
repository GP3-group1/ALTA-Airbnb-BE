package service

import (
	"alta-airbnb-be/features/reviews"
	"alta-airbnb-be/features/rooms"
	"alta-airbnb-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		limit        int
		offset       int
		queryParams  map[string][]string
		roomEntity   rooms.RoomEntity
		reviewEntity reviews.ReviewEntity
	}
	Output struct {
		IsError      bool
		roomEntity   interface{}
		reviewEntity interface{}
		errResult    error
	}
}

func CreateRoomTestTable() []TestTable {
	tname := "test create room"
	return []TestTable{
		{
			Name: tname + "expect failed - empty field name",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field overview",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field description",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field location",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field price",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       0,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				errResult: nil,
			},
		},
	}
}

func ChangeRoomTestTable() []TestTable {
	tname := "test change room"
	return []TestTable{
		{
			Name: tname + "expect failed - empty field name",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field overview",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field description",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field location",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed - empty field price",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       0,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(consts.ROOM_InvalidInput),
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				errResult: nil,
			},
		},
	}
}

func RemoveRoomTestTable() []TestTable {
	tname := "test remove room"
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					ID:          1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					ID:          1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				roomEntity: rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				errResult: nil,
			},
		},
	}
}

func GetRoomsTestTable() []TestTable {
	tname := "test get rooms"
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				limit:  1,
				offset: 1,
				queryParams: map[string][]string{
					"name": []string{"Tamara"},
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				limit:  1,
				offset: 1,
				queryParams: map[string][]string{
					"name": []string{"Tamara"},
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				}},
				errResult: nil,
			},
		},
	}
}

func GetRoomByRoomIdTestTable() []TestTable {
	tname := "test get room by room id"
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					ID:          1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				roomEntity: &rooms.RoomEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					ID:          1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				roomEntity: &rooms.RoomEntity{
					ID:          1,
					UserID:      1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
				errResult: nil,
			},
		},
	}
}

func GetRoomsByUserIdTestTable() []TestTable {
	tname := "test get rooms by user id"
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					UserID:      1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{}},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				roomEntity: rooms.RoomEntity{
					UserID:      1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				roomEntity: []*rooms.RoomEntity{&rooms.RoomEntity{
					ID:          1,
					UserID:      1,
					Name:        "Villa Tamara",
					Overview:    "Villa dengan view yang bagus",
					Description: "Villa ini mempunyai pemandangan luar yang bagus",
					Location:    "Jakarta",
					Price:       100,
				}},
				errResult: nil,
			},
		},
	}
}

func CreateReviewTestTable() []TestTable {
	tname := "test create review"
	return []TestTable{
		{
			Name: tname + "expect failed - rating is smaller than 1",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					UserID:  1,
					RoomID:  1,
					Rating:  -1,
					Comment: "Rekomen sih ini, bagus banget. Bukan villa kaleng-kaleng",
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				reviewEntity: &reviews.ReviewEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect failed - rating is greater than 5",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					UserID:  1,
					RoomID:  1,
					Rating:  6,
					Comment: "Rekomen sih ini, bagus banget. Bukan villa kaleng-kaleng",
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				reviewEntity: &reviews.ReviewEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect failed - rating is empty",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					UserID:  1,
					RoomID:  1,
					Rating:  0,
					Comment: "Rekomen sih ini, bagus banget. Bukan villa kaleng-kaleng",
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				reviewEntity: &reviews.ReviewEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect failed - comment is empty",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					UserID:  1,
					RoomID:  1,
					Rating:  4,
					Comment: "",
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				reviewEntity: &reviews.ReviewEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					UserID:  1,
					RoomID:  1,
					Rating:  4,
					Comment: "Rekomen sih ini, bagus banget. Bukan villa kaleng-kaleng",
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError:   true,
				reviewEntity: &reviews.ReviewEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					UserID:  1,
					RoomID:  1,
					Rating:  4,
					Comment: "Rekomen sih ini, bagus banget. Bukan villa kaleng-kaleng",
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				reviewEntity: &reviews.ReviewEntity{
					UserID:  1,
					RoomID:  1,
					Rating:  4,
					Comment: "Rekomen sih ini, bagus banget. Bukan villa kaleng-kaleng",
				},
				errResult: nil,
			},
		},
	}
}

func GetReviewsByRoomIdTestTable() []TestTable {
	tname := "test get reviews by room id"
	return []TestTable{
		{
			Name: tname + "expect failed",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: true,
				reviewEntity: []*reviews.ReviewEntity{&reviews.ReviewEntity{}},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				limit        int
				offset       int
				queryParams  map[string][]string
				roomEntity   rooms.RoomEntity
				reviewEntity reviews.ReviewEntity
			}{
				reviewEntity: reviews.ReviewEntity{
					RoomID: 1,
				},
			},
			Output: struct {
				IsError      bool
				roomEntity   interface{}
				reviewEntity interface{}
				errResult    error
			}{
				IsError: false,
				reviewEntity: []*reviews.ReviewEntity{&reviews.ReviewEntity{
					UserID:   1,
					Username: "Mike Tyson",
					RoomID:   1,
					Rating:   4,
					Comment:  "Rekomen sih ini, bagus banget. Bukan villa kaleng-kaleng",
				}},
				errResult: nil,
			},
		},
	}
}
