package service

import (
	"alta-airbnb-be/app/config"
	"alta-airbnb-be/features/reservations"
	"alta-airbnb-be/features/users"
	"alta-airbnb-be/utils/consts"
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type reservationService struct {
	reservationData reservations.ReservationData_
	validate        *validator.Validate
}

// CheckReservation implements reservations.ReservationService_
func (reservationService *reservationService) CheckReservation(input reservations.ReservationEntity, roomID uint) ([]reservations.ReservationEntity, error) {
	errValidate := reservationService.validate.StructExcept(input, "Room", "User")
	if errValidate != nil {
		return nil, errValidate
	}

	// date validation: check in date must lower than check out date
	diff := input.CheckOutDate.Sub(input.CheckInDate)
	totalNight := int(diff.Hours() / 24)
	if totalNight < 1 {
		return nil, errors.New(consts.RESERVATION_InvalidInput)
	}

	reservationEntity, errSelect := reservationService.reservationData.CheckReservation(input, roomID)
	if errSelect != nil {
		return nil, errSelect
	}
	return reservationEntity, nil
}

// GetAll implements reservations.ReservationServiceInterface_
func (reservationService *reservationService) GetAll(page, limit int, userID uint) ([]reservations.ReservationEntity, error) {
	offset := (page - 1) * limit
	reservationEntity, errSelect := reservationService.reservationData.SelectAll(limit, offset, userID)
	if errSelect != nil {
		return []reservations.ReservationEntity{}, errSelect
	}
	return reservationEntity, nil
}

// Create implements reservations.ReservationServiceInterface_
func (reservationService *reservationService) Create(userID, idParam uint, inputReservation reservations.ReservationEntity) (reservations.MidtransResponse, error) {
	errValidate := reservationService.validate.StructExcept(inputReservation, "Room", "User")
	if errValidate != nil {
		return reservations.MidtransResponse{}, errValidate
	}

	// input user id and room id
	inputReservation.UserID = userID
	inputReservation.RoomID = idParam

	// count total night
	diff := inputReservation.CheckOutDate.Sub(inputReservation.CheckInDate)
	inputReservation.TotalNight = int(diff.Hours() / 24)

	// input date validation
	if inputReservation.TotalNight < 1 {
		return reservations.MidtransResponse{}, errors.New(consts.RESERVATION_InvalidInput)
	}

	// get room data by id
	selectRoom, errSelectRoom := reservationService.reservationData.SelectRoom(idParam)
	if errSelectRoom != nil {
		return reservations.MidtransResponse{}, errSelectRoom
	}

	// count total price
	inputReservation.TotalPrice = float64(selectRoom.Room.Price) * float64(inputReservation.TotalNight)

	// get user data by id
	selectUser, errSelectUser := reservationService.reservationData.SelectUser(inputReservation.UserID)
	if errSelectUser != nil {
		return reservations.MidtransResponse{}, errSelectUser
	}

	// balance validation
	if selectUser.User.Balance < inputReservation.TotalPrice {
		return reservations.MidtransResponse{}, errors.New(consts.RESERVATION_InsertFailed)
	}

	// decrease user balance
	inputUser := users.UserEntity{}
	inputUser.Balance = selectUser.User.Balance - inputReservation.TotalPrice

	// insert to reservations and update user balance
	errInsert := reservationService.reservationData.Insert(inputReservation, inputUser, userID)
	if errInsert != nil {
		return reservations.MidtransResponse{}, errInsert
	}

	// get reservation id
	selectReservation, errSelectReservation := reservationService.reservationData.SelectReservation()
	if errSelectReservation != nil {
		return reservations.MidtransResponse{}, errSelectReservation
	}

	// request midtrans snap
	var snapClient = snap.Client{}
	snapClient.New(config.MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	// parsing user id and item id
	user_id := strconv.Itoa(int(userID))
	item_id := strconv.Itoa(int(selectReservation.ID))

	// customer
	custAddress := &midtrans.CustomerAddress{
		FName:       selectUser.User.Name,
		Phone:       selectUser.User.PhoneNumber,
		Address:     selectUser.User.Address,
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "ALTA-Airbnb-" + user_id + "-" + item_id,
			GrossAmt: int64(inputReservation.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    selectUser.User.Name,
			Email:    selectUser.User.Email,
			Phone:    selectUser.User.PhoneNumber,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Room-" + item_id,
				Qty:   int32(inputReservation.TotalNight),
				Price: int64(selectRoom.Room.Price),
				Name:  selectRoom.Room.Name,
			},
		},
	}

	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		return reservations.MidtransResponse{}, errSnap
	}

	midtransResponse := reservations.MidtransResponse{
		Token:       response.Token,
		RedirectUrl: response.RedirectURL,
	}

	return midtransResponse, nil
}

func New(reservationData reservations.ReservationData_) reservations.ReservationService_ {
	return &reservationService{
		reservationData: reservationData,
		validate:        validator.New(),
	}
}
