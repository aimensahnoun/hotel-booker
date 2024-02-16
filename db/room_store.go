package db

import (
	"context"

	"github.com/aimensahnoun/hotel-booker/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCol = "room"

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
  GetRooms(context.Context , string ) ([]*types.Room, error)
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

	s.hotelStore.AddHotelRoom(ctx, room.HotelID, room.ID)

	return room, nil

} 


func (s *MongoRoomStore) GetRooms(ctx context.Context , hotelID string) ([]*types.Room , error) {

  oid , err := primitive.ObjectIDFromHex(hotelID)

  if err != nil {
    return nil , err
  }

  cur , err := s.col.Find(ctx, bson.M{"hotelID" : oid}) 

  if err != nil {
    return nil , err
  }

  var rooms  []*types.Room

  if err :=cur.All(ctx, &rooms); err != nil {
    return nil, err
  }


  return rooms , nil
}
