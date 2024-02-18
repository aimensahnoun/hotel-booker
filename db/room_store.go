package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aimensahnoun/hotel-booker/types"
)

const roomCol = "room"

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, string) ([]*types.Room, error)
	GetRoomByID(context.Context, string) (*types.Room, error)
	InsertBookingToRoom(context.Context, *types.Booking) error
}

type MongoRoomStore struct {
	client     *mongo.Client
	col        *mongo.Collection
	hotelStore HotelStore
}

func NewMongoRoomStore(client *mongo.Client, dbname string, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		col:        client.Database(dbname).Collection(roomCol),
		hotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := s.col.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = res.InsertedID.(primitive.ObjectID).String()

	s.hotelStore.AddHotelRoom(ctx, room.HotelID, room.ID)

	return room, nil
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, hotelID string) ([]*types.Room, error) {
	oid, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		return nil, err
	}

	cur, err := s.col.Find(ctx, bson.M{"hotelID": oid})
	if err != nil {
		return nil, err
	}

	var rooms []*types.Room

	if err := cur.All(ctx, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *MongoRoomStore) GetRoomByID(ctx context.Context, roomId string) (*types.Room, error) {
	oid, err := primitive.ObjectIDFromHex(roomId)
	if err != nil {
		return nil, fmt.Errorf("Invalid room ID")
	}

	var room types.Room

	err = s.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&room)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (s *MongoRoomStore) InsertBookingToRoom(ctx context.Context, booking *types.Booking) error {
	roomOID, err := primitive.ObjectIDFromHex(booking.RoomID)
	if err != nil {
		return err
	}

	bookingOID, err := primitive.ObjectIDFromHex(booking.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": roomOID}
	update := bson.M{"$push": bson.M{"bookings": bookingOID}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("Invalid room id")
	}

	return nil
}
