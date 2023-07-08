package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

var devices []Device

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type Device struct {
	ID         int8 `json:"id"`
	systemInfo SystemInfo
	status     bool
}

func main() {
	app := fiber.New()

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

	//listen
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
