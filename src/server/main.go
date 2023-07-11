package main

import (
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

	setupRoutes(app)
	serverListen(app)

}

func setupDotenv() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
