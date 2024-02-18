package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/aimensahnoun/hotel-booker/api"
	"github.com/aimensahnoun/hotel-booker/api/middleware"
	"github.com/aimensahnoun/hotel-booker/db"
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
	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Init Mongo handlers
	var (
		userStore    = db.NewMongoUserStore(client, db.DBNAME)
		hotelStore   = db.NewMongoHotelStore(client, db.DBNAME)
		roomStore    = db.NewMongoRoomStore(client, db.DBNAME, hotelStore)
		bookingStore = db.NewMongoBookingStore(client, db.DBNAME, roomStore)
		store        = db.Store{
			UserStore:    userStore,
			HotelStore:   hotelStore,
			RoomStore:    roomStore,
			BookingStore: bookingStore,
		}

		userHandler    = api.NewUserHandler(store)
		authHandler    = api.NewAuthHandler(store)
		hotelHandler   = api.NewHotelHandler(store)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		app            = fiber.New(config)
		apiv1          = app.Group("/api/v1", middleware.JWTAuthentication)
	)

	// Auth
	apiv1.Post("/auth/login", authHandler.HandleAuthenticateUser)
	apiv1.Post("/auth/register", authHandler.HandleRegister)

	// User
	apiv1.Get("/user/:id", userHandler.HandleGetUserByID)
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

	// Booking
	apiv1.Post("/room/:id/book", bookingHandler.HandleCreateBooking)
	app.Listen(*listenAddr)
}
