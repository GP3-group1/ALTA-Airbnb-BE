package consts

// Bind Error
const (
	RESERVATION_ErrorBindReservationData string = "error bind reservation data"
)

// Response Success
const (
	// Insert
	RESERVATION_InsertSuccess string = "succesfully insert reservation data"

	RESERVATION_RoomAvailable string = "room available"
)

// Response Error
const (
	// Insert
	RESERVATION_InsertFailed string = "you have insufficient balance"

	RESERVATION_RoomNotAvailable string = "room not available on inputed date"
)
