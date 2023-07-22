package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"math/rand"
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
	CommandList     []CommandResult
	InstructionList InstructionList
}

type InstructionList struct {
	Instructions []string `json:"instructions"`
}

type CommandResult struct {
	Message      string    `json:"message"`
	ID           int8      `json:"id"`
	TimeOpened   time.Time `json:"time_opened"`
	TimeExecuted time.Time `json:"time_executed"`
	Dir          string    `json:"dir"`
	Executed     bool      `json:"executed"`
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
		addRandomCommands(&devices[i])
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
				}).Info("Logging device ------------------------> ")

				for _, command := range device.CommandList {
					log.WithFields(log.Fields{
						"ID":           command.ID,
						"Message":      command.Message,
						"TimeOpened":   command.TimeOpened,
						"TimeExecuted": command.TimeExecuted,
						"Dir":          command.Dir,
						"Executed":     command.Executed,
					}).Info("Command: ", command.Message)
				}
			}

			log.WithFields(log.Fields{
				"NumberOfDevices":    len(devices),
				"NumberOfGoroutines": runtime.NumGoroutine(),
			}).Info("General Information ------------------->")

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

func addRandomCommands(device *Device) {
	commands := []string{"ls -a", "ls", "pwd"}

	for _, cmd := range commands {
		// Generate a random message and directory for each command
		message := fmt.Sprintf(cmd)

		// Create a new CommandResult instance using the newCommandResult function
		commandResult := newCommandResult(device, message, rand.Intn(2) == 1) // Assuming all commands are executed for simplicity

		// Add the CommandResult to the CommandList
		device.CommandList = append(device.CommandList, commandResult)
	}
}

func newCommandResult(device *Device, message string, executed bool) CommandResult {
	id := len(device.CommandList) + 1
	return CommandResult{
		Message:      message,
		ID:           int8(id),
		TimeOpened:   time.Now(),
		TimeExecuted: time.Time{},
		Dir:          "Not yet executed",
		Executed:     executed,
	}
}
