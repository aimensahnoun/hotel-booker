package api

import (
	"context"

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
		id = c.Params("id")
	)
	user, err := h.userStore.GetUserByID(c.Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleInsertUser(c *fiber.Ctx) error {

	var params types.InsertUserParams

	err := c.BodyParser(&params)

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	if err != nil {
		return err
	}

	user, err := types.NewUserFromParams(params)

	if err != nil {
		return err
	}

	res, err := h.userStore.InsertUser(c.Context(), user)

	if err != nil {
		return err
	}

	return c.JSON(res)

}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	users, err := h.userStore.GetUsers(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(users)
}
