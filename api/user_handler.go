package api

import (
	"context"
	"log"

	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)

	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(user)
}

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
