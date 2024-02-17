package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aimensahnoun/hotel-booker/types"
)

const hotelCol = "hotels"

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	AddHotelRoom(context.Context, string, string) (string, error)
	GetHotels(context.Context) ([]*types.Hotel, error)
	GetHotelByID(context.Context, string) (*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		col:    client.Database(dbname).Collection(hotelCol),
	}
}

func (s *MongoHotelStore) InsertHotel(
	ctx context.Context,
	hotel *types.Hotel,
) (*types.Hotel, error) {
	res, err := s.col.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	hotel.ID = res.InsertedID.(primitive.ObjectID)

	return hotel, nil
}

func (s *MongoHotelStore) AddHotelRoom(
	ctx context.Context,
	hotelID string,
	roomID string,
) (string, error) {
	hotelOid, err := primitive.ObjectIDFromHex(hotelID)
	if err != nil {
		return "", err
	}

	roomOid, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return "", err
	}

	filter := bson.M{"_id": hotelOid}
	update := bson.M{"$push": bson.M{"rooms": roomOid}}

	res, err := s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}

	if res.MatchedCount == 0 {
		return "", errors.New("hotel does not exist")
	}

	return fmt.Sprintf("Room %s has been added to hotel %s", roomID, hotelID), nil
}

func (s *MongoHotelStore) GetHotels(ctx context.Context) ([]*types.Hotel, error) {
	cur, err := s.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var hotels []*types.Hotel

	if err := cur.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (s *MongoHotelStore) GetHotelByID(ctx context.Context, ID string) (*types.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	var hotel types.Hotel

	err = s.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel)
	if err != nil {
		return nil, err
	}

	return &hotel, nil
}
