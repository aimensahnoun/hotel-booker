package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	minHotelNameLength = 4
	minLocationLength  = 5
)

type InsertHotelParams struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (params InsertHotelParams) Validate() map[string]string {
	errors := map[string]string{}

	if len(params.Name) < minHotelNameLength {
		errors["name"] = fmt.Sprintf("Hotel name must be at least %d characters.", minHotelNameLength)
	}

	if len(params.Location) < minHotelNameLength {
		errors["location"] = fmt.Sprintf("Hotel location must be at least %d characters.", minLocationLength)
	}

	return errors

}

func NewHotelFromParams(params *InsertHotelParams) *Hotel {
	return &Hotel{
		Name:     params.Name,
		Location: params.Location,
		Rooms:    []primitive.ObjectID{},
	}
}

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
  Rating float32 `bson:"rating" json:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type InsertRoomParams struct {
	Type    RoomType           `json:"type"`
	Price   float64            `json:"price"`
	HotelID primitive.ObjectID `json:"hotelID"`
}

func (r RoomType) String() string {
	switch r {
	case SingleRoomType:
		return "Single room"
	case DoubleRoomType:
		return "Double room"
	case SeaSideRoomType:
		return "Sea side room"
	case DeluxeRoomType:
		return "Deluxe room"
	default:
		return "Unknown"
	}
}

func (params *InsertRoomParams) Validate() map[string]string {
	errors := map[string]string{}

	if params.Type > DeluxeRoomType || params.Type < SingleRoomType {
		errors["type"] = "Invalid room type"
	}

	if params.Price < 0 {
		errors["price"] = "Invalid price"
	}

	if params.HotelID.IsZero() || len(params.HotelID) != 12 {
		errors["hotelID"] = "Invalid hotel id"
	}

	return errors
}

func NewRoomFromParams(params *InsertRoomParams) *Room {
	return &Room{
		Type:    params.Type,
		HotelID: params.HotelID,
		Price:   params.Price,
	}
}

type Room struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type    RoomType           `bson:"type" json:"type"`
	Price   float64            `bson:"price" json:"price"`
	HotelID primitive.ObjectID `bson:"hotelID" json:"hotelId"`
}
