package api

import (
	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	hotelStore db.HotelStore
}

func NewHotelHandler(hotelStore db.HotelStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hotelStore,
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

	res, err := h.hotelStore.InsertHotel(c.Context(), hotel)

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(res)

}

func (h *HotelHandler) HandleGetAllHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelStore.GetHotels(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(hotels)

}


func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
  id := c.Params("id")
  
  hotel , err := h.hotelStore.GetHotelByID(c.Context(), id)

  if err != nil {
    return err
  }

  return c.JSON(hotel)

}
