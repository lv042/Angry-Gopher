package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var devices []Device
var app = fiber.New()

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

	setupDotenv()
	//print out one working token for testing
	token, _ := GenerateToken("admin")
	log.Println("Token for admin: ", token)

	generateFakeData()
	updateApplication(app)
	logDevices()
	setupRoutes(app)
	serverListen(app)
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
