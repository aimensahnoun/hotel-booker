package api

import (
	"github.com/aimensahnoun/hotel-booker/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	users := []types.User{
		{
			FirstName: "Hmida",
			LastName:  "Genawi",
		},
		{
			FirstName: "Rabie",
			LastName:  "Hasnawi",
		},
	}

	return c.JSON(users)
}

func HandleGetUserByID(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Hmida",
		LastName:  "Genawi",
	}

	return c.JSON(user)
}
