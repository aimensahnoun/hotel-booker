package main

import (
	"context"
	"log"

	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))

	if err != nil {
		log.Fatal(err)
	}

	client.Database(db.DBNAME).Drop(ctx)

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME, hotelStore)

	hotel := types.Hotel{
		Name:     "Teknik Yapi Concord",
		Location: "Kadikoy",
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []types.Room{
		{
			Type:  types.DeluxeRoomType,
			Price: 136.88,
		},
		{
			Type:  types.DoubleRoomType,
			Price: 116.74,
		}, {
			Type:  types.SingleRoomType,
			Price: 84.22,
		}, {
			Type:  types.SeaSideRoomType,
			Price: 244.99,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Hotel : %v", *insertedHotel)

	for i, room := range rooms {

		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Inserted room %d : %v", i+1, insertedRoom)

	}

}
