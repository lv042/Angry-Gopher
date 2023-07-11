package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func setupRoutes(app *fiber.App) {
	app.Post("/register", func(c *fiber.Ctx) error {
		var systemInfo SystemInfo

		if err := json.Unmarshal(c.Body(), &systemInfo); err != nil {
			return fiber.ErrBadRequest
		}
		//get highest id and increment
		id := len(devices) + 1
		device := Device{
			ID:         int8(id),
			systemInfo: systemInfo,
			status:     true,
		}
		devices = append(devices, device)

		//return their id
		return c.JSON(fiber.Map{"id": id})
	})

}

func serverListen(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
