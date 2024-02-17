package api

import (
	"fmt"

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



func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	users, err := h.userStore.GetUsers(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleDeleteuser(c *fiber.Ctx) error {
	var (
		ID = c.Params("id")
	)

	err := h.userStore.DeleteUser(c.Context(), ID)

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

	res, err := h.userStore.UpdateUser(c.Context(), &params, ID)

	if err != nil {
		return err
	}

	return c.JSON(fmt.Sprintf("Updated user: %s", res))

}


