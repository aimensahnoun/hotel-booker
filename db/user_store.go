package db

import (
	"context"
	"errors"

	"github.com/aimensahnoun/hotel-booker/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCol = "users"

type DbDropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	DbDropper

	GetUserByID(context.Context, string) (*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, *types.UpdateUserParams, string) (primitive.ObjectID, error)
}

type MongoUserStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {

	return &MongoUserStore{
		client: client,
		col:    client.Database(dbname).Collection(userCol),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	println("--------DROPPING USERS COLLECTION----------")
	return s.col.Drop(ctx)
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

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.col.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil

}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {

	cur, err := s.col.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []*types.User

	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, err
	}

	return users, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, ID string) error {
	oid, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	res, err := s.col.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, values *types.UpdateUserParams, userId string) (primitive.ObjectID, error) {

	oid, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	res, err := s.col.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{
		"$set": values,
	})

	if err != nil {
		return primitive.ObjectID{}, err
	}

	if res.MatchedCount == 0 {
		return primitive.ObjectID{}, errors.New("user does not exist")
	}

	return oid, nil

}
