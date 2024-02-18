package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aimensahnoun/hotel-booker/types"
)

const bookingCol = "booking"

type BookingStore interface {
	CreateBooking(context.Context, *types.Booking) (*types.Booking, error)
}

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

func (s *MongoBookingStore) CreateBooking(
	ctx context.Context,
	booking *types.Booking,
) (*types.Booking, error) {
	room, err := s.roomStore.GetRoomByID(ctx, booking.RoomID)
	if err != nil || room == nil {
		return nil, err
	}

	res, err := s.col.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}

	booking.ID = res.InsertedID.(primitive.ObjectID).Hex()

	err = s.roomStore.InsertBookingToRoom(ctx, booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}
