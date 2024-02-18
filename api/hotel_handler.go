package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
)

type HotelHandler struct {
	store db.Store
}

func NewHotelHandler(store db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleInsertHotel(c *fiber.Ctx) error {
	var params types.InsertHotelParams

	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(err)
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	hotel := types.NewHotelFromParams(&params)

	res, err := h.store.HotelStore.InsertHotel(c.Context(), hotel)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(res)
}

func (h *HotelHandler) HandleGetAllHotels(c *fiber.Ctx) error {
	hotels, err := h.store.HotelStore.GetHotels(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	id := c.Params("id")

	hotel, err := h.store.HotelStore.GetHotelByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(hotel)
}
