package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

func JWTAuthentication(c *fiber.Ctx) error {
	if strings.Contains(c.Path(), "/auth") {
		return c.Next()
	}

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	tokenString := strings.Split(authHeader, " ")[1]

	claims, err := parseJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	println("Authenticated email:", claims.Email)

	if time.Now().Unix() > claims.ExpiresAt {
		return fmt.Errorf("Token expired")
	}

	userID := claims.ID

	c.Context().SetUserValue("id", userID)

	return c.Next()
}

func GenerateJWT(email string, id string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	expiration := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		Email: email,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func parseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
