package dto

// GuestPayload is the payload used to create a Guest
type GuestPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
}

// ReservationPayload is the payload used to create a Reservation
type ReservationPayload struct {
	GuestUUID    string `json:"guest_uuid"`
	HotelUUID    string `json:"hotel_uuid"`
	RoomTypeUUID string `json:"roomtype_uuid"`
}

// CancelReservationPayload is the payload used to cancel a Reservation
type CancelReservationPayload struct {
	GuestUUID    string `json:"guest_uuid"`
	RoomTypeUUID string `json:"roomtype_uuid"`
}
