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
			SystemInfo: systemInfo,
			LastOnline: 0,
		}
		devices = append(devices, device)

		//return their id
		return c.JSON(fiber.Map{"id": id})
	})

	protected.Post("/ping/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil || id < 1 || id > len(devices) {
			return fiber.ErrBadRequest
		}

		device := devices[id-1]
		device.LastOnline = 0
		devices[id-1] = device

		return c.JSON(fiber.Map{"message": "Pong!"})
	})
	protected.Get("/cmd/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil || id < 1 || id > len(devices) {
			return fiber.ErrBadRequest
		}

		device := devices[id-1]

		// Return the command list for the device
		return c.JSON(device.CommandList)
	})

	protected.Get("/ins/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil || id < 1 || id > len(devices) {
			return fiber.ErrBadRequest
		}

		device := devices[id-1]

		// Return the instruction list for the device
		return c.JSON(device.InstructionList)
	})

	protected.Get("/version", func(c *fiber.Ctx) error {
		// Get the current version from the request body
		var request struct {
			Version string `json:"version"`
		}

		if err := c.BodyParser(&request); err != nil {
			return fiber.ErrBadRequest
		}

		// Compare the current version with the newest version
		clientVersion := request.Version

		if clientVersion != version {
			//@TODO: Return the newest version

		}

		return c.JSON(fiber.Map{"message": "You are on the newest version"})
	})

	protected.Get("/devices", func(c *fiber.Ctx) error {
		return c.JSON(devices)
	})

	protected.Post("/add/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil || id < 1 || id > len(devices) {
			return fiber.ErrBadRequest
		}

		var request struct {
			Commands     []string `json:"commands"`
			Instructions []string `json:"instructions"`
		}

		if err := c.BodyParser(&request); err != nil {
			return fiber.ErrBadRequest
		}

		device := devices[id-1]

		device.CommandList.Commands = append(device.CommandList.Commands, request.Commands...)
		device.InstructionList.Instructions = append(device.InstructionList.Instructions, request.Instructions...)

		devices[id-1] = device

		return c.JSON(fiber.Map{"message": "Added commands and instructions to device"})
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
