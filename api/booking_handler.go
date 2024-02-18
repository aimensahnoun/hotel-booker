package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
)

type BookingHandler struct {
	store db.Store
}

func NewBookingHandler(store db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) HandleCreateBooking(c *fiber.Ctx) error {
	var params types.CreateBookingParams
	roomID := c.Params("id")
	userID := c.Context().UserValue("id").(string)
	if err := c.BodyParser(&params); err != nil {
		return c.JSON(err)
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	newBooking := types.Booking{
		NumberOfGuests: params.NumberOfGuests,
		From:           params.From,
		Until:          params.Until,
		RoomID:         roomID,
		UserID:         userID,
	}

	createdBooking, err := h.store.BookingStore.CreateBooking(c.Context(), &newBooking)
	if err != nil {
		return err
	}

	return c.JSON(createdBooking)
}
