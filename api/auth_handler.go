package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/aimensahnoun/hotel-booker/api/middleware"
	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

func (h *AuthHandler) HandleAuthenticateUser(c *fiber.Ctx) error {
	params := types.AuthenticateUserParams{}

	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	existingUser, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil || existingUser == nil {
		return fmt.Errorf("User does not exist")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(existingUser.EncryptedPassword),
		[]byte(params.Password),
	)
	if err != nil {
		return fmt.Errorf("Invalid password")
	}

	jwtToken, err := middleware.GenerateJWT(params.Email, existingUser.ID)
	if err != nil {
		return err
	}

	return c.JSON(jwtToken)
}

func (h *AuthHandler) HandleRegister(c *fiber.Ctx) error {
	var params types.InsertUserParams

	err := c.BodyParser(&params)

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	if err != nil {
		return err
	}

	existingUser, _ := h.userStore.GetUserByEmail(c.Context(), params.Email)

	if existingUser != nil {
		return fmt.Errorf("User already exists")
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	res, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	jwtToken, err := middleware.GenerateJWT(params.Email, res.ID)
	if err != nil {
		return err
	}

	response := struct {
		User  *types.User `json:"user"`
		Token string      `json:"token"`
	}{
		User:  res,
		Token: jwtToken,
	}

	return c.JSON(response)
}
