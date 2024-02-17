package main

import (
	"context"
	"flag"
	"log"

	"github.com/aimensahnoun/hotel-booker/api"
	"github.com/aimensahnoun/hotel-booker/api/middleware"
	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{

	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{
			"Error: ": err.Error(),
		})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":4321", "The port of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
if err != nil {
		log.Fatal(err)
	}

	// Init Mongo handlers
	var (
		userHandler  = api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))
		hotelStore   = db.NewMongoHotelStore(client, db.DBNAME)
		hotelHandler = api.NewHotelHandler(hotelStore)
		roomHandler  = api.NewRoomHandler(db.NewMongoRoomStore(client, db.DBNAME, hotelStore))
		app          = fiber.New(config)
		apiv1        = app.Group("/api/v1", middleware.JWTAuthentication)
	)


  // Auth
  apiv1.Get("/auth/login", userHandler.HandleAuthenticateUser)

  // User
	apiv1.Get("/user/:id", userHandler.HandleGetUserByID)
	apiv1.Post("/user", userHandler.HandleInsertUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteuser)
	apiv1.Put("/user/:id", userHandler.HandleUpdateUser)

	// Hotel
	apiv1.Post("/hotel", hotelHandler.HandleInsertHotel)
	apiv1.Get("/hotel", hotelHandler.HandleGetAllHotels)
  apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotelByID)
  apiv1.Get("/hotel/:id/rooms", roomHandler.HanderGetRooms)

	// Room
	apiv1.Post("/room", roomHandler.HandleInsertRooms)

	app.Listen(*listenAddr)

}
