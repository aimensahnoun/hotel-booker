package types

import "time"

type RoomBooking struct {
  ID             string    `bson:"_id,omitempty"  json:"id,omitempty"`
  RoomID         string    `bson:"roomID"         json:"roomID"`
  HotelID        string    `bson:"hotelID"        json:"hotelID"`
  NumberOfGuests int8      `bson:"numberOfGuests" json:"numberOfGuests"`
  From           time.Time `bson:"from"           json:"from"`
  Until          time.Time `bson:"until"          json:"until"`
}

