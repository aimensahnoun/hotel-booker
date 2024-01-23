package main

import (
	"flag"

	"github.com/aimensahnoun/hotel-booker/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4321", "The port of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUserByID)
	app.Listen(*listenAddr)

}
