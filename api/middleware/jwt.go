package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func JWTAuthentication(c *fiber.Ctx) error {
	if strings.Contains(c.Path(), "/auth") {
		println("Skipping JWT for auth purposes")

		return c.Next()
	}

	println("--- JWT AUTH ---")

	return c.Next()
}

func GenerateJWT(email string) (string ,error){
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	expiration := time.Now().Add(7 * 24 * time.Hour)

claims:= &Claims{
    Email : email,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: expiration.Unix(),
    },
  } 

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  tokenString , err := token.SignedString(jwtSecret)
  
  if err != nil {
    return "" ,err
  }


  return tokenString , nil
}
