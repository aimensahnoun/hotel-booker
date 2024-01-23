package db

import (
	"context"

	"github.com/aimensahnoun/hotel-booker/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCol = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
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
	var user types.User
	err := s.col.FindOne(ctx, bson.M{"_id": ToObjectID(id)}).Decode(user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}
