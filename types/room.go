package types

type InsertRoomParams struct {
	Type    string  `json:"type"`
	Price   float64 `json:"price"`
	HotelID string  `json:"hotelID"`
	Seaside bool    `json:"seaside"`
}

func (params *InsertRoomParams) Validate() map[string]string {
	errors := map[string]string{}

	_, exists := RoomTypes[params.Type]

	if exists == false {
		errors["type"] = "Invalid room type"
	}

	if params.Price < 0 {
		errors["price"] = "Invalid price"
	}

	if len(params.HotelID) != 12 {
		errors["hotelID"] = "Invalid hotel id"
	}

	return errors
}

func NewRoomFromParams(params *InsertRoomParams) *Room {
	return &Room{
		Type:    params.Type,
		HotelID: params.HotelID,
		Price:   params.Price,
		SeaSide: params.Seaside,
	}
}

var RoomTypes = map[string]bool{
	"single": true,
	"double": true,
	"king":   true,
	"suite":  true,
}

type Room struct {
	ID      string  `bson:"_id,omitempty" json:"id,omitempty"`
	Type    string  `bson:"type"          json:"type"`
	SeaSide bool    `bson:"seaside"       json:"seaside"`
	Price   float64 `bson:"price"         json:"price"`
	HotelID string  `bson:"hotelID"       json:"hotelId"`
}
