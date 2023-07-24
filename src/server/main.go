package main

import (
	"github.com/gofiber/fiber/v2"
)

var devices []Device
var app = fiber.New()
var appConfig AppConfig

func main() {
	//setup config
	appConfig = newAppConfig()

	//initialize the dotenv file
	setupDotenv()
	checkForSecret()

	//all tasks that need to be done while the server is running
	updateApplication(app)

	//setup all the routes
	setupRoutes(app)

	//start the server and listen on port 3000
	serverListen(app)
}

func updateApplication(app *fiber.App) {
	updateLastOnline()
	logDevices()

	go func() {
		for {
			//TODO: update this to use a database
		}
	}()
}
