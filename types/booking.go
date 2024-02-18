package types

import "time"

type CreateBookingParams struct {
	From           time.Time `json:"from"`
	Until          time.Time `json:"until"`
	NumberOfGuests int8      `json:"numberOfGuests"`
}

func (params *CreateBookingParams) Validate() map[string]string {
	errors := map[string]string{}

	if time.Now().After(params.From) {
		errors["from"] = "Booking date cannot be in the past."
	}

	if time.Now().After(params.Until) {
		errors["unitl"] = "Until cannot be in the past."
	}

	if params.NumberOfGuests <= 0 {
		errors["numberOfGuests"] = "Number of guests invalid."
	}

	return errors
}

type Booking struct {
	ID             string    `bson:"_id,omitempty"  json:"id,omitempty"`
	RoomID         string    `bson:"roomID"         json:"roomID"`
	NumberOfGuests int8      `bson:"numberOfGuests" json:"numberOfGuests"`
	From           time.Time `bson:"from"           json:"from"`
	Until          time.Time `bson:"until"          json:"until"`
	UserID         string    `bson:"userID"         json:"userID"`
}
