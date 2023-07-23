package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func setupRoutes(app *fiber.App) {
	protected := app.Group("/").Use(authMiddleware)

	protected.Post("/register", handleRegister)

	protected.Post("/ping/:id", handlePing)

	protected.Get("/cmd/:id", handleGetCommandList)

	protected.Post("/cmd/:id", handleUpdateCommandResult)

	protected.Get("/ins/:id", handleGetInstructionList)

	protected.Get("/version", handleCheckVersion)

	protected.Get("/devices", handleGetDevices)

	protected.Post("/add/:id", handleAddCommandsAndInstructions)
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
