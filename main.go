package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4321", "The port of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/", randomGetter)
	app.Listen(*listenAddr)

}

func randomGetter(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"Random bullshit": "GO!",
	})
}
