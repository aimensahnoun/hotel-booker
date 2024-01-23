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
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)

	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleInsertUser(c *fiber.Ctx) error {

	var ctx = context.Background()

	var params types.InsertUserParams

	err := c.BodyParser(&params)

	if err != nil {
		return err
	}

	user, err := types.NewUserFromParams(params)

	if err != nil {
		return err
	}

	res, err := h.userStore.InsertUser(ctx, user)

	if err != nil {
		return err
	}

	return c.JSON(res)

}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	var (
		ctx = context.Background()
	)

	users, err := h.userStore.GetUsers(ctx)

	if err != nil {
		return err
	}

	return c.JSON(users)
}
