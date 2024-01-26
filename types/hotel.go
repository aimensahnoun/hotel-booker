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
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type      RoomType           `bson:"type" json:"type"`
	BasePrice float64            `bson:"basePrice" json:"basePrice"`
	Price     float64            `bson:"price" json:"price"`
	HotelID   primitive.ObjectID `bson:"hotelID" json:"hotelId"`
}
