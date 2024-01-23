package db

import (
	"context"
	"log"

	"github.com/aimensahnoun/hotel-booker/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCol = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	InsertUser(context.Context, types.User) (*types.User, error)
	GetUsers(context.Context) (*mongo.Cursor, error)
}

type MongoUserStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {

	return &MongoUserStore{
		client: client,
		col:    client.Database(DBNAME).Collection(userCol),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var user types.User
	err = s.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (s *MongoUserStore) InsertUser(ctx context.Context, user types.User) (*types.User, error) {
	res, err := s.col.InsertOne(ctx, user)

	if err != nil {
		log.Fatal(err)
	}

	insertedID, ok := res.InsertedID.(string)
	if !ok {
		log.Fatal("InsertedID is not a string")
	}

	return &types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		ID:        insertedID,
	}, nil

}

func (s *MongoUserStore) GetUsers(ctx context.Context) (*mongo.Cursor, error) {

	res, err := s.col.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	return res, nil
}
