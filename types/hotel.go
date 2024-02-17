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
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rating   float32 `json:"rating"`
}

func (params InsertHotelParams) Validate() map[string]string {
	errors := map[string]string{}

	if len(params.Name) < minHotelNameLength {
		errors["name"] = fmt.Sprintf(
			"Hotel name must be at least %d characters.",
			minHotelNameLength,
		)
	}

	if len(params.Location) < minHotelNameLength {
		errors["location"] = fmt.Sprintf(
			"Hotel location must be at least %d characters.",
			minLocationLength,
		)
	}

	if params.Rating <= 0 {
		errors["rating"] = "Invalid rating"
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
	Name     string               `bson:"name"          json:"name"`
	Location string               `bson:"location"      json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms"         json:"rooms"`
	Rating   float32              `bson:"rating"        json:"rating"`
}
