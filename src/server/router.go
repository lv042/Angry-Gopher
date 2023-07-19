package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func setupRoutes(app *fiber.App) {

	protected := app.Group("/").Use(authMiddleware)

	protected.Post("/register", func(c *fiber.Ctx) error {
		var systemInfo SystemInfo

		if err := json.Unmarshal(c.Body(), &systemInfo); err != nil {
			return fiber.ErrBadRequest
		}
		//get highest id and increment
		id := len(devices) + 1
		device := Device{
			ID:         int8(id),
			systemInfo: systemInfo,
			lastOnline: 0,
		}
		devices = append(devices, device)

		//return their id
		return c.JSON(fiber.Map{"id": id})
	})
	protected.Get("/ping:id", func(c *fiber.Ctx) error {
		var id, err = strconv.Atoi(c.Params("id"))
		if err != nil {
			return fiber.ErrBadRequest
		}
		devices[int(id)-1].lastOnline = 0

		return c.SendString("pong")
	})

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

func serverListen(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
