package helpers

import (
	"alta-airbnb-be/app/config"
	"alta-airbnb-be/features/reservations"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func RequestSnapMidtrans(user reservations.ReservationEntity, room reservations.ReservationEntity, input reservations.ReservationEntity) (reservations.MidtransResponse, error) {
	// request midtrans snap
	var snapClient = snap.Client{}
	snapClient.New(config.MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	// parsing user id, room id, uuid
	user_id := strconv.Itoa(int(input.UserID))
	room_id := strconv.Itoa(int(input.RoomID))
	uuid := uuid.New()
	order_id := fmt.Sprintf("%v", uuid)

	// customer
	custAddress := &midtrans.CustomerAddress{
		FName:       user.User.Name,
		Phone:       user.User.PhoneNumber,
		Address:     user.User.Address,
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "ALTA-Airbnb-" + user_id + "-" + order_id,
			GrossAmt: int64(input.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    user.User.Name,
			Email:    user.User.Email,
			Phone:    user.User.PhoneNumber,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Room-" + room_id,
				Qty:   int32(input.TotalNight),
				Price: int64(room.Room.Price),
				Name:  room.Room.Name,
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
