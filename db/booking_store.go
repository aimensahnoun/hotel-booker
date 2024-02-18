package db

import "go.mongodb.org/mongo-driver/mongo"

const bookingCol = "booking"

type BookingStore interface{}

type MongoBookingStore struct {
	client    *mongo.Client
	col       *mongo.Collection
	roomStore RoomStore
}

func NewMongoBookingStore(
	client *mongo.Client,
	dbName string,
	roomStore RoomStore,
) *MongoBookingStore {
	return &MongoBookingStore{
		client:    client,
		col:       client.Database(dbName).Collection(bookingCol),
		roomStore: roomStore,
	}
}
