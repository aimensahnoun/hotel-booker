package api

import (
	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
	"github.com/gofiber/fiber/v2"
)

type RoomHandler struct {
	roomStore db.RoomStore
}

func NewRoomHandler(roomStore db.RoomStore) *RoomHandler {
	return &RoomHandler{
		roomStore: roomStore,
	}
}

func (h *RoomHandler) HandleInsertRooms(c *fiber.Ctx) error {
	var (
		params types.InsertRoomParams
	)

	if err := c.BodyParser(&params); err != nil {
		return c.JSON(err)
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	room := types.NewRoomFromParams(&params)

	res, err := h.roomStore.InsertRoom(c.Context(), room)

	if err != nil {
		c.JSON(err)
	}

	return c.JSON(res)
}

func (h *RoomHandler) HanderGetRooms(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)

	rooms, err := h.roomStore.GetRooms(c.Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(rooms)
}
