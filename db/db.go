package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const DBNAME = "hotel_booker"
const DBURI = "mongodb://localhost:27017"

func ToObjectID(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	return oid
}
