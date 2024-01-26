package db

import (
	"context"

	"github.com/aimensahnoun/hotel-booker/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCol = "room"

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
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

	room.ID = res.InsertedID.(primitive.ObjectID)

	return room, nil

}
