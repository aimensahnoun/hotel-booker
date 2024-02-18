package db

const (
	DBNAME     = "hotel_booker"
	TESTDBNAME = "hotel_tester"
	DBURI      = "mongodb://localhost:27017"
	TESTDBURI  = "mongodb://localhost:27020"
)

type Store struct {
	RoomStore    RoomStore
	UserStore    UserStore
	HotelStore   HotelStore
	BookingStore BookingStore
}
