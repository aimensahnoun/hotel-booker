package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()
	app.Get("/", randomGetter)
	app.Listen(":4321")

}

func randomGetter(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"Random bullshit": "GO!",
	})
}
