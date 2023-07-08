package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	dotenv "github.com/joho/godotenv"
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
	setupDotenv()

	token, _ := CreateToken("test")
	fmt.Print("Token works:", VerifyToken(token))

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

func setupDotenv() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
