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

const version = "0.0.1"

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type Device struct {
	ID              int8       `json:"id"`
	SystemInfo      SystemInfo // Updated: Made it public
	LastOnline      int64
	CommandList     CommandList
	InstructionList InstructionList
}

type InstructionList struct {
	Instructions []string `json:"instructions"`
}

type CommandList struct {
	Commands []string `json:"commands"`
}

func main() {
	app := fiber.New()
	setupDotenv()
	//print out one working token for testing
	token, _ := GenerateToken("admin")
	log.Println("Token for admin: ", token)

	devices = []Device{
		{
			ID: 1,
			SystemInfo: SystemInfo{
				Hostname:     "Device1",
				OS:           "Linux",
				Architecture: "x86_64",
			},
		},
		{
			ID: 2,
			SystemInfo: SystemInfo{
				Hostname:     "Device2",
				OS:           "Windows",
				Architecture: "amd64",
			},
		},
	}

	for i := range devices {
		addRandomCommands(&devices[i].CommandList)
	}

	updateApplication(app)
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
					"SystemInfo": device.SystemInfo,
					"LastOnline": device.LastOnline,
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

func updateApplication(app *fiber.App) {
	updateLastOnline()

	go func() {
		for {
			//TODO: update this to use a database
		}
	}()
}

func updateLastOnline() {
	go func() {
		for {
			for i := range devices {
				// Increment the LastOnline count by 1
				devices[i].LastOnline++
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

func addRandomCommands(commandList *CommandList) {
	commands := []string{"ls -a", "ls", "pwd"}
	commandList.Commands = commands
}
