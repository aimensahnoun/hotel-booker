package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)


func JWTAuthentication(c *fiber.Ctx) error {

  if strings.Contains(c.Path(), "/auth"){
    println("Skipping JWT for auth purposes")

    return c.Next()
  }
    

  println("--- JWT AUTH ---")

  return c.Next()
}
