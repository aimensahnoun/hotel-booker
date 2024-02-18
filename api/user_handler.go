package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
)

type UserHandler struct {
	store db.Store
}

func NewUserHandler(store db.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.store.UserStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.UserStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleDeleteuser(c *fiber.Ctx) error {
	ID := c.Params("id")

	err := h.store.UserStore.DeleteUser(c.Context(), ID)
	if err != nil {
		return err
	}

	return c.JSON("User deleted")
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		ID     = c.Params("id")
	)

	err := c.BodyParser(&params)

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	if err != nil {
		return err
	}

	res, err := h.store.UserStore.UpdateUser(c.Context(), &params, ID)
	if err != nil {
		return err
	}

	return c.JSON(fmt.Sprintf("Updated user: %s", res))
}
