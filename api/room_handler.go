package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
)

type RoomHandler struct {
	store db.Store
}

func NewRoomHandler(store db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (h *RoomHandler) HandleInsertRooms(c *fiber.Ctx) error {
	var params types.InsertRoomParams

	if err := c.BodyParser(&params); err != nil {
		return c.JSON(err)
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	room := types.NewRoomFromParams(&params)

	res, err := h.store.RoomStore.InsertRoom(c.Context(), room)
	if err != nil {
		c.JSON(err)
	}

	return c.JSON(res)
}

func (h *RoomHandler) HanderGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")

	rooms, err := h.store.RoomStore.GetRooms(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	userID := c.Context().UserValue("id")

	fmt.Println("User id from context :", userID)
	return nil
}
