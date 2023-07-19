package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"

	"runtime"
	"time"
)

var devices []Device

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type Device struct {
	ID         int8       `json:"id"`
	systemInfo SystemInfo // Updated: Made it public
	lastOnline int64
}

func main() {
	app := fiber.New()
	setupDotenv()
	//print out one working token for testing
	token, _ := GenerateToken("admin")
	log.Println("Token for admin: ", token)

	logDevices()
	setupRoutes(app)
	serverListen(app)
}

func logDevices() {
	go func() {
		for {
			for _, device := range devices {
				log.WithFields(log.Fields{
					"ID":         device.ID,
					"SystemInfo": device.systemInfo,
					"LastOnline": device.lastOnline,
				}).Info("Logging device -> ")
			}

			log.WithFields(log.Fields{
				"NumberOfDevices":    len(devices),
				"NumberOfGoroutines": runtime.NumGoroutine(),
			}).Info("Logging device information -> ")

			time.Sleep(10 * time.Second)
		}
	}()
}
func setupDotenv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// print all the env variables
	log.Println("Environment variables:")
	for _, e := range os.Environ() {
		log.Println(e)
	}
}
