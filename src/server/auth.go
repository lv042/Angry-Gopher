package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

func VerifyToken(tokenString string) bool {

	if appConfig.SecretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(appConfig.SecretKey), nil
	})

	if err != nil {
		return false
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		return time.Now().Before(expirationTime)
	}

	return false
}

func GenerateToken(host string, id int, validity time.Duration) (string, error) {

	if appConfig.SecretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	if host == "" {
		return "", fmt.Errorf("host cannot be empty")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["host"] = host
	claims["exp"] = time.Now().Add(validity).Unix()

	tokenString, err := token.SignedString([]byte(appConfig.SecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func authMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	// Check if the token is present and properly formatted
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or invalid token",
		})
	}

	//Search if Bearer is in the string
	if len(token) > 6 && token[:7] == "Bearer " {
		token = token[7:]
	}

	tokenValid := VerifyToken(token)

	if !tokenValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	return c.Next()
}
